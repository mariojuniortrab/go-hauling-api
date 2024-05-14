package user_usecase

import (
	protocol_usecase "github.com/mariojuniortrab/hauling-api/internal/domain/usecase/protocol"
)

type LoginInputDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserDto struct {
	ID       string
	Name     string
	Active   bool
	Email    string
	Password string
}

type LoginOutputDto struct {
	Token  string `json:"token"`
	ID     string `json:"id"`
	Name   string `json:"name"`
	Active bool   `json:"active"`
	Email  string `json:"email"`
}

type Login struct {
	userRepository protocol_usecase.UserRepository
	tokenizer      protocol_usecase.Tokenizer
}

func NewLoginUseCase(userRepository protocol_usecase.UserRepository,
	tokenizer protocol_usecase.Tokenizer) *Login {
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
		Token:  token,
		ID:     input.ID,
		Name:   input.Name,
		Active: input.Active,
		Email:  input.Email,
	}, nil
}

func (u *Login) GetByEmail(input *LoginInputDto) (*UserDto, error) {
	user, err := u.userRepository.GetByEmail(input.Email, "")

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, nil
	}

	return &UserDto{
		ID:       user.ID,
		Name:     user.Name,
		Active:   user.Active,
		Email:    user.Email,
		Password: user.Password,
	}, nil
}
