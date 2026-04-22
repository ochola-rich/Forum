package handlers

import (
	"forum/database"
	"forum/utils"
	"html/template"
	"net/http"
)

func Root(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// Get logged in user name
	var data utils.PageData
	userID, _ := getIDFromSession(r)
	if userID != 0 {
		database.DB.QueryRow("SELECT username FROM users WHERE id = ?", userID).Scan(&data.Username)
	} else {
		data.Username = "Guest"
	}

	// Fetch Posts
	query := `
		SELECT p.id, u.username, p.title, p.content, p.created_at
		FROM posts p
		JOIN users u ON p.user_id = u.id
		ORDER BY p.created_at DESC`
	
	rows, err := database.DB.Query(query)
	if err != nil {
		http.Error(w, "Failed to fetch posts", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var p utils.Post
		rows.Scan(&p.ID, &p.AuthorName, &p.Title, &p.Content, &p.CreatedAt)
		
		// Fetch Categories for each post
		catRows, _ := database.DB.Query(`
			SELECT c.name FROM categories c 
			JOIN post_categories pc ON c.id = pc.category_id 
			WHERE pc.post_id = ?`, p.ID)
		
		for catRows.Next() {
			var catName string
			catRows.Scan(&catName)
			p.Categories = append(p.Categories, catName)
		}
		catRows.Close()

		data.Posts = append(data.Posts, p)
	}

	// Render Template
	tmpl, _ := template.ParseFiles("ui/templates/home.html")
	tmpl.Execute(w, data)
}
