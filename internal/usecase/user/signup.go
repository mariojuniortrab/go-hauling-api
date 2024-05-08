package user_usecase

import (
	user_entity "github.com/mariojuniortrab/hauling-api/internal/entity/user"
	infra_errors "github.com/mariojuniortrab/hauling-api/internal/infra/errors"
)

type SignUpValidation interface {
	Validate(signUpDto SignUpInputDto) *infra_errors.CustomError
}

type SignUpInputDto struct {
	Username             string `validate:"required" json:"username"`
	Password             string `validate:"required" json:"password"`
	Name                 string `validate:"required" json:"name"`
	PasswordConfirmation string `validate:"required" json:"password_confirmation"`
}

type signUpOutputDto struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
}

type SignUpUseCase struct {
	userRepository user_entity.UserRepository
}

func NewSignUpUseCase(userRepository user_entity.UserRepository) *SignUpUseCase {
	return &SignUpUseCase{
		userRepository,
	}
}

func (s *SignUpUseCase) Execute(input SignUpInputDto) (*signUpOutputDto, error) {
	user := user_entity.NewUser(input.Username, input.Name, input.Password)

	err := s.userRepository.Create(user)
	if err != nil {
		return nil, err
	}

	return &signUpOutputDto{
		Username: user.Username,
		Name:     user.Name,
		ID:       user.ID,
	}, nil
}
