package repository_database

import (
	"chat_application/internal/domain"
	"context"
	"time"

	"gorm.io/gorm"
)

type roomRepository struct {
	DB *gorm.DB
}

func NewRoomRepository(db *gorm.DB) *roomRepository {
	return &roomRepository{
		DB: db,
	}
}

func (rRepo *roomRepository) FindOneRoom(ctx context.Context, roomId int) (*domain.Room, error) {
	var room domain.Room
	if result := rRepo.DB.Where("id=?", roomId).First(&room); result.Error != nil {
		return nil, domain.ErrRoomNotFound
	}

	return &room, nil

}

func (rRepo *roomRepository) CreateRoom(ctx context.Context, r *domain.Room) error {
	if result := rRepo.DB.Save(&r); result.Error != nil {
		return result.Error
	}

	return nil
}

func (rRepo *roomRepository) UpdateRoom(ctx context.Context, roomId int, r *domain.Room) error {
	room, err := rRepo.FindOneRoom(ctx, roomId)

	if err != nil {
		return domain.ErrRoomNotFound
	}

	r.ID = room.ID

	if result := rRepo.DB.Save(&r); result.Error != nil {
		return result.Error
	}

	return nil
}

func (rRepo *roomRepository) DeleteRoom(ctx context.Context, roomId int) error {
	room, err := rRepo.FindOneRoom(ctx, roomId)

	if err != nil {
		return domain.ErrRoomNotFound
	}

	room.DeletedAt = time.Now()

	if result := rRepo.DB.Save(&room); result.Error != nil {
		return result.Error
	}

	return nil
}
