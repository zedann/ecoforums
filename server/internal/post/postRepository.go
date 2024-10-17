package post

import (
	"context"
	"database/sql"
)

type PostRepository struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) *PostRepository {
	return &PostRepository{
		db: db,
	}
}

func (r *PostRepository) CreatePost(ctx context.Context, post *Post) (*Post, error) {
	var lastInsertedID int64
	query := `INSERT INTO posts (title , content , image , user_id) VALUES ($1,$2,$3,$4) returning id;`
	err := r.db.QueryRowContext(ctx, query, post.Title, post.Content, post.Image, post.UserID).Scan(&lastInsertedID)
	if err != nil {
		return nil, err
	}
	post.ID = lastInsertedID
	return post, nil
}
