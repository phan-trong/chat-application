package repository_database

import (
	"chat_application/internal/domain"
	"context"

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

func (ur *userRepository) FindOneByEmail(ctx context.Context, email string) (*domain.User, error) {
	var user domain.User
	if result := ur.DB.Where("email = ?", email).First(&user); result.Error != nil {
		return nil, domain.ErrUserNotFound
	}

	return &user, nil
}

func (ur *userRepository) FindOneById(ctx context.Context, id int) (*domain.User, error) {
	var user domain.User
	if result := ur.DB.Where("id = ?", id).First(&user); result.Error != nil {
		return nil, domain.ErrUserNotFound
	}

	return &user, nil
}

func (ur *userRepository) CreateUser(ctx context.Context, user *domain.User) error {
	if result := ur.DB.Create(&user); result.Error != nil {
		return result.Error
	}
	return nil
}

func (ur *userRepository) UpdateUser(ctx context.Context, id int, u *domain.User) error {
	user, err := ur.FindOneById(ctx, id)

	if err != nil {
		return domain.ErrUserNotFound
	}

	u.ID = user.ID

	if result := ur.DB.Save(&u); result.Error != nil {
		return result.Error
	}

	return nil
}

func (ur *userRepository) DeleteUser(ctx context.Context, id int) error {
	user, err := ur.FindOneById(ctx, id)

	if err != nil {
		return domain.ErrUserNotFound
	}

	if result := ur.DB.Delete(&user.ID); result.Error != nil {
		return result.Error
	}

	return nil
}
