package user

import "chat_application/internal/domain"

// Reader interface
type Reader interface {
	Get(id domain.ID) (*domain.User, error)
	GetByEmail(email string) (*domain.User, error)
	Search(query string) ([]*domain.User, error)
	List() ([]*domain.User, error)
}

// Writer interface
type Writer interface {
	Create(e *domain.User) (string, error)
	Update(e *domain.User) error
	Delete(id domain.ID) error
}

// Repository Interface
type Repository interface {
	Reader
	Writer
}

// UseCase interface
type UseCase interface {
	GetUser(id domain.ID) (*domain.User, error)
	SearchUsers(query string) ([]*domain.User, error)
	ListUsers() ([]*domain.User, error)
	CreateUser(name, email, password string) (string, error)
	UpdateUser(e *domain.User) error
	DeleteUser(id domain.ID) error
	Login(email, password string) (string, error)
	SignUp(name, email, password string) (string, error)
}

type LoginRequest struct {
	Email    string
	Password string
}

type SignUpRequest struct {
	FullName string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
