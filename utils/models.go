package utils

import "time"

type Post struct {
	ID         int
	AuthorName string
	Title      string
	Content    string
	CreatedAt  time.Time
	Categories []string // Multiple categories per post
	Likes      int
	Dislikes   int
}

type PageData struct {
	Username string
	Posts    []Post
}
