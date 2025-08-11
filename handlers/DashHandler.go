package handlers

import (
	"forum/auth"
	"forum/utils"
	"html/template"
	"net/http"
)

type UserData struct {
	Name          string
	Authenticated bool
}

type Post struct {
	Img          string
	Title        string
	Desctiption  string
	LikeCount    string
	Comments     map[string]string // k: user; v: comment
	CommentCount int
}

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		utils.ErrorHandler(w, "This page does not exist.", 404)
		return
	}

	tmpt, err := template.ParseFiles("./templates/dashboard.html")
	if err != nil {
		utils.ErrorHandler(w, "Internal server error", 500)
		return
	}

	var user UserData

	Sess, Auth := auth.GetSession(r)

	user.Authenticated = Auth
	user.Name = Sess.Username
	tmpt.Execute(w, user)
}

func StaticHandler(w http.ResponseWriter, r *http.Request) {
	http.StripPrefix("/static/",
		http.FileServer(http.Dir("./templates/static")),
	).ServeHTTP(w, r)
}
