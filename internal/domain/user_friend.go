package domain

import "time"

type UserFriend struct {
	ID        int
	UserID    int
	FriendID  int
	DeletedAt time.Time
	CreateAt  time.Time
	UpdatedAt time.Time
}
