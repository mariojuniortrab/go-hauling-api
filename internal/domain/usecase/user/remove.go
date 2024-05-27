package user_usecase

import (
	"errors"

	protocol_data "github.com/mariojuniortrab/hauling-api/internal/domain/usecase/protocol/data"
)

type RemoveUserUseCase struct {
	repository protocol_data.RemoveUserRepository
}

func NewRemoveUserUsecase(repository protocol_data.RemoveUserRepository) *RemoveUserUseCase {
	return &RemoveUserUseCase{repository}
}

func (u *RemoveUserUseCase) Execute(id string) (error, error) {
	user, err := u.repository.GetById(id)

	if err != nil {
		return err, nil
	}

	if user == nil {
		return nil, errors.New("not found")
	}

	err = u.repository.Remove(id)
	if err != nil {
		return err, nil
	}

	return nil, nil
}
