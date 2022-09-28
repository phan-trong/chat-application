package domain

import "time"

type Room struct {
	ID         int
	RoomName   string
	RoomAuthor int // User Id created room
	DeletedAt  time.Time
	CreateAt   time.Time
	UpdatedAt  time.Time
}
