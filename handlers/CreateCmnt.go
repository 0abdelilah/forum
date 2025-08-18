package handlers

import (
	"encoding/json"
	"fmt"
	"forum/auth"
	"forum/database"
	"net/http"
	"time"
)

// save a comment to database (comment, date, postuuid, username)
func CmntHandler(w http.ResponseWriter, r *http.Request) {
	sess, exist := auth.GetSession(r)
	if !exist {
		w.WriteHeader(403)
		json.NewEncoder(w).Encode(map[string]string{
			"success": "false",
			"error":   "Please log in",
		})
		return
	}

	var CommentData struct {
		PostUUID string
		Content  string
	}

	err := json.NewDecoder(r.Body).Decode(&CommentData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"success": "false",
			"error":   "Invalid JSON",
		})
		return
	}

	PostUUID := CommentData.PostUUID
	content := Sanitize(CommentData.Content)
	Username := sess.Username
	commentDate := time.Now().Format("2006-01-02T15:04:05.000000000Z07:00")

	if valid, msg := isValidComment(content); !valid {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"success": "false",
			"error":   msg,
		})
		return
	}

	err = database.CreateComment(PostUUID, Username, commentDate, content)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"success": "false",
			"error":   err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"success": "true",
	})

}

func isValidComment(Comment string) (bool, string) {
	if !(len(Comment) >= 1 && len(Comment) <= 100) {
		return false, "Comment must be between 1 and 100 characters"
	}
	return true, ""
}
