package model

import (
	"database/sql"
	"errors"
	"time"

	"01.kood.tech/git/rols55/social-network/pkg/logger"
)

type Comment struct {
	Id           int64     `json:"id,omitempty"`
	UserId       string    `json:"uuid,omitempty"`
	PostId       int64     `json:"post_id,omitempty"`
	Author       string    `json:"author,omitempty"`
	Title        string    `json:"title,omitempty"`
	Text         string    `json:"text,omitempty"`
	ImageComment string    `json:"image,omitempty"`
	CreationDate time.Time `json:"creation_date,omitempty"`
}

// Create a new post in the database and returns the post with the id in the database
func (c *Comment) Create(db *sql.DB) (*Comment, error) {
	result, err := db.Exec(`INSERT INTO comments(
		user_id,
		post_id,
		title,
		text,
		image,
		creation_date
		) values( ?,?,?,?,?,? )`,
		c.UserId,
		c.PostId,
		c.Title,
		c.Text,
		c.ImageComment,
		c.CreationDate,
	)
	if err != nil {
		return nil, errors.New(logger.GetCurrentFuncName() + " " + err.Error())
	}

	if c.Id, err = result.LastInsertId(); err != nil {
		return nil, errors.New(logger.GetCurrentFuncName() + " " + err.Error())
	}

	return c, nil
}

// Delete a post from the database (this is "hard" delete, we should use "soft" delete instead)
func (p *Comment) Delete(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM comments WHERE id = ?", p.Id)
	if err != nil {
		return errors.New(logger.GetCurrentFuncName() + " " + err.Error())
	}
	p = nil
	return nil
}

// Updates the current post (we should add a update date also)
func (c *Comment) Update(db *sql.DB) error {
	if _, err := db.Exec(`UPDATE comments SET
		user_id = ?,
		post_id = ?,
		title = ?,
		text = ?,
		image = ?,
		creation_date = ?, 
		WHERE id = ?`,
		c.UserId,
		c.PostId,
		c.Title,
		c.Text,
		c.ImageComment,
		c.CreationDate,
		c.Id,
	); err != nil {
		return errors.New(logger.GetCurrentFuncName() + " " + err.Error())
	}
	return nil
}

// Returns all the posts (not comments) in the database in ascending order
func GetComments(db *sql.DB) ([]*Comment, error) {
	rows, err := db.Query(`SELECT
		comments.id,
		comments.user_id,
		comments.post_id,
		comments.title,
		comments.text,
		comments.image,
		comments.creation_date,
		users.username
		FROM comments INNER JOIN users ON comments.user_id = users.uuid ORDER BY creation_date ASC`)
	if err != nil {
		return nil, errors.New(logger.GetCurrentFuncName() + " " + err.Error())
	}
	defer rows.Close()

	comments := make([]*Comment, 0)
	for rows.Next() {
		var comment Comment
		err = rows.Scan(
			&comment.Id,
			&comment.UserId,
			&comment.PostId,
			&comment.Title,
			&comment.Text,
			&comment.ImageComment,
			&comment.CreationDate,
			&comment.Author,
		)
		if err != nil {
			return nil, errors.New(logger.GetCurrentFuncName() + " " + err.Error())
		}
		comments = append(comments, &comment)
	}
	if len(comments) == 0 {
		return nil, ErrNotFound
	}

	return comments, nil
}

// Returns comment by its id
func GetCommentById(db *sql.DB, Id int64) (*Comment, error) {
	row := db.QueryRow(`SELECT
		comments.id,
		comments.user_id,
		comments.post_id,
		comments.title,
		comments.text,
		comments.image,
		comments.creation_date,
		users.username
		FROM comments INNER JOIN users ON comments.user_id = users.uuid WHERE comments.id =?`, Id)
	comment := &Comment{}
	if err := row.Scan(
		&comment.Id,
		&comment.UserId,
		&comment.PostId,
		&comment.Title,
		&comment.Text,
		&comment.ImageComment,
		&comment.CreationDate,
		&comment.Author,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, errors.New(logger.GetCurrentFuncName() + " " + err.Error())
	}
	return comment, nil
}

