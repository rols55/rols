package controllers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"01.kood.tech/git/rols55/social-network/pkg/api/model"
	"01.kood.tech/git/rols55/social-network/pkg/api/route/middleware/acl"
	"01.kood.tech/git/rols55/social-network/pkg/logger"

	"github.com/gofrs/uuid"
)

var imgLocation = "media/posts"

type PostResponse struct {
	Status   string           `json:"status"`
	User     *model.User      `json:"user,omitempty"`
	Post     *model.Post      `json:"post,omitempty"`
	Comments []*model.Comment `json:"comments,omitempty"`
	Posts    []*model.Post    `json:"posts,omitempty"`
}

// Respond to GET request with specific post
func (h *BaseController) PostGET(w http.ResponseWriter, r *http.Request) {
	var err error
	var user *model.User

	//check if the user is logged in
	if err = h.getKeyVal(r, acl.UserKey, &user); err != nil {
		logger.Error(err)
		h.statusJSON(w, InternalError)
	}

	postId, err := strconv.ParseInt(path.Base(r.URL.Path), 10, 64)
	if err != nil {
		logger.Error(err)
		h.statusJSON(w, NotFound)
		return
	}

	//get the post
	post, err := model.GetPostById(h.db, postId)
	if err != nil || post == nil {
		logger.Error(err)
		h.statusJSON(w, NotFound)
		return
	}

	//get the post's comments
	comments, err := post.GetComments(h.db)
	if err != nil && err != model.ErrNotFound {
		logger.Error(err)
		h.statusJSON(w, InternalError)
		return
	}

	res := PostResponse{
		Status:   "OK",
		User:     user,
		Post:     post,
		Comments: comments,
	}

	h.writeJSON(w, res)
}

// Respond to GET request with all the none group posts
func (h *BaseController) PostsGET(w http.ResponseWriter, r *http.Request) {
	var err error
	var user *model.User

	//check if the user is logged in
	if err = h.getKeyVal(r, acl.UserKey, &user); err != nil {
		logger.Error(err)
		h.statusJSON(w, InternalError)
		return
	}

	//get the posts
	posts, err := model.GetPosts(h.db)
	if err != nil && err != model.ErrNotFound {
		logger.Error(err)
		h.statusJSON(w, NotFound)
		return
	}

	res := PostResponse{
		Status: "OK",
		User:   user,
		Posts:  posts,
	}

	h.writeJSON(w, res)
}

