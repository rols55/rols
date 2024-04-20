package controllers

import (
	"net/http"
	"strconv"

	"01.kood.tech/git/rols55/social-network/pkg/api/model"
	"01.kood.tech/git/rols55/social-network/pkg/api/route/middleware/acl"
	"01.kood.tech/git/rols55/social-network/pkg/logger"
)

type NotificationsResponse struct {
	Status        string                `json:"status"`
	User          *model.User           `json:"user,omitempty"`
	Notification  *model.Notification   `json:"notification,omitempty"`
	Notifications []*model.Notification `json:"notifications,omitempty"`
}

// Delete the post
func (h *BaseController) NotificationGET(w http.ResponseWriter, r *http.Request) {
	var err error
	var user *model.User

	//check if the user is logged in
	if err = h.getKeyVal(r, acl.UserKey, &user); err != nil && err != model.ErrNotFound {
		logger.Error(err)
		h.statusJSON(w, Unauthorized)
		return
	}

	//get the post's id form url path
	notificationId, err := strconv.Atoi(r.URL.Path[len("/api/notification/"):])
	if err != nil {
		logger.Error(err)
		h.statusJSON(w, NotFound)
		return
	}

	notification := model.Notification{}

	// Update IsRead field to true
	err = notification.UpdateNotificationIsRead(h.db, notificationId, true)
	if err != nil {
		logger.Error(err)
		h.statusJSON(w, NotFound)
		return
	}

	res := NotificationsResponse{
		Status: "OK",
	}

	h.writeJSON(w, res)
}

// Delete the post
func (h *BaseController) NotificationsGET(w http.ResponseWriter, r *http.Request) {

	//var id int
	var err error
	var user *model.User

	cookie := r.Header.Get("Cookie")
	logger.Info("Cookie:", cookie)

	//check if the user is logged in
	if err = h.getKeyVal(r, acl.UserKey, &user); err != nil {
		logger.Error(err)
		h.statusJSON(w, InternalError)
	}

	notification := model.Notification{}
	notifications, err := notification.GetNotificationsForUser(h.db, user.UUID)
	if err != nil {
		panic(err)
	}

	for i, j := 0, len(notifications)-1; i < j; i, j = i+1, j-1 {
		notifications[i], notifications[j] = notifications[j], notifications[i]
	}

	res := NotificationsResponse{
		Status:        "OK",
		User:          user,
		Notifications: notifications,
	}

	h.writeJSON(w, res)
}
