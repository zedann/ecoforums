package user

import (
	"context"
	"database/sql"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(ctx context.Context, user *User) (*User, error) {
	var lastInsertedID int64
	query := `INSERT INTO users
             (username , email , password , picture) 
             VALUES ($1,$2,$3,$4)
             returning id;
    `
	err := r.db.QueryRowContext(ctx, query, user.Username, user.Email, user.Password, "/public/imgs/avatar.jpeg").Scan(&lastInsertedID)
	if err != nil {
		return nil, err
	}

	user.ID = lastInsertedID
	return user, nil
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*User, error) {

	var user User
	query := ` SELECT id ,
					username ,
					email ,
					role 
			FROM users
			WHERE email=$1;
	`
	err := r.db.QueryRowContext(ctx, query, email).Scan(&user.ID, &user.Username, &user.Email, &user.Role)

	if err != nil {
		return nil, err
	}

	return &user, err
}
