package models

import "time"

type Post struct {
	ID        int       `json:"id" db:"id"`
	UserID    int       `json:"user_id" db:"user_id"`
	Title     string    `json:"title" db:"title"`
	Content   string    `json:"content" db:"content"`
	Category  string    `json:"category" db:"category"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	Author    string    `json:"author"`
}

// PostInput
type PostInput struct {
	Title    string `json:"title"`
	Content  string `json:"content"`
	Category string `json:"category"`
}
