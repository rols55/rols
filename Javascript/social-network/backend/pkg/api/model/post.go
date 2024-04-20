package model

import (
	"database/sql"
	"errors"
	"time"

	"01.kood.tech/git/rols55/social-network/pkg/logger"
)

type Post struct {
	Id           int64     `json:"id,omitempty"`
	UserId       string    `json:"uuid,omitempty"`
	Author       string    `json:"author,omitempty"`
	Title        string    `json:"title,omitempty"`
	Text         string    `json:"text,omitempty"`
	ImagePost    string    `json:"image,omitempty"`
	Privacy      string    `json:"privacy,omitempty"`
	Followers    string    `json:"followers,omitempty"`
	CreationDate time.Time `json:"creation_date,omitempty"`
	Comments     int64     `json:"comments_count,omitempty"`
	GroupId      int64     `json:"group_id,omitempty"`
}

// Create a new post in the database and returns the post with the id in the database
func (p *Post) Create(db *sql.DB) (*Post, error) {
	result, err := db.Exec(`INSERT INTO posts(
		user_id,
		author,
		title,
		text,
		image,
		privacy,
		followers,
		creation_date,
		group_id
		) values( ?,?,?,?,?,?,?,?,? )`,
		p.UserId,
		p.Author,
		p.Title,
		p.Text,
		p.ImagePost,
		p.Privacy,
		p.Followers,
		p.CreationDate,
		p.GroupId,
	)
	if err != nil {
		return nil, errors.New(logger.GetCurrentFuncName() + " " + err.Error())
	}

	if p.Id, err = result.LastInsertId(); err != nil {
		return nil, errors.New(logger.GetCurrentFuncName() + " " + err.Error())
	}

	return p, nil
}

// Delete a post from the database (this is "hard" delete, we should use "soft" delete instead)
func (p *Post) Delete(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM posts WHERE id = ?", p.Id)
	if err != nil {
		return errors.New(logger.GetCurrentFuncName() + " " + err.Error())
	}
	p = nil
	return nil
}

// Updates the current post (we should add a update date also)
func (p *Post) Update(db *sql.DB) error {
	if _, err := db.Exec(`UPDATE posts SET
		user_id = ?,
		author = ?,
		title = ?,
		text = ?,
		image =?,
		privacy =?,
		followers =?,
		creation_date = ?,
		group_id = ?,
		WHERE id = ?`,
		p.UserId,
		p.Author,
		p.Title,
		p.Text,
		p.ImagePost,
		p.Privacy,
		p.Followers,
		p.CreationDate,
		p.GroupId,
		p.Id,
	); err != nil {
		return errors.New(logger.GetCurrentFuncName() + " " + err.Error())
	}
	return nil
}

// Return the amount of comments for a post
func (p *Post) GetAmountOfComments(db *sql.DB) (commentsAmount int64, err error) {
	err = db.QueryRow(`SELECT COUNT(*) FROM comments WHERE post_id = ?`, p.Id).Scan(&commentsAmount)
	if err != nil {
		return 0, errors.New(logger.GetCurrentFuncName() + " " + err.Error())
	}
	return commentsAmount, nil
}

// Returns all the current posts comments
func (p *Post) GetComments(db *sql.DB) ([]*Comment, error) {
	return GetCommentByPostId(db, p.Id)
}

// Returns all the none group posts in the database, descending order
func GetPosts(db *sql.DB) ([]*Post, error) {
	rows, err := db.Query(`SELECT * FROM posts WHERE group_id < 1 ORDER BY creation_date DESC`)
	if err != nil {
		return nil, errors.New(logger.GetCurrentFuncName() + " " + err.Error())
	}
	defer rows.Close()
	posts := make([]*Post, 0)
	for rows.Next() {
		var post Post
		err = rows.Scan(
			&post.Id,
			&post.UserId,
			&post.Author,
			&post.Title,
			&post.Text,
			&post.ImagePost,
			&post.Privacy,
			&post.Followers,
			&post.CreationDate,
			&post.GroupId,
		)
		if err != nil {
			return nil, errors.New(logger.GetCurrentFuncName() + " " + err.Error())
		}
		post.Comments, _ = post.GetAmountOfComments(db)
		posts = append(posts, &post)
	}
	if len(posts) == 0 {
		return nil, ErrNotFound
	}

	return posts, nil
}

