package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"01.kood.tech/git/rols55/social-network/pkg/api/model"
	"01.kood.tech/git/rols55/social-network/pkg/logger"
	"01.kood.tech/git/rols55/social-network/pkg/session"

	"golang.org/x/crypto/bcrypt"
)

type LoginResponse struct {
	Status         string        `json:"status"`
	User           *model.User   `json:"user,omitempty"`
	Posts          []*model.Post `json:"posts,omitempty"`
	CookieName     string
	SessionToken   string
	ExpirationTime time.Time `json:"categories,omitempty"`
}

type RequestBody struct {
	SessionToken string `json:"session_token"`
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
		w.WriteHeader(http.StatusBadRequest)
		h.statusJSON(w, InvalidInput(errMsg))
		return
		//Does the password match?
	} else if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err == nil {

		var posts []*model.Post

		logger.Info("User %v signed in", username)

		//make the session for this user
		sess := session.New(&w, user.Id, user.UUID)

		// get all the posts from the database, throw an http.error if there is error and its not ErrNotFound (no posts)
		if posts, err = model.GetPosts(h.db); err != nil && err != model.ErrNotFound {
			logger.Error(err)
			h.statusJSON(w, NotFound)
		}

		// set the data for response
		res := LoginResponse{
			Status:         "OK",
			User:           user,
			Posts:          posts,
			CookieName:     session.CookieName,
			SessionToken:   sess.Token,
			ExpirationTime: sess.Expires,
		}

		//Login was successful, send successful JSON response
		h.writeJSON(w, res)
		return
	}

	//Login failed, send unsuccessful JSON response
	errMsg := "Wrong password"
	logger.Error(err)
	w.WriteHeader(http.StatusBadRequest)
	h.statusJSON(w, InvalidInput(errMsg))
}

// Logging out handler/controller
func (h *BaseController) CheckSession(w http.ResponseWriter, r *http.Request) {
	// Get the session token from the request body
	var requestBody RequestBody
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Access the session_token value from the RequestBody struct
	sessionToken := requestBody.SessionToken

	// Check if the session token exists in the session manager
	sess, err := session.GetSession(sessionToken)
	if err != nil {
		if err == session.ErrNotFound {
			// Handle case where session is not found
			logger.Warning("Session not found")
			http.Error(w, "Session not found", http.StatusUnauthorized)
			return
		}
		if err == session.ErrExpired {
			// Handle case where session is expired
			logger.Warning("Session expired")
			http.Error(w, "Session expired", http.StatusUnauthorized)
			return
		}
		// Handle other errors
		logger.Error(err)
		http.Error(w, "Error retrieving session", http.StatusInternalServerError)
		return
	}

	// Session found and not expired
	fmt.Fprintln(w, "Session found. User ID:", sess)

}

// Logging out handler/controller
func (h *BaseController) Logout(w http.ResponseWriter, r *http.Request) {

	session.Delete(&w, r)

	//Logout was successful and send them back to home page
	//http.Redirect(w, r, "/", http.StatusFound)
	h.statusJSON(w, Redirect("/"))
}
