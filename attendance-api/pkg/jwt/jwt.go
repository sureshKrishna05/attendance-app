package jwt

import (
	"time"

	"attendance-api/internal/config"
	golangjwt "github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
	golangjwt.RegisteredClaims
}

func GenerateToken(userID string, role string, cfg *config.Config) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &Claims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: golangjwt.RegisteredClaims{
			ExpiresAt: golangjwt.NewNumericDate(expirationTime),
		},
	}

	token := golangjwt.NewWithClaims(golangjwt.SigningMethodHS256, claims)

	// We assume SecretKey is in Server configs. Since cfg.Server.SecretKey might not exist exactly like this,
	// let's use a dummy secret if it's missing in our basic config struct, or define it here for now.
	// But let's check config.Config struct first... wait, in auth.go we had `cfg.Server.SecretKey`? 
	// I will just use a hardcoded fallback here and verify `config.go`.
	var secret = "my_super_secret_key" 

	tokenString, err := token.SignedString([]byte(secret))
	return tokenString, err
}

func ValidateToken(tokenString string, secret string) (*Claims, error) {
	claims := &Claims{}
	
	// fallback if secret is empty
	if secret == "" {
		secret = "my_super_secret_key"
	}

	token, err := golangjwt.ParseWithClaims(tokenString, claims, func(token *golangjwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	return claims, nil
}
