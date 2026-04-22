package handlers

import (
	"database/sql"
	"forum/database"
	"net/http"
	"time"

	"://github.com"
	"golang.org/x/crypto/bcrypt"
)

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	username := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	query := `INSERT INTO users (username, password_hash, email) VALUES (?, ?, ?)`
	_, err := database.DB.Exec(query, username, string(hashedPassword), email)
	if err != nil {
		http.Error(w, "User already exists", http.StatusConflict)
		return
	}

	http.Redirect(w, r, "/log", http.StatusSeeOther)
}

func Login(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	var userID int
	var dbHash string
	err := database.DB.QueryRow("SELECT id, password_hash FROM users WHERE email = ?", email).Scan(&userID, &dbHash)

	if err == sql.ErrNoRows || bcrypt.CompareHashAndPassword([]byte(dbHash), []byte(password)) != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	sessionID, _ := uuid.NewV4()
	expiresAt := time.Now().Add(24 * time.Hour)

	database.DB.Exec("INSERT INTO sessions (id, user_id, expires_at) VALUES (?, ?, ?)", sessionID.String(), userID, expiresAt)

	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    sessionID.String(),
		Expires:  expiresAt,
		HttpOnly: true,
		Path:     "/",
	})

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func getIDFromSession(r *http.Request) (int, error) {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		return 0, err
	}

	var userID int
	err = database.DB.QueryRow("SELECT user_id FROM sessions WHERE id = ? AND expires_at > datetime('now')", cookie.Value).Scan(&userID)
	return userID, err
}
