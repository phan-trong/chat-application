package domain

import "time"

type User_Friend struct {
	ID        int
	UserID    int
	FriendId  int
	DeletedAt time.Time
	CreateAt  time.Time
	UpdatedAt time.Time
}
