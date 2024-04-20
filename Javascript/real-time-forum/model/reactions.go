package model

import (
	"database/sql"
	"errors"

	"forum/shared/logger"
)

type Reaction struct {
	Id       int64 `json:"id,omitempty"`
	UserId   int64 `json:"user_id,omitempty"`
	PostId   int64 `json:"post_id,omitempty"`
	Reaction bool  `json:"reaction"` //True if liked, False if disliked
	IsPost   bool  `json:"is_post"`  //True if its post false if its comment
}

type PostWithReactionsAndCommentQty struct {
	Post
	Likes        int     `json:"likes"`
	Dislikes     int     `json:"dislikes"`
	UserLiked    bool    `json:"user_liked"`
	UserDisliked bool    `json:"user_disliked"`
	CommentsQty  int     `json:"comments_qty"`
	CategoryIds  []int64 `json:"category_ids,omitempty"`
}

type CommentWithReactions struct {
	Comment
	Likes        int  `json:"likes"`
	Dislikes     int  `json:"dislikes"`
	UserLiked    bool `json:"user_liked"`
	UserDisliked bool `json:"user_disliked"`
}

// Adds a reaction (like or dislike)
func AddReaction(db *sql.DB, userID int64, postID int64, reaction bool, isPost bool) (updatedReactionAmount int, err error) {
	_, err = db.Exec(`INSERT INTO reactions (user_id, post_id, reaction, is_post)
	VALUES (?, ?, ?, ?);`, userID, postID, reaction, isPost)
	if err != nil {
		return 0, errors.New(logger.GetCurrentFuncName() + " " + err.Error())
	}
	err = db.QueryRow(`SELECT COUNT(*) FROM reactions WHERE post_id = ? AND Reaction = ? AND is_post = ?;`, postID, reaction, isPost).Scan(&updatedReactionAmount)
	if err != nil {
		return 0, errors.New(logger.GetCurrentFuncName() + " " + err.Error())
	}
	return updatedReactionAmount, nil
}

// Removes a reaction (like or dislike)
func RemoveReaction(db *sql.DB, userID int64, postID int64, reaction bool, isPost bool) (updatedReactionAmount int, err error) {
	_, err = db.Exec(`DELETE FROM reactions WHERE user_id = ? AND post_id = ? AND reaction = ? AND is_post = ?;`, userID, postID, reaction, isPost)
	if err != nil {
		return 0, errors.New(logger.GetCurrentFuncName() + " " + err.Error())
	}
	err = db.QueryRow(`SELECT COUNT(*) FROM reactions WHERE post_id = ? AND Reaction = ? AND is_post = ?;`, postID, reaction, isPost).Scan(&updatedReactionAmount)
	if err != nil {
		return 0, errors.New(logger.GetCurrentFuncName() + " " + err.Error())
	}
	return updatedReactionAmount, nil
}

func RemoveReactionByPostId(db *sql.DB, postId int64, isPost bool) error {
	_, err := db.Exec(`DELETE FROM reactions WHERE post_id = ? AND is_post = ?;`, postId, isPost)
	if err != nil {
		return errors.New(logger.GetCurrentFuncName() + " " + err.Error())
	}
	return nil
}

func GetReactionAmountsByPostId(db *sql.DB, postID int64, isPost bool) (likes int, dislikes int, err error) {
	err = db.QueryRow(`SELECT COUNT(*) FROM reactions WHERE post_id = ? AND Reaction = true AND is_post = ?;`, postID, isPost).Scan(&likes)
	if err != nil {
		return 0, 0, errors.New(logger.GetCurrentFuncName() + " " + err.Error())
	}
	err = db.QueryRow(`SELECT COUNT(*) FROM reactions WHERE post_id = ? AND Reaction = false AND is_post = ?;`, postID, isPost).Scan(&dislikes)
	if err != nil {
		return 0, 0, errors.New(logger.GetCurrentFuncName() + " " + err.Error())
	}
	return likes, dislikes, nil
}

func GetUserReactionsByPostId(db *sql.DB, userID int64, postID int64, isPost bool) (*bool, error) {
	var reaction bool
	err := db.QueryRow(`SELECT reaction FROM reactions WHERE user_id = ? AND post_id = ? AND is_post = ?;`, userID, postID, isPost).Scan(&reaction)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No entry was found, return nil as result
		}
		return nil, err
	}
	return &reaction, nil // Return true for like, false for dislike
}

