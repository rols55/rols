package controllers

import (
	"fmt"
	"net/http"

	"forum/model"
	"forum/shared/logger"
	"forum/shared/session"

	"golang.org/x/crypto/bcrypt"
)

type LoginResponse struct {
	Status     string                                  `json:"status"`
	User       *model.User                             `json:"user,omitempty"`
	Posts      []*model.PostWithReactionsAndCommentQty `json:"posts,omitempty"`
	Categories []*model.Category                       `json:"categories,omitempty"`
}

// Process the login form
func (h *BaseController) LoginPOST(w http.ResponseWriter, r *http.Request) {
	var err error
	username := r.FormValue("username")
	password := r.FormValue("password")

	//if the form values are empty, send them back to login page (needs rework for better validation)
	if r.FormValue("username") == "" || r.FormValue("username") == "undefined" {
		errStr := fmt.Sprintf("Failed login attempt: No username received, Error: %v", err)
		errMsg := "Username missing"
		logger.Error(errStr)
		h.statusJSON(w, InvalidInput(errMsg))
		return
	}
	if r.FormValue("password") == "" || r.FormValue("password") == "undefined" {
		errStr := fmt.Sprintf("Failed login attempt: No password received, Error: %v", err)
		errMsg := "No password entered"
		logger.Error(errStr)
		h.statusJSON(w, InvalidInput(errMsg))
		return
	}

	//Does the given username exist in database?
	if user, err := model.GetUserByUsername(h.db, username); err != nil {
		errStr := fmt.Sprintf("Failed login attempt: No such user %s, Error: %v", username, err)
		errMsg := fmt.Sprintf("User \"%s\" doesn't exist", username)
		logger.Error(errStr)
		h.statusJSON(w, InvalidInput(errMsg))
		return
		//Does the password match?
	} else if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err == nil {

		var posts []*model.Post
		var postsWithReactionsAndCommentQty []*model.PostWithReactionsAndCommentQty
		var categories []*model.Category
		var userID int64

		logger.Info("User %v signed in", username)

		//make the session for this user
		session.New(&w, user.Id)

		//get all categories that have posts
		categories, err = model.GetCategoriesWithPosts(h.db)
		if err != nil && err != model.ErrNotFound {
			logger.Error(err)
		}

		// get all the posts from the database, throw an http.error if there is error and its not ErrNotFound (no posts)
		if posts, err = model.GetPosts(h.db); err != nil && err != model.ErrNotFound {
			logger.Error(err)
			h.statusJSON(w, NotFound)
		}
		// get reactions and amount of comment for posts
		if postsWithReactionsAndCommentQty, err = model.GetPostsWithReactionsAndCommentQty(h.db, posts, userID); err != nil && err != model.ErrNotFound {
			logger.Error(err)
			h.statusJSON(w, NotFound)
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

		// set the data for response
		res := LoginResponse{
			Status:     "OK",
			User:       user,
			Posts:      postsWithReactionsAndCommentQty,
			Categories: categories,
		}

		//Login was successful, send successful JSON response
		h.writeJSON(w, res)
		return
	}

	//Login failed, send unsuccessful JSON response
	errMsg := "Wrong password"
	logger.Error(err)
	h.statusJSON(w, InvalidInput(errMsg))
}

// Logging out handler/controller
func (h *BaseController) Logout(w http.ResponseWriter, r *http.Request) {

	session.Delete(&w, r)

	//Logout was successful and send them back to home page
	http.Redirect(w, r, "/", http.StatusFound)
	//h.statusJSON(w, Redirect("/"))
}
