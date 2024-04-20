package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"01.kood.tech/git/rols55/social-network/pkg/api/model"
	"01.kood.tech/git/rols55/social-network/pkg/api/route/middleware/acl"
	"01.kood.tech/git/rols55/social-network/pkg/logger"
)

type ProfileResponse struct {
	Status string        `json:"status"`
	User   *model.User   `json:"user"`
	Posts  []*model.Post `json:"posts,omitempty"`
	Users  []*model.User `json:"users,omitempty"`
	Show   bool          `json:"show,omitempty"`
}

func (h *BaseController) ProfileGET(w http.ResponseWriter, r *http.Request) {
	var err error
	var posts []*model.Post
	var user *model.User
	if err = h.getKeyVal(r, acl.UserKey, &user); err != nil {
		logger.Error(err)
		h.statusJSON(w, Unauthorized)
		return
	}

	// get user's created posts from the database, throw an http.error if there is error and its not ErrNotFound (no posts)
	if posts, err = model.GetPostsByUserUuid(h.db, user.UUID); err != nil && err != model.ErrNotFound {
		logger.Error(err)
		h.statusJSON(w, InternalError)
		return
	}

	res := ProfileResponse{
		Status: "OK",
		User:   user,
		Posts:  posts,
		Show:   true,
	}

	h.writeJSON(w, res)
}

func (h *BaseController) UsersGET(w http.ResponseWriter, r *http.Request) {
	var err error
	var user *model.User
	var users []*model.User

	if err = h.getKeyVal(r, acl.UserKey, &user); err != nil {
		logger.Error(err)
		h.statusJSON(w, Unauthorized)
		return
	}
	/*
		if users, err = model.GetUsers(h.db, user.Id); err != nil && err != model.ErrNotFound {
			logger.Error(err)
			h.statusJSON(w, InternalError)
		}
	*/
	if users, err = model.GetUsersSorted(h.db, user); err != nil && err != model.ErrNotFound {
		logger.Error(err)
		h.statusJSON(w, InternalError)
		return
	}

	res := ProfileResponse{
		Status: "OK",
		User:   user,
		Users:  users,
	}

	h.writeJSON(w, res)
}

func (h *BaseController) ChatUsersGET(w http.ResponseWriter, r *http.Request) {
	var err error
	var user *model.User
	var users []*model.User

	if err = h.getKeyVal(r, acl.UserKey, &user); err != nil {
		logger.Error(err)
		h.statusJSON(w, Unauthorized)
		return
	}

	if users, err = model.GetUsersSorted(h.db, user); err != nil && err != model.ErrNotFound {
		logger.Error(err)
		h.statusJSON(w, InternalError)
		return
	}

	filteredChat := make([]*model.User, 0)
	for _, user2 := range users {
		if model.ChatFilter(h.db, user.UUID, user2.UUID) {
			filteredChat = append(filteredChat, user2)
		}
	}

	res := ProfileResponse{
		Status: "OK",
		User:   user,
		Users:  filteredChat,
	}

	h.writeJSON(w, res)
}

func (h *BaseController) GetImageUser(w http.ResponseWriter, r *http.Request) {
	http.StripPrefix("/api/getimageuser/", http.FileServer(http.Dir("media/profile"))).ServeHTTP(w, r)
}

func (h *BaseController) OtherUserGET(w http.ResponseWriter, r *http.Request) {
	var err error
	var user *model.User
	var posts []*model.Post

	//check if the user is logged in
	if err = h.getKeyVal(r, acl.UserKey, &user); err != nil {
		logger.Error(err)
		h.statusJSON(w, InternalError)
	}

	targetUuid := r.URL.Query().Get("uuid")
	if err != nil {
		logger.Error(err)
		h.statusJSON(w, NotFound)
		return
	}

	target, err := model.GetUserByUUID(h.db, targetUuid)
	if err != nil {
		logger.Error(err)
		h.statusJSON(w, NotFound)
		return
	}

	if posts, err = model.GetPostsByUserUuid(h.db, targetUuid); err != nil && err != model.ErrNotFound {
		logger.Error(err)
		h.statusJSON(w, InternalError)
		return
	}

	var show bool
	if targetUuid == user.UUID {
		show = true
	} else {
		show = model.AllowShow(h.db, user.UUID, targetUuid)
	}

	res := ProfileResponse{
		Status: "OK",
		User:   target,
		Posts:  posts,
		Show:   show,
	}
	h.writeJSON(w, res)
}

func (h *BaseController) ProfileUpdatePOST(w http.ResponseWriter, r *http.Request) {
	var err error
	var user *model.User
	//check if the user is logged in
	if err = h.getKeyVal(r, acl.UserKey, &user); err != nil {
		logger.Error(err)
		h.statusJSON(w, InternalError)
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		logger.Error(err)
		h.statusJSON(w, InternalError)
		return
	}
	// Define a struct to hold the JSON data
	var target *model.User
	// Parse the JSON data into the struct
	err = json.Unmarshal(body, &target)
	if err != nil {
		logger.Error(err)
		h.statusJSON(w, InternalError)
		return
	}
	err = target.UpdateUser(h.db)
	if err != nil {
		logger.Error(err)
		h.statusJSON(w, InternalError)
		return
	}
	h.statusJSON(w, Success)
}
