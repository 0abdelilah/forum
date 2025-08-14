package auth

import (
	"encoding/json"
	"fmt"
	"forum/database"
	"net/http"
	"regexp"
	"strings"

	"github.com/gofrs/uuid"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {

	var creds struct {
		User  string `json:"username"`
		Email string `json:"email"`
		Pass  string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"success": "false",
			"error":   "Invalid JSON",
		})
		return
	}

	username := creds.User
	email := creds.Email
	password := creds.Pass

	// verify feilds //

	if msg, ok := isValidUser(username); !ok {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"success": "false",
			"error":   msg,
		})
		return
	}

	if msg, ok := isValidEmail(email); !ok {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"success": "false",
			"error":   msg,
		})
		return
	}

	if msg, ok := isValidPassword(password); !ok {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"success": "false",
			"error":   msg,
		})
		return
	}

	uuidV4, err := uuid.NewV4()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"success": "false",
			"error":   "Internal server error",
		})
		return
	}
	UserUUID := "user_" + uuidV4.String()

	// register user
	err = database.Register(UserUUID, username, email, password)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"success": "false",
			"error":   "Internal server error",
		})
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"success": "true",
	})
}

func isValidUser(username string) (string, bool) {
	if strings.ContainsAny(username, `'"<>\/;:&%$#@!(){}[]|=+?.~^,*`) {
		return "Username must Only have letters, numbers, _ and -", false
	}

	// check lenght
	if !(len(username) >= 4 && len(username) <= 20) {
		return "Username must be between 4 and 20 characters", false
	}
	return "", true
}

func isValidEmail(email string) (string, bool) {
	// chcek if already on db
	if database.EmailExists(email) {
		return "Email already used", false
	}

	// check regex
	emailRegex := `^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`
	re := regexp.MustCompile(emailRegex)

	if !re.MatchString(email) {
		return "Invalid email format", false
	}

	return "", true
}

func isValidPassword(password string) (string, bool) {
	if !(len(password) >= 4 && len(password) <= 20) {
		return "Password must be between 4 and 20 characters", false
	}
	return "", true
}
