package handlers

import (
	"encoding/json"
	"fmt"
	"forum/auth"
	"forum/database"
	"net/http"
)

func ReactHandler(w http.ResponseWriter, r *http.Request) {
	sess, exist := auth.GetSession(r)
	if !exist {
		w.WriteHeader(403)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   "Please log in",
			"success": "false",
		})
		return
	}

	var ReactData struct {
		PostUUID string
		Reaction string
	}

	err := json.NewDecoder(r.Body).Decode(&ReactData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"success": "false",
			"error":   "Invalid json",
		})
		return
	}

	if ReactData.Reaction != "like" && ReactData.Reaction != "dislike" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"success": "false",
			"error":   "Invalid reaction",
		})
		return
	}

	err = database.SaveReact(ReactData.PostUUID, sess.Username, ReactData.Reaction)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"success": "false",
			"error":   "Failed to save react",
		})
		fmt.Println(err)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"success": "true",
	})
}
