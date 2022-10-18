package user

import (
	"chat_application/internal/domain"
	"chat_application/internal/ports/middlewares"
	"chat_application/internal/usecase/auth"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// User service implements
type Service struct {
	repo        Repository
	authService auth.UseCase
}

func NewService(r Repository, as auth.UseCase) *Service {
	return &Service{
		repo:        r,
		authService: as,
	}
}

// GetUser Get an User
func (s *Service) GetUser(id domain.ID) (*domain.User, error) {
	return s.repo.Get(id)
}

// SearchUsers Search Users
func (s *Service) SearchUsers(query string) ([]*domain.User, error) {
	return s.repo.Search(strings.ToLower(query))
}

// CreateUser Create an user
func (s *Service) CreateUser(name, email, password string) (string, error) {
	e, err := domain.NewUser(name, email, password)
	if err != nil {
		return e.ID.String(), err
	}
	return s.repo.Create(e)
}

// ListUsers List Users
func (s *Service) ListUsers() ([]*domain.User, error) {
	return s.repo.List()
}

// DeleteUser Delete an user
func (s *Service) DeleteUser(id domain.ID) error {
	u, err := s.GetUser(id)
	if u != nil {
		return err
	}
	if err != nil {
		return err
	}

	return s.repo.Delete(id)
}

// UpdateUser Update an user
func (s *Service) UpdateUser(e *domain.User) error {
	err := e.Validate()
	if err != nil {
		return domain.ErrInvalidEntity
	}
	e.UpdatedAt = time.Now()
	return s.repo.Update(e)
}

// Login
func (s *Service) Login(email, password string) (string, error) {
	user, err := s.repo.GetByEmail(email)
	if err != nil {
		return "", err
	}

	err = middlewares.VerifyPassword(password, user.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := s.authService.CreateJWTToken(user)

	if err != nil {
		return "", err
	}

	return token, nil
}

// SignUp
func (s *Service) SignUp(name, email, password string) (string, error) {
	u, _ := s.repo.GetByEmail(email)

	if u != nil {
		return "", domain.ErrExists
	}

	user, err := domain.NewUser(name, email, password)
	if err != nil {
		return "", err
	}

	return s.repo.Create(user)
}
