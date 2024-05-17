package user_usecase

import (
	"fmt"
	"strings"

	user_entity "github.com/mariojuniortrab/hauling-api/internal/domain/entity/user"
	protocol_usecase "github.com/mariojuniortrab/hauling-api/internal/domain/usecase/protocol"
	util_usecase "github.com/mariojuniortrab/hauling-api/internal/domain/usecase/util"
)

type UserUpdateInputDto struct {
	ID string
}

type UpdateFields struct {
	Name                 string `json:"name"`
	Birth                string `json:"birth"`
	Active               string `json:"active"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"passwordConfirmation"`
}

type userUpdateOutputDto struct {
	ID     string `json:"id"`
	Email  string `json:"email"`
	Name   string `json:"name"`
	Birth  string `json:"birth"`
	Active bool   `json:"active"`
}

func NewUserUpdateOutputDto(user *user_entity.User) *userUpdateOutputDto {
	birth := util_usecase.GetStringFromDate(user.Birth)

	return &userUpdateOutputDto{
		ID:     user.ID,
		Name:   user.Name,
		Email:  user.Email,
		Birth:  birth,
		Active: user.Active,
	}
}

type UpdateUserUseCase struct {
	userRepository protocol_usecase.UserRepository
}

func NewUpdateUserUsecase(userRepository protocol_usecase.UserRepository) *UpdateUserUseCase {
	return &UpdateUserUseCase{userRepository}
}

func (u *UpdateUserUseCase) Execute(input *UserUpdateInputDto, updateValues *UpdateFields) (*userUpdateOutputDto, error) {
	user, err := u.userRepository.GetById(input.ID)
	fmt.Println("[user_usecase > update > Execute] user:", user)

	if err != nil {
		fmt.Println("[user_usecase > update > Execute] err:", err)
		return nil, err
	}

	if user == nil {
		return nil, nil
	}

	updatedUser := getUpdatedUser(user, updateValues)

	err = u.userRepository.Update(input.ID, updatedUser)
	if err != nil {
		return nil, err
	}

	output := NewUserUpdateOutputDto(updatedUser)
	return output, nil
}

func getUpdatedUser(savedUser *user_entity.User,
	updateValues *UpdateFields) *user_entity.User {
	updatedUser := savedUser

	if updateValues.Name != "" {
		updatedUser.Name = updateValues.Name
	}

	if updateValues.Birth != "" {
		birth, _ := util_usecase.GetDateFromString(updateValues.Birth)
		updatedUser.Birth = birth
	}

	if updateValues.Active != "" {
		updatedUser.Active = strings.ToLower(updateValues.Active) == "true"
	}

	if updateValues.Password != "" {
		updatedUser.Password = updateValues.Password
	}

	return updatedUser
}
