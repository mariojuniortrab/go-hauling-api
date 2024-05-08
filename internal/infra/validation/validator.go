package infra_validation

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	infra_errors "github.com/mariojuniortrab/hauling-api/internal/infra/errors"
)

type validatorAdapter struct{}

func NewValidator() *validatorAdapter {
	return &validatorAdapter{}
}

func (v *validatorAdapter) Validate(s interface{}) *infra_errors.CustomError {
	validate := validator.New()

	err := validate.Struct(s)
	if err != nil {
		var messages []*infra_errors.CustomErrorMessage

		for _, validationErr := range err.(validator.ValidationErrors) {
			adaptedError := infra_errors.ValidationErrorAdapter(validationErr)
			messages = append(messages, infra_errors.NewCustomErrorMessage(adaptedError, validationErr.Field()))
		}

		return infra_errors.NewCustomErrorArray(messages, http.StatusBadRequest)
	}

	return nil
}
