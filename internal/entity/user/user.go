package user_entity

import (
	"time"

	"github.com/google/uuid"
)

type UserRepository interface {
	Create(user *User) error
	GetByUsername(username string) (*User, error)
	Login(username, password string) (*User, error)
}

type User struct {
	ID       string
	Username string
	Name     string
	Password string
	Active   bool
	Email    string
	Birth    time.Time
}

func NewUser(username, name, password, email string, birth time.Time) *User {

	return &User{
		ID:       uuid.New().String(),
		Username: username,
		Name:     name,
		Password: password,
		Active:   true,
		Email:    email,
		Birth:    birth,
	}
}
