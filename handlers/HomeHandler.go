package handlers

import (
	"fmt"
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

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		utils.ErrorHandler(w, "Internal server error", 500)
		return
	}

	tmpt, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		print("templates/index.html")
		fmt.Println("cant parse")
		utils.ErrorHandler(w, "Internal server error", 500)
		return
	}

	fmt.Println("serving", r.URL.Path)
	var user UserData
	user.Authenticated = auth.IsAuthenticated()
	user.Name = "Abdelilah"
	tmpt.Execute(w, user)
}
