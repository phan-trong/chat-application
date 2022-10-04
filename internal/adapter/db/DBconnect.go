package db

import (
	"chat_application/internal/domain"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(url string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&domain.User{}, &domain.UserFriend{}, &domain.Room{}, &domain.RoomUser{}, &domain.Message{})

	return db
}
