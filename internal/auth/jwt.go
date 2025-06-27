package auth

import (
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/umdalecs/todo-list-api/config"
)

type CustomClaims struct {
	jwt.RegisteredClaims
}

func GenerateToken(userID int) (string, error) {
	claims := CustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   strconv.Itoa(userID),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(config.Envs.JwtSecret)
}

func ValidateToken(token string) (*CustomClaims, error) {
	claims := &CustomClaims{}

	parsed, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", t.Header["alg"])
		}

		return []byte(config.Envs.JwtSecret), nil

	})

	if err != nil {
		return nil, fmt.Errorf("error parsing token: %w", err)
	}

	claims, ok := parsed.Claims.(*CustomClaims)
	if !ok || !parsed.Valid {
		return nil, fmt.Errorf("token is not valid")
	}

	return claims, nil
}
