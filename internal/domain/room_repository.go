package domain

import "context"

type RoomReposity interface {
	FindOneRoom(ctx context.Context, roomId int) (*Room, error)
	CreateRoom(ctx context.Context, r *Room) error
	DeleteRoom(ctx context.Context, roomId int) error
}
