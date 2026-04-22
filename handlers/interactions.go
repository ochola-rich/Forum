package handlers

import (
	"forum/database"
	"net/http"
	"strconv"
)

func HandleLike(w http.ResponseWriter, r *http.Request) {
	vote(w, r, 1)
}

func HandleDislike(w http.ResponseWriter, r *http.Request) {
	vote(w, r, -1)
}

func vote(w http.ResponseWriter, r *http.Request, value int) {
	userID, err := GetIDFromSession(r)
	if err != nil {
		http.Redirect(w, r, "/log", http.StatusSeeOther)
		return
	}

	postIDStr := r.URL.Query().Get("id")
	postID, _ := strconv.Atoi(postIDStr)

	// "INSERT OR REPLACE" logic: 
	// If the user already voted on this post, SQLite will update the existing row.
	query := `
		INSERT INTO interactions (user_id, post_id, is_like) 
		VALUES (?, ?, ?)
		ON CONFLICT(user_id, post_id) DO UPDATE SET is_like = excluded.is_like`

	_, err = database.DB.Exec(query, userID, postID, value)
	if err != nil {
		http.Error(w, "Failed to record vote", http.StatusInternalServerError)
		return
	}

	// Redirect back to the page the user was on
	http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
}
