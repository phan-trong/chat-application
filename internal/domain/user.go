package domain

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	FullName  string    `gorm:"size:255;not null" json:"full_name"`
	Avatar    string    `json:"avatar"`
	Email     string    `gorm:"size:255;not null;unique" json:"email"`
	Password  string    `gorm:"size:255;not null;" json:"password"`
	DeletedAt *time.Time
	CreateAt  time.Time
	UpdatedAt time.Time
}

func (u *User) GetID() uuid.UUID {
	return u.ID
}

func (u *User) GetFullName() string {
	return u.FullName
}

func (u *User) GetAvatarURL() string {
	return u.Avatar
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) ComparePassword(password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return false
	}
	return true
}
