package domain

import (
	"context"
)

type MessageRepository interface {
	FindAllMessageOfRoom(ctx context.Context, roomId int) ([]*Message, error)
	AddMessage(context context.Context, m *Message) error
	DeleteMessage(ctx context.Context, messageId int) error
}
