package model

import (
	"database/sql"
	"errors"

	"forum/shared/logger"
)

type Category struct {
	Id       int64  `json:"id,omitempty"`
	Category string `json:"category,omitempty"`
}

func (c *Category) Create(db *sql.DB) (*Category, error) {
	result, err := db.Exec("INSERT INTO categories(category) values( ? )", c.Category)
	if err != nil {
		return nil, errors.New(logger.GetCurrentFuncName() + " " + err.Error())
	}
	if c.Id, err = result.LastInsertId(); err != nil {
		return nil, errors.New(logger.GetCurrentFuncName() + " " + err.Error())
	}
	return c, nil
}

func GetCategories(db *sql.DB) ([]*Category, error) {
	rows, err := db.Query("SELECT id, category FROM categories")
	if err != nil {
		return nil, errors.New(logger.GetCurrentFuncName() + " " + err.Error())
	}
	categories := make([]*Category, 0)
	for rows.Next() {
		var category Category
		err = rows.Scan(&category.Id, &category.Category)
		if err != nil {
			return nil, errors.New(logger.GetCurrentFuncName() + " " + err.Error())
		}
		categories = append(categories, &category)
	}
	if len(categories) == 0 {
		return nil, ErrNotFound
	}
	return categories, nil
}

func GetCategoryById(db *sql.DB, Id int64) (*Category, error) {
	row := db.QueryRow("SELECT id, category FROM categories WHERE id =?", Id)
	category := &Category{}
	if err := row.Scan(&category.Id, &category.Category); err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, errors.New(logger.GetCurrentFuncName() + " " + err.Error())
	}
	return category, nil
}

//some functions form category controller (if we need it) to be able create new category
/*func CreateCategory(w http.ResponseWriter, r *http.Request) {
	logger.Info("Request URL: %v method: %v", r.URL, r.Method)

	view := view.New(".html")

	if r.Method == "POST" {
		CreateMyCategoryPOST(w, r, view)
		return
	}
	CreateCategoryGET(w, r, view)
}

func CreateCategoryGET(w http.ResponseWriter, r *http.Request, view *view.View) {
	if err := view.Execute(w); err != nil {
		logger.Error(err)
		http.Error(w, "Internal Error", http.StatusInternalServerError)
	}
}

func CreateCategoryPOST(w http.ResponseWriter, r *http.Request, view *view.View) {
	var err error

	if r.FormValue("category") == "" {
		logger.Info("No category provided")
		CreatePostGET(w, r, view)
		return
	}

	category := &model.Category{
		Category: r.FormValue("category"),
	}

	if category, err = category.Create(); err != nil {
		logger.Error(err)
		http.Error(w, "Internal Error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/category=%v", category.Id), http.StatusFound)
}
*/
