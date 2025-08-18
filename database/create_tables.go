package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB
var usersStmt *sql.Stmt
var postStmt *sql.Stmt

// create & prepare tables

func InitDB() error {
	var err error
	db, err = sql.Open("sqlite3", "./sqlite.db")
	if err != nil {
		return err
	}
	return nil
}

func CreateUsersTable() error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		UserUUID TEXT NOT NULL UNIQUE,
		username TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE,
		password TEXT
	)`)
	if err != nil {
		return err
	}
	usersStmt, err = db.Prepare("INSERT INTO users(UserUUID, username, email, password) VALUES (?, ?, ?, ?)")
	return err
}

func CreatePostsTable() error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS posts (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			PostUUID TEXT NOT NULL UNIQUE,
			username TEXT NOT NULL,
			creationDate TEXT NOT NULL,
			title TEXT NOT NULL,
			content TEXT NOT NULL,
			category TEXT NOT NULL
		)
	`)
	if err != nil {
		return err
	}

	postStmt, err = db.Prepare(`
		INSERT INTO posts(PostUUID, username, creationDate, title, content, category)
		VALUES (?, ?, ?, ?, ?, ?)
	`)
	return err
}

func CreateCommentsTable() error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS comments (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        postCmntId INTEGER NOT NULL,
        username TEXT NOT NULL,
        date TEXT NOT NULL,
        content TEXT NOT NULL,
        FOREIGN KEY (postCmntId) REFERENCES posts(id) ON DELETE CASCADE
    )`)

	return err
}

func CreateReactsTable() error {
	_, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS reacts (
		PostUUID TEXT NOT NULL,
		username TEXT NOT NULL,
		reactType TEXT NOT NULL,
		UNIQUE(PostUUID, username)
	)
    `)
	return err
}
