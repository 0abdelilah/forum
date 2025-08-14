package database

import (
	"database/sql"
	"fmt"
)

func Register(UserUUID, username, email, password string) error {
	_, err := usersStmt.Exec(UserUUID, username, email, password)
	return err
}

func Login(username, password string) (bool, error) {
	var id int
	err := db.QueryRow("SELECT id FROM users WHERE username = ? AND password = ?", username, password).Scan(&id)
	if err == sql.ErrNoRows {
		// utilisateur non trouvé
		return false, nil
	}
	if err != nil {
		// autre erreur SQL
		return false, err
	}
	return true, nil
}

func EmailExists(email string) bool {
	var exists bool
	err := db.QueryRow(
		"SELECT EXISTS(SELECT 1 FROM users WHERE email = ?)",
		email,
	).Scan(&exists)

	if err != nil {
		fmt.Println(err)
		return true
	}
	return exists
}
