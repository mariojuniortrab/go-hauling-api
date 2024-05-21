package user_usecase

import (
	"errors"

	protocol_usecase "github.com/mariojuniortrab/hauling-api/internal/domain/usecase/protocol"
)

type RemoveUserUseCase struct {
	userRepository protocol_usecase.UserRepository
}

func NewRemoveUserUsecase(userRepository protocol_usecase.UserRepository) *RemoveUserUseCase {
	return &RemoveUserUseCase{userRepository}
}

func (u *RemoveUserUseCase) Execute(id string) (error, error) {
	user, err := u.userRepository.GetById(id)

	if err != nil {
		return err, nil
	}

	if user == nil {
		return nil, errors.New("not found")
	}

	err = u.userRepository.Remove(id)
	if err != nil {
		return err, nil
	}

	return nil, nil
}
