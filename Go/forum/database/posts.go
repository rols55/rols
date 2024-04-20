package database

import (
	"database/sql"
	"log"
	"time"
)

type Post struct {
	Id         int
	User_id    int
	Username   string
	Title      string
	Body       string
	Categories string
	Likes      int
	Dislikes   int
	Status     string
	Created_at string
}

type PostToSend struct {
	Id         int    `json:"Id"`
	User_id    int    `json:"User_id"`
	Username   string `json:"Username"`
	Title      string `json:"Title"`
	Body       string `json:"Body"`
	Categories string `json:"Categories"`
	Likes      int    `json:"Likes"`
	Dislikes   int    `json:"Dislikes"`
	Status     string `json:"Status"`
	Created_at string `json:"Created_at"`
}

var defaultPosts []Post

func makeDefaultPosts() {
	id := 0
	titles := []string{"Rats", "Bars on cold planets?", "Radio isotopes"}
	bodies := []string{"Can anybody reccomend some shops for rat clothes. The cuurent ones for my rat are getting too small and the ones for opposums are a tad large.",
		"Hey I'm looking for some bars that are men only specifically on cold planets, anybody has any recommendations?",
		"Hey my source let's discuss your favourite radioisotopes. Mine is radium-223 'cause it helped my friend get rid of prostate cancer!"}
	user_ids := []int{}
	categories := []string{"outdoors", "pets", "isotopes"}
	for i := range defaultUsers {
		user_ids = append(user_ids, defaultUsers[i].id)
	}
	statuses := []string{"approved", "approved", "approved"}
	var post Post

	for i := range titles {
		post = Post{
			id,
			user_ids[i],
			GetNameById(user_ids[i]),
			titles[i],
			bodies[i],
			categories[i],
			0,
			0,
			statuses[i],
			time.Now().Format("02/01/2006"),
		}
		defaultPosts = append(defaultPosts, post)
		id++
	}
}

func populateDefaultPosts() {
	statement, err := db.Prepare(`INSERT INTO posts (
		id,
		user_id,
		title,
		body,
		categories,
		status,
		likes,
		dislikes,
		created_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		log.Println(err)
	}
	for _, v := range defaultPosts {
		_, err = statement.Exec(v.Id, v.User_id, v.Title, v.Body, v.Categories, v.Status, v.Likes, v.Dislikes, v.Created_at)
		if err != nil {
			log.Fatal(err)
		}
	}
	log.Println("Populated db with default posts")

	err = statement.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func CreatePost(title, body, categories, userId string) {
	statement, err := db.Prepare(`INSERT INTO posts (
		user_id,
		title,
		body,
		categories,
		status,
		likes,
		dislikes,
		created_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		log.Println(err)
	}

	_, err = statement.Exec(userId, title, body, categories, "approved", 0, 0, time.Now().Format("02/01/2006"))
	if err != nil {
		log.Println(err)
	}

	log.Println("Populated db with new post")

	err = statement.Close()
	if err != nil {
		log.Fatal(err)
	}
}

// start and end are determine range of posts database a query should fetch
func GetAllPosts() []PostToSend {
	var post PostToSend
	var posts []PostToSend
	statement, err := db.Prepare("SELECT * FROM posts")
	if err != nil {
		log.Fatal(err)
	}

	rows, err := statement.Query()
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		err = rows.Scan(&post.Id, &post.User_id, &post.Title, &post.Body, &post.Categories, &post.Status, &post.Likes, &post.Dislikes, &post.Created_at)
		if err != nil {
			log.Fatal(err)
		}
		post.Username = GetNameById(post.User_id)
		post.Likes = getPostLikes(post.Id)
		post.Dislikes = getPostDislikes(post.Id)
		posts = append(posts, post)
	}

	log.Println("Posts fetched")
	err = statement.Close()
	if err != nil {
		log.Fatal(err)
	}
	return posts
}

// start and end are determine range of posts database a query should fetch
func GetPosts(start, end int) []PostToSend {
	var post PostToSend
	var posts []PostToSend
	statement, err := db.Prepare("SELECT * FROM posts limit ?, ?")
	if err != nil {
		log.Fatal(err)
	}

	rows, err := statement.Query(start, end)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		err = rows.Scan(&post.Id, &post.Username, &post.Title, &post.Body, &post.Categories, &post.Status, &post.Likes, &post.Dislikes, &post.Created_at)
		if err != nil {
			log.Fatal(err)
		}
		post.Username = GetNameById(post.Id)
		post.Likes = getPostLikes(post.Id)
		post.Dislikes = getPostDislikes(post.Id)
		posts = append(posts, post)
	}

	log.Println("Posts fetched")
	err = statement.Close()
	if err != nil {
		log.Fatal(err)
	}
	return posts
}

