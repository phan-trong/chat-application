package domain

import (
	"context"
	"errors"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type UserRepository interface {
	FindOneByEmail(ctx context.Context, email string) (*User, error)
	FindOneById(ctx context.Context, id string) (*User, error)
	CreateUser(ctx context.Context, u *User) error
	UpdateUser(ctx context.Context, id string, u *User) error
	DeleteUser(ctx context.Context, id string) error
}
