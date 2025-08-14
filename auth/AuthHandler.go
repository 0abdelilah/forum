package auth

import (
	"html/template"
	"net/http"
)

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	tmpt, err := template.ParseFiles("./templates/auth.html")
	if err != nil {
		http.Error(w, "Internal server error", 500)
		return
	}

	tmpt.Execute(w, nil)
}
