package route

import (
	"database/sql"
	"forum/controllers"
	"forum/route/middleware/acl"
	"forum/route/middleware/log"
	"forum/route/middleware/method"
	"forum/shared/logger"
	"forum/ws"
	"net/http"
)

var (
	GET  = method.Method("GET")  // Only allow GET requests
	POST = method.Method("POST") // Only allow POST requests
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

func Load() http.Handler {
	return middleware(routes())
}

func routes() *http.ServeMux {

	r := Mux{Mux: http.NewServeMux()}
	db, _ := sql.Open("sqlite3", "./forum.db?_foreign_keys=on")
	cp := ws.New(db)
	controllers := controllers.New(db, cp)

	go cp.Run()
	r.Handle("/ws", GET).ThenFunc(controllers.WebSocket)

	r.Handle("/", GET).ThenFunc(controllers.Index)
	r.Handle("/post/", GET).ThenFunc(controllers.PostGET)
	r.Handle("/post/create", POST).ThenFunc(controllers.CreatePostPOST) //no alc, controller handles it, we need unified way to report errors (session expired)
	r.Handle("/post/delete", POST).ThenFunc(controllers.DeletePost)

	r.Handle("/comment/create", POST).ThenFunc(controllers.CreateCommentPOST) //no alc, controller handles it, we need unified way to report errors (session expired)
	r.Handle("/comment/delete", POST, acl.DisallowAnon).ThenFunc(controllers.DeleteComment)

	r.Handle("/login", POST).ThenFunc(controllers.LoginPOST) //no alc, controller handles it, we need unified way to report errors (session expired)
	r.Handle("/logout", GET).ThenFunc(controllers.Logout)
	r.Handle("/register", POST, acl.DisallowAuth).ThenFunc(controllers.RegisterPOST)

	r.Handle("/users", GET, acl.DisallowAnon).ThenFunc(controllers.UsersGET)

	r.Handle("/history", GET).ThenFunc(controllers.GetHistory)

	staticDirectory := http.Dir("static")
	staticServer := http.FileServer(staticDirectory)
	r.Handle("/static/").Then(http.StripPrefix("/static/", staticServer))

	return r.Mux
}

// this will applay middleware to all controllers
func middleware(h http.Handler) http.Handler {
	//Print info about the request
	h = log.Log(h)
	h = acl.AddUser(h)
	return h
}
