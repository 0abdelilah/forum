package auth

import (
	"forum/database"
	"forum/utils"
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
		utils.ToastError(w, "Username too long", "register")
		return
	}

	if !isValidEmail(email) {
		utils.ToastError(w, "Invalid email", "register")
		return
	}

	if !isValidPassword(password) {
		utils.ToastError(w, "Invalid password", "register")
		return
	}

	if password != re_password {
		utils.ToastError(w, "Passwords dont match", "register")
		return
	}

	// register user
	err := database.Register(username, email, password)
	if err != nil {
		utils.ErrorHandler(w, "Internal server error", 500)
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
