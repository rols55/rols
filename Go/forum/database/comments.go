package database

import (
	"database/sql"
	"log"
	"time"
)

type Comment struct {
	Id         int
	Post_id    int
	User_id    int
	Username   string
	Body       string
	Status     string
	Likes      int
	Dislikes   int
	Created_at string
}

var defaultComments []Comment

func makeDefaultComments() {
	var id int
	postId := []int{0, 1, 2}
	names := []string{"John Doe", "Spike Spiegel", "Marie Curie"}
	status := "approved"
	bodies := []string{"Yeppers", "Yup", "Eh", "Why not", "Helloooo!?", "Not is not the time for that"}

	for i := 0; i < len(postId); i++ {
		for j := 0; j < 2; j++ {
			defaultComments = append(defaultComments, Comment{
				id,
				postId[i],
				j,
				names[j],
				bodies[j],
				status,
				0,
				0,
				time.Now().Format("02/01/2006"),
			})
			id++
		}
	}
}

func populateDefaultComments() {
	statement, err := db.Prepare(`INSERT INTO comments (
		id,
		post_id,
		user_id,
		body,
		status,
		likes,
		dislikes,
		created_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		log.Println(err)
	}
	for _, v := range defaultComments {
		statement.Exec(v.Id, v.Post_id, v.User_id, v.Body, v.Status, v.Likes, v.Dislikes, v.Created_at)
	}
	log.Println("Populated db with default comments")

	err = statement.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func CreateComment(body, postId, userId string) {
	statement, err := db.Prepare(`INSERT INTO comments (
		post_id,
		user_id,
		body,
		status,
		likes,
		dislikes,
		created_at) VALUES (?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		log.Println(err)
	}

	_, err = statement.Exec(postId, userId, body, "approved", 0, 0, time.Now().Format("02/01/2006"))
	if err != nil {
		log.Println(err)
	}

	log.Println("Populated db with new comment")

	err = statement.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func GetComments(postId string, start, end int) []Comment {
	var comment Comment
	var comments []Comment
	var statement *sql.Stmt
	var rows *sql.Rows

	if end != 0 {
		statement, err = db.Prepare("SELECT * FROM comments WHERE post_id=? limit ?, ?")
		if err != nil {
			log.Fatal(err)
		}

		rows, err = statement.Query(postId, start, end)
		if err != nil {
			log.Fatal(err)

		}

	} else {
		statement, err = db.Prepare("SELECT * FROM comments WHERE post_id=?")
		if err != nil {
			log.Fatal(err)
		}

		rows, err = statement.Query(postId)
		if err != nil {
			log.Fatal(err)
		}

	}

	for rows.Next() {
		rows.Scan(&comment.Id, &comment.Post_id, &comment.User_id, &comment.Body, &comment.Status, &comment.Likes, &comment.Dislikes, &comment.Created_at)
		comment.Username = GetNameById(comment.User_id)
		comment.Likes = getCommentLikes(comment.Id)
		comment.Dislikes = getCommentDislikes(comment.Id)
		comments = append(comments, comment)
	}

	err = statement.Close()
	if err != nil {
		log.Fatal(err)
	}
	return comments
}

func LikeComment(commentId, userId string) {
	//check whether user liked that post or disliked it if not then like it
	if !checkIfCommentIsLiked(commentId, userId) && !checkIfCommentIsDisliked(commentId, userId) {
		likeComment(commentId, userId)
	} else if !checkIfCommentIsLiked(commentId, userId) && checkIfCommentIsDisliked(commentId, userId) {
		deleteCommentFromDislikes(commentId, userId)
		likeComment(commentId, userId)
	}
}

func DislikeComment(commentId, userId string) {
	if !checkIfCommentIsLiked(commentId, userId) && !checkIfCommentIsDisliked(commentId, userId) {
		dislikeComment(commentId, userId)
	} else if checkIfCommentIsLiked(commentId, userId) && !checkIfCommentIsDisliked(commentId, userId) {
		deleteCommentFromLikes(commentId, userId)
		dislikeComment(commentId, userId)
	}
}

func likeComment(commentId, userId string) {
	statement, err := db.Prepare("INSERT INTO likes VALUES (NULL, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}

	_, err = statement.Exec(commentId, userId)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("\nComment %v liked by %v", commentId, userId)
	err = statement.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func dislikeComment(commentId, userId string) {
	statement, err := db.Prepare("INSERT INTO dislikes VALUES (NULL, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}

	_, err = statement.Exec(commentId, userId)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("\nComment %v disliked by %v", commentId, userId)
	err = statement.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func deleteCommentFromLikes(commentId, userId string) {
	statement, err := db.Prepare("DELETE FROM likes WHERE comment_id=? AND user_id=?")
	if err != nil {
		log.Fatal(err)
	}

	_, err = statement.Exec(commentId, userId)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Like deleted")
	err = statement.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func deleteCommentFromDislikes(commentId, userId string) {
	statement, err := db.Prepare("DELETE FROM dislikes WHERE comment_id=? AND user_id=?")
	if err != nil {
		log.Fatal(err)
	}

	_, err = statement.Exec(commentId, userId)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Dislike deleted")
	err = statement.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func checkIfCommentIsLiked(commentId, userId string) bool {
	statement, err := db.Prepare("SELECT comment_id FROM likes WHERE comment_id=? AND user_id=?")
	if err != nil {
		log.Fatal(err)
	}

	row := statement.QueryRow(commentId, userId).Scan(nil)
	if row != nil {
		if row == sql.ErrNoRows {
			return false
		}
	}
	err = statement.Close()
	if err != nil {
		log.Fatal(err)
	}
	return true
}

func checkIfCommentIsDisliked(commentId, userId string) bool {
	statement, err := db.Prepare("SELECT comment_id FROM dislikes WHERE comment_id=? AND user_id=?")
	if err != nil {
		log.Fatal(err)
	}

	row := statement.QueryRow(commentId, userId).Scan(nil)
	if row != nil {
		if row == sql.ErrNoRows {
			return false
		}
	}
	err = statement.Close()
	if err != nil {
		log.Fatal(err)
	}
	return true
}

func getCommentLikes(commentId int) int {
	query := "SELECT COUNT(*) FROM likes WHERE comment_id=?"
	row := db.QueryRow(query, commentId)

	var result int
	err := row.Scan(&result)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Likes fetched")
	}

	return result
}

func getCommentDislikes(commentId int) int {
	query := "SELECT COUNT(*) FROM dislikes WHERE comment_id=?"
	row := db.QueryRow(query, commentId)

	var result int
	err := row.Scan(&result)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Dislikes fetched")
	}

	return result
}
