package user_usecase

import (
	user_entity "github.com/mariojuniortrab/hauling-api/internal/domain/entity/user"
	protocol_data "github.com/mariojuniortrab/hauling-api/internal/domain/usecase/protocol/data"
)

type List struct {
	repository protocol_data.ListUserRepository
}

func NewListUseCase(repository protocol_data.ListUserRepository) *List {
	return &List{repository}
}

func (u *List) Execute(input *user_entity.ListUserInputDto) (*user_entity.ListOutputDto, error) {
	listUserParams, err := user_entity.NewListUserParams(input)
	if err != nil {
		return nil, err
	}

	users, err := u.repository.List(listUserParams)
	if err != nil {
		return nil, err
	}

	result := &user_entity.ListOutputDto{}

	for _, user := range users {
		result.Items = append(result.Items, user_entity.NewUserDetailOutputDto(user))
	}

	return result, nil
}
