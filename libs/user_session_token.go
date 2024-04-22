package libs

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenearteJwtToken(email string) (string, error) {
	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 24 * 40).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return "", fmt.Errorf("libs: jwt secret is required")
	}

	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", fmt.Errorf("libs: signed string, details: %w", err)
	}

	return tokenString, nil
}

func ValidateJwtToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		jwtSecret := os.Getenv("JWT_SECRET")
		if jwtSecret == "" {
			return nil, fmt.Errorf("libs: jwt secret is required")
		}
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return "", fmt.Errorf("libs: parse token, details: %w", err)
	}
	if !token.Valid {
		return "", fmt.Errorf("libs: invalid token")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("libs: invalid claims")
	}
	email, ok := claims["email"].(string)
	if !ok {
		return "", fmt.Errorf("libs: invalid email claim")
	}

	return email, nil
}
