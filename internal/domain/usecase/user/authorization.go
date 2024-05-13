package user_usecase

import (
	"errors"

	protocol_usecase "github.com/mariojuniortrab/hauling-api/internal/domain/usecase/protocol"
)

type AuthInputDto struct {
	Token string `json:"token"`
}

type Authorization struct {
	tokenizer protocol_usecase.Tokenizer
}

func NewAuthorization(tokenizer protocol_usecase.Tokenizer) *Authorization {
	return &Authorization{
		tokenizer,
	}
}

func (a *Authorization) Execute(input *AuthInputDto) (*protocol_usecase.TokenOutputDto, error) {
	if input.Token == "" {
		return nil, errors.New("token is empty")
	}

	output, err := a.tokenizer.ParseToken(input.Token)
	if err != nil {
		return nil, err
	}

	return output, nil
}
