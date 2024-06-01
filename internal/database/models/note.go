package models

import "time"

// const (
// 	public = "public"
// 	private = "private"
// )

type Note struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	HexID     string    `json:"hex_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Scope    string    `json:"scope"`
} 