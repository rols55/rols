package main

import (
	"net/http"

	"01.kood.tech/git/rols55/social-network/pkg/api/route"
	"01.kood.tech/git/rols55/social-network/pkg/database"
	"01.kood.tech/git/rols55/social-network/pkg/logger"
)

func main() {
	logger.Info("Starting server ... ")

	logger.Info("Connetcing to Database")
	db, err := database.OpenDBConnection()
	if db == nil {
		logger.Error(err)
	} else if err != nil {
		logger.Error(err)
	}

	defer db.Close()

	logger.Info("Server started at localhost:8080")
	if err := http.ListenAndServe(":8080", route.Load(db)); err != nil {
		logger.Error(err)
		return
	}
}
