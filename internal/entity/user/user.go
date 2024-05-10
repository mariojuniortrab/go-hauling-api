package user_entity

import (
	"time"

	"github.com/google/uuid"
)

type UserRepository interface {
	Create(*User) error
	GetByEmail(string, string) (*User, error)
	Login(string, string) (*User, error)
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
