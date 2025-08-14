package main

import (
	"fmt"
	"forum/auth"
	"forum/database"
	"forum/handlers"

	"log"
	"net/http"
)

func main() {
	// prepare database
	err := database.InitDB()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = database.CreateUsersTable()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = database.CreatePostsTable()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = database.CreateCommentsTable()
	if err != nil {
		fmt.Println(err)
		return
	}

	// start mux
	mux := http.NewServeMux()

	// handle authentication patterns
	mux.HandleFunc("GET /static/", handlers.StaticHandler)

	mux.HandleFunc("GET /", handlers.DashboardHandler)
	mux.HandleFunc("GET /auth", auth.AuthHandler)
	mux.HandleFunc("POST /register", auth.RegisterHandler)
	mux.HandleFunc("POST /login", auth.LoginHandler)
	mux.HandleFunc("GET /logout", auth.LogoutHandler)

	// authenticated patterns
	mux.HandleFunc("POST /create-post", handlers.CreatePHandler)
	mux.HandleFunc("GET /load-posts", handlers.LoadPostsHandler)
	mux.HandleFunc("POST /comment", handlers.CmntHandler)

	// start server
	fmt.Println("started listening at http://localhost:8081/auth#register")
	log.Fatal(http.ListenAndServe(":8081", mux))
}
