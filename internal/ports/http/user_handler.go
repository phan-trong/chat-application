package http

import (
	"chat_application/internal/domain"
	"chat_application/internal/usecase"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	uc       usecase.UseCase
	userRepo domain.UserRepository
}

func NewUserHander(uc usecase.UseCase, userRepo domain.UserRepository) *userHandler {
	return &userHandler{
		uc:       uc,
		userRepo: userRepo,
	}
}

func (u *userHandler) Login(c *gin.Context) {
	ctx := context.TODO()
	body := usecase.LoginRequest{}

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
