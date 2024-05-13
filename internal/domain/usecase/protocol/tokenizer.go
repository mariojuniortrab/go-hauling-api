package protocol_usecase

type Tokenizer interface {
	GenerateToken(id string, email string) (string, error)
	ParseToken(token string) (*TokenOutputDto, error)
}

type TokenOutputDto struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}
