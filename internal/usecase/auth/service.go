package auth

import (
	"chat_application/internal/adapter/infrastructure/repository"
	"chat_application/internal/domain"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const hmacSecret = "SecretValueReplaceThis"
const defaulExpireTime = 604800 // 1 week

type AuthService struct {
	userRepo repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) *AuthService {
	return &AuthService{
		userRepo: userRepo,
	}
}

type Claims struct {
	User *domain.User `json:"user"`
	jwt.StandardClaims
}

// CreateJWTToken generates a JWT signed token for for the given user
func (as *AuthService) CreateJWTToken(user *domain.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user":      user,
		"ExpiresAt": time.Now().Unix() + defaulExpireTime,
	})
	tokenString, err := token.SignedString([]byte(hmacSecret))

	return tokenString, err
}

func (as *AuthService) ValidateToken(tokenString string) (*Claims, error) {
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
		user, err := as.userRepo.Get(claims.User.GetID())
		if err != nil {
			return nil, err
		}
		claims.User = user

		return claims, nil
	} else {
		return nil, err
	}
}
