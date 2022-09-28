package usecase

import (
	"chat_application/internal/domain"
	"context"
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
	_userRepo domain.UserRepository
}

func (lg *loginUseCase) Handle(ctx context.Context, req *LoginRequest) (string, error) {
	_, err := lg._userRepo.FindOneByEmail(ctx, req.Email)
	if err != nil {
		return "", err
	}

	return "abccccc", nil
}

func NewLoginUseCase(userRepo domain.UserRepository) *loginUseCase {
	return &loginUseCase{
		_userRepo: userRepo,
	}
}
