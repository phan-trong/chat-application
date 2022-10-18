package http

import (
	"chat_application/internal/ports/middlewares"
	"chat_application/internal/usecase/auth"
	"chat_application/internal/usecase/user"
	"context"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserHandler struct {
	gRouter     gin.IRouter
	service     user.UseCase
	authService auth.UseCase
}

func NewUserHander(gRouter gin.IRouter, s user.UseCase, authService auth.UseCase) *UserHandler {
	h := &UserHandler{
		gRouter:     gRouter,
		service:     s,
		authService: authService,
	}
	gRouter.POST("/login", h.Login)
	gRouter.POST("/sign-up", h.SignUp)
	gRouter.GET("/me", middlewares.JwtAuthMiddleware(authService), h.GetInfor)
	gRouter.POST("/upload-avatar", middlewares.JwtAuthMiddleware(authService), h.UploadAvatar)
	// protected.Use(middlewares.JwtAuthMiddleware())
	// protected.GET("/user",controllers.CurrentUser)
	return h
}

func (u *UserHandler) Login(c *gin.Context) {
	_, cfn := context.WithCancel(c)
	defer cfn()
	var body user.LoginRequest

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, &ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	token, err := u.service.Login(body.Email, body.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &SuccessResponse{
		Code: http.StatusOK,
		Data: map[string]interface{}{"token": token},
	})
}

func (u *UserHandler) SignUp(c *gin.Context) {
	_, cfn := context.WithCancel(c)
	defer cfn()
	var body user.SignUpRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, &ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	_, err := u.service.SignUp(body.FullName, body.Email, body.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &SuccessResponse{
		Code: http.StatusOK,
		Data: map[string]interface{}{"status": "Create user success"},
	})
}

func (u *UserHandler) GetInfor(c *gin.Context) {
	_, cfn := context.WithCancel(c)
	defer cfn()

	user := middlewares.GetCurrentUser(c)

	c.JSON(http.StatusOK, &SuccessResponse{
		Code: http.StatusOK,
		Data: map[string]interface{}{"user": user},
	})
}

func (u *UserHandler) UploadAvatar(c *gin.Context) {
	_, cfn := context.WithCancel(c)
	defer cfn()

	user := middlewares.GetCurrentUser(c)

	file, err := c.FormFile("file")

	if err != nil {
		c.JSON(http.StatusBadRequest, &ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	// Retrieve file information
	extension := filepath.Ext(file.Filename)
	// Generate random file name for the new uploaded file so it doesn't override the old file with same name
	newFileName := uuid.New().String() + extension

	// The file is received, so let's save it
	if err := c.SaveUploadedFile(file, "./public/avatars/"+newFileName); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to save the file",
		})
		return
	}

	user.Avatar = "./public/avatars/" + newFileName
	err = u.service.UpdateUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, &ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your file has been successfully uploaded.",
	})
}
