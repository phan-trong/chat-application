package http

import (
	"chat_application/internal/adapter/middlewares"
	"chat_application/internal/domain"
	"chat_application/internal/usecase"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	uc       usecase.UseCase
	userRepo domain.UserRepository
	gRouter  gin.IRouter
}

func NewUserHander(uc usecase.UseCase, userRepo domain.UserRepository, gRouter gin.IRouter) *userHandler {
	h := &userHandler{
		uc:       uc,
		userRepo: userRepo,
		gRouter:  gRouter,
	}
	gRouter.POST("/login", h.Login)
	gRouter.POST("/sign-up", h.SignUp)
	// rotected.Use(middlewares.JwtAuthMiddleware())
	// protected.GET("/user",controllers.CurrentUser)
	gRouter.GET("/me", middlewares.JwtAuthMiddleware(), h.GetInfor)
	return h
}

func (u *userHandler) Login(c *gin.Context) {
	ctx, cfn := context.WithCancel(c)
	defer cfn()
	var body usecase.LoginRequest

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, &ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	resp, err := u.uc.LoginUC.Handle(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &SuccessResponse{
		Code: http.StatusOK,
		Data: map[string]interface{}{"token": resp},
	})
}

func (u *userHandler) SignUp(c *gin.Context) {
	ctx, cfn := context.WithCancel(c)
	defer cfn()
	var body usecase.SignUpRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, &ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	err := u.uc.SignUpUC.Handle(ctx, &body)
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

func (u *userHandler) GetInfor(c *gin.Context) {
	_, cfn := context.WithCancel(c)
	defer cfn()

	user := middlewares.GetCurrentUser(c)

	c.JSON(http.StatusOK, &SuccessResponse{
		Code: http.StatusOK,
		Data: map[string]interface{}{"user": user},
	})
}
