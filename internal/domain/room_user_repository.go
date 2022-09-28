package domain

import "context"

type RoomUserRepository interface {
	FindAllMemberOfRoom(ctx context.Context, roomId int) ([]*User, error)
	AddMemberRoom(ctx context.Context, roomId int, memberId int) error
	DeleteMemberRoom(ctx context.Context, roomId int, memberId int) error
	InvitationMember(ctx context.Context, roomId int, memberId int) error
	LeaveRoom(ctx context.Context, roomId int, memberid int) error
}
