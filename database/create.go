package database

import (
	"database/sql"
	"errors"

	_ "github.com/mattn/go-sqlite3"
)

var stmt *sql.Stmt
var db *sql.DB

func CreateDB() (*sql.Stmt, error) {
	var err error
	db, err = sql.Open("sqlite3", "./sqlite.db")
	if err != nil {
		return nil, err
	}
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, username TEXT, email TEXT, password TEXT)")
	if err != nil {
		return nil, err
	}
	stmt, err = db.Prepare("INSERT INTO users(username, email, password) VALUES (?, ?, ?)")
	return stmt, err
}

func Register(username, email, password string) error {
	if stmt == nil {
		return errors.New("insert statement not prepared")
	}

	_, err := stmt.Exec(username, email, password)
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
