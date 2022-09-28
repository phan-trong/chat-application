package domain

import "time"

type RoomUser struct {
	ID        int
	RoomID    int
	MemberID  int // UserID member of room
	DeletedAt time.Time
	CreateAt  time.Time
	UpdatedAt time.Time
}
