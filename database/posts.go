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
}

var Posts []Post

type Comment struct {
	Username string
	Date     string
	Content  string
}

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
		WHERE (? = '' OR category = ?)
		ORDER BY creationDate DESC
	`, category, category)
	if err != nil {
		return err
	}
	defer rows.Close()
	fmt.Println("tt")

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

		Posts = append(Posts, p)
		fmt.Println(p.Category)
	}

	return rows.Err()
}

func CreateComment(username, date, content, PostUUID string) error {
	_, err := db.Exec(`
        INSERT INTO comments (postCmntId, username, date, content)
        VALUES ((SELECT id FROM posts WHERE PostUUID = ?), ?, ?, ?)`,
		PostUUID, username, date, content,
	)
	return err
}

func LoadComments(PostUUID string) ([]Comment, error) {
	rows, err := db.Query(`
        SELECT comments.username, comments.date, comments.content
		FROM comments
		JOIN posts ON comments.postCmntId = posts.id
		WHERE posts.PostUUID = ?
		`,
		PostUUID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []Comment
	for rows.Next() {
		var cmnt Comment

		if err := rows.Scan(&cmnt.Username, &cmnt.Date, &cmnt.Content); err != nil {
			return nil, err
		}
		comments = append(comments, cmnt)
	}
	return comments, nil
}
