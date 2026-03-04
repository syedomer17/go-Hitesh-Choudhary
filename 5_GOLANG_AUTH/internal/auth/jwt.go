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

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signed , err := token.SignedString([]byte(jwtSecret))

	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}
	return signed, nil
}

func ParseToken(jwtSecret string, tokenString string) (Claims, error){
	var claims Claims

	parsed, err := jwt.ParseWithClaims(tokenString, &claims, func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, fmt.Errorf("Unexpected signin method: %v",t.Header["alg"])
		}
		return []byte(jwtSecret), nil
	},
	jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}),
)
	if err != nil {
		return Claims{}, fmt.Errorf("failed to parse token: %w", err)
	}

	if !parsed.Valid {
		return Claims{}, fmt.Errorf("Invalid token")
	}

	if claims.Subject == "" {
		return Claims{}, fmt.Errorf("token missing subject")
	}
	return claims, nil
}
