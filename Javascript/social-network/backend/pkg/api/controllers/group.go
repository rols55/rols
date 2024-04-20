package controllers

import (
	"net/http"
	"path"
	"strconv"
	"strings"
	"time"

	"01.kood.tech/git/rols55/social-network/pkg/api/model"
	"01.kood.tech/git/rols55/social-network/pkg/api/route/middleware/acl"
	"01.kood.tech/git/rols55/social-network/pkg/api/ws"
	"01.kood.tech/git/rols55/social-network/pkg/logger"
)

type GroupsResponse struct {
	Status        string           `json:"status"`
	User          *model.User      `json:"user,omitempty"`
	Groups        []*model.Group   `json:"groups,omitempty"`
	Group         *model.Group     `json:"group,omitempty"`
	Members       []*model.User    `json:"members,omitempty"`
	Posts         []*model.Post    `json:"posts,omitempty"`
	GroupRequests []*GroupRequests `json:"group_requests,omitempty"`
	GroupState    string           `json:"group_state,omitempty"`
}

type GroupRequests struct {
	GroupId    int64    `json:"group_id,omitempty"`
	GroupTitle string   `json:"group_title,omitempty"`
	Usernames  []string `json:"usernames,omitempty"`
	UserUUIDs  []string `json:"user_uuids,omitempty"`
}

func (h *BaseController) GroupsGET(w http.ResponseWriter, r *http.Request) {
	var err error
	var user *model.User
	var groups []*model.Group

	if err = h.getKeyVal(r, acl.UserKey, &user); err != nil {
		logger.Error(err)
		h.statusJSON(w, InternalError)
		return
	}

	uuid := r.URL.Query().Get("uuid")
	if uuid != "" {
		filterUser, err := model.GetUserByUUID(h.db, uuid)
		if err != nil {
			logger.Error(err)
			h.statusJSON(w, NotFound)
			return
		}

		groups, err = filterUser.GetGroups(h.db)
		if err != nil && err != model.ErrNotFound {
			logger.Error(err)
			h.statusJSON(w, InternalError)
			return
		}
	} else {

		if groups, err = model.GetGroups(h.db); err != nil && err != model.ErrNotFound {
			logger.Error(err)
			h.statusJSON(w, InternalError)
			return
		}
	}

	res := GroupsResponse{
		Status: "OK",
		User:   user,
		Groups: groups,
	}

	h.writeJSON(w, res)
}

func (h *BaseController) GroupPOST(w http.ResponseWriter, r *http.Request) {
	var err error
	var user *model.User

	//check if the user is logged in
	if err = h.getKeyVal(r, acl.UserKey, &user); err != nil {
		logger.Warning("%v", err)
		h.statusJSON(w, Unauthorized)
		return
	}

	if r.FormValue("title") == "" {
		logger.Info("Failed to create group: No title provided")
		errMsg := "Title missing"
		h.statusJSON(w, InvalidInput(errMsg))
		return
	}

	if len(r.FormValue("title")) > 60 {
		logger.Info("Failed to create group: Title too long")
		errMsg := "Title too long"
		h.statusJSON(w, InvalidInput(errMsg))
		return
	}

	if r.FormValue("description") == "" {
		logger.Info("Failed to create group: No description provided")
		errMsg := "Description missing"
		h.statusJSON(w, InvalidInput(errMsg))
		return
	}

	group := &model.Group{
		UserId:      user.Id,
		Title:       r.FormValue("title"),
		Description: r.FormValue("description"),
	}
	group.Create(h.db)
	logger.Info("Group %v created, id = %v", group.Title, group.Id)

	if r.FormValue("invite") != "" {
		inviteUsers := strings.Split(r.FormValue("invite"), ",")
		logger.Info("Inviting users to group: %v", inviteUsers)
		if err = group.InviteUsers(h.db, user.Id, inviteUsers); err != nil {
			logger.Info("Failed to invite users to group: %v", err)
			h.statusJSON(w, InternalError)
			return
		}

		for _, uuid := range inviteUsers {
			notification := model.Notification{}
			currentTime := time.Now()

			err = notification.InsertNotification(h.db, user.UUID, uuid, "You have new group invite from "+user.Username, "invite", int(group.Id), currentTime, false)
			if err != nil {
				panic(err)
			}

			notificationData := ws.DataJSON{
				Type:         "notification", // Or any identifier for notifications
				ReciverUUID:  uuid,
				Notification: "You have new group invite from " + user.Username,
				Timestamp:    currentTime,
			}

			h.cp.Notify(notificationData) // Correct way to call the Notify method.
		}
	}

	res := GroupsResponse{
		Status: "OK",
		User:   user,
		Group:  group,
	}

	h.writeJSON(w, res)
}

