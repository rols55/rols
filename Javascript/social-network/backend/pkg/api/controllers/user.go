package controllers

import (
	"net/http"

	"01.kood.tech/git/rols55/social-network/pkg/api/model"
	"01.kood.tech/git/rols55/social-network/pkg/api/route/middleware/acl"
	"01.kood.tech/git/rols55/social-network/pkg/logger"
)

func (h *BaseController) TogglePublic(w http.ResponseWriter, r *http.Request) {
	var user *model.User
	if err := h.getKeyVal(r, acl.UserKey, &user); err != nil {
		logger.Error(err)
		h.statusJSON(w, Unauthorized)
		return
	}
	err := user.TogglePublic(h.db)
	if err != nil {
		logger.Error(err)
		h.statusJSON(w, InternalError)
		return
	}

	h.statusJSON(w, Success)
}
