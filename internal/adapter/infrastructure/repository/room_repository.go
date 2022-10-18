package repository

import (
	"chat_application/internal/domain"
	"time"

	"gorm.io/gorm"
)

type RoomRepository struct {
	DB *gorm.DB
}

func NewRoomRepository(db *gorm.DB) *RoomRepository {
	return &RoomRepository{
		DB: db,
	}
}

func (rRepo *RoomRepository) Get(roomId int) (*domain.Room, error) {
	var room domain.Room
	if result := rRepo.DB.Where("id=?", roomId).First(&room); result.Error != nil {
		return nil, domain.ErrRoomNotFound
	}

	return &room, nil

}

func (rRepo *RoomRepository) Create(r *domain.Room) error {
	if result := rRepo.DB.Save(&r); result.Error != nil {
		return result.Error
	}

	return nil
}

func (rRepo *RoomRepository) Update(roomId int, r *domain.Room) error {
	room, err := rRepo.Get(roomId)

	if err != nil {
		return domain.ErrRoomNotFound
	}

	r.ID = room.ID

	if result := rRepo.DB.Save(&r); result.Error != nil {
		return result.Error
	}

	return nil
}

func (rRepo *RoomRepository) Delete(roomId int) error {
	room, err := rRepo.Get(roomId)

	if err != nil {
		return domain.ErrRoomNotFound
	}

	room.DeletedAt = time.Now()

	if result := rRepo.DB.Save(&room); result.Error != nil {
		return result.Error
	}

	return nil
}
