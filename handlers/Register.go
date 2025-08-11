package handlers

import (
	"fmt"
	"forum/database"
	"html/template"
	"net/http"
	"regexp"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")
	re_password := r.FormValue("re-password")

	// verify feilds //

	if !isValidUser(username) {
		ToastError(w, "Username too long", "medium")
		return
	}

	if !isValidEmail(email) {
		ToastError(w, "Invalid email", "medium")
		return
	}

	if !isValidPassword(password) {
		ToastError(w, "Invalid password", "medium")
		return
	}

	if password != re_password {
		ToastError(w, "Passwords dont match", "register")
		return
	}

	// register user
	err := database.Register(username, email, password)
	if err != nil {
		fmt.Println("Error registering user")
		return
	}
}

func isValidEmail(email string) bool {
	// todo: chcek if already on db

	// check regex
	emailRegex := `^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}

func isValidUser(username string) bool {
	// todo: chcek if already on db

	/*
		// check regex
			userRegex := ``
			re := regexp.MustCompile(username)
			return re.MatchString(userRegex)
	*/

	// check lenght
	return len(username) > 4 && len(username) < 20
}

func isValidPassword(password string) bool {
	/*
		passRegex := `^(?=.*\d)(?=.*[a-z])(?=.*[A-Z])(?=.*[a-zA-Z]).{8,}$`
		re := regexp.MustCompile(password)
		return re.MatchString(passRegex)
	*/
	return len(password) > 4 && len(password) < 20
}

func ToastError(w http.ResponseWriter, text string, hash string) {
	tmpt, err := template.ParseFiles("./templates/auth.html")
	if err != nil {
		http.Error(w, "Template parsing error", http.StatusInternalServerError)
		return
	}

	err = tmpt.Execute(w, struct {
		Text string
		Hash string
	}{
		Text: text,
		Hash: hash,
	})
	if err != nil {
		http.Error(w, "Template execution error", http.StatusInternalServerError)
		return
	}
}
