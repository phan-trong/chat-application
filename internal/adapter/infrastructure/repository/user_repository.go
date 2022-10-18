package repository

import (
	"chat_application/internal/domain"
	"time"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (r *UserRepository) Get(id domain.ID) (*domain.User, error) {
	var user domain.User
	if result := r.DB.Where("id = ?", id).First(&user); result.Error != nil {
		return nil, domain.ErrNotFound
	}

	return &user, nil
}

func (r *UserRepository) GetByEmail(email string) (*domain.User, error) {
	var user domain.User
	if result := r.DB.Where("email = ?", email).First(&user); result.Error != nil {
		return nil, domain.ErrNotFound
	}
	return &user, nil
}

func (r *UserRepository) Search(query string) ([]*domain.User, error) {
	var users []*domain.User
	// TODO: add query search in sql statement
	if result := r.DB.Where("deleted_at IS NOT NULL ").Find(&users); result.Error != nil {
		return nil, domain.ErrNotFound
	}
	return users, nil
}

func (r *UserRepository) List() ([]*domain.User, error) {
	var users []*domain.User
	if result := r.DB.Where("deleted_at IS NOT NULL").Find(&users); result.Error != nil {
		return nil, domain.ErrNotFound
	}
	return users, nil
}

func (r *UserRepository) Create(user *domain.User) (string, error) {
	if result := r.DB.Create(&user); result.Error != nil {
		return "", result.Error
	}
	return "", nil
}

func (r *UserRepository) Update(e *domain.User) error {
	_, err := r.Get(e.ID)
	if err != nil {
		return domain.ErrNotFound
	}
	e.UpdatedAt = time.Now()

	if result := r.DB.Save(&e); result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *UserRepository) Delete(id domain.ID) error {
	user, err := r.Get(id)
	if err != nil {
		return domain.ErrNotFound
	}
	currentTime := time.Now()
	user.DeletedAt = &currentTime
	if result := r.DB.Save(&user); result.Error != nil {
		return result.Error
	}
	return nil
}
