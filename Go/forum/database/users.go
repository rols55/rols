package database

import (
	"database/sql"
	"fmt"
	"forum/login"
	"log"
	"time"
)

type defaultUser struct {
	id         int
	username   string
	password   string
	email      string
	role       string
	created_at string
}

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Hash     string
}

// global and exportableto be used in posts
var defaultUsers []defaultUser

func makeDefaultUsers() {
	names := []string{"John Doe", "Spike Spiegel", "Marie Curie"}
	passwords := []string{"0000000000", "1234567890", "radiumIsCoolAndDangerous"}
	emails := []string{"ordinary.guy@mail.com", "tank@space.com", "element.lover@ra.don"}
	roles := []string{"user", "user", "moderator"}
	id := 0
	for i, v := range names {
		user := defaultUser{
			id,
			v,
			passwords[i],
			emails[i],
			roles[i],
			time.Now().String(),
		}
		defaultUsers = append(defaultUsers, user)
		id++
	}
	log.Println("Made default users")
}

func populateDefaultUsers() {
	statement, err := db.Prepare(`
		INSERT INTO users (
		id,
		username,
		email,
		password,
		role,
		created_at) VALUES (?, ?, ?, ?, ?, ?)
	`)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range defaultUsers {
		_, err = statement.Exec(v.id, v.username, login.HashPassword(v.password), v.email, v.role, v.created_at)
		if err != nil {
			log.Fatal(err)
		}
	}
	log.Println("Populated db with default users")

	err = statement.Close()
	if err != nil {
		log.Fatal(err)
	}
}

// used only within database package
func GetNameById(user_id int) string {
	query := "SELECT username FROM users WHERE id=?"
	row := db.QueryRow(query, user_id)

	var result string
	err := row.Scan(&result)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Username fetched")
	}

	return result
}

func InsertUser(user User) error {
	statement, err := db.Prepare(`
		INSERT INTO users (
			username,
			password,
			email
		) VALUES (?, ?, ?)
	`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = statement.Exec(user.Username, login.HashPassword(user.Password), user.Email)
	if err != nil {
		return err
	} else {
		log.Println("New user added")
	}

	err = statement.Close()
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func FetchUser(username, password string) (*User, error) {
	statement, err := db.Prepare(`SELECT id, username, email, password FROM users WHERE username =?`)
	if err != nil {
		log.Fatal(err)
	}

	row := statement.QueryRow(username)

	user := &User{}
	err = row.Scan(&user.Id, &user.Username, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("User not found")
		}
		return nil, err
	}

	err = login.CheckPasswordHash(user.Password, password)
	if err != nil {
		return nil, err
	}

	log.Println("User fetched")
	err = statement.Close()
	if err != nil {
		log.Fatal(err)
	}
	return user, nil
}
