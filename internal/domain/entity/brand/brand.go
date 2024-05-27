package brand_entity

import (
	"github.com/google/uuid"
)

type Brand struct {
	ID   string
	Name string
}

func NewBrand(name string) *Brand {

	return &Brand{
		ID:   uuid.New().String(),
		Name: name,
	}
}
