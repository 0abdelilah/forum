package handlers

import (
	"fmt"
	"forum/database"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	Exist, err := database.Login(username, password)
	if err != nil {
		panic(err)
	}

	if !Exist {
		ToastError(w, "User not found", "#login")
		fmt.Println("User not found:", username, password)
		return
	}
	fmt.Println("User logged", username, password)
}