package route

import (
	"database/sql"
	"net/http"

	"01.kood.tech/git/rols55/social-network/pkg/api/controllers"
	"01.kood.tech/git/rols55/social-network/pkg/api/route/middleware/acl"
	"01.kood.tech/git/rols55/social-network/pkg/api/route/middleware/cors"
	"01.kood.tech/git/rols55/social-network/pkg/api/route/middleware/log"
	"01.kood.tech/git/rols55/social-network/pkg/api/route/middleware/method"
	"01.kood.tech/git/rols55/social-network/pkg/api/ws"
	"01.kood.tech/git/rols55/social-network/pkg/logger"
)

var (
	GET          = method.Method("GET")             // Only allow GET requests
	POST         = method.Method("POST")            // Only allow POST requests
	POST_OPTIONS = method.Method("POST", "OPTIONS") // Only allow POST & OPTIONS requests
)

type Mux struct {
	Mux        *http.ServeMux
	Middleware []func(http.Handler) http.Handler
	Pattern    string
}

func (m *Mux) Then(handler http.Handler) {
	if handler == nil {
		logger.Error("HTTP: nil handler")
		return
	}
	result := handler

	for idx := range m.Middleware {
		result = m.Middleware[len(m.Middleware)-idx-1](result)
	}
	m.Mux.Handle(m.Pattern, result)
}

func (m *Mux) ThenFunc(fn http.HandlerFunc) {
	if fn == nil {
		m.Then(nil)
		return
	}
	m.Then(fn)
}

func (m *Mux) Handle(pattern string, middleware ...func(http.Handler) http.Handler) *Mux {
	return &Mux{Mux: m.Mux, Pattern: pattern, Middleware: middleware}
}

func Load(db *sql.DB) http.Handler {
	return middleware(routes(db))
}

func routes(db *sql.DB) *http.ServeMux {

	r := Mux{Mux: http.NewServeMux()}
	//db, _ := sql.Open("sqlite3", "./forum.db?_foreign_keys=on")
	cp := ws.New(db)
	controllers := controllers.New(db, cp)

	go cp.Run()
	r.Handle("/api/ws", GET).ThenFunc(controllers.WebSocket)

	r.Handle("/api", GET).ThenFunc(controllers.Index)

	r.Handle("/api/check-session", POST).ThenFunc(controllers.CheckSession)

	//Posts
	r.Handle("/api/feed", GET).ThenFunc(controllers.FeedGET)
	r.Handle("/api/post/", GET).ThenFunc(controllers.PostGET)
	r.Handle("/api/posts", GET).ThenFunc(controllers.PostsGET)
	r.Handle("/api/post/create", POST).ThenFunc(controllers.CreatePostPOST)
	r.Handle("/api/post/delete", POST).ThenFunc(controllers.DeletePost)
	r.Handle("/api/getimage/", GET).ThenFunc(controllers.GetImagePosts)
	//Notifications
	r.Handle("/api/notification/", GET).ThenFunc(controllers.NotificationGET)
	r.Handle("/api/notifications/", GET).ThenFunc(controllers.NotificationsGET)

	r.Handle("/api/comment/create", POST).ThenFunc(controllers.CreateCommentPOST)
	r.Handle("/api/comment/delete", POST, acl.DisallowAnon).ThenFunc(controllers.DeleteComment)

	r.Handle("/api/login", POST).ThenFunc(controllers.LoginPOST)
	r.Handle("/api/logout", GET).ThenFunc(controllers.Logout)
	r.Handle("/api/register", POST).ThenFunc(controllers.RegisterPOST)

	//Profile
	r.Handle("/api/users", GET).ThenFunc(controllers.UsersGET)
	r.Handle("/api/chatusers", GET).ThenFunc(controllers.ChatUsersGET)
	r.Handle("/api/profile", GET).ThenFunc(controllers.ProfileGET)
	r.Handle("/api/otherprofile", GET).ThenFunc(controllers.OtherUserGET)
	r.Handle("/api/profile/edit", POST).ThenFunc(controllers.ProfileUpdatePOST)
	r.Handle("/api/user/publictoggle", method.Method("POST", "OPTIONS")).ThenFunc(controllers.TogglePublic)
	r.Handle("/api/getimageuser/", GET).ThenFunc(controllers.GetImageUser)

	r.Handle("/api/history", GET).ThenFunc(controllers.GetHistory)

	//Follow
	r.Handle("/api/followers", GET).ThenFunc(controllers.RequestFollowInfo)
	r.Handle("/api/otherfollowers", GET).ThenFunc(controllers.RequestOtherFollowInfo)
	r.Handle("/api/followers/accept", POST).ThenFunc(controllers.AcceptFollow)
	r.Handle("/api/followers/request", POST).ThenFunc(controllers.RequestFollow)
	r.Handle("/api/followers/unfollow", POST).ThenFunc(controllers.CancelFollow)
	r.Handle("/api/followers/cancel", POST).ThenFunc(controllers.DeclineFollow)

	//Groups
	r.Handle("/api/groups", GET).ThenFunc(controllers.GroupsGET)
	r.Handle("/api/group", POST_OPTIONS).ThenFunc(controllers.GroupPOST)
	r.Handle("/api/group/", GET).ThenFunc(controllers.GroupGET)
	r.Handle("/api/groups/invites", GET).ThenFunc(controllers.InvitesGET)
	r.Handle("/api/groups/invite", POST_OPTIONS).ThenFunc(controllers.InvitePOST)
	r.Handle("/api/groups/accept", POST_OPTIONS).ThenFunc(controllers.AcceptInvitePOST)
	r.Handle("/api/groups/reject", POST_OPTIONS).ThenFunc(controllers.RejectInvitePOST)
	r.Handle("/api/groups/request", POST_OPTIONS).ThenFunc(controllers.RequestInvitePOST)
	r.Handle("/api/groups/requests", GET).ThenFunc(controllers.RequestsGET)
	r.Handle("/api/groups/request/approve", POST_OPTIONS).ThenFunc(controllers.ApproveRequestPOST)
	r.Handle("/api/groups/request/reject", POST_OPTIONS).ThenFunc(controllers.RejectRequestPOST)

	//Events
	//r.Handle("/api/events", GET).ThenFunc(controllers.EventsGET)
	r.Handle("/api/events", GET).ThenFunc(controllers.GroupEventsGET)
	r.Handle("/api/event", POST_OPTIONS).ThenFunc(controllers.EventPOST)
	r.Handle("/api/event/", GET).ThenFunc(controllers.EventGET)
	r.Handle("/api/events/attend", POST_OPTIONS).ThenFunc(controllers.EventAttendPOST)
	r.Handle("/api/events/attendance", GET).ThenFunc(controllers.EventAttendGET)

	//Group Chat
	r.Handle("/api/history/group", GET).ThenFunc(controllers.GetGroupHistory)

	return r.Mux
}

// this will applay middleware to all controllers
func middleware(h http.Handler) http.Handler {
	//Print info about the request
	h = log.Log(h)
	h = cors.Cors(h)
	h = acl.AddUser(h)
	return h
}
