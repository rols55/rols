package controllers

import (
	"fmt"
	"net/http"

	"forum/model"
	"forum/shared/logger"

	"golang.org/x/crypto/bcrypt"
)

type RegisterResponse struct {
}

func (h *BaseController) HandleRegisterError(w http.ResponseWriter, errStr string, errMsg string) {
	logger.Error(errStr)
	h.statusJSON(w, InvalidInput(errMsg))
}

// Process the register form
func (h *BaseController) RegisterPOST(w http.ResponseWriter, r *http.Request) {

	//if the form values are empty, send them back to login page
	if r.FormValue("username") == "" {
		errStr := "Failed register attempt: No username received"
		errMsg := "Username missing"
		h.HandleRegisterError(w, errStr, errMsg)
		return
	} else if r.FormValue("age") == "" {
		errStr := "Failed register attempt: No age received"
		errMsg := "Age missing"
		h.HandleRegisterError(w, errStr, errMsg)
		return
	} else if r.FormValue("gender") == "" {
		errStr := "Failed register attempt: No gender received"
		errMsg := "Gender missing"
		h.HandleRegisterError(w, errStr, errMsg)
		return
	} else if r.FormValue("firstName") == "" {
		errStr := "Failed register attempt: No first name received"
		errMsg := "First name missing"
		h.HandleRegisterError(w, errStr, errMsg)
		return
	} else if r.FormValue("lastName") == "" {
		errStr := "Failed register attempt: No last name received"
		errMsg := "Last name missing"
		h.HandleRegisterError(w, errStr, errMsg)
		return
	} else if r.FormValue("email") == "" {
		errStr := "Failed register attempt: No email received"
		errMsg := "Email address missing"
		h.HandleRegisterError(w, errStr, errMsg)
		return
	}  else if r.FormValue("password") == "" {
		errStr := "Failed register attempt: No password received"
		errMsg := "Password missing"
		h.HandleRegisterError(w, errStr, errMsg)
		return
	}

	//username
	username := r.FormValue("username")
	if _, err := model.GetUserByUsername(h.db, username); err == nil {
		errStr := fmt.Sprintf("Failed register attempt: User \"%s\" already exists", username)
		errMsg := fmt.Sprintf("Username \"%s\" is already taken", username)
		h.HandleRegisterError(w, errStr, errMsg)
		return
	}
	//email
	email := r.FormValue("email")
	if _, err := model.GetUserByEmail(h.db, email); err == nil {
		errStr := fmt.Sprintf("Failed register attempt: %v already exists", email)
		errMsg := fmt.Sprintf("Account with %s already exists", email)
		h.HandleRegisterError(w, errStr, errMsg)
		return
	}
	//password
	password, err := bcrypt.GenerateFromPassword([]byte(r.FormValue("password")), bcrypt.DefaultCost) //hash the password
	if err != nil {
		logger.Error(err.Error())
		h.statusJSON(w, InternalError)
		return
	}

	age := r.FormValue("age")
	gender := r.FormValue("gender")
	firstName := r.FormValue("firstName")
	lastName := r.FormValue("lastName")

	//Create the user
	if user, err := model.CreateUser(h.db, username, firstName, lastName, gender, age, email, string(password)); err != nil || user == nil {
		// we should be more specific with the error here, either our server fd up or something that the cilent did
		// check aganst specific error for sending a server error or return the client to registration page (and say what was wrong)
		logger.Error(err.Error())
		h.statusJSON(w, InternalError)
		return
	}
	//New user successfully created
	logger.Info("Registered new user: %v", username)
	h.statusJSON(w, Success)
}
