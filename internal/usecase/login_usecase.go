package usecase

import (
	"chat_application/internal/adapter/middlewares"
	"chat_application/internal/adapter/services"
	"chat_application/internal/domain"
	"context"

	"golang.org/x/crypto/bcrypt"
)

type LoginUseCase interface {
	Handle(ctx context.Context, req *LoginRequest) (string, error)
}

type LoginRequest struct {
	Email    string
	Password string
}

var _ LoginUseCase = &loginUseCase{}

type loginUseCase struct {
	userRepo domain.UserRepository
}

func (lg *loginUseCase) Handle(ctx context.Context, req *LoginRequest) (string, error) {
	user, err := lg.userRepo.FindOneByEmail(ctx, req.Email)
	if err != nil {
		return "", err
	}

	err = middlewares.VerifyPassword(req.Password, user.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := services.CreateJWTToken(user)

	if err != nil {
		return "", err
	}

	return token, nil
}

func NewLoginUseCase(_userRepo domain.UserRepository) *loginUseCase {
	return &loginUseCase{
		userRepo: _userRepo,
	}
}
