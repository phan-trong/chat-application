package main

import (
	"chat_application/internal/adapter/db"
	"chat_application/internal/adapter/infrastructure/repository"
	"chat_application/internal/ports/http"
	"chat_application/internal/ports/middlewares"
	"chat_application/internal/ports/websocket"
	"chat_application/internal/usecase/auth"
	"chat_application/internal/usecase/user"
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

	uRepo := repository.NewUserRepository(database)
	authService := auth.NewAuthService(*uRepo)
	uService := user.NewService(uRepo, authService)

	mRepo := repository.NewMessageRepository(database)

	r := gin.Default()
	r.Use(middlewares.Logger())
	r.Use(middlewares.CORSMiddleware())

	api := r.Group("/api")
	api.Static("/images", "./public/avatars")
	api.Static("/test_chat", "./public")

	// Run server Websocket
	wsServer := websocket.NewWebsocketServer(uRepo)
	go wsServer.Run()

	// Handle Socket Webserver Request
	api.GET("/ws", middlewares.JwtAuthMiddleware(authService), func(c *gin.Context) {
		websocket.ServeWs(wsServer, c, mRepo)
	})

	// Set Max File Size Upload
	r.MaxMultipartMemory = 10 << 20 // 10 MiB

	passportGroup := api.Group("/v1/passport")
	_ = http.NewUserHander(passportGroup, uService, authService)
	chatGroup := api.Group("/v1/chat-application")
	_ = http.NewMessageHander(chatGroup, uService)
	api.GET("/tab", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "nine",
		})
	})

	r.Run()
}
