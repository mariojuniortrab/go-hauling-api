package infra_adapters

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type BcryptAdapter struct{}

func NewBcryptAdapter() *BcryptAdapter {
	return &BcryptAdapter{}
}

func (b *BcryptAdapter) Hash(password string) (string, error) {
	fmt.Println("[infra_adapters > BcryptAdapter > Hash] password:", password)

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (b *BcryptAdapter) CheckPasswordHash(hash, password string) bool {
	fmt.Println("[infra_adapters > BcryptAdapter > CheckPasswordHash] hash:", hash)
	fmt.Println("[infra_adapters > BcryptAdapter > CheckPasswordHash] password:", password)

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