func GetCommentByPostId(db *sql.DB, Id int64) ([]*Comment, error) {
	rows, err := db.Query(`SELECT
		comments.id,
		comments.user_id,
		comments.post_id,
		comments.title,
		comments.text,
		comments.image,
		comments.creation_date,
		users.username
		FROM comments INNER JOIN users ON comments.user_id = users.uuid WHERE comments.post_id = ? ORDER BY creation_date ASC`, Id)
	if err != nil {
		return nil, errors.New(logger.GetCurrentFuncName() + " " + err.Error())
	}
	defer rows.Close()

	comments := make([]*Comment, 0)
	for rows.Next() {
		var comment Comment
		err = rows.Scan(
			&comment.Id,
			&comment.UserId,
			&comment.PostId,
			&comment.Title,
			&comment.Text,
			&comment.ImageComment,
			&comment.CreationDate,
			&comment.Author,
		)
		if err != nil {
			return nil, errors.New(logger.GetCurrentFuncName() + " " + err.Error())
		}
		comments = append(comments, &comment)
	}
	if len(comments) == 0 {
		return nil, ErrNotFound
	}

	return comments, nil
}

func GetCommentsByUserUuid(db *sql.DB, userId string) ([]*Comment, error) {
	rows, err := db.Query(`SELECT
	comments.id,
	comments.user_id,
	comments.post_id,
	comments.title,
	comments.text,
	comments.image,
	comments.creation_date,
	users.username
	FROM comments INNER JOIN users ON comments.user_id = users.uuid WHERE comments.user_id = ? ORDER BY creation_date ASC`, userId)
	if err != nil {
		return nil, errors.New(logger.GetCurrentFuncName() + " " + err.Error())
	}
	defer rows.Close()

	comments := make([]*Comment, 0)
	for rows.Next() {
		var comment Comment
		err = rows.Scan(
			&comment.Id,
			&comment.UserId,
			&comment.PostId,
			&comment.Title,
			&comment.Text,
			&comment.ImageComment,
			&comment.CreationDate,
			&comment.Author,
		)
		if err != nil {
			return nil, errors.New(logger.GetCurrentFuncName() + " " + err.Error())
		}
		comments = append(comments, &comment)
	}
	if len(comments) == 0 {
		return nil, ErrNotFound
	}

	return comments, nil
}

func GetUserLikedComments(db *sql.DB, userId string) ([]*Comment, error) {
	rows, err := db.Query(`SELECT
	comments.id,
	comments.user_id,
	comments.post_id,
	comments.title,
	comments.text
	FROM comments JOIN reactions ON comments.id = reactions.post_id WHERE reactions.user_id = ?  AND reactions.is_post = false AND reactions.reaction = true;`, userId)
	if err != nil {
		return nil, errors.New(logger.GetCurrentFuncName() + " " + err.Error())
	}
	defer rows.Close()

	comments := make([]*Comment, 0)
	for rows.Next() {
		var comment Comment
		err = rows.Scan(&comment.Id, &comment.UserId, &comment.PostId, &comment.Title, &comment.Text)
		if err != nil {
			return nil, errors.New(logger.GetCurrentFuncName() + " " + err.Error())
		}
		comments = append(comments, &comment)
	}
	if len(comments) == 0 {
		return nil, ErrNotFound
	}

	return comments, nil
}

// Delete post by its id and any comments that the post has (this is "hard" delete, we should use "soft" delete instead)
func DeleteCommentById(db *sql.DB, Id int) error {
	_, err := db.Exec("DELETE FROM comments WHERE id = ?", Id)
	if err != nil {
		return errors.New(logger.GetCurrentFuncName() + " " + err.Error())
	}

	return nil
}

// Deletes all posts (this is "hard" delete, we should use "soft" delete instead)
func DeleteAllComments(db *sql.DB) error {
	if _, err := db.Exec("DELETE FROM comments"); err != nil {
		return errors.New(logger.GetCurrentFuncName() + " " + err.Error())
	}

	return nil
}

// Returns the number of posts (not comments) in the database
func GetNumCommentsByPostId(db *sql.DB, postId int) (int64, error) {
	var count int64
	row := db.QueryRow("SELECT COUNT(*) FROM comments WHERE post_id = ?", postId)
	if err := row.Scan(&count); err != nil {
		return 0, errors.New(logger.GetCurrentFuncName() + " " + err.Error())
	}

	return count, nil
}
