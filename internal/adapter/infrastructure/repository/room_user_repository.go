package repository

import (
	"chat_application/internal/domain"
	"context"
	"time"

	"gorm.io/gorm"
)

type roomUserReposity struct {
	DB *gorm.DB
}

func NewRoomUserRepository(db *gorm.DB) *roomUserReposity {
	return &roomUserReposity{
		DB: db,
	}
}

func (ruRepo *roomUserReposity) FindAllMemberOfRoom(ctx context.Context, roomId int) ([]*domain.RoomUser, error) {
	var memberOfRoom []*domain.RoomUser

	if result := ruRepo.DB.Where("room_id=? AND deleted_at IS NULL", roomId).Find(&memberOfRoom); result.Error != nil {
		return nil, result.Error
	}

	return memberOfRoom, nil
}
func (ruRepo *roomUserReposity) AddMemberRoom(ctx context.Context, roomId int, memberId int) error {
	roomUser := domain.RoomUser{
		RoomID:   roomId,
		MemberID: memberId,
	}

	if result := ruRepo.DB.Save(&roomUser); result.Error != nil {
		return result.Error
	}

	return nil
}
func (ruRepo *roomUserReposity) DeleteMemberRoom(ctx context.Context, roomId int, memberId int) error {
	var roomUser *domain.RoomUser

	if result := ruRepo.DB.Where("room_id = ? AND member_id = ?", roomId, memberId).First(&roomUser); result.Error != nil {
		return result.Error
	}

	roomUser.DeletedAt = time.Now()

	if resultSave := ruRepo.DB.Save(&roomUser); resultSave.Error != nil {
		return resultSave.Error
	}

	return nil
}
