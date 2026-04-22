package handlers

import (
	"fmt"
	"forum/database"
	"net/http"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query("SELECT username, email FROM users")
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var username, email string
		if err := rows.Scan(&username, &email); err == nil {
			fmt.Fprintf(w, "User: %s | Email: %s\n", username, email)
		}
	}
}