func (h *BaseController) GroupGET(w http.ResponseWriter, r *http.Request) {
	var err error
	var user *model.User
	var groupId int64
	var group *model.Group
	var members []*model.User
	var posts []*model.Post

	if err = h.getKeyVal(r, acl.UserKey, &user); err != nil {
		logger.Warning("%v", err)
		h.statusJSON(w, Unauthorized)
		return
	}

	groupId, err = strconv.ParseInt(path.Base(r.URL.Path), 10, 64)
	if err != nil {
		logger.Error(err)
		h.statusJSON(w, NotFound)
		return
	}

	group, err = model.GetGroupById(h.db, groupId)
	if err != nil || group == nil {
		logger.Error(err)
		h.statusJSON(w, NotFound)
		return
	}

	members, err = group.GetMemebers(h.db)
	if err != nil {
		logger.Error(err)
		h.statusJSON(w, InternalError)
		return
	}

	var groupState string
	if group.UserId == user.Id {
		groupState = "accepted"
	} else {
		groupState, err = group.GetUserInviteState(h.db, user.Id)
		if err != nil && err != model.ErrNotFound {
			logger.Error(err)
			h.statusJSON(w, InternalError)
			return
		}
	}

	if groupState == "accepted" {
		posts, err = group.GetPosts(h.db)
		if err != nil && err != model.ErrNotFound {
			logger.Error(err)
			h.statusJSON(w, InternalError)
			return
		}
	}

	res := GroupsResponse{
		Status:     "OK",
		User:       user,
		Group:      group,
		Members:    members,
		Posts:      posts,
		GroupState: groupState,
	}

	h.writeJSON(w, res)
}

func (h *BaseController) InvitesGET(w http.ResponseWriter, r *http.Request) {
	var err error
	var user *model.User
	var groups []*model.Group

	if err = h.getKeyVal(r, acl.UserKey, &user); err != nil {
		logger.Warning("%v", err)
		h.statusJSON(w, Unauthorized)
		return
	}

	groups, err = user.GetGroupInvites(h.db)
	if err != nil && err != model.ErrNotFound {
		logger.Error(err)
		h.statusJSON(w, InternalError)
		return
	}

	res := GroupsResponse{
		Status: "OK",
		User:   user,
		Groups: groups,
	}

	h.writeJSON(w, res)
}

func (h *BaseController) AcceptInvitePOST(w http.ResponseWriter, r *http.Request) {
	var err error
	var user *model.User

	if err = h.getKeyVal(r, acl.UserKey, &user); err != nil {
		logger.Warning("%v", err)
		h.statusJSON(w, Unauthorized)
		return
	}

	if r.FormValue("group_id") == "" {
		logger.Info("Failed to accept group invite: No group id provided")
		errMsg := "Group Id missing"
		h.statusJSON(w, InvalidInput(errMsg))
		return
	}

	groupId, err := strconv.ParseInt(r.FormValue("group_id"), 10, 64)
	if err != nil {
		logger.Error(err)
		h.statusJSON(w, NotFound)
		return
	}

	if err = user.AcceptGroupInvite(h.db, groupId); err != nil {
		logger.Error(err)
		h.statusJSON(w, NotFound)
	}

	h.writeJSON(w, Success)
}

func (h *BaseController) RejectInvitePOST(w http.ResponseWriter, r *http.Request) {
	var err error
	var user *model.User

	if err = h.getKeyVal(r, acl.UserKey, &user); err != nil {
		logger.Warning("%v", err)
		h.statusJSON(w, Unauthorized)
		return
	}

	if r.FormValue("group_id") == "" {
		logger.Info("Failed to reject group invite: No group id provided")
		errMsg := "Group Id missing"
		h.statusJSON(w, InvalidInput(errMsg))
		return
	}

	groupId, err := strconv.ParseInt(r.FormValue("group_id"), 10, 64)
	if err != nil {
		logger.Error(err)
		h.statusJSON(w, NotFound)
		return
	}

	if err = user.RejectGroupInvite(h.db, groupId); err != nil {
		logger.Error(err)
		h.statusJSON(w, NotFound)
	}

	h.writeJSON(w, Success)
}

