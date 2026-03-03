package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	jwt.RegisteredClaims  

	Role string `json:"role"`
}

func CreateToken(jwtSecret string, userID string, role string) (string, error){
	now := time.Now().UTC()
	exp := now.Add(7 * 24 * time.Hour)

	claims := Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject: userID,
			IssuedAt: jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(exp),
		},
		Role: role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

	signed , err := token.SignedString([]byte(jwtSecret))

	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}
	return signed, nil
}