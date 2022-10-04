package usecase

import (
	"chat_application/internal/domain"
	"context"
	"errors"
	"time"
)

type SignUpUseCase interface {
	Handle(ctx context.Context, req *SignUpRequest) error
}

type SignUpRequest struct {
	FullName string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type signupUsecase struct {
	userRepo domain.UserRepository
}

func (su *signupUsecase) Handle(ctx context.Context, req *SignUpRequest) error {
	_, isExists := su.userRepo.FindOneByEmail(ctx, req.Email)

	if isExists != nil {
		user := &domain.User{
			FullName:  req.FullName,
			Email:     req.Email,
			Password:  req.Password,
			Avatar:    "https://source.unsplash.com/random",
			CreateAt:  time.Now(),
			UpdatedAt: time.Now(),
		}
		err := user.HashPassword()
		if err != nil {
			return err
		}

		err = su.userRepo.CreateUser(ctx, user)
		if err != nil {
			return err
		}
		return nil
	} else {
		return errors.New("Email Already Exists")
	}

}

func NewSignUpUseCase(_userRepo domain.UserRepository) *signupUsecase {
	return &signupUsecase{
		userRepo: _userRepo,
	}
}
