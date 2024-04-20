package controllers

import (
	"net/http"
	"time"

	"01.kood.tech/git/rols55/social-network/pkg/api/model"
	"01.kood.tech/git/rols55/social-network/pkg/api/route/middleware/acl"
	"01.kood.tech/git/rols55/social-network/pkg/api/ws"
	"01.kood.tech/git/rols55/social-network/pkg/logger"
)

type FollowInfo struct {
	Status    string            `json:"status"`
	Following []*model.Follower `json:"following,omitempty"`
	Followers []*model.Follower `json:"followers,omitempty"`
}

func (h *BaseController) RequestFollowInfo(w http.ResponseWriter, r *http.Request) {
	var user *model.User
	var err error
	if err = h.getKeyVal(r, acl.UserKey, &user); err != nil {
		logger.Error(err)
		h.statusJSON(w, Unauthorized)
		return
	}

	followers, err := model.GetFollowers(h.db, user.UUID)
	if err != nil {
		logger.Error(err)
		h.statusJSON(w, InternalError)
		return
	}

	following, err := model.GetFollowing(h.db, user.UUID)
	if err != nil {
		logger.Error(err)
		h.statusJSON(w, InternalError)
		return
	}

	followInfo := FollowInfo{
		Status:    "OK",
		Following: following,
		Followers: followers,
	}
	h.writeJSON(w, followInfo)
}

func (h *BaseController) RequestFollow(w http.ResponseWriter, r *http.Request) {
	var user *model.User
	var err error
	if err = h.getKeyVal(r, acl.UserKey, &user); err != nil {
		logger.Error(err)
		h.statusJSON(w, Unauthorized)
		return
	}
	targetUuid := r.FormValue("uuid")
	targetUser, err := model.GetUserByUUID(h.db, targetUuid)
	if err != nil {
		logger.Error("User does not exists")
		h.statusJSON(w, StatusResponse{StatusCode: 404, Msg: "User does not exists"})
		return
	}
	junction := model.FollowJunction{
		Follower: user.UUID,
		Followed: targetUuid,
		Allowed:  targetUser.Public,
	}
	err = junction.Create(h.db)
	if err != nil {
		logger.Error(err)
		h.statusJSON(w, StatusResponse{StatusCode: 409, Msg: "Already sent"})
		return
	}

	if !targetUser.Public {
		notification := model.Notification{}
		currentTime := time.Now()

		err = notification.InsertNotification(h.db, user.UUID, targetUuid, "You have new follower request from "+user.Username, "follow", 1, currentTime, false)
		if err != nil {
			panic(err)
		}

		notificationData := ws.DataJSON{
			Type:         "notification", // Or any identifier for notifications
			ReciverUUID:  targetUuid,
			Notification: "You have a new follower request from " + user.Username,
			Timestamp:    time.Now(),
		}
		// Assuming you have access to the connection pool instance, `cp`
		h.cp.Notify(notificationData) // Correct way to call the Notify method.
	}

	if targetUser.Public {
		h.statusJSON(w, StatusResponse{Status: "OK", StatusCode: 201, Msg: "Following"})
	} else {
		h.statusJSON(w, StatusResponse{Status: "OK", StatusCode: 200, Msg: "Follow request sent"})
	}

}

func (h *BaseController) AcceptFollow(w http.ResponseWriter, r *http.Request) {
	var user *model.User
	var err error
	if err = h.getKeyVal(r, acl.UserKey, &user); err != nil {
		logger.Error(err)
		h.statusJSON(w, Unauthorized)
	}
	junction := model.FollowJunction{
		Follower: r.FormValue("uuid"),
		Followed: user.UUID,
	}
	err = junction.Allow(h.db)
	if err != nil {
		logger.Error(err)
		h.statusJSON(w, InternalError)
		return
	}
	h.statusJSON(w, Success)
}

func (h *BaseController) RequestOtherFollowInfo(w http.ResponseWriter, r *http.Request) {
	var user *model.User
	var err error
	if err = h.getKeyVal(r, acl.UserKey, &user); err != nil {
		logger.Error(err)
		h.statusJSON(w, Unauthorized)
		return
	}

	targetUuid := r.URL.Query().Get("uuid")

	followers, err := model.GetFollowers(h.db, targetUuid)
	if err != nil {
		logger.Error(err)
		h.statusJSON(w, InternalError)
		return
	}

	following, err := model.GetFollowing(h.db, targetUuid)
	if err != nil {
		logger.Error(err)
		h.statusJSON(w, InternalError)
		return
	}

	followInfo := FollowInfo{
		Status:    "OK",
		Following: following,
		Followers: followers,
	}
	h.writeJSON(w, followInfo)
}

func (h *BaseController) CancelFollow(w http.ResponseWriter, r *http.Request) {
	var user *model.User
	var err error
	if err = h.getKeyVal(r, acl.UserKey, &user); err != nil {
		logger.Error(err)
		h.statusJSON(w, Unauthorized)
	}
	junction := model.FollowJunction{
		Follower: user.UUID,
		Followed: r.FormValue("uuid"),
		Allowed:  false,
	}
	err = junction.Delete(h.db)
	if err != nil {
		logger.Error(err)
		h.statusJSON(w, InternalError)
		return
	}
	h.statusJSON(w, Success)
}

func (h *BaseController) DeclineFollow(w http.ResponseWriter, r *http.Request) {
	var user *model.User
	var err error
	if err = h.getKeyVal(r, acl.UserKey, &user); err != nil {
		logger.Error(err)
		h.statusJSON(w, Unauthorized)
	}
	junction := model.FollowJunction{
		Follower: r.FormValue("uuid"),
		Followed: user.UUID,
		Allowed:  false,
	}
	err = junction.Delete(h.db)
	if err != nil {
		logger.Error(err)
		h.statusJSON(w, InternalError)
		return
	}
	h.statusJSON(w, Success)
}
