package ws

import (
	"database/sql"
	"encoding/json"
	"time"

	"01.kood.tech/git/rols55/social-network/pkg/api/model"
	"01.kood.tech/git/rols55/social-network/pkg/logger"

	"github.com/gorilla/websocket"
)

const (
	CONNECTED    = "connected"
	DISCONNECTED = "disconnected"
	MESSAGE      = "message"
	BROADCAST    = "broadcast"
	GET_ONLINE   = "get_online"
)

type DataJSON struct {
	Type         string    `json:"type"` //"connected", "disconnected", "message", "broadcast", "get_online"
	SenderUUID   string    `json:"sender_uuid,omitempty"`
	ReciverUUID  string    `json:"reciver_uuid,omitempty"`
	MessageText  string    `json:"message_text,omitempty"`
	Timestamp    time.Time `json:"timestamp,omitempty"`
	Online       []string  `json:"online,omitempty"`
	GroupId      int64     `json:"group_id,omitempty"`
	Notification string    `json:"notification,omitempty"`
}
type ConnetionPool struct {
	// The DB connection
	db *sql.DB
	// The connected clinets
	clients map[string]*WsClient
	// Channel for adding clients
	add chan *WsClient
	// Channel for removing clients
	remove chan *WsClient
	// Channel for broadcasting data to all clients
	broadcast chan DataJSON
	// Channel for private messages
	message chan DataJSON
	// Channel for notifications
	notify chan DataJSON
}
type WsClient struct {
	User *model.User
	Conn *websocket.Conn
	send chan DataJSON
	pool *ConnetionPool
}

// WebSocket Connetion Pool
func New(db *sql.DB) *ConnetionPool {
	return &ConnetionPool{
		db:        db,
		clients:   make(map[string]*WsClient),
		add:       make(chan *WsClient),
		remove:    make(chan *WsClient),
		broadcast: make(chan DataJSON),
		message:   make(chan DataJSON),
		notify:    make(chan DataJSON),
	}
}

func (cp *ConnetionPool) Run() {
	for {
		select {
		case client := <-cp.add:
			cp.clients[client.User.UUID] = client
			logger.Info("Added WsClient %v to connetion pool", client.User.Username)
			data := DataJSON{
				Type:       CONNECTED,
				SenderUUID: client.User.UUID,
			}
			cp.Broadcast(data)
		case client := <-cp.remove:
			logger.Info("Removed WsClient %v to connetion pool", client.User.Username)
			data := DataJSON{
				Type:       DISCONNECTED,
				SenderUUID: client.User.UUID,
			}
			cp.Broadcast(data)
			if _, ok := cp.clients[client.User.UUID]; ok {
				delete(cp.clients, client.User.UUID)
				close(client.send)
			}
		case data := <-cp.broadcast:
			cp.Broadcast(data)
		case message := <-cp.message:
			cp.Message(message)
		case notification := <-cp.notify:
			cp.Notify(notification)
		}
	}
}

func (cp *ConnetionPool) Notify(notification DataJSON) {
	if client, ok := cp.clients[notification.ReciverUUID]; ok {
		select {
		case client.send <- notification:
			/*
				default:
					// Handle failed send, if necessary
			*/
		}
	}
}

