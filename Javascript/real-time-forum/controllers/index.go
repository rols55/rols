package controllers

import (
	"net/http"
	"strconv"

	"forum/model"
	"forum/route/middleware/acl"
	"forum/shared/logger"
	"forum/view"
)

// Home/Index page handler/controller
func (h *BaseController) Index(w http.ResponseWriter, r *http.Request) {

	var err error
	var posts []*model.Post
	var postsWithReactionsAndCommentQty []*model.PostWithReactionsAndCommentQty
	var categories []*model.Category
	var user *model.User
	var userID int64

	if r.URL.Path != "/" {
		http.Error(w, "Not Found 404", http.StatusNotFound)
		return
	}

	//view to be rendered
	view := view.New(h.db, "index.html")

	// Is the user logged in
	if err = h.getKeyVal(r, acl.UserKey, &user); err == nil {
		userID = user.Id
		view.Authenticated = true
		view.Vars["User"] = user
	} else if err != model.ErrNotFound {
		logger.Error(err)
		h.statusJSON(w, InternalError)
	}

	//get all categories that have posts
	categories, err = model.GetCategoriesWithPosts(h.db)
	if err != nil && err != model.ErrNotFound {
		logger.Error(err)
	}
	view.Vars["Categories"] = categories

	//filtering by category
	selectedCategory := r.URL.Query().Get("category")
	if len(selectedCategory) == 0 {
		// get all the posts from the database, throw an http.error if there is error and its not ErrNotFound (no posts)
		if posts, err = model.GetPosts(h.db); err != nil && err != model.ErrNotFound {
			logger.Error(err)
			http.NotFound(w, r) // ???
		}
		// get reactions and amount of comment for posts
		if postsWithReactionsAndCommentQty, err = model.GetPostsWithReactionsAndCommentQty(h.db, posts, userID); err != nil && err != model.ErrNotFound {
			logger.Error(err)
			http.NotFound(w, r) // ???
		}
		view.Vars["CurrentCategory"] = "All"
	} else {
		selectedCategoryId, err := strconv.Atoi(selectedCategory)
		if err != nil {
			logger.Error(err)
		}
		//get id of posts from the relationship postscategories table
		postsId, err := model.GetPostIdByCategoryId(h.db, selectedCategoryId)
		if err != nil && err != model.ErrNotFound {
			logger.Error(err)
			http.NotFound(w, r) // ???
			return
		}
		if len(postsId) > 0 {
			for i := 0; i < len(postsId); i++ {
				//get every post by id from post table
				post, err := model.GetPostById(h.db, int64(postsId[i]))
				if err != nil || post == nil {
					logger.Error(err)
					http.NotFound(w, r) // ???
					return
				}
				posts = append(posts, post)
			}
			// get reactions for posts
			if postsWithReactionsAndCommentQty, err = model.GetPostsWithReactionsAndCommentQty(h.db, posts, userID); err != nil && err != model.ErrNotFound {
				logger.Error(err)
				http.Error(w, "Not Found 404", http.StatusNotFound) // ???
			}
		}
		//Get chosen category name for template
		var currentCatId *model.Category
		if currentCatId, err = model.GetCategoryById(h.db, int64(selectedCategoryId)); err != nil {
			logger.Error(err)
			http.Error(w, "Internal Error", http.StatusInternalServerError) // ???
		}
		view.Vars["CurrentCategory"] = currentCatId.Category
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
			} else if i > 460 {
				post.Text = post.Text[0:i] + "..."
				break
			}
		}
	}

	// set the data to the view
	view.Vars["Posts"] = postsWithReactionsAndCommentQty

	// render the view
	if err = view.Execute(w); err != nil {
		logger.Error(err)
		http.Error(w, "Internal Error", http.StatusInternalServerError)
	}
}
