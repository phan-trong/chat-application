package domain

import "context"

type UserFriendRepository interface {
	FindAllFriendOfUserId(ctx context.Context, userId int) ([]*User, error)
	AddFriend(ctx context.Context, userId int, friendId int) error
	DeleteFriend(ctx context.Context, userId int, friendId int) error
}
