package message

import (
	"chat_application/internal/domain"
)

// Reader interface
type Reader interface {
	List(roomId domain.ID) ([]*domain.Message, error)
}

// Writer interface
type Writer interface {
	Create(e *domain.Message) error
	Delete(id domain.ID) error
}

// Repository Interface
type Repository interface {
	Reader
	Writer
}
