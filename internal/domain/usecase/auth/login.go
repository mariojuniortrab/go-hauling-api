package auth_usecase

import (
	auth_entity "github.com/mariojuniortrab/hauling-api/internal/domain/entity/auth"
	protocol_usecase "github.com/mariojuniortrab/hauling-api/internal/domain/usecase/protocol"
)

type Login struct {
	userRepository protocol_usecase.UserRepository
	tokenizer      protocol_usecase.Tokenizer
}

func NewLoginUseCase(userRepository protocol_usecase.UserRepository,
	tokenizer protocol_usecase.Tokenizer) *Login {
	return &Login{userRepository, tokenizer}
}

func (u *Login) Execute(input *auth_entity.LoginDto) (*auth_entity.LoginOutputDto, error) {
	token, err := u.tokenizer.GenerateToken(input.ID, input.Email)
	if err != nil {
		return nil, err
	}

	return auth_entity.NewLoginOutputDto(input, token), nil
}

func (u *Login) GetByEmail(input *auth_entity.LoginInputDto) (*auth_entity.LoginDto, error) {
	user, err := u.userRepository.GetByEmail(input.Email, "")
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, nil
	}

	return auth_entity.NewLoginDto(user), nil
}
