package ws

import (
	"database/sql"
	"encoding/json"
	"forum/model"
	"forum/shared/logger"
	"time"

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
	Type        string    `json:"type"` //"connected", "disconnected", "message", "broadcast", "get_online"
	SenderUUID  string    `json:"sender_uuid,omitempty"`
	ReciverUUID string    `json:"reciver_uuid,omitempty"`
	MessageText string    `json:"message_text,omitempty"`
	Timestamp   time.Time `json:"timestamp,omitempty"`
	Online      []string  `json:"online,omitempty"`
}

type ConnetionPool struct {
	// The DB connection
	db *sql.DB
	// The connected clinets
	clients map[string]*WsClient
	// Channel for adding clients
	add chan *WsClient
	// Channel for removeing clients
	remove chan *WsClient
	// Channel for broadcasting data to all clients
	broadcast chan DataJSON
	// Channel for private messages
	message chan DataJSON
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
			delete(cp.clients, client.User.UUID)
			close(client.send)
		case data := <-cp.broadcast:
			cp.Broadcast(data)
		case message := <-cp.message:
			cp.Message(message)
		}
	}
}

func (cp *ConnetionPool) Broadcast(data DataJSON) {
	// For now, even the sender gets the message
	for _, client := range cp.clients {
		select {
		case client.send <- data:
		default:
			delete(cp.clients, client.User.UUID)
			close(client.send)
		}
	}
}

func (cp *ConnetionPool) Message(data DataJSON) {
	toSave := model.Message{
		Sender_UUID:  data.SenderUUID,
		Reciver_UUID: data.ReciverUUID,
		Message_text: data.MessageText,
	}
	toSave.Create(cp.db)
	client, ok := cp.clients[data.ReciverUUID]
	if !ok {
		logger.Error("Unable to find client with data: %v", data)
	} else {
		select {
		case client.send <- data:
		default:
			delete(cp.clients, client.User.UUID)
			close(client.send)
		}
	}

	sender := cp.clients[data.SenderUUID]
	data.Timestamp = toSave.Timestamp
	select {
	case sender.send <- data:
	default:
		delete(cp.clients, data.SenderUUID)
		close(sender.send)
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
		logger.Info("%v's WebSocket closed", c.User.Username)
	}()
	for {
		_, data, err := c.Conn.ReadMessage()
		if err != nil {
			logger.Error(err)
			return
		}
		if err := json.Unmarshal(data, &dataJSON); err != nil {
			logger.Error(err)
		} else {
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
	}
}

func (c *WsClient) Write() {
	defer func() {
		c.Conn.Close()
		logger.Info("%v's WebSocket closed", c.User.Username)
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