func GetPostsWithReactionsAndCommentQty(db *sql.DB, posts []*Post, userID int64) ([]*PostWithReactionsAndCommentQty, error) {
	postsWithReactionsAndCommentQty := make([]*PostWithReactionsAndCommentQty, 0)
	for _, post := range posts {
		likes, dislikes, err := GetReactionAmountsByPostId(db, post.Id, true)
		if err != nil {
			return nil, errors.New(logger.GetCurrentFuncName() + " " + err.Error())
		}
		commentsQty, err := post.GetAmountOfComments(db)
		if err != nil {
			return nil, errors.New(logger.GetCurrentFuncName() + " " + err.Error())
		}

		userLiked := false
		userDisliked := false
		if userID != 0 {
			reaction, err := GetUserReactionsByPostId(db, userID, post.Id, true)
			if err != nil {
				return nil, errors.New(logger.GetCurrentFuncName() + " " + err.Error())
			} else if reaction != nil {
				if *reaction {
					userLiked = true
				} else {
					userDisliked = true
				}
			}
		}

		var categoryIds []int64
		if categoryIds, err = GetCategoryIdByPostId(db, post.Id); err != nil {
			return nil, errors.New(logger.GetCurrentFuncName() + " " + err.Error())
		}

		postsWithReactionsAndCommentQty = append(postsWithReactionsAndCommentQty, &PostWithReactionsAndCommentQty{
			Post:         *post,
			Likes:        likes,
			Dislikes:     dislikes,
			UserLiked:    userLiked,
			UserDisliked: userDisliked,
			CommentsQty:  commentsQty,
			CategoryIds:  categoryIds,
		})
	}
	return postsWithReactionsAndCommentQty, nil
}

func GetPostWithReactionsAndCommentQty(db *sql.DB, post *Post, userID int64) (*PostWithReactionsAndCommentQty, error) {
	likes, dislikes, err := GetReactionAmountsByPostId(db, post.Id, true)
	if err != nil {
		return nil, errors.New(logger.GetCurrentFuncName() + " " + err.Error())
	}
	commentsQty, err := post.GetAmountOfComments(db)
	if err != nil {
		return nil, errors.New(logger.GetCurrentFuncName() + " " + err.Error())
	}
	userLiked := false
	userDisliked := false
	if userID != 0 {
		reaction, err := GetUserReactionsByPostId(db, userID, post.Id, true)
		if err != nil {
			return nil, errors.New(logger.GetCurrentFuncName() + " " + err.Error())
		} else if reaction != nil {
			if *reaction {
				userLiked = true
			} else {
				userDisliked = true
			}
		}
	}

	var categoryIds []int64
	if categoryIds, err = GetCategoryIdByPostId(db, post.Id); err != nil {
		return nil, errors.New(logger.GetCurrentFuncName() + " " + err.Error())
	}

	postWithReactionsAndCommentQty := &PostWithReactionsAndCommentQty{
		Post:         *post,
		Likes:        likes,
		Dislikes:     dislikes,
		UserLiked:    userLiked,
		UserDisliked: userDisliked,
		CommentsQty:  commentsQty,
		CategoryIds:  categoryIds,
	}
	return postWithReactionsAndCommentQty, nil
}

func GetCommentsWithReactions(db *sql.DB, comments []*Comment, userID int64) ([]*CommentWithReactions, error) {
	commentsWithReactions := make([]*CommentWithReactions, 0)
	for _, comment := range comments {
		likes, dislikes, err := GetReactionAmountsByPostId(db, comment.Id, false)
		if err != nil {
			return nil, errors.New(logger.GetCurrentFuncName() + " " + err.Error())
		}

		userLiked := false
		userDisliked := false
		if userID != 0 {
			reaction, err := GetUserReactionsByPostId(db, userID, comment.Id, false)
			if err != nil {
				return nil, errors.New(logger.GetCurrentFuncName() + " " + err.Error())
			} else if reaction != nil {
				if *reaction {
					userLiked = true
				} else {
					userDisliked = true
				}
			}
		}
		commentsWithReactions = append(commentsWithReactions, &CommentWithReactions{
			Comment:      *comment,
			Likes:        likes,
			Dislikes:     dislikes,
			UserLiked:    userLiked,
			UserDisliked: userDisliked,
		})
	}
	return commentsWithReactions, nil
}

func GetCommentWithReactions(db *sql.DB, comment *Comment, userID int64) (*CommentWithReactions, error) {
	likes, dislikes, err := GetReactionAmountsByPostId(db, comment.Id, false)
	if err != nil {
		return nil, errors.New(logger.GetCurrentFuncName() + " " + err.Error())
	}
	userLiked := false
	userDisliked := false
	if userID != 0 {
		reaction, err := GetUserReactionsByPostId(db, userID, comment.Id, false)
		if err != nil {
			return nil, errors.New(logger.GetCurrentFuncName() + " " + err.Error())
		} else if reaction != nil {
			if *reaction {
				userLiked = true
			} else {
				userDisliked = true
			}
		}
	}
	commentWithReactions := &CommentWithReactions{
		Comment:      *comment,
		Likes:        likes,
		Dislikes:     dislikes,
		UserLiked:    userLiked,
		UserDisliked: userDisliked,
	}
	return commentWithReactions, nil
}
