package entity

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	infra_errors "github.com/mariojuniortrab/hauling-api/internal/infra/errors"
)

type Brand struct {
	ID   string
	Name string `validate:"required"`
}

func (b *Brand) Validate() error {
	validate := validator.New(validator.WithRequiredStructEnabled())

	error := validate.Struct(b)
	if error != nil {
		for _, err := range error.(validator.ValidationErrors) {
			return infra_errors.GetErrorMessage(err)
		}
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
