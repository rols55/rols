package main

import (
	"net/http"

	"server/pkg/logger"
	"server/pkg/game"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	logger.Info("Starting server ... ")
	cp := game.New()
	go cp.Run()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Serve the index.html file
		http.ServeFile(w, r, "static/index.html")
	})

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		upgrader.CheckOrigin = func(r *http.Request) bool { return true }
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			logger.Error(err)
			return
		}
		name := r.URL.Query().Get("name")
		game.NewClient(name, conn, cp)
	})

	staticDirectory := http.Dir("static")
	staticServer := http.FileServer(staticDirectory)
	http.Handle("/static/", http.StripPrefix("/static/", staticServer))

	logger.Info("Server started at localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		logger.Error(err)
		return
	}

}
