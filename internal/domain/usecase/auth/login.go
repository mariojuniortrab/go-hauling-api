package auth_usecase

import (
	auth_entity "github.com/mariojuniortrab/hauling-api/internal/domain/entity/auth"
	protocol_application "github.com/mariojuniortrab/hauling-api/internal/domain/usecase/protocol/application"
	protocol_data "github.com/mariojuniortrab/hauling-api/internal/domain/usecase/protocol/data"
)

const timeToExpire = 4

type Login struct {
	repository protocol_data.LoginRepository
	tokenizer  protocol_application.Tokenizer
}

func NewLoginUseCase(repository protocol_data.LoginRepository,
	tokenizer protocol_application.Tokenizer) *Login {
	return &Login{repository, tokenizer}
}

func (u *Login) Execute(input *auth_entity.LoginDto) (*auth_entity.LoginOutputDto, error) {
	token, err := u.tokenizer.GenerateToken(input.ID, input.Email, timeToExpire)
	if err != nil {
		return nil, err
	}

	return auth_entity.NewLoginOutputDto(input, token), nil
}

func (u *Login) GetByEmail(input *auth_entity.LoginInputDto) (*auth_entity.LoginDto, error) {
	user, err := u.repository.GetByEmail(*input.Email, "")
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, nil
	}

	return auth_entity.NewLoginDto(user), nil
}
