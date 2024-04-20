package controllers

import (
	"net/http"
	"strings"

	"forum/model"
	"forum/route/middleware/acl"
	"forum/shared/logger"
	"forum/view"
)

type ProfileResponse struct {
	User  *model.User                             `json:"user"`
	Posts []*model.PostWithReactionsAndCommentQty `json:"posts,omitempty"`
	Users []*model.User                           `json:"users,omitempty"`
}

func (h *BaseController) ProfileGET(w http.ResponseWriter, r *http.Request) {
	var err error
	var posts []*model.Post
	var postsWithReactionsAndCommentQty []*model.PostWithReactionsAndCommentQty
	var user *model.User

	view := view.New(h.db, "profile.html")

	if strings.HasSuffix(r.URL.Path, "/liked") {
		h.ProfileGetLikedPosts(w, r, view)
		return
	}

	if err = h.getKeyVal(r, acl.UserKey, &user); err != nil {
		logger.Error(err)
		h.statusJSON(w, Unauthorized)
	}

	// get user's created posts from the database, throw an http.error if there is error and its not ErrNotFound (no posts)
	if posts, err = model.GetPostsByUserId(h.db, user.Id); err != nil && err != model.ErrNotFound {
		logger.Error(err)
		h.statusJSON(w, InternalError)
	}
	// get reactions and amount of comment for posts
	if postsWithReactionsAndCommentQty, err = model.GetPostsWithReactionsAndCommentQty(h.db, posts, user.Id); err != nil && err != model.ErrNotFound {
		logger.Error(err)
		h.statusJSON(w, InternalError)
	}

	// shorten the longer posts so only preview would be seen on main page
	for _, post := range postsWithReactionsAndCommentQty {
		if len(post.Post.Text) < 450 {
			continue
		}
		for i := 440; i < len(post.Text); i++ { // find first space after 440 char
			if string(post.Text[i]) == " " {
				post.Text = post.Text[0:i] + "..."
				break
			}
		}
	}

	res := ProfileResponse{
		User:  user,
		Posts: postsWithReactionsAndCommentQty,
	}

	h.writeJSON(w, res)
}

func (h *BaseController) ProfileGetLikedPosts(w http.ResponseWriter, r *http.Request, view *view.View) {
	var err error
	var posts []*model.Post
	var postsWithReactionsAndCommentQty []*model.PostWithReactionsAndCommentQty
	var user *model.User

	if err = h.getKeyVal(r, acl.UserKey, &user); err != nil {
		logger.Error(err)
		h.statusJSON(w, Unauthorized)
	}

	// get user liked posts from the database, throw an http.error if there is error and its not ErrNotFound (no posts)
	if posts, err = model.GetUserLikedPosts(h.db, user.Id); err != nil && err != model.ErrNotFound {
		logger.Error(err)
		h.statusJSON(w, InternalError)
	}
	// get reactions and amount of comment for posts
	if postsWithReactionsAndCommentQty, err = model.GetPostsWithReactionsAndCommentQty(h.db, posts, user.Id); err != nil && err != model.ErrNotFound {
		logger.Error(err)
		h.statusJSON(w, InternalError)
	}

	// shorten the longer posts so only preview would be seen on main page
	for _, post := range postsWithReactionsAndCommentQty {
		if len(post.Text) < 450 {
			continue
		}
		for i := 440; i < len(post.Text); i++ { // find first space after 440 char
			if string(post.Text[i]) == " " {
				post.Text = post.Text[0:i] + "..."
				break
			}
		}
	}

	res := ProfileResponse{
		User:  user,
		Posts: postsWithReactionsAndCommentQty,
	}

	h.writeJSON(w, res)
}

func (h *BaseController) UsersGET(w http.ResponseWriter, r *http.Request) {
	var err error
	var user *model.User
	var users []*model.User

	if err = h.getKeyVal(r, acl.UserKey, &user); err != nil {
		logger.Error(err)
		h.statusJSON(w, Unauthorized)
		return
	}

	/*
		if users, err = model.GetUsers(h.db, user.Id); err != nil && err != model.ErrNotFound {
			logger.Error(err)
			h.statusJSON(w, InternalError)
		}
	*/

	if users, err = model.GetUsersSorted(h.db, user); err != nil && err != model.ErrNotFound {
		logger.Error(err)
		h.statusJSON(w, InternalError)
	}

	res := ProfileResponse{
		User:  user,
		Users: users,
	}

	h.writeJSON(w, res)
}
