package auth

import "chat_application/internal/domain"

type UseCase interface {
	CreateJWTToken(user *domain.User) (string, error)
	ValidateToken(string) (*Claims, error)
}