// Returns post by its id
func GetPostById(db *sql.DB, Id int64) (*Post, error) {
	row := db.QueryRow("SELECT * FROM posts WHERE id =?", Id)
	post := &Post{}
	if err := row.Scan(
		&post.Id,
		&post.UserId,
		&post.Author,
		&post.Title,
		&post.Text,
		&post.ImagePost,
		&post.Privacy,
		&post.Followers,
		&post.CreationDate,
		&post.GroupId,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, errors.New(logger.GetCurrentFuncName() + " " + err.Error())
	}
	return post, nil
}

// Delete post by its id and any comments that the post has (this is "hard" delete, we should use "soft" delete instead)
func DeletePostById(db *sql.DB, Id int) error {
	_, err := db.Exec("DELETE FROM posts WHERE id = ?", Id)
	if err != nil {
		return errors.New(logger.GetCurrentFuncName() + " " + err.Error())
	}

	return nil
}

// Deletes all posts (this is "hard" delete, we should use "soft" delete instead)
func DeleteAllPosts(db *sql.DB) error {
	if _, err := db.Exec("DELETE FROM posts"); err != nil {
		return errors.New(logger.GetCurrentFuncName() + " " + err.Error())
	}

	return nil
}

// Returns the number of posts (not comments) in the database
func GetNumPosts(db *sql.DB) (int, error) {
	var count int
	row := db.QueryRow("SELECT COUNT(*) FROM posts")
	if err := row.Scan(&count); err != nil {
		return 0, errors.New(logger.GetCurrentFuncName() + " " + err.Error())
	}

	return count, nil
}

func GetPostsByUserId(db *sql.DB, userId int64) ([]*Post, error) {
	rows, err := db.Query(`SELECT * FROM posts WHERE user_id = ? ORDER BY creation_date ASC`, userId)
	if err != nil {
		return nil, errors.New(logger.GetCurrentFuncName() + " " + err.Error())
	}
	defer rows.Close()

	posts := make([]*Post, 0)
	for rows.Next() {
		var post Post
		err = rows.Scan(
			&post.Id,
			&post.UserId,
			&post.Author,
			&post.Title,
			&post.Text,
			&post.ImagePost,
			&post.Privacy,
			&post.Followers,
			&post.CreationDate,
			&post.GroupId,
		)
		if err != nil {
			return nil, errors.New(logger.GetCurrentFuncName() + " " + err.Error())
		}
		posts = append(posts, &post)
	}
	if len(posts) == 0 {
		return nil, ErrNotFound
	}

	return posts, nil
}

func GetPostsByUserUuid(db *sql.DB, userId string) ([]*Post, error) {
	rows, err := db.Query(`SELECT
	*
	FROM posts WHERE user_id = ? ORDER BY creation_date ASC`, userId)
	if err != nil {
		return nil, errors.New(logger.GetCurrentFuncName() + " " + err.Error())
	}
	defer rows.Close()

	posts := make([]*Post, 0)
	for rows.Next() {
		var post Post
		err = rows.Scan(
			&post.Id,
			&post.UserId,
			&post.Author,
			&post.Title,
			&post.Text,
			&post.ImagePost,
			&post.Privacy,
			&post.Followers,
			&post.CreationDate,
			&post.GroupId,
		)
		if err != nil {
			return nil, errors.New(logger.GetCurrentFuncName() + " " + err.Error())
		}
		posts = append(posts, &post)
	}
	if len(posts) == 0 {
		return nil, ErrNotFound
	}

	return posts, nil
}
