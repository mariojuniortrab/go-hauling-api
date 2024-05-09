package util_usecase

type Encrypter interface {
	Hash(string) (string, error)
	CheckPasswordHash(string, string) bool
}
