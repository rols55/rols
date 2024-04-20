package controllers

import (
	"net/http"
	"strconv"

	"01.kood.tech/git/rols55/social-network/pkg/api/model"
	"01.kood.tech/git/rols55/social-network/pkg/api/route/middleware/acl"
	"01.kood.tech/git/rols55/social-network/pkg/logger"
)

// not in use
type HistoryResponse struct {
	Status  string           `json:"status"`
	User    *model.User      `json:"user,omitempty"`
	Target  string           `json:"target,omitempty"`
	Offset  int              `json:"offset,omitempty"`
	History []*model.Message `json:"history,omitempty"`
}

func (h *BaseController) GetHistory(w http.ResponseWriter, r *http.Request) {
	var err error
	var user *model.User
	// Is the user logged in
	if err = h.getKeyVal(r, acl.UserKey, &user); err != nil {
		logger.Error(err)
		h.statusJSON(w, Unauthorized)
		return
	}

	// Retrive values from r
	target := r.URL.Query().Get("target")
	offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))
	if target == "" {
		logger.Error(err)
		h.statusJSON(w, InvalidInput("Missing user target"))
		return
	}

	history, err := model.GetMessages(h.db, user.UUID, target, offset)
	if err != nil && err != model.ErrNotFound {
		logger.Error(err)
		h.statusJSON(w, InvalidInput(err.Error()))
		return
	}

	resp := HistoryResponse{
		Status:  "OK",
		History: history,
	}

	h.writeJSON(w, resp)
}

func (h *BaseController) GetGroupHistory(w http.ResponseWriter, r *http.Request) {
	var err error
	var user *model.User
	// Is the user logged in
	if err = h.getKeyVal(r, acl.UserKey, &user); err != nil {
		logger.Error(err)
		h.statusJSON(w, Unauthorized)
		return
	}

	groupId, err := strconv.ParseInt(r.URL.Query().Get("target"), 10, 64)
	if err != nil {
		logger.Error(err)
		h.statusJSON(w, InvalidInput(err.Error()))
		return
	}

	group, err := model.GetGroupById(h.db, groupId)
	if err != nil {
		logger.Error(err)
		h.statusJSON(w, InvalidInput(err.Error()))
		return
	}

	if !group.IsMember(h.db, user.Id) {
		errStr := "User is not a member of the group"
		logger.Error(errStr)
		h.statusJSON(w, NotFound)
		return
	}

	offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
	if err != nil {
		logger.Error(err)
		h.statusJSON(w, InvalidInput(err.Error()))
		return
	}

	history, err := model.GetGroupMessages(h.db, groupId, offset)
	if err != nil && err != model.ErrNotFound {
		logger.Error(err)
		h.statusJSON(w, InternalError)
		return
	}

	resp := HistoryResponse{
		Status:  "OK",
		History: history,
	}

	h.writeJSON(w, resp)
}
