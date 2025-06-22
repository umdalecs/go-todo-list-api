package auth

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/umdalecs/todo-list-api/internal"
)

func GenerateToken(userID int64) (string, error) {
	claims := jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(time.Hour * 48).Unix(),
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(internal.Envs.JwtSecret)
}
