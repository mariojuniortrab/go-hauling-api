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

func (u *List) Execute(input *user_entity.ListUserInputDto) (*user_entity.ListOutputDto, int, error) {
	listUserDto, err := user_entity.NewListUserDto(input)
	if err != nil {
		return nil, 0, err
	}

	users, total, err := u.repository.List(listUserDto)
	if err != nil {
		return nil, 0, err
	}

	result := &user_entity.ListOutputDto{}

	for _, user := range users {
		result.Items = append(result.Items, user_entity.NewUserDetailOutputDto(user))
	}

	return result, total, nil
}
