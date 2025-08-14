package handlers

import (
	"encoding/json"
	"fmt"
	"forum/auth"
	"forum/database"
	"html"
	"net/http"
	"time"

	"github.com/gofrs/uuid"
)

func CreatePHandler(w http.ResponseWriter, r *http.Request) {
	sess, exist := auth.GetSession(r)
	if !exist {
		w.WriteHeader(403)
		json.NewEncoder(w).Encode(map[string]string{
			"success": "false",
			"error":   "Plesae log in",
		})
		return
	}

	var postData struct {
		Title    string `json:"title"`
		Content  string `json:"content"`
		Category string `json:"category"`
	}

	err := json.NewDecoder(r.Body).Decode(&postData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"success": "false",
			"error":   "Invalid JSON",
		})
		return
	}

	uuidV4, err := uuid.NewV4()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"success": "false",
			"error":   "Internal server error",
		})
		fmt.Println(err)
		return
	}
	PostUUID := "post_" + uuidV4.String()

	title := Sanitize(postData.Title)
	content := Sanitize(postData.Content)
	category := postData.Category
	username := sess.Username
	creationDate := time.Now().Format("2006-01-02T15:04:05.000000000Z07:00")

	if valid, msg := isValidPost(title, content, category); !valid {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"success": "false",
			"error":   msg,
		})
		return
	}

	err = database.CreatePost(PostUUID, username, creationDate, category, title, content)
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

func isValidPost(title, content, category string) (bool, string) {
	if category != "" && category != "programming" && category != "music" && category != "gaming" {
		return false, "Invalid category"
	}

	if !(len(title) >= 1 && len(title) <= 300) {
		return false, "Title must be between 1 and 300 characters"
	}

	if !(len(content) >= 1 && len(content) <= 300) {
		return false, "Content must be between 1 and 300 characters"
	}
	return true, ""
}

func Sanitize(text string) string {
	return html.EscapeString(text)
}
