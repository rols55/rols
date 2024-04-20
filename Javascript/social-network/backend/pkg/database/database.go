package database

import (
	"database/sql"
	"errors"
	"os"
	"strings"

	"01.kood.tech/git/rols55/social-network/pkg/logger"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
)

var addTestData = func() bool {
	str := os.Getenv("ADD_TEST_DATA")
	if strings.ToLower(str) == "false" {
		return false
	}
	if str == "0" {
		return false
	}
	return true
}()

var runMigrations = func() bool {
	str := os.Getenv("RUN_MIGRATIONS")
	if strings.ToLower(str) == "false" {
		return false
	}
	if str == "0" {
		return false
	}
	return true
}()

func OpenDBConnection() (*sql.DB, error) {

	file, ok := os.LookupEnv("DATABASE_DIR")
	if !ok {
		file = "forum.db"
	}

	db, err := sql.Open("sqlite3", file+"?_foreign_keys=on")
	if err != nil {
		return nil, errors.New(logger.GetCurrentFuncName() + " " + err.Error())
	}

	if runMigrations {
		if err := doMigrations(db); err != nil {
			return nil, err
		}
	}

	if addTestData {
		if err := addData(db); err != nil {
			return db, err
		}
	}

	return db, nil
}

func addData(db *sql.DB) error {
	logger.Info("Adding test data to database ...")

	file, ok := os.LookupEnv("TEST_DATA")
	if !ok {
		file = "pkg/database/testData.sql"
	}

	logger.Info("Test data path: %v", file)

	// Read the schema.sql file
	schema, err := os.ReadFile(file)
	if err != nil {
		return errors.New(logger.GetCurrentFuncName() + " " + err.Error())
	}

	// Execute the schema
	_, err = db.Exec(string(schema))
	if err != nil {
		return errors.New(logger.GetCurrentFuncName() + " " + err.Error())
	}

	logger.Info("Adding test data done")

	return nil
}

func doMigrations(db *sql.DB) error {
	logger.Info("Doing database migrations ...")

	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		return errors.New(logger.GetCurrentFuncName() + " " + err.Error())
	}

	file, ok := os.LookupEnv("MIGRATIONS")
	if !ok {
		file = "file://pkg/database/migrations"
	}

	logger.Info("Migrations path: %v", file)

	m, err := migrate.NewWithDatabaseInstance(file, "sqlite3", driver)
	if err != nil {
		return errors.New(logger.GetCurrentFuncName() + " " + err.Error())
	}

	if err := m.Down(); err != nil && err != migrate.ErrNoChange {
		return errors.New(logger.GetCurrentFuncName() + " " + err.Error())
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return errors.New(logger.GetCurrentFuncName() + " " + err.Error())
	}

	// Run migrations twice for errors
	if err := m.Down(); err != nil && err != migrate.ErrNoChange {
		return errors.New(logger.GetCurrentFuncName() + " " + err.Error())
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return errors.New(logger.GetCurrentFuncName() + " " + err.Error())
	}

	logger.Info("Database migrations done")

	return nil
}
