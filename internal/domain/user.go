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
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
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

//NewUser create a new user
func NewUser(name, email, password string) (*User, error) {
	u := &User{
		ID:        NewID(),
		FullName:  name,
		Avatar:    name,
		Email:     email,
		CreatedAt: time.Now(),
	}
	err := u.HashPassword(password)
	if err != nil {
		return nil, err
	}

	err = u.Validate()
	if err != nil {
		return nil, ErrInvalidEntity
	}
	return u, nil
}

//Validate validate data
func (u *User) Validate() error {
	if u.Email == "" || u.FullName == "" || u.Password == "" {
		return ErrInvalidEntity
	}

	return nil
}

func (u *User) HashPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
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
