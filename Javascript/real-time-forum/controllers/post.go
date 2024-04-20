package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"forum/model"
	"forum/route/middleware/acl"
	"forum/shared/logger"
)

type PostResponse struct {
	Status     string                                `json:"status"`
	User       *model.User                           `json:"user,omitempty"`
	Post       *model.PostWithReactionsAndCommentQty `json:"post,omitempty"`
	Categories []*model.Category                     `json:"categories,omitempty"`
	Comments   []*model.CommentWithReactions         `json:"comments,omitempty"`
}

// Respond to GET request with specific post
func (h *BaseController) PostGET(w http.ResponseWriter, r *http.Request) {
	var err error
	var postWithReactionsAndCommentQty *model.PostWithReactionsAndCommentQty

	//add user to view vars if logged in
	var user *model.User
	if err = h.getKeyVal(r, acl.UserKey, &user); err != nil && err != model.ErrNotFound {
		logger.Error(err)
		h.statusJSON(w, InternalError)
	}

	postId, err := strconv.Atoi(r.URL.Path[len("/post/"):])
	if err != nil {
		logger.Error(err)
		h.statusJSON(w, NotFound)
		return
	}

	//get the post to be rendered in the view
	post, err := model.GetPostById(h.db, int64(postId))
	if err != nil || post == nil {
		logger.Error(err)
		h.statusJSON(w, NotFound)
		return
	}

	categoriesId, err := model.GetCategoryIdByPostId(h.db, int64(postId))
	if err != nil || categoriesId == nil {
		logger.Error(err)
		h.statusJSON(w, NotFound)
		return
	}

	categories := make([]*model.Category, 0)
	for i := 0; i < len(categoriesId); i++ {
		category, err := model.GetCategoryById(h.db, categoriesId[i])
		if err != nil || post == nil {
			logger.Error(err)
			h.statusJSON(w, NotFound)
			return
		}
		categories = append(categories, category)
	}

	var commentsWithReactions []*model.CommentWithReactions
	//get the post's comments to be rendered in the view
	if comments, err := post.GetComments(h.db); err == nil {
		// get reactions for comments
		if commentsWithReactions, err = model.GetCommentsWithReactions(h.db, comments, user.Id); err != nil && err != model.ErrNotFound {
			logger.Error(err)
			h.statusJSON(w, NotFound)
		}
		// if err is ErrNotFound, then the post does not have any comments and we do nothing, on other type of errors. we have a internal problem
	} else if err != model.ErrNotFound {
		logger.Error(err)
		h.statusJSON(w, InternalError)
	}

	//get reaction amounts and user reactions for post
	if postWithReactionsAndCommentQty, err = model.GetPostWithReactionsAndCommentQty(h.db, post, user.Id); err != nil && err != model.ErrNotFound {
		logger.Error(err)
		h.statusJSON(w, NotFound)
	}

	res := PostResponse{
		Status:     "OK",
		User:       user,
		Post:       postWithReactionsAndCommentQty,
		Categories: categories,
		Comments:   commentsWithReactions,
	}

	h.writeJSON(w, res)
}

