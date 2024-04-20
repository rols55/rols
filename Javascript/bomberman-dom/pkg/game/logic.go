package game

import (
	"math/rand"
	"strings"
	"time"
)

type Game struct {
	Board   Board `json:"board,omitempty"`
	WsPool  *ConnetionPool
	BombMap map[Coords]bool
}

type Action struct {
	Name string `json:"name,omitempty"`
	Id   int    `json:"id,omitempty"`
	X    int    `json:"x"`
	Y    int    `json:"y"`
}

type Bomb struct {
	Flame int
	Y     int
	X     int
}

type Coords struct {
	Y int
	X int
}

func InitGame(ws *ConnetionPool) *Game {
	return &Game{
		Board:   *CreateMap(),
		WsPool:  ws,
		BombMap: make(map[Coords]bool),
	}
}

func (g *Game) Move(move Action) {
	mover := g.WsPool.clients[move.Name].Player
	field := &g.Board[move.Y][move.X]
	powerUp := strings.Contains(*field, "power")
	_, isBomb := g.BombMap[Coords{Y: move.Y, X: move.X}]
	time.Sleep(time.Duration(500-mover.Speed*100) * time.Millisecond)
	if powerUp || (*field == "empty" && !isBomb) {
		g.Board[mover.Y][mover.X] = "empty"
		mover.X = move.X
		mover.Y = move.Y
		if powerUp {
			switch *field {
			case "powerFlame":
				mover.Flame++
			case "powerSpeed":
				mover.Speed++
			case "powerBomb":
				mover.Bombs++
			}
			*field = mover.Name
			data := DataJSON{
				Type:  UPDATE,
				Board: &g.Board,
			}
			g.WsPool.Broadcast(data)
		}
		*field = mover.Name
		data := DataJSON{
			Type:   MOVE,
			Sender: move.Name,
			Action: &move,
		}
		g.WsPool.Broadcast(data)
	}
	mover.Moving = false
}

func (g *Game) rollPowerUp(x, y int) {
	if rand.Intn(5) == 0 {
		switch rand.Intn(3) + 1 {
		case 1:
			g.Board[x][y] = "powerFlame"
		case 2:
			g.Board[x][y] = "powerSpeed"
		case 3:
			g.Board[x][y] = "powerBomb"
		}
		g.WsPool.Broadcast(DataJSON{
			Type:  UPDATE,
			Board: &g.Board,
		})
	}
}

func (g *Game) Explode(bomb *Bomb) {
	right := true
	up := true
	left := true
	down := true
	g.Damage(g.Board[bomb.Y][bomb.X])
	for i := 1; i <= bomb.Flame; i++ {
		if right {
			target := &g.Board[bomb.Y][bomb.X+i]
			if *target == "box" {
				*target = "empty"
				right = false
				g.rollPowerUp(bomb.Y, bomb.X+i)
			}
			if *target == "wall" {
				right = false
			}
			g.Damage(*target)
		}
		if up {
			target := &g.Board[bomb.Y-i][bomb.X]
			if *target == "box" {
				*target = "empty"
				up = false
				g.rollPowerUp(bomb.Y-i, bomb.X)
			}
			if *target == "wall" {
				up = false
			}
			g.Damage(*target)
		}
		if left {
			target := &g.Board[bomb.Y][bomb.X-i]
			if *target == "box" {
				*target = "empty"
				left = false
				g.rollPowerUp(bomb.Y, bomb.X-i)
			}
			if *target == "wall" {
				left = false
			}
			g.Damage(*target)
		}
		if down {
			target := &g.Board[bomb.Y+i][bomb.X]
			if *target == "box" {
				*target = "empty"
				down = false
				g.rollPowerUp(bomb.Y+i, bomb.X)
			}
			if *target == "wall" {
				down = false
			}
			g.Damage(*target)
		}
	}
	g.WsPool.Broadcast(DataJSON{
		Type:  UPDATE,
		Board: &g.Board,
	})
}

func (g *Game) Damage(target string) {
	if strings.Contains(target, "player") {
		nameOfGamer := strings.Split(target, "_")[1]
		client := g.WsPool.clients[nameOfGamer]
		client.Player.Life--
		data := DataJSON{
			Type:        PUPDATE,
			MessageText: "damaged",
		}
		client.send <- data
		if client.Player.Life == 0 {
			g.WsPool.Broadcast(DataJSON{
				Type:        PUPDATE,
				Sender:      nameOfGamer,
				MessageText: "died",
			})
			client.Player.Alive = false
			g.Board[client.Player.Y][client.Player.X] = "empty"
			winnerChecker(g.WsPool)
		}
	}
}

func winnerChecker(ws *ConnetionPool) {
	aliveCount := 0
	var winner *WsClient
	for _, player := range ws.clients {
		if player.Player.Alive {
			aliveCount++
			winner = player
		}
	}
	if aliveCount == 1 {
		winner.send <- DataJSON{
			Type:        PUPDATE,
			MessageText: "won",
		}
		//ENDGAME
	}
}
