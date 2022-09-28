package main

import (
	"chat_application/internal/adapter/db"
	"chat_application/internal/adapter/middlewares"
	"chat_application/internal/adapter/repository_database"
	httpport "chat_application/internal/ports/http"
	"chat_application/internal/usecase"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUrl := os.Getenv("DB_URL")
	database := db.Init(dbUrl)

	ur := repository_database.NewUserRepository(database)
	loginUC := usecase.NewLoginUseCase(ur)

	uc := usecase.UseCase{
		LoginUC: loginUC,
	}

	r := gin.Default()
	r.Use(middlewares.Logger())
	passportGroup := r.Group("/v1/passport")
	_ = httpport.NewUserHander(uc, ur, passportGroup)
	r.GET("/tab", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "nine",
		})
	})
	r.Run()
}
