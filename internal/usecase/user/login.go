package user_usecase

import (
	user_entity "github.com/mariojuniortrab/hauling-api/internal/entity/user"
	infra_errors "github.com/mariojuniortrab/hauling-api/internal/infra/errors"
	util_usecase "github.com/mariojuniortrab/hauling-api/internal/usecase/util"
)

type LoginValidation interface {
	Validate(*LoginInputDto) *infra_errors.CustomError
	ValidateCredentials(*UserDto, string) *infra_errors.CustomError
}

type LoginInputDto struct {
	Username string `validate:"required" json:"username"`
	Password string `validate:"required" json:"password"`
}

type UserDto struct {
	ID       string
	Name     string
	Active   bool
	Email    string
	Username string
	Password string
}

type LoginOutputDto struct {
	Token    string `json:"token"`
	ID       string `json:"id"`
	Name     string `json:"name"`
	Active   bool   `json:"active"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type Login struct {
	userRepository user_entity.UserRepository
	tokenizer      util_usecase.Tokenizer
}

func NewLoginUseCase(userRepository user_entity.UserRepository,
	tokenizer util_usecase.Tokenizer) *Login {
	return &Login{
		userRepository,
		tokenizer,
	}
}

func (u *Login) Execute(input *UserDto) (*LoginOutputDto, error) {
	token, err := u.tokenizer.GenerateToken(input.ID, input.Email)
	if err != nil {
		return nil, err
	}

	return &LoginOutputDto{
		Token:    token,
		ID:       input.ID,
		Name:     input.Name,
		Active:   input.Active,
		Email:    input.Email,
		Username: input.Username,
	}, nil
}

func (u *Login) GetByUsername(input *LoginInputDto) (*UserDto, error) {
	user, err := u.userRepository.GetByUsername(input.Username)
	if err != nil {
		return nil, err
	}

	return &UserDto{
		ID:       user.ID,
		Name:     user.Name,
		Active:   user.Active,
		Email:    user.Email,
		Username: user.Username,
	}, nil
}
