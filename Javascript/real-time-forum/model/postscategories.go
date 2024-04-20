package model

import (
	"database/sql"
	"errors"

	"forum/shared/logger"
)

type PostsCategories struct {
	Id         int64 `json:"id,omitempty"`
	PostId     int64 `json:"post_id,omitempty"`
	CategoryId int64 `json:"category_id,omitempty"`
}

func (pc *PostsCategories) Create(db *sql.DB) (*PostsCategories, error) {
	result, err := db.Exec("INSERT INTO postscategories(post_id, category_id) values( ?, ? )", pc.PostId, pc.CategoryId)
	if err != nil {
		return nil, errors.New(logger.GetCurrentFuncName() + " " + err.Error())
	}
	if pc.Id, err = result.LastInsertId(); err != nil {
		return nil, errors.New(logger.GetCurrentFuncName() + " " + err.Error())
	}
	return pc, nil
}

func GetCategoryIdByPostId(db *sql.DB, Id int64) ([]int64, error) {
	rows, err := db.Query("SELECT category_id FROM postscategories WHERE post_id =?", Id)
	if err != nil {
		return nil, errors.New(logger.GetCurrentFuncName() + " " + err.Error())
	}
	categoriesId := make([]int64, 0)
	for rows.Next() {
		var categoryId PostsCategories
		err = rows.Scan(&categoryId.CategoryId)
		if err != nil {
			return nil, errors.New(logger.GetCurrentFuncName() + " " + err.Error())
		}
		categoriesId = append(categoriesId, categoryId.CategoryId)
	}
	if len(categoriesId) == 0 {
		return nil, ErrNotFound
	}
	return categoriesId, nil
}

func GetPostIdByCategoryId(db *sql.DB, Id int) ([]int, error) {
	rows, err := db.Query("SELECT post_id FROM postscategories WHERE category_id =?", Id)
	if err != nil {
		return nil, errors.New(logger.GetCurrentFuncName() + " " + err.Error())
	}
	postsId := make([]int, 0)
	for rows.Next() {
		var postId PostsCategories
		err = rows.Scan(&postId.PostId)
		if err != nil {
			return nil, errors.New(logger.GetCurrentFuncName() + " " + err.Error())
		}
		postsId = append(postsId, int(postId.PostId))

	}
	if len(postsId) == 0 {
		return nil, ErrNotFound
	}

	return postsId, nil
}

// Returns categories that have posts associated
func GetCategoriesWithPosts(db *sql.DB) ([]*Category, error) {
	rows, err := db.Query("SELECT DISTINCT category_id, categories.category FROM postscategories INNER JOIN categories ON postscategories.category_id=categories.id")
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
