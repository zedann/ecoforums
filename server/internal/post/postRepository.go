package post

import (
	"context"
	"database/sql"

	"github.com/zedann/ecoforum/server/types"
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

func (r *PostRepository) GetPosts(ctx context.Context, reqConfig *types.ReqConfig) ([]*Post, error) {
	query := ` SELECT p.id , p.title , p.content , p.image , p.ups_number , p.downs_number , p.created_at , u.username
				FROM posts AS p
				INNER JOIN users as u
				ON p.user_id = u.id

				LIMIT $1 
				OFFSET $1
	`

	rows, err := r.db.QueryContext(ctx, query, reqConfig.Limit, reqConfig.Offset)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []*Post

	for rows.Next() {
		var post Post
		if err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.Image, &post.UpsNumber, &post.DownsNumber, &post.CreatedAt, &post.Username); err != nil {
			return nil, err
		}
		posts = append(posts, &post)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}
