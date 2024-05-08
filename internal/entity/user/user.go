package user_entity

import "github.com/google/uuid"

type UserRepository interface {
	Create(user *User) error
	GetByUsername(username string) (*User, error)
}

type User struct {
	ID       string
	Username string
	Name     string
	Password string
	Active   bool
}

func NewUser(username, name, password string) *User {
	return &User{
		ID:       uuid.New().String(),
		Username: username,
		Name:     name,
		Password: password,
		Active:   true,
	}
}
