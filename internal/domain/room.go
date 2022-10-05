package domain

import (
	"time"

	"github.com/google/uuid"
)

type Room struct {
	ID         uuid.UUID `json:"id"`
	RoomName   string    `json:"room_name"`
	RoomAuthor int       // User Id created room
	DeletedAt  time.Time
	CreateAt   time.Time
	UpdatedAt  time.Time
}
