package model

import (
	"database/sql"
	"errors"
	"os"

	"forum/shared/logger"

	_ "github.com/mattn/go-sqlite3"
)

var (
	// No records found
	ErrNotFound = errors.New("not found in database")
)

func OpenDBConnection(file string) (*sql.DB, error) {
	init := false

	// check if we already have a database
	if _, err := os.Stat(file); err != nil {
		init = true
	}
	// Open db connection
	db, err := sql.Open("sqlite3", file+"?_foreign_keys=on")
	if err != nil {
		return nil, errors.New(logger.GetCurrentFuncName() + " " + err.Error())
	}
	if init {
		if err := inti(db); err != nil {
			logger.Error("Failed to initialize db with error: %v", err)
		}
	}
	return db, nil
}

func inti(db *sql.DB) error {
	logger.Info("Initializing database")

	// Read the schema.sql file
	schema, err := os.ReadFile("./schema.sql")
	if err != nil {
		return errors.New(logger.GetCurrentFuncName() + " " + err.Error())
	}

	// Execute the schema
	_, err = db.Exec(string(schema))
	if err != nil {
		return errors.New(logger.GetCurrentFuncName() + " " + err.Error())
	}

	return nil
}
