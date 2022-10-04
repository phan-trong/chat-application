package middlewares

import (
	"chat_application/internal/adapter/services"
	"chat_application/internal/domain"
	"chat_application/pkg/globals"
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		u, err := TokenValid(c)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
		c.Set(globals.UserKey, u)
		c.Next()
	}
}

func TokenValid(c *gin.Context) (*domain.User, error) {
	tokenString := ExtractToken(c)
	jwtData, err := services.ValidateToken(tokenString)

	if err != nil {
		return nil, err
	}
	return jwtData.User, err
}

func ExtractToken(c *gin.Context) string {
	token := c.Query("token")
	if token != "" {
		return token
	}
	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func GetCurrentUser(ctx context.Context) *domain.User {
	cu := ctx.Value(globals.UserKey).(*domain.User)
	return cu
}