func (cp *ConnetionPool) Broadcast(data DataJSON) {
	// For now, even the sender gets the message
	for _, client := range cp.clients {
		if client.User.UUID == data.SenderUUID {
			continue
		}
		select {
		case client.send <- data:
			/*
				default:
					delete(cp.clients, client.User.UUID)
					close(client.send)
			*/
		}
	}
}
func (cp *ConnetionPool) Message(data DataJSON) {
	toSave := model.Message{
		Sender_UUID:  data.SenderUUID,
		Reciver_UUID: data.ReciverUUID,
		Message_text: data.MessageText,
		GroupId:      data.GroupId,
	}
	_, err := toSave.Create(cp.db)
	if err != nil {
		logger.Error(err)
		return
	}

	data.Timestamp = toSave.Timestamp

	sender, ok := cp.clients[data.SenderUUID]
	if !ok {
		logger.Error("Unable to find sender with data: %v", data)
		return
	}

	if data.GroupId > 0 {
		group, err := model.GetGroupById(cp.db, data.GroupId)
		if err != nil {
			logger.Error(err)
			return
		}

		if !group.IsMember(cp.db, sender.User.Id) {
			logger.Error("User with id: %v is not a member of group: %v", sender.User.Id, group.Id)
			return
		}

		members, err := group.GetMemebers(cp.db)
		if err != nil {
			logger.Error(err)
			return
		}

		for _, m := range members {
			if client, ok := cp.clients[m.UUID]; ok {
				select {
				case client.send <- data:
				}
			}
		}
	} else {
		client, ok := cp.clients[data.ReciverUUID]
		if !ok {
			logger.Info("Client: %v is not online (data: %v)", data.ReciverUUID, data)
		} else {
			if model.IsFollowing(cp.db, data.ReciverUUID, data.SenderUUID) || sender.User.Public {
				select {
				case client.send <- data:
					/*
						default:
							delete(cp.clients, client.User.UUID)
							close(client.send)
					*/
				}
			}
		}

		select {
		case sender.send <- data:
			/*
				default:
					delete(cp.clients, data.SenderUUID)
					close(sender.send)
			*/
		}
	}
}

// WebSocket Client
// Creates a new Client instance
func NewClient(user *model.User, conn *websocket.Conn, cp *ConnetionPool) *WsClient {

	client := &WsClient{
		User: user,
		Conn: conn,
		send: make(chan DataJSON),
		pool: cp,
	}

	go client.Write()
	go client.Read()

	client.pool.add <- client

	return client
}
func (c *WsClient) Read() {
	var dataJSON DataJSON
	defer func() {
		c.pool.remove <- c
		c.Conn.Close()
	}()
	for {
		msgType, data, err := c.Conn.ReadMessage()
		if err != nil {
			if !websocket.IsCloseError(err, websocket.CloseGoingAway) {
				logger.Error(err)
			}
			return
		}
		if err := json.Unmarshal(data, &dataJSON); err != nil {
			logger.Error(err)
		} else {
			logger.Info("Reading form webSocket type: %v, data: %v", msgType, dataJSON)
			dataJSON.SenderUUID = c.User.UUID
			switch dataJSON.Type {
			case MESSAGE:
				c.pool.message <- dataJSON
			case BROADCAST:
				c.pool.broadcast <- dataJSON
			case GET_ONLINE:
				c.send <- c.GetOnline()
			default:
				logger.Info("WsClient dont know how to handle data with type %s [%v]", dataJSON.Type, dataJSON)
			}
		}
		dataJSON = DataJSON{}
	}
}
func (c *WsClient) Write() {
	defer func() {
		c.pool.remove <- c
		c.Conn.Close()
	}()
	for {
		/*
			for data := range c.send {
				c.Conn.Write(data)
			}
		*/
		select {
		case dataJSON, ok := <-c.send:
			if !ok {
				logger.Info("Trying to read form cahnnel but it has been closed")
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			w, err := c.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				logger.Error(err)
				return
			}
			if data, err := json.Marshal(dataJSON); err != nil {
				logger.Error(err)
			} else {
				logger.Info("Writing to websocket with data: %v", dataJSON)
				w.Write(data)
				// Add queued chat messages to the current websocket message.
				n := len(c.send)
				for i := 0; i < n; i++ {
					if data, err := json.Marshal(<-c.send); err != nil {
						logger.Error(err)
					} else {
						w.Write([]byte{'\n'})
						w.Write(data)
					}
				}
			}
			if err := w.Close(); err != nil {
				logger.Error(err)
				return
			}
		}
	}
}
func (c *WsClient) GetOnline() DataJSON {
	users := make([]string, 0, len(c.pool.clients))
	for key := range c.pool.clients {
		users = append(users, key)
	}
	return DataJSON{
		Type:   GET_ONLINE,
		Online: users,
	}
}
