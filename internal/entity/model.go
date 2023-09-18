package entity

import (
	"github.com/go-playground/validator/v10"
	infra_errors "github.com/mariojuniortrab/hauling-api/internal/infra/errors"
)

type Model struct {
	ID    string
	Name  string `validate:"required"`
	Brand Brand  `validate:"required"`
}

func (m *Model) Validate() error {
	validate := validator.New(validator.WithRequiredStructEnabled())

	error := validate.Struct(m)
	if error != nil {
		for _, err := range error.(validator.ValidationErrors) {
			return infra_errors.GetErrorMessage(err)
		}
	}

	return nil
}
