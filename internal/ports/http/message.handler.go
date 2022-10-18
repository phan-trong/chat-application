package http

import (
	"chat_application/internal/domain"
	"chat_application/internal/usecase/auth"
	"chat_application/internal/usecase/user"
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MessageHandler struct {
	gRouter     gin.IRouter
	repo        domain.Message
	authService auth.UseCase
}

func NewMessageHander(gRouter gin.IRouter, s user.UseCase) *UserHandler {
	h := &UserHandler{
		gRouter: gRouter,
		service: s,
	}
	gRouter.GET("/messages/:roomId", h.GetMessages)

	return h
}

//GeneratePaginationFromRequest ..
func generatePaginationFromRequest(c *gin.Context) domain.Pagination {
	// Initializing default
	//	var mode string
	limit := 10
	page := 1
	sort := "created_at asc"
	query := c.Request.URL.Query()
	for key, value := range query {
		queryValue := value[len(value)-1]
		switch key {
		case "limit":
			limit, _ = strconv.Atoi(queryValue)
			break
		case "page":
			page, _ = strconv.Atoi(queryValue)
			break
		case "sort":
			sort = queryValue
			break

		}
	}
	return domain.Pagination{
		Limit: limit,
		Page:  page,
		Sort:  sort,
	}

}
func (u *UserHandler) GetMessages(c *gin.Context) {
	_, cfn := context.WithCancel(c)
	defer cfn()
	roomId := c.Param("roomId")
	pagination := generatePaginationFromRequest(c)
	fmt.Println(roomId)
	fmt.Println(pagination)

	c.JSON(http.StatusOK, &SuccessResponse{
		Code: http.StatusOK,
		Data: map[string]interface{}{"pagination": pagination},
	})
}
