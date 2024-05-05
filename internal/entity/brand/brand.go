package brand_entity

import (
	"github.com/google/uuid"
)

type BrandRepository interface {
	Create(brand *Brand) error
	ListAll() ([]*Brand, error)
	GetById(id string) (*Brand, error)
	GetByName(name string) (*Brand, error)
	GetByNameForEdition(name string, id string) (*Brand, error)
}

type Brand struct {
	ID   string
	Name string
}

func NewBrand(name string) *Brand {
	brand := Brand{
		ID:   uuid.New().String(),
		Name: name,
	}

	return &brand
}
