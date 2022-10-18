package domain

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	ID         uuid.UUID
	Action     string
	SenderId   uuid.UUID
	RoomId     uuid.UUID
	Message    string
	DeletedAt  time.Time
	CreatedAt  time.Time
	UpdateddAt time.Time
}

type Target struct {
}

//NewMessage create a new Message
func NewMessage(senderId, roomId uuid.UUID, message string) (*Message, error) {
	u := &Message{
		ID:        NewID(),
		SenderId:  senderId,
		RoomId:    roomId,
		Message:   message,
		CreatedAt: time.Now(),
	}

	err := u.Validate()
	if err != nil {
		return nil, ErrInvalidEntity
	}

	return u, nil
}

//Validate validate data
func (m *Message) Validate() error {
	if m.Message == "" {
		return ErrInvalidEntity
	}

	return nil
}
