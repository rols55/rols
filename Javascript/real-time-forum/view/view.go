package view

import (
	"database/sql"
	"errors"
	"html/template"
	"net/http"

	"forum/model"
	"forum/shared/logger"
)

var (
	Files = []string{"templates/base.html",
		"templates/menu.html",
		"templates/categories.html",
		"templates/register.html",
		"templates/login.html",
		"templates/createpost.html",
		"templates/post.html",
		"templates/chat.html",
		"templates/chatUsers.html",
	} // Files that are always served.
	BasePath = "templates/"
)

type View struct {
	Files         []string
	Vars          map[string]interface{}
	Authenticated bool
	db            *sql.DB
}

// Creates a new template view
func New(db *sql.DB, path string) *View {
	view := &View{
		Files:         append(make([]string, 0), Files...),
		Vars:          make(map[string]interface{}),
		Authenticated: false,
		db:            db,
	}
	view.Files = append(view.Files, BasePath+path)
	view.Vars["Authenticated"] = false
	return view
}

// Executes the template and sends it to the client
func (v *View) Execute(w http.ResponseWriter) error {

	//Load the templates
	ts, err := template.New("").Funcs(template.FuncMap{
		"GetCategories": GetCategories(v.db),
	}).ParseFiles(v.Files...)

	if err != nil {
		return errors.New(logger.GetCurrentFuncName() + ": " + err.Error())
	}

	//Set authenticated for the template
	if v.Authenticated {
		v.Vars["Authenticated"] = true
	}

	//Execute the template with our custom data (Vars)
	return ts.ExecuteTemplate(w, "base", v.Vars)
}

func GetCategories(db *sql.DB) func() []*model.Category {
	return func() []*model.Category {
		allCategories, err := model.GetCategories(db)
		if err != nil {
			logger.Error(err)
		}
		return allCategories
	}
}
