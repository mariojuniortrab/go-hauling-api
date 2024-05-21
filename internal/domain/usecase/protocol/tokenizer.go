package protocol_usecase

import auth_entity "github.com/mariojuniortrab/hauling-api/internal/domain/entity/auth"

type Tokenizer interface {
	GenerateToken(id string, email string) (string, error)
	ParseToken(token string) (*auth_entity.TokenOutputDto, error)
}
