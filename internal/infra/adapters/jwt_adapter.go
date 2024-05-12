package infra_adapters

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type jwtAdapter struct{}

const secretKey = "abliblablubla"

func NewJwtAdapter() *jwtAdapter {
	return &jwtAdapter{}
}

func (j *jwtAdapter) GenerateToken(id string, email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"id":    id,
		"exp":   time.Now().Add(time.Hour + 1).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}
