package game

import (
	"time"
)

type Lobby struct {
	WsPool *ConnetionPool
}

func UpdateTimer(ws *ConnetionPool) {
	timer := 20      // Initial 20 seconds countdown
	finaltimer := 10 // Final countdown
	for timer > 0 {
		//ws.update <- DataJSON{
		ws.Broadcast(DataJSON{
			Type:  "timer",
			Timer: timer,
			//Players: players,
		})
		if len(ws.clients) == 4 {
			break
		}
		if len(ws.clients) < 2 {
			break
		}
		time.Sleep(1 * time.Second)
		timer--
	}

	if len(ws.clients) > 1 && len(ws.clients) < 5 {
		for finaltimer > 0 {
			//ws.update <- DataJSON{
			ws.Broadcast(DataJSON{
				Type:       "timer",
				TimerFinal: finaltimer,
				//Players:    players,
			})
			time.Sleep(1 * time.Second)
			finaltimer--
		}

		startGame := DataJSON{
			Type: START,
		}
		ws.start <- startGame
	}
}
