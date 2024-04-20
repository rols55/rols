package main

import (
	"encoding/json"
	"fmt"
	"forum/database"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"strings"

	//"strings"
	"time"

	"github.com/google/uuid"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	database.RunDatabase()

	/*
		pass := []byte("pass")
		hash, _ := bcrypt.GenerateFromPassword(pass, bcrypt.MinCost)
		statement, _ := db.Prepare("INSERT INTO users (username, password, email) VALUES (?, ?, ?)")
		statement.Exec("user", hash, "rols@example.com")
		rows, _ := db.Query("SELECT username, password FROM users")
		var username string
		var password string
		for rows.Next() {
			rows.Scan(&username, &password)
			fmt.Println(username, password)
		}
	*/

	// Apply the checkSession middleware to all routes
	//router.Use(checkSessionMiddleware)

	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/", forumHandler)
	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/logout", logoutHandler)
	http.HandleFunc("/morePosts", getMorePosts)
	http.HandleFunc("/post", showPost)
	http.HandleFunc("/likePost", likePost)
	http.HandleFunc("/dislikePost", dislikePost)
	http.HandleFunc("/likeComment", likeComment)
	http.HandleFunc("/dislikeComment", dislikeComment)
	http.HandleFunc("/createComment", createComment)
	http.HandleFunc("/createPost", createPost)
	http.HandleFunc("/savePost", savePost)
	http.HandleFunc("/myPosts", serveUserPosts)

	// Serve static files (CSS, JavaScript, etc.)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Start the server
	log.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

/*
func checkSessionMiddleware(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Skip middleware for static files
			if strings.HasPrefix(r.URL.Path, "/static/") {
				next.ServeHTTP(w, r)
				return
			}

			// Skip middleware for the login and register route
			if r.URL.Path == "/login" || r.URL.Path == "/register" {
				next.ServeHTTP(w, r)
				return
			}

			// Retrieve the session ID from the request cookie
			cookie, err := r.Cookie("sessionID")
			if err != nil || cookie.Value == "" {
				// No session ID found, redirect to the login page
				http.Redirect(w, r, "/login", http.StatusFound)
				return
			}

			// Parse the session ID cookie value into a UUID
			_, err = uuid.Parse(cookie.Value)
			if err != nil {
				// Invalid session ID, redirect to the login page
				http.Redirect(w, r, "/login", http.StatusFound)
				return
			}

			// Session ID is valid, proceed to the next handler
			next.ServeHTTP(w, r)
		})
	}
*/

func showPost(w http.ResponseWriter, r *http.Request) {
	var postId string

	if r.Method == "GET" {
		postId = r.FormValue("postId")

		data := struct {
			Post     database.Post
			Comments []database.Comment
		}{
			database.GetPost(postId),
			database.GetComments(postId, 0, 0),
		}
		fmt.Println(data.Comments)
		tmpl, err := template.ParseFiles("static/pages/post.html", "static/pages/components/header.html")
		if err != nil {
			http.Error(w, "Error parsing template", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, data)
		u, err := url.Parse("http://localhost:8080/post")
		if err != nil {
			log.Fatal(err)
		}
		q := u.Query()
		q.Set("postId", postId)
		u.RawQuery = q.Encode()
		r.URL = u
	}
}

func forumHandler(w http.ResponseWriter, r *http.Request) {

	// posts := database.GetPosts(1, 10)

	if r.Method == "GET" {
		tmpl, err := template.ParseFiles("static/pages/index.html", "static/pages/components/header.html")
		if err != nil {
			http.Error(w, "Error parsing template", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, nil)
	}

}

func getMorePosts(w http.ResponseWriter, r *http.Request) {

	posts := database.GetAllPosts()
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(posts)
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the request body into a User struct
	var user database.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Println(err)
		fmt.Println(user)
		return
	}
	// Insert the user into the database
	err = database.InsertUser(user)
	if err != nil {
		w.WriteHeader(http.StatusConflict)
		http.Redirect(w, r, "/login", http.StatusFound)
	} else {
		// Return a success response
		w.WriteHeader(http.StatusOK)
		http.Redirect(w, r, "/login", http.StatusFound)
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	// Handle the login form submission
	var user database.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Error parsing JSON data", http.StatusBadRequest)
		fmt.Println(err)
		return
	}
	fmt.Println("query", user.Username, user.Password)
	// Perform authentication logic here
	// Fetch the user from the database based on username or email
	userFromDb, err := database.FetchUser(user.Username, user.Password)
	if err != nil {
		log.Println("Error fetching user:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	// Generate a UUID for the session
	sessionID := uuid.New()
	cookieValue := fmt.Sprintf("%v, %v, %v", sessionID.String(), userFromDb.Username, userFromDb.Id)
	// Store the session ID in the session cookie for 24hrs
	cookie := &http.Cookie{
		Name:    "sessionID",
		Value:   cookieValue,
		Path:    "/",
		Expires: time.Now().Add(24 * time.Hour),
	}
	// Set the cookie in the response
	http.SetCookie(w, cookie)
	// Redirect the user to the forum page
	http.Redirect(w, r, "/", http.StatusFound)
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	// Delete the session cookie by setting its MaxAge to -1
	cookie := &http.Cookie{
		Name:     "sessionID",
		Value:    "",
		HttpOnly: true,
		Path:     "/",
		MaxAge:   -1,
	}
	http.SetCookie(w, cookie)

	// Redirect the user to the login page or any other desired page
	http.Redirect(w, r, "/", http.StatusFound)
}

func likePost(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Header.Get("Referer"))

	if len(r.Cookies()) > 0 {
		postId := r.FormValue("like")
		cookies := r.Cookies()
		userId := strings.Split(cookies[0].Value, ", ")[2]
		database.LikePost(postId, userId)
		http.Redirect(w, r, r.Header.Get("Referer"), http.StatusFound)
	} else {
		http.Redirect(w, r, r.Header.Get("Referer"), http.StatusFound)
	}

}

func dislikePost(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Header.Get("Referer"))

	if len(r.Cookies()) > 0 {
		cookies := r.Cookies()
		userId := strings.Split(cookies[0].Value, ", ")[2]
		database.DislikePost(r.FormValue("dislike"), userId)
		http.Redirect(w, r, r.Header.Get("Referer"), http.StatusFound)
	} else {
		http.Redirect(w, r, r.Header.Get("Referer"), http.StatusFound)
	}
}

func likeComment(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Header.Get("Referer"))

	if len(r.Cookies()) > 0 {
		cookies := r.Cookies()
		userId := strings.Split(cookies[0].Value, ", ")[2]
		database.LikeComment(r.FormValue("like"), userId)
		http.Redirect(w, r, r.Header.Get("Referer"), http.StatusFound)
	} else {
		http.Redirect(w, r, r.Header.Get("Referer"), http.StatusFound)
	}
}

func dislikeComment(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Header.Get("Referer"))

	if len(r.Cookies()) > 0 {
		cookies := r.Cookies()
		userId := strings.Split(cookies[0].Value, ", ")[2]
		database.DislikeComment(r.FormValue("dislike"), userId)
		http.Redirect(w, r, r.Header.Get("Referer"), http.StatusFound)
	} else {
		http.Redirect(w, r, r.Header.Get("Referer"), http.StatusFound)
	}
}

