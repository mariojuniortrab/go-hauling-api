package infra_adapters

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	protocol_usecase "github.com/mariojuniortrab/hauling-api/internal/domain/usecase/protocol"
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

func (j *jwtAdapter) ParseToken(token string) (*protocol_usecase.TokenOutputDto, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if !parsedToken.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token")
	}

	return &protocol_usecase.TokenOutputDto{
			Email: claims["email"].(string),
			ID:    claims["id"].(string),
		},
		nil
}
