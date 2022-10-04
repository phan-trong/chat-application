package domain

import "context"

type RoomUserRepository interface {
	FindAllMemberOfRoom(ctx context.Context, roomId int) ([]*RoomUser, error)
	AddMemberRoom(ctx context.Context, roomId int, memberId int) error
	DeleteMemberRoom(ctx context.Context, roomId int, memberId int) error
}