// Proccess the post creation form
func (h *BaseController) CreatePostPOST(w http.ResponseWriter, r *http.Request) {
	var err error
	var user *model.User

	//check if the user is logged in
	if err = h.getKeyVal(r, acl.UserKey, &user); err != nil {
		logger.Error(err)
		h.statusJSON(w, Unauthorized)
		return
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

	var groupId int64
	if r.FormValue("group_id") != "" {
		groupId, err = strconv.ParseInt(r.FormValue("group_id"), 10, 64)
		if err != nil {
			logger.Info("Failed to create event: Invalid group")
			errMsg := "Invalid group"
			h.statusJSON(w, InvalidInput(errMsg))
			return
		}

		group, err := model.GetGroupById(h.db, groupId)
		if err != nil || group == nil {
			logger.Info("Failed to create event: Invalid group")
			errMsg := "Invalid group"
			h.statusJSON(w, InvalidInput(errMsg))
			return
		}

		if !group.IsMember(h.db, user.Id) {
			logger.Info("Failed to create event: User is not a member of the group")
			errMsg := "Not a member of the group"
			h.statusJSON(w, InvalidInput(errMsg))
			return
		}
	}

	imgName := ""

	file, handler, err := r.FormFile("image")
	if err != nil {
		if err != http.ErrMissingFile {
			// Handle other errors if necessary
			logger.Error("Error retrieving form file:", err)
			http.Error(w, "Error retrieving form file", http.StatusBadRequest)
			return
		}
		// No file was uploaded, handle this case accordingly
		imgName = ""
	} else {
		defer file.Close()

		imgName = uuid.Must(uuid.NewV4()).String() + filepath.Ext(handler.Filename)

		err = os.MkdirAll(imgLocation, 0755)
		if err != nil {
			logger.Error("Error creating directory '%s': %v", imgLocation, err)
			http.Error(w, "Error creating directory", http.StatusInternalServerError)
			return
		}

		newFile, err := os.Create(imgLocation + "/" + imgName)
		if err != nil {
			logger.Error("Error creating new file '%s': %v", newFile, err)
			http.Error(w, "Error creating new file", http.StatusInternalServerError)
			return
		}
		defer newFile.Close()

		_, err = io.Copy(newFile, file)
		if err != nil {
			logger.Error("Error copying file data:", err)
			http.Error(w, "Error copying file data", http.StatusInternalServerError)
			return
		}
	}

	post := &model.Post{
		UserId:       user.UUID,
		Author:       user.Username,
		Title:        r.FormValue("title"),
		Text:         r.FormValue("text"),
		ImagePost:    imgName,
		Privacy:      r.FormValue("privacy"),
		Followers:    r.FormValue("followers"),
		CreationDate: time.Now(),
		GroupId:      groupId,
	}

	if post, err = post.Create(h.db); err != nil {
		logger.Error(err)
		h.statusJSON(w, InternalError)
		return
	}

	logger.Info(fmt.Sprintf("New post created - %v", post.Id))

	res := PostResponse{
		Status: "OK",
		User:   user,
		Post:   post,
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
		return
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
	if post.UserId != user.UUID {
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

	logger.Info("Post deleted successfully %v:%v", post.Id, post.Title)
	h.statusJSON(w, Redirect("/"))
}

type PostsResponse struct {
	Posts []*model.Post `json:"posts,omitempty"`
}

func (h *BaseController) FeedGET(w http.ResponseWriter, r *http.Request) {
	var posts []*model.Post

	var user *model.User
	var err error

	//check if the user is logged in
	if err = h.getKeyVal(r, acl.UserKey, &user); err != nil {
		logger.Error(err)
		h.statusJSON(w, InternalError)
	}

	//get the posts
	posts, err = model.GetPosts(h.db)
	if err != nil || posts == nil {
		logger.Error(err)
		h.statusJSON(w, NotFound)
		return
	}

	// posts filtered by following and privacy
	filteredbyFollowing := make([]*model.Post, 0)

	for _, post := range posts {
		author, err := model.GetUserByUsername(h.db, post.Author)
		if err != nil {
			log.Printf("Error fetching user by username %s: %v\n", post.Author, err)
		}
		if user.Username == author.Username {
			filteredbyFollowing = append(filteredbyFollowing, post)
		} else {
			authorFollowers, err := model.GetFollowers(h.db, author.UUID)
			if err != nil {
				log.Printf("Error fetching user by followers %s: %v\n", post.Author, err)
			}
			for _, follower := range authorFollowers {
				if follower.Username == user.Username {
					filteredbyFollowing = append(filteredbyFollowing, post)
					break
				}
			}
		}
	}

	//filtered by post privacy
	filteredPosts := make([]*model.Post, 0)

	for _, post := range filteredbyFollowing {
		if post.Privacy == "private" {
			if post.Author == user.Username {
				filteredPosts = append(filteredPosts, post)
			} else {
				followers := strings.Split(post.Followers, ",")
				for _, follower := range followers {
					if follower == user.Username {
						filteredPosts = append(filteredPosts, post)
					}
				}
			}
		} else {
			filteredPosts = append(filteredPosts, post)
		}
	}

	res := PostsResponse{
		Posts: filteredPosts,
	}
	h.writeJSON(w, res)
}

func (h *BaseController) GetImagePosts(w http.ResponseWriter, r *http.Request) {
	http.StripPrefix("/api/getimage/", http.FileServer(http.Dir(imgLocation))).ServeHTTP(w, r)
}
