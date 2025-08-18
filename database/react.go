package database

import "database/sql"

type Reacts struct {
	Likes    ReactLike
	Dislikes ReactDislike
	PostUUID string
}

type ReactLike struct {
	Amount    int
	Usernames []string
}

type ReactDislike struct {
	Amount    int
	Usernames []string
}

func SaveReact(PostUUID, username, reactionType string) error {
	// check if user already reacted
	var existing string
	err := db.QueryRow(`SELECT reactType FROM reacts WHERE PostUUID = ? AND username = ?`, PostUUID, username).Scan(&existing)

	if err == sql.ErrNoRows {
		// no existing reaction → just insert
		_, err = db.Exec(`
            INSERT INTO reacts (PostUUID, username, reactType)
            VALUES (?, ?, ?)
        `, PostUUID, username, reactionType)
		return err
	}
	if err != nil {
		return err
	}

	// if same reaction → remove it
	if existing == reactionType {
		_, err = db.Exec(`DELETE FROM reacts WHERE PostUUID = ? AND username = ?`, PostUUID, username)
		return err
	}

	// otherwise update reaction
	_, err = db.Exec(`UPDATE reacts SET reactType = ? WHERE PostUUID = ? AND username = ?`, reactionType, PostUUID, username)
	return err
}

func LoadReacts(PostUUID string) (Reacts, error) {
	rows, err := db.Query(`
		SELECT username, reactType
		FROM reacts
		WHERE PostUUID = ?
	`, PostUUID)
	if err != nil {
		return Reacts{}, err
	}
	defer rows.Close()

	var reacts Reacts
	reacts.PostUUID = PostUUID

	for rows.Next() {
		var username, reactType string
		if err := rows.Scan(&username, &reactType); err != nil {
			return Reacts{}, err
		}

		switch reactType {
		case "like":
			reacts.Likes.Amount++
			reacts.Likes.Usernames = append(reacts.Likes.Usernames, username)
		case "dislike":
			reacts.Dislikes.Amount++
			reacts.Dislikes.Usernames = append(reacts.Dislikes.Usernames, username)
		}
	}

	return reacts, rows.Err()
}
