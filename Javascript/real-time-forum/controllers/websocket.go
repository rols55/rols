package controllers

import (
	"net/http"

	"forum/model"
	"forum/route/middleware/acl"
	"forum/shared/logger"
	"forum/ws"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (h *BaseController) WebSocket(w http.ResponseWriter, r *http.Request) {

	var err error
	var userID int64
	var user *model.User

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logger.Error(err)
		return
	}

	closeMessage := websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Connection refused")

	if r.Context().Value(acl.UserKey) != nil {
		userID = r.Context().Value(acl.UserKey).(int64)
		if user, err = model.GetUserById(h.db, userID); err != nil {
			logger.Error(err)
			if err = conn.WriteMessage(websocket.CloseMessage, closeMessage); err != nil {
				logger.Error("Error sending close message:", err)
			}
			conn.Close()
			return
		}
	}

	if user == nil {
		logger.Info("Authentication needed")
		if err = conn.WriteMessage(websocket.CloseMessage, closeMessage); err != nil {
			logger.Error("Error sending close message:", err)
		}
		conn.Close()
		return
	}

	ws.NewClient(user, conn, h.cp)
}
