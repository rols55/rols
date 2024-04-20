package controllers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"01.kood.tech/git/rols55/social-network/pkg/api/model"
	"01.kood.tech/git/rols55/social-network/pkg/logger"

	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

var imgProfile = "media/profile"

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
		w.WriteHeader(http.StatusBadRequest)
		h.HandleRegisterError(w, errStr, errMsg)
		return
	} else if r.FormValue("age") == "" {
		errStr := "Failed register attempt: No age received"
		errMsg := "Age missing"
		w.WriteHeader(http.StatusBadRequest)
		h.HandleRegisterError(w, errStr, errMsg)
		return
	} else if r.FormValue("gender") == "" {
		errStr := "Failed register attempt: No gender received"
		errMsg := "Gender missing"
		w.WriteHeader(http.StatusBadRequest)
		h.HandleRegisterError(w, errStr, errMsg)
		return
	} else if r.FormValue("firstName") == "" {
		errStr := "Failed register attempt: No first name received"
		errMsg := "First name missing"
		w.WriteHeader(http.StatusBadRequest)
		h.HandleRegisterError(w, errStr, errMsg)
		return
	} else if r.FormValue("lastName") == "" {
		errStr := "Failed register attempt: No last name received"
		errMsg := "Last name missing"
		w.WriteHeader(http.StatusBadRequest)
		h.HandleRegisterError(w, errStr, errMsg)
		return
	} else if r.FormValue("email") == "" {
		errStr := "Failed register attempt: No email received"
		errMsg := "Email address missing"
		w.WriteHeader(http.StatusBadRequest)
		h.HandleRegisterError(w, errStr, errMsg)
		return
	} else if r.FormValue("password") == "" {
		errStr := "Failed register attempt: No password received"
		errMsg := "Password missing"
		w.WriteHeader(http.StatusBadRequest)
		h.HandleRegisterError(w, errStr, errMsg)
		return
	}

	//username
	username := r.FormValue("username")
	if _, err := model.GetUserByUsername(h.db, username); err == nil {
		errStr := fmt.Sprintf("Failed register attempt: User \"%s\" already exists", username)
		errMsg := fmt.Sprintf("Username \"%s\" is already taken", username)
		w.WriteHeader(http.StatusBadRequest)
		h.HandleRegisterError(w, errStr, errMsg)
		return
	}
	//email
	email := r.FormValue("email")
	if _, err := model.GetUserByEmail(h.db, email); err == nil {
		errStr := fmt.Sprintf("Failed register attempt: %v email already exists", email)
		errMsg := fmt.Sprintf("Account with %s email already exists", email)
		w.WriteHeader(http.StatusBadRequest)
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
	nickName := r.FormValue("nickName")
	aboutMe := r.FormValue("aboutMe")

	imgName := "defaultImage.png"

	file, handler, err := r.FormFile("image")
	if err != nil {
		if err != http.ErrMissingFile {
			// Handle other errors if necessary
			log.Println("Error retrieving form file:", err)
			http.Error(w, "Error retrieving form file", http.StatusBadRequest)
			return
		}
		// No file was uploaded, handle this case accordingly
		imgName = "defaultImage.png"
	} else {
		defer file.Close()

		imgName = uuid.Must(uuid.NewV4()).String() + filepath.Ext(handler.Filename)

		err = os.MkdirAll(imgProfile, 0755)
		if err != nil {
			log.Printf("Error creating directory '%s': %v", imgProfile, err)
			http.Error(w, "Error creating directory", http.StatusInternalServerError)
			return
		}

		newFile, err := os.Create(imgProfile + "/" + imgName)
		if err != nil {
			log.Printf("Error creating new file '%s': %v", newFile.Name(), err)
			http.Error(w, "Error creating new file", http.StatusInternalServerError)
			return
		}
		defer newFile.Close()

		_, err = io.Copy(newFile, file)
		if err != nil {
			log.Println("Error copying file data:", err)
			http.Error(w, "Error copying file data", http.StatusInternalServerError)
			return
		}
	}

	//var public = true;

	fmt.Println("pssword", string(password))

	//Create the user
	if user, err := model.CreateUser(h.db, username, firstName, lastName, gender, age, email, nickName, aboutMe, imgName, false, string(password)); err != nil || user == nil {
		// we should be more specific with the error here, either our server fd up or something that the cilent did
		// check aganst specific error for sending a server error or return the client to registration page (and say what was wrong)
		logger.Error(err.Error())
		h.statusJSON(w, InternalError)
		return
	}

	if err != nil {
		// Handle error
		fmt.Println("Error:", err)
		return
	}
	//New user successfully created
	logger.Info("Registered new user: %v", username)
	h.statusJSON(w, Success)
}
