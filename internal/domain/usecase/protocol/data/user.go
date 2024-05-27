package protocol_data

import (
	user_entity "github.com/mariojuniortrab/hauling-api/internal/domain/entity/user"
)

type SignupRepository interface {
	GetByEmail(email string, id string) (*user_entity.User, error)
	Create(*user_entity.User) error
}

type LoginRepository interface {
	GetByEmail(email string, id string) (*user_entity.User, error)
	Login(email string, password string) (*user_entity.User, error)
}

type ListUserRepository interface {
	List(input *user_entity.ListUserParams) ([]*user_entity.User, error)
}

type DetailUserRepository interface {
	GetById(id string) (*user_entity.User, error)
}

type RemoveUserRepository interface {
	GetById(id string) (*user_entity.User, error)
	Remove(id string) error
}

type UpdateRepository interface {
	GetForUpdate(id string) (*user_entity.User, error)
	Update(id string, editedUser *user_entity.User) (*user_entity.User, error)
}