func (h *BaseController) ApproveRequestPOST(w http.ResponseWriter, r *http.Request) {
	var err error
	var user *model.User

	if err = h.getKeyVal(r, acl.UserKey, &user); err != nil {
		logger.Warning("%v", err)
		h.statusJSON(w, Unauthorized)
		return
	}

	if r.FormValue("group_id") == "" {
		logger.Info("Failed to approve group invite: No group id provided")
		errMsg := "Group Id missing"
		h.statusJSON(w, InvalidInput(errMsg))
		return
	}

	if r.FormValue("requester_uuid") == "" {
		logger.Info("Failed to approve group invite: No requester id provided")
		errMsg := "Invite requester Id missing"
		h.statusJSON(w, InvalidInput(errMsg))
		return
	}

	groupId, err := strconv.ParseInt(r.FormValue("group_id"), 10, 64)
	if err != nil {
		logger.Error(err)
		h.statusJSON(w, InternalError)
		return
	}

	reqUser, err := model.GetUserByUUID(h.db, r.FormValue("requester_uuid"))
	if err != nil {
		logger.Error(err)
		h.statusJSON(w, NotFound)
		return
	}

	if err = user.ApproveRequest(h.db, groupId, reqUser.Id); err != nil {
		logger.Error(err)
		h.statusJSON(w, InternalError)
		return
	}

	h.writeJSON(w, Success)
}

func (h *BaseController) RejectRequestPOST(w http.ResponseWriter, r *http.Request) {
	var err error
	var user *model.User

	if err = h.getKeyVal(r, acl.UserKey, &user); err != nil {
		logger.Warning("%v", err)
		h.statusJSON(w, Unauthorized)
		return
	}

	if r.FormValue("group_id") == "" {
		logger.Info("Failed to approve group invite: No group id provided")
		errMsg := "Group Id missing"
		h.statusJSON(w, InvalidInput(errMsg))
		return
	}

	if r.FormValue("requester_uuid") == "" {
		logger.Info("Failed to approve group invite: No requester id provided")
		errMsg := "Invite requester Id missing"
		h.statusJSON(w, InvalidInput(errMsg))
		return
	}

	groupId, err := strconv.ParseInt(r.FormValue("group_id"), 10, 64)
	if err != nil {
		logger.Error(err)
		h.statusJSON(w, InternalError)
		return
	}

	reqUser, err := model.GetUserByUUID(h.db, r.FormValue("requester_uuid"))
	if err != nil {
		logger.Error(err)
		h.statusJSON(w, NotFound)
		return
	}

	if err = user.RejectRequest(h.db, groupId, reqUser.Id); err != nil {
		logger.Error(err)
		h.statusJSON(w, InternalError)
		return
	}

	h.writeJSON(w, Success)
}

func (h *BaseController) RequestsGET(w http.ResponseWriter, r *http.Request) {
	var err error
	var user *model.User
	var groups []*model.Group

	if err = h.getKeyVal(r, acl.UserKey, &user); err != nil {
		logger.Warning("%v", err)
		h.statusJSON(w, Unauthorized)
		return
	}

	groupIdStr := r.URL.Query().Get("group")
	if groupIdStr != "" {

		groupId, err := strconv.ParseInt(groupIdStr, 10, 64)
		if err != nil {
			logger.Error(err)
			h.statusJSON(w, InvalidInput("Invalid Group ID"))
			return
		}

		group, err := model.GetGroupById(h.db, groupId)
		if err != nil {
			logger.Error(err)
			h.statusJSON(w, NotFound)
			return
		}

		groups = append(groups, group)
	} else {
		groups, err = user.GetGroups(h.db)
		if err != nil && err != model.ErrNotFound {
			logger.Error(err)
			h.statusJSON(w, InternalError)
			return
		}
	}

	groupRequests := make([]*GroupRequests, 0)
	for _, group := range groups {
		var request GroupRequests
		users, err := group.GetRequests(h.db)
		if err != nil && err != model.ErrNotFound {
			logger.Error(err)
		}
		if len(users) > 0 {
			request.GroupId = group.Id
			request.GroupTitle = group.Title
			for _, reqUser := range users {
				request.Usernames = append(request.Usernames, reqUser.Username)
				request.UserUUIDs = append(request.UserUUIDs, reqUser.UUID)
			}
			groupRequests = append(groupRequests, &request)
		}
	}

	res := GroupsResponse{
		Status:        "OK",
		User:          user,
		GroupRequests: groupRequests,
	}

	h.writeJSON(w, res)

}

