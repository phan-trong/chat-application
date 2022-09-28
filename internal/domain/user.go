package domain

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int
	FullName  string
	Avatar    string
	Email     string
	Password  string
	DeletedAt time.Time
	CreateAt  time.Time
	UpdatedAt time.Time
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
