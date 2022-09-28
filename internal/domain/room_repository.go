package domain

import (
	"context"
	"errors"
)

var (
	ErrRoomNotFound = errors.New("room not found")
)

type RoomReposity interface {
	FindOneRoom(ctx context.Context, roomId int) (*Room, error)
	CreateRoom(ctx context.Context, r *Room) error
	UpdateRoom(ctx context.Context, roomId int, r *Room) error
	DeleteRoom(ctx context.Context, roomId int) error
}