// Proccess the post creation form
func (h *BaseController) CreatePostPOST(w http.ResponseWriter, r *http.Request) {
	var err error
	var user *model.User
	var categories []*model.Category

	//check if the user is logged in
	if err = h.getKeyVal(r, acl.UserKey, &user); err != nil {
		logger.Error(err)
		h.statusJSON(w, Unauthorized)
	}

	if r.FormValue("title") == "" {
		logger.Info("Failed create post attempt: No text or title provided")
		errMsg := "Title missing"
		h.statusJSON(w, InvalidInput(errMsg))
		return
	}

	if len(r.FormValue("title")) > 60 {
		logger.Info("Failed create post attempt: Title too long")
		errMsg := "Title too long"
		h.statusJSON(w, InvalidInput(errMsg))
		return
	}

	if r.FormValue("text") == "" {
		logger.Info("Failed create post attempt: No text or title provided")
		errMsg := "Text missing"
		h.statusJSON(w, InvalidInput(errMsg))
		return
	}

	//here we get the array of selected categories from form
	for _, c := range strings.Split(r.FormValue("category"), ",") {
		categoryId, err := strconv.Atoi(c)
		if err != nil {
			logger.Error(err)
			errMsg := "Category missing"
			h.statusJSON(w, InvalidInput(errMsg))
			return
		}
		category, err := model.GetCategoryById(h.db, int64(categoryId))
		if err != nil {
			logger.Error(err)
			errMsg := "Category missing"
			h.statusJSON(w, InvalidInput(errMsg))
			return
		}
		categories = append(categories, category)
	}

	logger.Info("Category = %v", categories)

	post := &model.Post{
		UserId:       user.Id,
		Author:       user.Username,
		Title:        r.FormValue("title"),
		Text:         r.FormValue("text"),
		CreationDate: time.Now(),
	}

	if post, err = post.Create(h.db); err != nil {
		logger.Error(err)
		h.statusJSON(w, InternalError)
		return
	}

	//creating a relationship record between post and categories
	for i := 0; i < len(categories); i++ {
		relations := &model.PostsCategories{
			PostId:     post.Id,
			CategoryId: categories[i].Id,
		}
		if _, err = relations.Create(h.db); err != nil {
			logger.Error(err)
			h.statusJSON(w, InternalError)
			return
		}
	}

	var commentsWithReactions []*model.CommentWithReactions
	//get the post's comments to be rendered in the view
	if comments, err := post.GetComments(h.db); err == nil {
		// get reactions for comments
		if commentsWithReactions, err = model.GetCommentsWithReactions(h.db, comments, user.Id); err != nil && err != model.ErrNotFound {
			logger.Error(err)
			h.statusJSON(w, NotFound)
		}
		// if err is ErrNotFound, then the post does not have any comments and we do nothing, on other type of errors. we have a internal problem
	} else if err != model.ErrNotFound {
		logger.Error(err)
		h.statusJSON(w, InternalError)
	}

	//get reaction amounts and user reactions for post
	var postWithReactionsAndCommentQty *model.PostWithReactionsAndCommentQty
	if postWithReactionsAndCommentQty, err = model.GetPostWithReactionsAndCommentQty(h.db, post, user.Id); err != nil && err != model.ErrNotFound {
		logger.Error(err)
		h.statusJSON(w, NotFound)
	}

	logger.Info(fmt.Sprintf("New post created - %v", post.Id))
	//h.statusJSON(w, Redirect(fmt.Sprintf("/post/%v", post.Id)))
	res := PostResponse{
		Status:     "OK",
		User:       user,
		Post:       postWithReactionsAndCommentQty,
		Categories: categories,
		Comments:   commentsWithReactions,
	}

	h.writeJSON(w, res)
}

// Delete the post
func (h *BaseController) DeletePost(w http.ResponseWriter, r *http.Request) {

	var id int
	var err error
	var post *model.Post
	var user *model.User

	//Only authenticate users can delete posts
	if err = h.getKeyVal(r, acl.UserKey, &user); err != nil {
		logger.Error(err)
		h.statusJSON(w, Unauthorized)
	}

	//Check if post id is provided
	if r.FormValue("id") == "" {
		errMsg := "Missing Post Id"
		logger.Info("Failed to delete post: %s", errMsg)
		h.statusJSON(w, InvalidInput(errMsg))
		return
	}

	if id, err = strconv.Atoi(r.FormValue("id")); err != nil {
		errMsg := "Invalid post Id"
		logger.Info("Failed to delete post: %s", errMsg)
		h.statusJSON(w, InvalidInput(errMsg))
		return
	}

	//Check if we have to post
	if post, err = model.GetPostById(h.db, int64(id)); err != nil {
		errMsg := "No such post in DB"
		logger.Info("Failed to delete post: %s", errMsg)
		h.statusJSON(w, NotFound)
		return
	}

	//Check if we are the author of this post
	if post.UserId != user.Id {
		errMsg := "User is not the author of this post"
		logger.Info("Failed to delete post: %s", errMsg)
		h.statusJSON(w, Unauthorized)
		return
	}

	if err = post.Delete(h.db); err != nil {
		logger.Error(err)
		h.statusJSON(w, InternalError)
		return
	}

	if err = model.RemoveReactionByPostId(h.db, post.Id, true); err != nil {
		logger.Error(err)
		h.statusJSON(w, InternalError)
		return
	}

	logger.Info("Post deleted successfully %v:%v", post.Id, post.Title)
	h.statusJSON(w, Redirect("/"))
}
