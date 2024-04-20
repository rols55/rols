package controllers

import (
	"forum/model"
	"forum/route/middleware/acl"
	"net/http"
	"strconv"
)

//not in use
type HistoryResponse struct {
	Status string      `json:"status"`
	User   *model.User `json:"user,omitempty"`
	Target string      `json:"target"`
	Offset int         `json:"offset"`
}

func (h *BaseController) GetHistory(w http.ResponseWriter, r *http.Request) {
	var err error
	var user *model.User
	// Is the user logged in
	if err = h.getKeyVal(r, acl.UserKey, &user); err != nil && err != model.ErrNotFound {
		h.statusJSON(w, InternalError)
		return
	}

	// Retrive values from r
	target := r.URL.Query().Get("target")
	offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))
	if err != nil {
		h.statusJSON(w, InvalidInput(err.Error()))
		return
	}

	history, err := model.GetMessages(h.db, user.UUID, target, offset)
	if err == nil {
		h.writeJSON(w, history)
	}
}
