package post

import (
	"context"
	"time"
)

type PostService struct {
	*PostRepository
	Timeout time.Duration
}

func NewPostService(postRepo *PostRepository) *PostService {
	return &PostService{
		PostRepository: postRepo,
	}
}

func (s *PostService) CreatePost(ctx context.Context, req *CreatePostReq) (*CreatePostRes, error) {
	ctx, cancel := context.WithTimeout(ctx, s.Timeout)
	defer cancel()

	p := &Post{
		Title:   req.Title,
		Content: req.Content,
		Image:   req.Image,
		UserID:  req.UserID,
	}

	post, err := s.PostRepository.CreatePost(ctx, p)
	if err != nil {
		return nil, err
	}

	return &CreatePostRes{
		ID:        post.ID,
		Title:     post.Title,
		Image:     post.Image,
		Content:   post.Content,
		UserID:    post.UserID,
		CreatedAt: post.CreatedAt,
	}, nil

}
