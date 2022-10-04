package repository_database

import (
	"chat_application/internal/domain"
	"context"
	"time"

	"gorm.io/gorm"
)

type messageRepository struct {
	DB *gorm.DB
}

func NewMessageRepository(db *gorm.DB) *messageRepository {
	return &messageRepository{
		DB: db,
	}
}

func (mRepo *messageRepository) FindAllMessageOfRoom(ctx context.Context, roomId int) ([]*domain.Message, error) {
	var messages []*domain.Message

	if result := mRepo.DB.Where("room_id=?", roomId).Find(&messages); result.Error != nil {
		return nil, result.Error
	}

	return messages, nil
}

func (mRepo *messageRepository) AddMessage(context context.Context, m *domain.Message) error {
	if result := mRepo.DB.Save(&m); result.Error != nil {
		return result.Error
	}
	return nil
}

func (mRepo *messageRepository) DeleteMessage(ctx context.Context, messageId int) error {
	var message *domain.Message

	if result := mRepo.DB.Where("id=?", messageId).First(&message); result.Error != nil {
		return result.Error
	}

	message.DeletedAt = time.Now()

	if result := mRepo.DB.Save(&message); result.Error != nil {
		return result.Error
	}

	return nil
}
