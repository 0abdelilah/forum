package database

type Comment struct {
	Username string
	Date     string
	Content  string
}

func CreateComment(PostUUID, username, commentDate, content string) error {
	_, err := db.Exec(`
        INSERT INTO comments (postCmntId, username, date, content)
        VALUES ((SELECT id FROM posts WHERE PostUUID = ?), ?, ?, ?)`,
		PostUUID, username, commentDate, content,
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
