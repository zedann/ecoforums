package user

import (
	"context"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/zedann/ecoforum/server/util"
)

type UserService struct {
	*UserRepository
	timeout time.Duration
}

func NewUserService(userRepo *UserRepository) *UserService {
	return &UserService{
		UserRepository: userRepo,
		timeout:        time.Duration(2) * time.Second,
	}
}

type MyJWTClaims struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func CreateToken(user *User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, MyJWTClaims{
		ID:       strconv.Itoa(int(user.ID)),
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    strconv.Itoa(int(user.ID)),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 10)),
		},
	})

	secretKey := os.Getenv("SECRET_KEY")
	return token.SignedString([]byte(secretKey))
}

func (s *UserService) CreateUser(ctx context.Context, req *CreateUserReq) (*CreateUserRes, error) {
	// context timeout for 2s
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
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

func (s *UserService) Login(ctx context.Context, req *LoginUserReq) (*LoginUserRes, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	u, err := s.UserRepository.GetUserByEmail(ctx, req.Email)

	if err != nil {
		return nil, err
	}

	//  create Token
	accessToken, err := CreateToken(u)
	if err != nil {
		return nil, err
	}
	user := &LoginUserRes{
		accessToken: accessToken,
		ID:          u.ID,
		Username:    u.Username,
	}

	return user, nil
}
