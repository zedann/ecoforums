package post

import (
	"time"
)

type Post struct {
	ID          int64     `json:"id" db:"id"`
	Title       string    `json:"title" db:"title"`
	Content     string    `json:"content" db:"content"`
	Image       string    `json:"image" db:"image"`
	UpsNumber   int       `json:"ups_number" db:"ups_number"`
	DownsNumber int       `json:"downs_number" db:"downs_number"`
	UserID      int64     `json:"user_id" db:"user_id"`
	Username    string    `json:"username" db:"username"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

type CreatePostReq struct {
	Title   string `json:"title" db:"title"`
	Content string `json:"content" db:"content"`
	Image   string `json:"image" db:"image"`
	UserID  int64  `json:"user_id" db:"user_id"`
}
type CreatePostRes struct {
	ID        int64     `json:"id" db:"id"`
	Title     string    `json:"title" db:"title"`
	Content   string    `json:"content" db:"content"`
	Image     string    `json:"image" db:"image"`
	UserID    int64     `json:"user_id" db:"user_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
