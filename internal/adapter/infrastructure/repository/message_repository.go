package repository

import (
	"chat_application/internal/domain"
	"time"

	"gorm.io/gorm"
)

type MessageRepository struct {
	DB *gorm.DB
}

func NewMessageRepository(db *gorm.DB) *MessageRepository {
	return &MessageRepository{
		DB: db,
	}
}

func (r *MessageRepository) List(roomId domain.ID) ([]*domain.Message, error) {
	var messages []*domain.Message

	if result := r.DB.Where("room_id=?", roomId).Find(&messages); result.Error != nil {
		return nil, result.Error
	}

	return messages, nil
}

func (r *MessageRepository) Create(m *domain.Message) error {
	if result := r.DB.Save(&m); result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *MessageRepository) Delete(messageId domain.ID) error {
	var message *domain.Message

	if result := r.DB.Where("id=?", messageId).First(&message); result.Error != nil {
		return result.Error
	}

	message.DeletedAt = time.Now()

	if result := r.DB.Save(&message); result.Error != nil {
		return result.Error
	}

	return nil
}
