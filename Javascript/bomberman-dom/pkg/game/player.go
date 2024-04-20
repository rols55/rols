package game

import (
	"time"
)

var players []*Player

type Player struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Life        int    `json:"life"`
	Speed       int    `json:"speed"`
	Bombs       int    `json:"bombs"`
	BombsPlaced int    `json:"bombsPlaced"`
	Flame       int    `json:"flame"`
	X           int    `json:"x"`
	Y           int    `json:"y"`
	Alive       bool   `json:"alive"`
	Moving      bool   `json:"moving"`
}

func NewPlayer(id int, name string) *Player {
	p := &Player{
		Id:          id,
		Life:        3,
		Speed:       1,
		Bombs:       1,
		BombsPlaced: 0,
		Flame:       1,
		Alive:       true,
		Moving:      false,
	}
	switch id {
	case 0:
		p.X = 1
		p.Y = 1
		p.Name = "player1_" + name
	case 1:
		p.X = 11
		p.Y = 1
		p.Name = "player2_" + name
	case 2:
		p.X = 1
		p.Y = 11
		p.Name = "player3_" + name
	case 3:
		p.X = 11
		p.Y = 11
		p.Name = "player4_" + name
	}
	players = append(players, p)
	return p
}

func RemovePlayer(playerId int) {
	for i, p := range players {
		if p.Id == playerId {
			players = append(players[:i], players[i+1:]...)
			break
		}
	}
}

func (p *Player) StartBombExploder(game *Game) {
	go func() {
		bomb := &Bomb{
			Flame: p.Flame,
			X:     p.X,
			Y:     p.Y,
		}
		game.BombMap[Coords{
			Y: bomb.Y,
			X: bomb.X,
		}] = true
		time.Sleep(2 * time.Second)
		delete(game.BombMap, Coords{
			Y: bomb.Y,
			X: bomb.X,
		})
		game.Explode(bomb)
		p.BombsPlaced--
	}()
}

func (p *Player) PlaceBomb(game *Game) bool {
	if p.BombsPlaced >= p.Bombs || !p.Alive {
		return false
	}
	p.BombsPlaced++
	p.StartBombExploder(game)
	return true
}
