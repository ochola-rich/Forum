package handlers

import (
	"forum/database"
	"net/http"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Get user ID from session 
	userID, err := getIDFromSession(r)
	if err != nil {
		http.Error(w, "You must be logged in", http.StatusUnauthorized)
		return
	}

	title := r.FormValue("title")
	content := r.FormValue("content")
	categories := r.Form["categories"] // This gets a slice of strings from checkboxes

	if title == "" || content == "" || len(categories) == 0 {
		http.Error(w, "Title, content, and at least one category required", http.StatusBadRequest)
		return
	}

	// Insert the Post
	res, err := database.DB.Exec("INSERT INTO posts (user_id, title, content) VALUES (?, ?, ?)", userID, title, content)
	if err != nil {
		http.Error(w, "Failed to create post", http.StatusInternalServerError)
		return
	}

	postID, _ := res.LastInsertId()

	// Link Categories (Post-Category association)
	for _, catName := range categories {
		var catID int
		// Find or Create the category ID
		err := database.DB.QueryRow("SELECT id FROM categories WHERE name = ?", catName).Scan(&catID)
		if err != nil {
			// If category doesn't exist, we skip or handle it
			continue 
		}
		database.DB.Exec("INSERT INTO post_categories (post_id, category_id) VALUES (?, ?)", postID, catID)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
