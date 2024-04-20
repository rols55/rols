package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"

	"01.kood.tech/git/rols55/social-network/pkg/api/model"
	"01.kood.tech/git/rols55/social-network/pkg/api/ws"
	"01.kood.tech/git/rols55/social-network/pkg/logger"
)

var (
	NotFound = StatusResponse{
		Status:     "ERROR",
		StatusCode: http.StatusNotFound,
		Msg:        "Not Found",
	}

	InternalError = StatusResponse{
		Status:     "ERROR",
		StatusCode: http.StatusInternalServerError,
		Msg:        "Internal Error",
	}

	Unauthorized = StatusResponse{
		Status:     "ERROR",
		StatusCode: http.StatusUnauthorized,
		Msg:        "Unauthorized access or session has expired",
	}

	InvalidInput = func(msg ...string) StatusResponse {
		str := "Invalid input"
		if len(msg) > 1 {
			str = strings.Join(msg, "\n")
		} else if len(msg) == 1 {
			str = msg[0]
		}
		return StatusResponse{
			Status:     "ERROR",
			StatusCode: http.StatusUnprocessableEntity,
			Msg:        str,
		}
	}

	Success = StatusResponse{
		Status: "OK",
	}

	Redirect = func(redirect string) StatusResponse {
		return StatusResponse{
			Status:     "OK",
			StatusCode: http.StatusPermanentRedirect,
			Redirect:   redirect,
		}
	}
)

type StatusResponse struct {
	Status     string `json:"status,omitempty"`
	StatusCode int    `json:"code,omitempty"`
	Msg        string `json:"msg,omitempty"`
	Redirect   string `json:"redirect,omitempty"`
}

type BaseController struct {
	db *sql.DB
	cp *ws.ConnetionPool
}

func New(db *sql.DB, cp *ws.ConnetionPool) *BaseController {
	return &BaseController{
		db: db,
		cp: cp,
	}
}

func (h *BaseController) writeJSON(w http.ResponseWriter, res interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(res); err != nil {
		errStr := "Error encoding JSON"
		logger.Error(errStr)
		http.Error(w, errStr, http.StatusInternalServerError)
	}
}

func (h *BaseController) statusJSON(w http.ResponseWriter, res StatusResponse) {
	//w.WriteHeader(res.StatusCode)
	h.writeJSON(w, res)
}

func (h *BaseController) getKeyVal(r *http.Request, key any, val any) error {
	if r.Context().Value(key) != nil {
		var err error
		switch v := val.(type) {
		case *int:
			*v = r.Context().Value(key).(int)
		case *string:
			*v = r.Context().Value(key).(string)
		case **model.User:
			userID := r.Context().Value(key).(int64)
			if *v, err = model.GetUserById(h.db, userID); err != nil {
				return err
			}
		default:
			return model.ErrNotFound
		}
		return nil
	}
	return model.ErrNotFound
}
