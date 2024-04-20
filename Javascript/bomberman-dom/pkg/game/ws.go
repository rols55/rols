package game

import (
	"encoding/json"
	"strconv"

	"server/pkg/logger"

	"github.com/gorilla/websocket"
)

const (
	CONNECTED    = "connected"
	DISCONNECTED = "disconnected"
	MESSAGE      = "message"
	BROADCAST    = "broadcast"
	JOIN         = "join"
	ACTION       = "action"
	MOVE         = "move"
	BOMB         = "BOMB"
	UPDATE       = "update"
	PUPDATE      = "pupdate"
	START        = "start"
	TIMER        = "timer"
)

type DataJSON struct {
	Type        string    `json:"type"` //"connected", "disconnected", "message", "action"
	Sender      string    `json:"sender_uuid,omitempty"`
	MessageText string    `json:"message_text,omitempty"`
	Action      *Action   `json:"action,omitempty"`
	Board       *Board    `json:"board,omitempty"`
	Timer       int       `json:"timer,omitempty"`
	TimerFinal  int       `json:"timerfinal,omitempty"`
	Players     []*Player `json:"players,omitempty"`
	Player      *Player   `json:"player,omitempty"`
}

type ConnetionPool struct {
	// The connected clinets
	clients map[string]*WsClient
	// Channel for adding clients
	add chan *WsClient
	// Channel for removing clients
	remove chan *WsClient
	// Channel for private messages
	message chan DataJSON
	// Channel for game actions
	action chan DataJSON
	// Channel for updates
	update chan DataJSON
	// Channel for timer
	timer chan DataJSON
	// Channel for start
	start chan DataJSON
}

type WsClient struct {
	User   string
	Id     int
	Player *Player
	Conn   *websocket.Conn
	send   chan DataJSON
	pool   *ConnetionPool
}

func New() *ConnetionPool {
	return &ConnetionPool{
		clients: make(map[string]*WsClient),
		add:     make(chan *WsClient),
		remove:  make(chan *WsClient),
		message: make(chan DataJSON),
		action:  make(chan DataJSON),
		start:   make(chan DataJSON),
		timer:   make(chan DataJSON),
	}
}

func (cp *ConnetionPool) Run() {
	var Game *Game
	for {
		//size = len(cp.clients)
		select {
		case <-cp.start:
			Game = InitGame(cp)
			data := DataJSON{
				Type:    START,
				Players: players,
				Board:   &Game.Board,
			}
			for _, v := range cp.clients {
				data.Player = v.Player
				v.send <- data
			}

		case client := <-cp.add:
			_, ok := cp.clients[client.User]
			if ok {
				logger.Error("User %s already exists", client.User)
				continue
			}
			if len(cp.clients) < 4 {
				client.Id = len(cp.clients)
				cp.clients[client.User] = client
				if len(cp.clients) == 2 {
					//if size == 0 {
					//Game = InitGame(cp)
					go UpdateTimer(cp)
				}
				logger.Info("Added WsClient %v to connetion pool", client.User)
				data := DataJSON{
					Type:    CONNECTED,
					Sender:  client.User,
					Players: players,
					//Board:   &Game.Board,
					/*
						Action: Action{
							X: client.Player.X,
							Y: client.Player.Y,
						},
					*/
					//Player: client.Player,
				}
				cp.Broadcast(data)
			} else {
				logger.Error("Too many clients")
			}

		case client := <-cp.remove:
			logger.Info("Removed WsClient %v to connetion pool", client.User)
			data := DataJSON{
				Type:   DISCONNECTED,
				Sender: client.User,
			}
			if _, ok := cp.clients[client.User]; ok {
				RemovePlayer(client.Id)
				delete(cp.clients, client.User)
				close(client.send)
				cp.Broadcast(data)
			}
		case message := <-cp.message:
			for _, client := range cp.clients {
				client.send <- message
			}
		case timer := <-cp.timer:
			for _, client := range cp.clients {
				client.send <- timer
			}
		case update := <-cp.update:
			//cp.clients[update.Sender].send <- update
			cp.Broadcast(update)
		case action := <-cp.action:
			switch action.Type {
			case MOVE:
				player := Game.WsPool.clients[action.Sender].Player
				if !player.Moving && player.Alive {
					player.Moving = true
					action.Action.Name = Game.WsPool.clients[action.Sender].User
					go Game.Move(*action.Action)
				}
			case BOMB:
				if Game.WsPool.clients[action.Sender].Player.PlaceBomb(Game) {
					data := DataJSON{
						Type:        BOMB,
						MessageText: strconv.Itoa(Game.WsPool.clients[action.Sender].Player.Flame),
						Sender:      action.Sender,
						Action:      action.Action,
					}
					cp.Broadcast(data)
				}
			}
		}
	}
}

func (cp *ConnetionPool) Broadcast(data DataJSON) {
	for _, client := range cp.clients {
		client.send <- data
	}
}

// WebSocket Client
// Creates a new Client instance
func NewClient(user string, conn *websocket.Conn, pool *ConnetionPool) *WsClient {
	client := &WsClient{
		User:   user,
		Player: NewPlayer(len(pool.clients), user),
		Conn:   conn,
		send:   make(chan DataJSON),
		pool:   pool,
	}
	go client.Write()
	go client.Read()
	client.pool.add <- client
	return client
}

// server POV
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
			logger.Info("Reading form webSocket type: %v, data: %v", msgType, string(data))
			dataJSON.Sender = c.User
			switch dataJSON.Type {
			case MESSAGE:
				c.pool.message <- dataJSON
			case ACTION, MOVE, BOMB:
				c.pool.action <- dataJSON
			case UPDATE:
				c.pool.update <- dataJSON
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
				logger.Info("Writing to websocket with data: %v", string(data))
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
