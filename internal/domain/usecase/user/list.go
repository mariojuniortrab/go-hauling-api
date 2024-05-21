package user_usecase

import (
	user_entity "github.com/mariojuniortrab/hauling-api/internal/domain/entity/user"
	protocol_usecase "github.com/mariojuniortrab/hauling-api/internal/domain/usecase/protocol"
)

type List struct {
	userRepository protocol_usecase.UserRepository
}

func NewListUseCase(userRepository protocol_usecase.UserRepository) *List {
	return &List{userRepository}
}

func (u *List) Execute(input *user_entity.ListUserInputDto) (*user_entity.ListOutputDto, error) {
	listUserParams, err := user_entity.NewListUserParams(input)
	if err != nil {
		return nil, err
	}

	users, err := u.userRepository.List(listUserParams)
	if err != nil {
		return nil, err
	}

	result := &user_entity.ListOutputDto{}

	for _, user := range users {
		result.Items = append(result.Items, user_entity.NewUserDetailOutputDto(user))
	}

	return result, nil
}
