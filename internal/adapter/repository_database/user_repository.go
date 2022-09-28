package repository_database

import (
	"chat_application/internal/domain"
	"context"
	"time"

	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{
		DB: db,
	}
}

func (uRepo *userRepository) FindOneByEmail(ctx context.Context, email string) (*domain.User, error) {
	var user domain.User
	if result := uRepo.DB.Where("email = ?", email).First(&user); result.Error != nil {
		return nil, domain.ErrUserNotFound
	}

	return &user, nil
}

func (uRepo *userRepository) FindOneById(ctx context.Context, id int) (*domain.User, error) {
	var user domain.User
	if result := uRepo.DB.Where("id = ?", id).First(&user); result.Error != nil {
		return nil, domain.ErrUserNotFound
	}

	return &user, nil
}

func (uRepo *userRepository) CreateUser(ctx context.Context, user *domain.User) error {
	if result := uRepo.DB.Create(&user); result.Error != nil {
		return result.Error
	}
	return nil
}

func (uRepo *userRepository) UpdateUser(ctx context.Context, id int, u *domain.User) error {
	user, err := uRepo.FindOneById(ctx, id)

	if err != nil {
		return domain.ErrUserNotFound
	}

	u.ID = user.ID

	if result := uRepo.DB.Save(&u); result.Error != nil {
		return result.Error
	}

	return nil
}

func (uRepo *userRepository) DeleteUser(ctx context.Context, id int) error {
	user, err := uRepo.FindOneById(ctx, id)

	if err != nil {
		return domain.ErrUserNotFound
	}

	user.DeletedAt = time.Now()

	if result := uRepo.DB.Save(&user); result.Error != nil {
		return result.Error
	}

	return nil
}
