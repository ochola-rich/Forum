package main

import (
	"database/sql"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"time"

	"forum/database" 
	"forum/handlers"

	"://github.com"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	// Initialize Database
	err := database.InitDB("./forum.db")
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer database.DB.Close()

	// Setup Router
	mux := http.NewServeMux()

	// Static Files
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./ui/static/"))))

	// Routes
	mux.HandleFunc("/", handlers.root)
	mux.HandleFunc("/registering", handlers.handleRegisterHtml)
	mux.HandleFunc("/register", handlers.register)
	mux.HandleFunc("/log", handlers.handleLoginPage)
	mux.HandleFunc("/login", handlers.login)
	mux.HandleFunc("/getusers", handlers.getUsers)
	mux.HandleFunc("/create-post", handlers.CreatePost)
	mux.HandleFunc("/create-post-page", handlers.HandleCreatePostPage)
	mux.HandleFunc("/like", handlers.HandleLike)
	mux.HandleFunc("/dislike", handlers.HandleDislike)

	// Start Server
	fmt.Println("Server running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
