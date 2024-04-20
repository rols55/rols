package model

import (
	"database/sql"
	"errors"

	"01.kood.tech/git/rols55/social-network/pkg/logger"
)

type FollowJunction struct {
	Follower string
	Followed string
	Allowed  bool
}

type Follower struct {
	Follower string `json:"follower"`
	Username string `json:"username"`
	Accepted bool   `json:"accepted"`
}

func (j *FollowJunction) Create(db *sql.DB) error {
	_, err := db.Exec(`INSERT INTO followers(follower, followed, allowed) 
		values( ?, ?, ?)`,
		j.Follower, j.Followed, j.Allowed)
	if err != nil {
		return errors.New(logger.GetCurrentFuncName() + " " + err.Error())
	}
	return nil
}

func (j *FollowJunction) Delete(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM followers WHERE follower = ? AND followed = ?", j.Follower, j.Followed)
	if err != nil {
		return errors.New(logger.GetCurrentFuncName() + " " + err.Error())
	}
	return nil
}

func (j *FollowJunction) Allow(db *sql.DB) error {
	_, err := db.Exec(`UPDATE followers SET allowed = TRUE WHERE follower = ? AND followed = ?`, j.Follower, j.Followed)
	if err != nil {
		return errors.New(logger.GetCurrentFuncName() + " " + err.Error())
	}
	return nil
}

func GetFollowing(db *sql.DB, uuid string) ([]*Follower, error) {
	rows, err := db.Query(`SELECT users.username, followers.followed, followers.allowed
	FROM users
	INNER JOIN followers ON users.uuid = followers.followed
	WHERE followers.follower = ?`, uuid)
	array := make([]*Follower, 0)
	if err != nil {
		return array, err
	}
	for rows.Next() {
		var follower Follower
		rows.Scan(
			&follower.Username,
			&follower.Follower,
			&follower.Accepted,
		)
		array = append(array, &follower)
	}
	return array, nil
}

func GetFollowers(db *sql.DB, uuid string) ([]*Follower, error) {
	rows, err := db.Query(`SELECT users.username, followers.follower, followers.allowed
	FROM users
	INNER JOIN followers ON users.uuid = followers.follower
	WHERE followers.followed = ?`, uuid)
	array := make([]*Follower, 0)
	if err != nil {
		return array, err
	}
	for rows.Next() {
		var follower Follower
		rows.Scan(
			&follower.Username,
			&follower.Follower,
			&follower.Accepted,
		)
		array = append(array, &follower)
	}
	return array, nil
}

func AllowShow(db *sql.DB, self string, target string) bool {
	var value bool
	err := db.QueryRow(`SELECT CASE WHEN EXISTS (
		SELECT 1
		FROM followers
		WHERE follower = ? AND followed = ? AND allowed = TRUE
	) THEN TRUE ELSE FALSE END AS pair_exists;`, self, target).Scan(&value)
	if err != nil {
		return false
	}
	return value
}

func ChatFilter(db *sql.DB, self string, target string) bool {
	if self == target {
		return false
	}

	var value bool
	err := db.QueryRow(`
        SELECT CASE 
            WHEN EXISTS (
                SELECT 1
                FROM followers
                WHERE follower = ? AND followed = ? AND allowed = TRUE
            ) OR EXISTS (
                SELECT 1
                FROM followers
                WHERE follower = ? AND followed = ? AND allowed = TRUE
            ) THEN TRUE 
            ELSE FALSE 
        END AS pair_exists;`, self, target, target, self).Scan(&value)
	if err != nil {
		return false
	}
	return value
}

func IsFollowing(db *sql.DB, self string, target string) bool {
	if self == target {
		return false
	}

	var value bool
	err := db.QueryRow(`
		SELECT CASE 
			WHEN EXISTS (
				SELECT 1
				FROM followers
				WHERE follower = ? AND followed = ? AND allowed = TRUE
			) THEN TRUE 
			ELSE FALSE 
		END AS pair_exists;`, self, target).Scan(&value)
	if err != nil {
		return false
	}
	return value
}
