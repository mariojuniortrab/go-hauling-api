package user_usecase

import (
	user_entity "github.com/mariojuniortrab/hauling-api/internal/domain/entity/user"
	protocol_usecase "github.com/mariojuniortrab/hauling-api/internal/domain/usecase/protocol"
)

type UpdateUserUseCase struct {
	userRepository protocol_usecase.UserRepository
}

func NewUpdateUserUsecase(userRepository protocol_usecase.UserRepository) *UpdateUserUseCase {
	return &UpdateUserUseCase{userRepository}
}

func (u *UpdateUserUseCase) Execute(id string, input *user_entity.UserUpdateInputDto) (*user_entity.UserDetailOutputDto, error) {
	user, err := user_entity.NewUserFromUpdateInputDto(input)
	if err != nil {
		return nil, err
	}

	user, err = u.userRepository.Update(id, user)
	if err != nil {
		return nil, err
	}

	output := user_entity.NewUserDetailOutputDto(user)
	return output, nil
}

func (u *UpdateUserUseCase) GetForUpdate(id string) (*user_entity.UserUpdateInputDto, error) {
	user, err := u.userRepository.GetById(id)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, nil
	}

	return user_entity.NewUserUpdateInputDto(user), err
}
