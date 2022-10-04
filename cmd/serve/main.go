package main

import (
	"chat_application/internal/adapter/db"
	"chat_application/internal/adapter/middlewares"
	"chat_application/internal/adapter/repository_database"
	"chat_application/internal/adapter/websocket"
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
	signUpUC := usecase.NewSignUpUseCase(ur)

	uc := usecase.UseCase{
		LoginUC:  loginUC,
		SignUpUC: signUpUC,
	}

	r := gin.Default()
	r.Use(middlewares.CORSMiddleware())
	r.Use(middlewares.Logger())
	r.Static("/test_chat", "./public")

	// Run server Websocket
	wsServer := websocket.NewWebsocketServer()
	go wsServer.Run()

	// Handle Socket Webserver Request
	r.GET("/ws", func(c *gin.Context) {
		websocket.ServeWs(wsServer, c)
	})

	passportGroup := r.Group("/v1/passport")
	_ = httpport.NewUserHander(uc, ur, passportGroup)
	r.GET("/tab", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "nine",
		})
	})
	r.Run()
}
