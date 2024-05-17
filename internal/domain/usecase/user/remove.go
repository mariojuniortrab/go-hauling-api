package user_usecase

import (
	"errors"
	"fmt"

	user_entity "github.com/mariojuniortrab/hauling-api/internal/domain/entity/user"
	protocol_usecase "github.com/mariojuniortrab/hauling-api/internal/domain/usecase/protocol"
	util_usecase "github.com/mariojuniortrab/hauling-api/internal/domain/usecase/util"
)

type UserRemoveInputDto struct {
	ID string
}

type userRemoveOutputDto struct {
	ID     string `json:"id"`
	Email  string `json:"email"`
	Name   string `json:"name"`
	Birth  string `json:"birth"`
	Active bool   `json:"active"`
}

func NewUserRemoveOutputDto(user *user_entity.User) *userRemoveOutputDto {
	birth := util_usecase.GetStringFromDate(user.Birth)

	return &userRemoveOutputDto{
		ID:     user.ID,
		Name:   user.Name,
		Email:  user.Email,
		Birth:  birth,
		Active: user.Active,
	}
}

type RemoveUserUseCase struct {
	userRepository protocol_usecase.UserRepository
}

func NewRemoveUserUsecase(userRepository protocol_usecase.UserRepository) *RemoveUserUseCase {
	return &RemoveUserUseCase{userRepository}
}

func (u *RemoveUserUseCase) Execute(input *UserRemoveInputDto) (error, error) {
	user, err := u.userRepository.GetById(input.ID)
	fmt.Println("[user_usecase > Remove > Execute] user:", user)

	if err != nil {
		fmt.Println("[user_usecase > Remove > Execute] err:", err)
		return err, nil
	}

	if user == nil {
		return nil, errors.New("not found")
	}

	err = u.userRepository.Remove(input.ID)
	if err != nil {
		fmt.Println("[user_usecase > Remove > Execute] err:", err)
		return nil, err
	}

	return nil, nil
}
