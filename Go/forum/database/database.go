package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB
var err error

// populate functions should be uncommented when the database gets flushed otherwise
// otherwise you'll get unique key errors from SQLite
func RunDatabase() {
	db, err = sql.Open("sqlite3", "./forum.db")
	if err != nil {
		log.Fatal(err)
	}

	// clearDB()
	// makeTables()
	makeDefaultUsers()
	// populateDefaultUsers()
	makeDefaultPosts()
	// populateDefaultPosts()
	makeDefaultComments()
	// populateDefaultComments()
}

// in case of table needing an update
func clearDB() {
	statement, err := db.Prepare("DROP TABLE users")
	if err != nil {
		log.Println(err)
	} else {
		statement.Exec()
		log.Println("Cleared users")
	}

	statement, err = db.Prepare("DROP TABLE posts")
	if err != nil {
		log.Println(err)
	} else {
		statement.Exec()
		log.Println("Cleared posts")
	}

	statement, err = db.Prepare("DROP TABLE comments")
	if err != nil {
		log.Println(err)
	} else {
		statement.Exec()
		log.Println("Cleared comments")
	}

	statement.Close()
}

func makeTables() {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username TEXT UNIQUE,
			password TEX,
			email TEXT UNIQUE,
			role TEXT,
			created_at TEXT
		);
	`)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Made users table")
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS posts (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			user_id INTEGER,
			title TEXT,
			body TEXT,
			categories TEXT,
			status TEXT,
			likes INTEGER,
			dislikes INTEGER,
			created_at TEXT,
			FOREIGN KEY(user_id) REFERENCES users(id)
		);
	`)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Made posts table")
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS comments (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			post_id INTEGER,
			user_id INTEGER,
			body TEXT,
			status TEXT,
			likes INTEGER,
			dislikes INTEGER,
			created_at TEXT,
			FOREIGN KEY(post_id) REFERENCES posts(id),
			FOREIGN KEY(user_id) REFERENCES users(id)
		);
	`)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Made comments table")
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS likes (
			post_id INTEGER,
			comment_id INTEGER,
			user_id INTEGER,
			FOREIGN KEY(post_id) REFERENCES posts(id),
			FOREIGN KEY(comment_id) REFERENCES comments(id),
			FOREIGN KEY(user_id) REFERENCES users(id)
		);
	`)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Made likes table")
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS dislikes (
			post_id INTEGER,
			comment_id INTEGER,
			user_id INTEGER,
			FOREIGN KEY(post_id) REFERENCES posts(id),
			FOREIGN KEY(comment_id) REFERENCES comments(id),
			FOREIGN KEY(user_id) REFERENCES users(id)
		);
	`)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Made dislikes table")
	}
}
