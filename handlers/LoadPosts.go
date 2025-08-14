package handlers

import (
	"encoding/json"
	"fmt"
	"forum/database"
	"net/http"
)

func LoadPostsHandler(w http.ResponseWriter, r *http.Request) {

	category := r.URL.Query().Get("category")

	if category != "" && category != "programming" && category != "music" && category != "gaming" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"success": "false",
			"error":   "Invalid category",
		})
		return
	}


	err := database.LoadPosts(category)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"success": "false",
			"error":   "Error loading database",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": "true",
		"posts":   database.Posts,
	})
}
