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

	// handle patterns
	mux.HandleFunc("GET /", handlers.HomeHandler)
	mux.HandleFunc("GET /logout", auth.LogoutHandler)
	mux.HandleFunc("GET /auth", handlers.AuthHandler)
	mux.HandleFunc("POST /register", handlers.RegisterHandler)
	mux.HandleFunc("POST /login", handlers.LoginHandler)

	// start server
	fmt.Println("started listening at http://localhost:8080/auth#register")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