func (h *BaseController) RequestInvitePOST(w http.ResponseWriter, r *http.Request) {
	var err error
	var user *model.User

	if err = h.getKeyVal(r, acl.UserKey, &user); err != nil {
		logger.Warning("%v", err)
		h.statusJSON(w, Unauthorized)
		return
	}

	if r.FormValue("group_id") == "" {
		logger.Info("Failed to request group invite: No group id provided")
		errMsg := "Group Id missing"
		h.statusJSON(w, InvalidInput(errMsg))
		return
	}

	groupId, err := strconv.ParseInt(r.FormValue("group_id"), 10, 64)
	if err != nil {
		logger.Error(err)
		h.statusJSON(w, InternalError)
		return
	}

	group, err := model.GetGroupById(h.db, groupId)
	if err != nil {
		logger.Error(err)
		h.statusJSON(w, NotFound)
		return
	}

	err = group.RequestInvite(h.db, user.Id)
	if err != nil {
		logger.Error(err)
		h.statusJSON(w, NotFound)
		return
	}

	author, err := model.GetUserById(h.db, group.UserId)
	if err != nil {
		logger.Error(err)
		h.statusJSON(w, InternalError)
		return
	}

	notification := model.Notification{}
	currentTime := time.Now()

	err = notification.InsertNotification(h.db, user.UUID, author.UUID, "You have new group request from "+user.Username, "request", int(group.Id), currentTime, false)
	if err != nil {
		panic(err)
	}

	notificationData := ws.DataJSON{
		Type:         "notification",
		ReciverUUID:  author.UUID,
		Notification: "You have new group request from " + user.Username,
		Timestamp:    currentTime,
	}

	h.cp.Notify(notificationData)

	h.writeJSON(w, Success)
}

func (h *BaseController) InvitePOST(w http.ResponseWriter, r *http.Request) {
	var err error
	var user *model.User

	if err = h.getKeyVal(r, acl.UserKey, &user); err != nil {
		logger.Warning("%v", err)
		h.statusJSON(w, Unauthorized)
		return
	}

	if r.FormValue("group_id") == "" {
		logger.Info("Failed to request group invite: No group id provided")
		errMsg := "Group Id missing"
		h.statusJSON(w, InvalidInput(errMsg))
		return
	}

	groupId, err := strconv.ParseInt(r.FormValue("group_id"), 10, 64)
	if err != nil {
		logger.Error(err)
		h.statusJSON(w, InternalError)
		return
	}

	group, err := model.GetGroupById(h.db, groupId)
	if err != nil {
		logger.Error(err)
		h.statusJSON(w, NotFound)
		return
	}

	if r.FormValue("invite") == "" {
		logger.Info("Failed to invite to group: No users provided")
		errMsg := "Users missing"
		h.statusJSON(w, InvalidInput(errMsg))
		return
	}

	inviteUsers := strings.Split(r.FormValue("invite"), ",")
	logger.Info("Inviting users to group (%v): %v", group.Id, inviteUsers)
	if err = group.InviteUsers(h.db, user.Id, inviteUsers); err != nil {
		logger.Info("Failed to invite users to group: %v", err)
		h.statusJSON(w, InternalError)
		return
	}

	for _, uuid := range inviteUsers {
		notification := model.Notification{}
		currentTime := time.Now()

		err = notification.InsertNotification(h.db, user.UUID, uuid, "You have new group invite from "+user.Username, "invite", int(group.Id), currentTime, false)
		if err != nil {
			panic(err)
		}

		notificationData := ws.DataJSON{
			Type:         "notification", // Or any identifier for notifications
			ReciverUUID:  uuid,
			Notification: "You have new group invite from " + user.Username,
			Timestamp:    currentTime,
		}

		h.cp.Notify(notificationData) // Correct way to call the Notify method.
	}

	res := GroupsResponse{
		Status: "OK",
		User:   user,
		Group:  group,
	}

	h.writeJSON(w, res)
}