func createComment(w http.ResponseWriter, r *http.Request) {
	body := r.FormValue("content")
	postId := r.FormValue("postId")
	cookies := r.Cookies()
	userId := strings.Split(cookies[0].Value, ", ")[2]

	database.CreateComment(body, postId, userId)

	http.Redirect(w, r, r.Header.Get("Referer"), http.StatusFound)

}

func createPost(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl, err := template.ParseFiles("static/pages/createPost.html", "static/pages/components/header.html")
		if err != nil {
			http.Error(w, "Error parsing template", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, nil)
	}
}

func savePost(w http.ResponseWriter, r *http.Request) {
	if len(r.Cookies()) > 0 {
		title := r.FormValue("title")
		body := r.FormValue("content")
		categories := r.FormValue("categories")
		cookies := r.Cookies()
		userId := strings.Split(cookies[0].Value, ", ")[2]

		database.CreatePost(title, body, categories, userId)

		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

}

func serveUserPosts(w http.ResponseWriter, r *http.Request) {
	cookies := r.Cookies()
	userId := strings.Split(cookies[0].Value, ", ")[2]

	userPosts := database.GetUserPosts(userId)
	if r.Method == "GET" {
		tmpl, err := template.ParseFiles("static/pages/myPosts.html", "static/pages/components/header.html")
		if err != nil {
			http.Error(w, "Error parsing template", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, userPosts)
	}
}
