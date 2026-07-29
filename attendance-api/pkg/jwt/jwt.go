package jwt

import (
	"errors"
	"time"

	"attendance-api/configs"

	jwt_lib "github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
	jwt_lib.RegisteredClaims
}

func GenerateToken(userID, role string, cfg *configs.Config) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt_lib.RegisteredClaims{
			ExpiresAt: jwt_lib.NewNumericDate(expirationTime),
		},
	}

	token := jwt_lib.NewWithClaims(jwt_lib.SigningMethodHS256, claims)
	return token.SignedString([]byte(cfg.JWTSecret))
}

func ValidateToken(tokenString string, cfg *configs.Config) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt_lib.ParseWithClaims(tokenString, claims, func(token *jwt_lib.Token) (interface{}, error) {
		return []byte(cfg.JWTSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