func GetUserPosts(id string) []PostToSend {
	var post PostToSend
	var posts []PostToSend
	statement, err := db.Prepare("SELECT * FROM posts WHERE user_id=?")
	if err != nil {
		log.Fatal(err)
	}

	rows, err := statement.Query(id)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		err = rows.Scan(&post.Id, &post.Username, &post.Title, &post.Body, &post.Categories, &post.Status, &post.Likes, &post.Dislikes, &post.Created_at)
		if err != nil {
			log.Fatal(err)
		}
		post.Username = GetNameById(post.Id)
		post.Likes = getPostLikes(post.Id)
		post.Dislikes = getPostDislikes(post.Id)
		posts = append(posts, post)
	}

	log.Println("Posts fetched")
	err = statement.Close()
	if err != nil {
		log.Fatal(err)
	}
	return posts
}

func GetPost(id string) Post {
	var post Post
	statement, err := db.Prepare("SELECT * FROM posts WHERE id=?")
	if err != nil {
		log.Fatal(err)
	}

	rows, err := statement.Query(id)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		err = rows.Scan(&post.Id, &post.User_id, &post.Title, &post.Body, &post.Categories, &post.Status, &post.Likes, &post.Dislikes, &post.Created_at)
		if err != nil {
			log.Fatal(err)
		}
		post.Username = GetNameById(post.User_id)
		post.Likes = getPostLikes(post.Id)
		post.Dislikes = getPostDislikes(post.Id)
	}

	log.Println("Post fetched")
	err = statement.Close()
	if err != nil {
		log.Fatal(err)
	}
	return post
}

func LikePost(postId, userId string) {
	//check whether user liked that post or disliked it if not then like it
	if !checkIfPostIsLiked(postId, userId) && !checkIfPostIsDisliked(postId, userId) {
		likePost(postId, userId)
	} else if !checkIfPostIsLiked(postId, userId) && checkIfPostIsDisliked(postId, userId) {
		deleteFromDislikes(postId, userId)
		likePost(postId, userId)
	}
}

func DislikePost(postId, userId string) {
	if !checkIfPostIsLiked(postId, userId) && !checkIfPostIsDisliked(postId, userId) {
		dislikePost(postId, userId)
	} else if checkIfPostIsLiked(postId, userId) && !checkIfPostIsDisliked(postId, userId) {
		deleteFromLikes(postId, userId)
		dislikePost(postId, userId)
	}
}

func likePost(postId, userId string) {
	statement, err := db.Prepare("INSERT INTO likes VALUES (?, NULL, ?)")
	if err != nil {
		log.Fatal(err)
	}

	_, err = statement.Exec(postId, userId)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("\nPost %v liked by %v", postId, userId)
	err = statement.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func dislikePost(postId, userId string) {
	statement, err := db.Prepare("INSERT INTO dislikes VALUES (?, NULL, ?)")
	if err != nil {
		log.Fatal(err)
	}

	_, err = statement.Exec(postId, userId)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("\nPost %v disliked by %v", postId, userId)
	err = statement.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func deleteFromLikes(postId, userId string) {
	statement, err := db.Prepare("DELETE FROM likes WHERE post_id=? AND user_id=?")
	if err != nil {
		log.Fatal(err)
	}

	_, err = statement.Exec(postId, userId)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Like deleted")
	err = statement.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func deleteFromDislikes(postId, userId string) {
	statement, err := db.Prepare("DELETE FROM dislikes WHERE post_id=? AND user_id=?")
	if err != nil {
		log.Fatal(err)
	}

	_, err = statement.Exec(postId, userId)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Dislike deleted")
	err = statement.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func checkIfPostIsLiked(postId, userId string) bool {
	statement, err := db.Prepare("SELECT post_id FROM likes WHERE post_id=? AND user_id=?")
	if err != nil {
		log.Fatal(err)
	}

	row := statement.QueryRow(postId, userId).Scan(nil)
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

func checkIfPostIsDisliked(postId, userId string) bool {
	statement, err := db.Prepare("SELECT post_id FROM dislikes WHERE post_id=? AND user_id=?")
	if err != nil {
		log.Fatal(err)
	}

	row := statement.QueryRow(postId, userId).Scan(nil)
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

func getPostLikes(postId int) int {
	query := "SELECT COUNT(*) FROM likes WHERE post_id=?"
	row := db.QueryRow(query, postId)

	var result int
	err := row.Scan(&result)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Likes fetched")
	}

	return result
}

func getPostDislikes(postId int) int {
	query := "SELECT COUNT(*) FROM dislikes WHERE post_id=?"
	row := db.QueryRow(query, postId)

	var result int
	err := row.Scan(&result)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Dislikes fetched")
	}

	return result
}
