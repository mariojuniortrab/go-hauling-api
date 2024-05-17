package protocol_usecase

import (
	user_entity "github.com/mariojuniortrab/hauling-api/internal/domain/entity/user"
)

type UserRepository interface {
	Create(*user_entity.User) error
	GetByEmail(email string, id string) (*user_entity.User, error)
	Login(email string, password string) (*user_entity.User, error)
	List(input *user_entity.ListUserParams) ([]*user_entity.User, error)
	GetById(id string) (*user_entity.User, error)
	Remove(id string) error
}
