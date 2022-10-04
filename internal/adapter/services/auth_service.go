package services

import (
	"chat_application/internal/domain"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const hmacSecret = "SecretValueReplaceThis"
const defaulExpireTime = 604800 // 1 week

type Claims struct {
	User *domain.User `json:"user"`
	jwt.StandardClaims
}

// CreateJWTToken generates a JWT signed token for for the given user
func CreateJWTToken(user *domain.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user":      user,
		"ExpiresAt": time.Now().Unix() + defaulExpireTime,
	})
	tokenString, err := token.SignedString([]byte(hmacSecret))

	return tokenString, err
}

func ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(hmacSecret), nil
	})

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		claims.User.Password = ""
		return claims, nil
	} else {
		return nil, err
	}
}
