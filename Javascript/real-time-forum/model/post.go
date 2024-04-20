package model

import (
	"database/sql"
	"errors"
	"time"

	"forum/shared/logger"
)

type Post struct {
	Id           int64     `json:"id,omitempty"`
	UserId       int64     `json:"user_id,omitempty"`
	Author       string    `json:"author,omitempty"`
	Title        string    `json:"title,omitempty"`
	Text         string    `json:"text,omitempty"`
	CreationDate time.Time `json:"creation_date,omitempty"`
}

// Create a new post in the database and returns the post with the id in the database
func (p *Post) Create(db *sql.DB) (*Post, error) {
	result, err := db.Exec(`INSERT INTO posts(
		user_id,
		author,
		title,
		text,
		creation_date
		) values( ?,?,?,?,? )`,
		p.UserId,
		p.Author,
		p.Title,
		p.Text,
		p.CreationDate,
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
		creation_date = ?,   
		WHERE id = ?`,
		p.UserId,
		p.Author,
		p.Title,
		p.Text,
		p.CreationDate,
		p.Id,
	); err != nil {
		return errors.New(logger.GetCurrentFuncName() + " " + err.Error())
	}
	return nil
}

// Return the amount of comments for a post
func (p *Post) GetAmountOfComments(db *sql.DB) (commentsAmount int, err error) {
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

// Returns all the posts (not comments) in the database in descending order
func GetPosts(db *sql.DB) ([]*Post, error) {
	rows, err := db.Query(`SELECT
		id,
		user_id,
		author,
		title,
		text,
		creation_date
		FROM posts ORDER BY creation_date DESC`)
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
			&post.CreationDate,
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
		&post.CreationDate,
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
	rows, err := db.Query(`SELECT
	id,
	user_id,
	author,
	title,
	text,
	creation_date
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
			&post.CreationDate,
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

func GetUserLikedPosts(db *sql.DB, userId int64) ([]*Post, error) {
	rows, err := db.Query("SELECT posts.id, posts.user_id, posts.author, posts.title, posts.text FROM posts JOIN reactions ON posts.id = reactions.post_id WHERE reactions.user_id = ?  AND reactions.is_post = true AND reactions.reaction = true;", userId)
	if err != nil {
		return nil, errors.New(logger.GetCurrentFuncName() + " " + err.Error())
	}
	defer rows.Close()

	posts := make([]*Post, 0)
	for rows.Next() {
		var post Post
		err = rows.Scan(&post.Id, &post.UserId, &post.Author, &post.Title, &post.Text)
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
