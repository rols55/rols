package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"forum/model"
	"forum/route/middleware/acl"
	"forum/shared/logger"
)

type CommentResponse struct {
	Status  string                      `json:"status"`
	User    *model.User                 `json:"user,omitempty"`
	Comment *model.CommentWithReactions `json:"comment"`
}

// Proccess the post creation form
func (h *BaseController) CreateCommentPOST(w http.ResponseWriter, r *http.Request) {
	var err error
	var user *model.User
	var postId int
	var post *model.Post

	// Is the user logged in
	if err = h.getKeyVal(r, acl.UserKey, &user); err != nil {
		h.statusJSON(w, Unauthorized)
		return
	}

	if r.FormValue("text") == "" {
		logger.Info("Failed create post attempt: No text or title provided")
		errMsg := "Text missing"
		h.statusJSON(w, InvalidInput(errMsg))
		return
	}

	if postId, err = strconv.Atoi(r.FormValue("postid")); err != nil {
		logger.Error(err)
		h.statusJSON(w, InternalError)
		return
	}

	if post, err = model.GetPostById(h.db, int64(postId)); err != nil {
		logger.Error(err)
		h.statusJSON(w, InternalError)
		return
	}

	user, err = model.GetUserById(h.db, user.Id)
	if err != nil {
		logger.Error(err)
		h.statusJSON(w, InternalError)
		return
	}

	comment := &model.Comment{
		UserId:       user.Id,
		PostId:       post.Id,
		Title:        fmt.Sprintf("RE: %s", post.Title),
		Text:         r.FormValue("text"),
		Author:       user.Username,
		CreationDate: time.Now(),
	}

	if _, err = comment.Create(h.db); err != nil {
		logger.Error(err)
		h.statusJSON(w, InternalError)
		return
	}

	var commentsWithReactions *model.CommentWithReactions
	if commentsWithReactions, err = model.GetCommentWithReactions(h.db, comment, user.Id); err != nil && err != model.ErrNotFound {
		logger.Error(err)
		h.statusJSON(w, InternalError)
	}

	res := CommentResponse{
		Status:  "OK",
		User:    user,
		Comment: commentsWithReactions,
	}

	logger.Info(fmt.Sprintf("New post created - %v", comment.Id))
	//h.statusJSON(w, Success)
	h.writeJSON(w, res)
}

// Delete the post
func (h *BaseController) DeleteComment(w http.ResponseWriter, r *http.Request) {

	var id int
	var err error
	var comment *model.Comment
	var user *model.User

	//Only authenticate users can delete comments
	if err = h.getKeyVal(r, acl.UserKey, &user); err != nil {
		h.statusJSON(w, Unauthorized)
		return
	}

	//Check if comment id is provided
	if r.FormValue("id") == "" {
		errMsg := "Missing Comment Id"
		h.statusJSON(w, InvalidInput(errMsg))
		logger.Info("Failed to delete Comment: %s", errMsg)
		return
	}

	if id, err = strconv.Atoi(r.FormValue("id")); err != nil {
		errMsg := "Invalid comment Id"
		h.statusJSON(w, InvalidInput(errMsg))
		logger.Info("Failed to delete comment: %s", errMsg)
		return
	}

	//Check if we have to post
	if comment, err = model.GetCommentById(h.db, int64(id)); err != nil {
		errMsg := "No such comment in DB"
		h.statusJSON(w, InvalidInput(errMsg))
		logger.Info("Failed to delete comment: %s", errMsg)
		return
	}

	//Check if we are the author of this post
	if comment.UserId != user.Id {
		errMsg := "User is not the author of this comment"
		h.statusJSON(w, InvalidInput(errMsg))
		logger.Info("Failed to delete comment: %s", errMsg)
		return
	}

	if err = comment.Delete(h.db); err != nil {
		logger.Error(err)
		h.statusJSON(w, InternalError)
		return
	}

	if err = model.RemoveReactionByPostId(h.db, comment.Id, false); err != nil {
		logger.Error(err)
		h.statusJSON(w, InternalError)
		return
	}

	logger.Info("Post deleted successfully %v:%v", comment.Id, comment.Title)
	h.writeJSON(w, map[string]any{"status": "OK", "post_id": comment.PostId})
}
