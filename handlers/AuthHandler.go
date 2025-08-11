package handlers

import (
	"fmt"
	"forum/utils"
	"html/template"
	"net/http"
)

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	tmpt, err := template.ParseFiles("./templates/auth.html")
	if err != nil {
		fmt.Println(err)
		utils.ErrorHandler(w, "Internal Server Error", 500)
		return
	}

	tmpt.Execute(w, nil)
}
