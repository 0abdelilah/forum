package database

import (
	"fmt"
)

type Post struct {
	PostUUID     string
	Username     string
	CreationDate string
	Title        string
	Content      string
	Category     string
	Comments     []Comment
	Reacts       Reacts
}

var Posts []Post

func CreatePost(PostUUID, username, creationDate, title, content, category string) error {
	_, err := postStmt.Exec(PostUUID, username, creationDate, title, content, category)
	if err == nil {
		fmt.Println("Saved a post to db")
	}

	return err
}

func LoadPosts(category string) error {
	rows, err := db.Query(`
		SELECT PostUUID, username, creationDate, title, content, category
		FROM posts
		ORDER BY creationDate DESC
	`, category, category)
	if err != nil {
		return err
	}
	defer rows.Close()

	Posts = nil // clear before loading

	for rows.Next() {
		var p Post
		if err := rows.Scan(
			&p.PostUUID,
			&p.Username,
			&p.CreationDate,
			&p.Title,
			&p.Content,
			&p.Category,
		); err != nil {
			return err
		}

		// Get comments for this post
		cmnts, err := LoadComments(p.PostUUID)
		if err != nil {
			return err
		}
		p.Comments = cmnts

		// Get reactions for this post
		reacts, err := LoadReacts(p.PostUUID)
		if err != nil {
			return err
		}

		p.Reacts = reacts

		Posts = append(Posts, p)
	}

	return rows.Err()
}
