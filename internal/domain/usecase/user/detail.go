package user_usecase

import (
	user_entity "github.com/mariojuniortrab/hauling-api/internal/domain/entity/user"
	protocol_usecase "github.com/mariojuniortrab/hauling-api/internal/domain/usecase/protocol"
)

type DetailUserUseCase struct {
	userRepository protocol_usecase.UserRepository
}

func NewDetailuserUsecase(userRepository protocol_usecase.UserRepository) *DetailUserUseCase {
	return &DetailUserUseCase{userRepository}
}

func (u *DetailUserUseCase) Execute(id string) (*user_entity.UserDetailOutputDto, error) {
	user, err := u.userRepository.GetById(id)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, nil
	}

	output := user_entity.NewUserDetailOutputDto(user)
	return output, nil
}
