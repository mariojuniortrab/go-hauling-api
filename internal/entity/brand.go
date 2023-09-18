package entity

import (
	"errors"

	"github.com/google/uuid"
)

type Brand struct {
	ID   string
	Name string
}

func (b *Brand) Validate() error {
	if b.Name == "" {
		return errors.New("name is required")
	}

	return nil
}

func NewBrand(name string) (*Brand, error) {
	brand := Brand{
		ID:   uuid.New().String(),
		Name: name,
	}
	error := brand.Validate()
	if error != nil {
		return nil, error
	}

	return &brand, nil
}
