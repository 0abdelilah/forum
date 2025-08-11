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
	database.CreateDB()

	// start mux
	mux := http.NewServeMux()

	// handle authentication patterns
	mux.HandleFunc("GET /static/", handlers.StaticHandler)

	mux.HandleFunc("GET /", handlers.DashboardHandler)
	mux.HandleFunc("GET /auth", auth.AuthHandler)
	mux.HandleFunc("POST /register", auth.RegisterHandler)
	mux.HandleFunc("POST /login", auth.LoginHandler)
	mux.HandleFunc("GET /logout", auth.LogoutHandler)

	// start server
	fmt.Println("started listening at http://localhost:8080/auth#register")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
