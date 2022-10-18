package repository

import (
	"chat_application/internal/domain"
	"context"
	"time"

	"gorm.io/gorm"
)

type userFriendRepository struct {
	DB *gorm.DB
}

func NewUserFriendRepository(db *gorm.DB) *userFriendRepository {
	return &userFriendRepository{
		DB: db,
	}
}

func (ufRepo *userFriendRepository) FindAllFriendOfUserId(ctx context.Context, userId int) ([]*domain.UserFriend, error) {
	var userFriends []*domain.UserFriend

	if result := ufRepo.DB.Where("user_id=?", userId).Find(&userFriends); result.Error != nil {
		return nil, result.Error
	}

	return userFriends, nil
}

func (ufRepo *userFriendRepository) AddFriend(ctx context.Context, userId int, friendId int) error {
	userFriend := &domain.UserFriend{
		UserID:   userId,
		FriendID: friendId,
	}

	if result := ufRepo.DB.Save(&userFriend); result.Error != nil {
		return result.Error
	}

	return nil
}

func (ufRepo *userFriendRepository) DeleteFriend(ctx context.Context, userId int, friendId int) error {
	var userFriend *domain.UserFriend

	if result := ufRepo.DB.Where("user_id = ? AND friend_id = ?", userId, friendId).First(&userFriend); result.Error != nil {
		return result.Error
	}

	userFriend.DeletedAt = time.Now()

	if result := ufRepo.DB.Save(&userFriend); result.Error != nil {
		return result.Error
	}

	return nil
}
