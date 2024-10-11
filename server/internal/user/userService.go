package user

import (
	"context"
	"time"

	"github.com/zedann/ecoforum/server/util"
)

type UserService struct {
	*UserRepository
}

func NewUserService(userRepo *UserRepository) *UserService {
	return &UserService{
		UserRepository: userRepo,
	}
}

func (s *UserService) CreateUser(ctx context.Context, req *CreateUserReq) (*CreateUserRes, error) {
	// context timeout for 2s
	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()
	// hash the password
	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}
	u := &User{
		Username: req.Username,
		Password: hashedPassword,
		Email:    req.Email,
	}
	user, err := s.UserRepository.CreateUser(ctx, u)
	if err != nil {
		return nil, err
	}

	return &CreateUserRes{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}, nil

}
