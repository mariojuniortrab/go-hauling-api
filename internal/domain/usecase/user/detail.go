package user_usecase

import (
	"fmt"

	user_entity "github.com/mariojuniortrab/hauling-api/internal/domain/entity/user"
	protocol_usecase "github.com/mariojuniortrab/hauling-api/internal/domain/usecase/protocol"
	util_usecase "github.com/mariojuniortrab/hauling-api/internal/domain/usecase/util"
)

type UserDetailInputDto struct {
	ID string
}

type userDetailOutputDto struct {
	ID     string `json:"id"`
	Email  string `json:"email"`
	Name   string `json:"name"`
	Birth  string `json:"birth"`
	Active bool   `json:"active"`
}

func NewUserDetailOutputDto(user *user_entity.User) *userDetailOutputDto {
	birth := util_usecase.GetStringFromDate(user.Birth)

	return &userDetailOutputDto{
		ID:     user.ID,
		Name:   user.Name,
		Email:  user.Email,
		Birth:  birth,
		Active: user.Active,
	}
}

type DetailUserUseCase struct {
	userRepository protocol_usecase.UserRepository
}

func NewDetailuserUsecase(userRepository protocol_usecase.UserRepository) *DetailUserUseCase {
	return &DetailUserUseCase{userRepository}
}

func (u *DetailUserUseCase) Execute(input *UserDetailInputDto) (*userDetailOutputDto, error) {
	user, err := u.userRepository.GetById(input.ID)
	fmt.Println("[user_usecase > detail > Execute] user:", user)
	if err != nil {
		fmt.Println("[user_usecase > detail > Execute] err:", err)
		return nil, err
	}

	output := NewUserDetailOutputDto(user)
	return output, nil
}
