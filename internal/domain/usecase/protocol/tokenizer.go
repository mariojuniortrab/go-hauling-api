package protocol_usecase

type Tokenizer interface {
	GenerateToken(string, string) (string, error)
}
