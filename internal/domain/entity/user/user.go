package user_entity

import (
	"time"

	"github.com/google/uuid"
	default_entity "github.com/mariojuniortrab/hauling-api/internal/domain/entity/default"
)

type ListUserParams struct {
	default_entity.List
	ID                string
	Name              string
	WillFilterActives bool
	Active            bool
	Email             string
}

type User struct {
	ID       string
	Name     string
	Password string
	Active   bool
	Email    string
	Birth    time.Time
}

func NewUser(name, password, email string, birth time.Time) *User {

	return &User{
		ID:       uuid.New().String(),
		Name:     name,
		Password: password,
		Active:   true,
		Email:    email,
		Birth:    birth,
	}
}
