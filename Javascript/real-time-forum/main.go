package main

import (
	"net/http"

	"forum/model"
	"forum/route"
	"forum/shared/logger"
)

const (
	CERT = "ssl/cert.pem"
	KEY  = "ssl/key.pem"
)

func main() {
	logger.Info("Starting server ... ")

	logger.Info("Connetcing to Database")
	db, err := model.OpenDBConnection("./forum.db")
	if err != nil {
		logger.Error(err)
		return
	}
	defer db.Close()

	logger.Info("Server started at localhost:8080")
	if err := http.ListenAndServe(":8080", route.Load()); err != nil {
		logger.Error(err)
		return
	}
	/* Ignore it for now

	tlsConfig := &tls.Config{
		MinVersion:               tls.VersionTLS12,
		PreferServerCipherSuites: true,
		CurvePreferences:         []tls.CurveID{tls.CurveP256},
	}

	server := &http.Server{
		Addr:      ":443",
		Handler:   route.Load(),
		TLSConfig: tlsConfig,
	}

	logger.Info("Server started at localhost%v", server.Addr)
	if err := server.ListenAndServeTLS(CERT, KEY); err != nil {
		logger.Error(err)
	}
	*/

}
