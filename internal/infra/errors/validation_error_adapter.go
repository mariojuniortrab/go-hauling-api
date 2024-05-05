package infra_errors

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

func ValidationErrorAdapter(v validator.FieldError) error {
	if v.Tag() == "required" {
		return IsRequired(v.Field())
	}

	return errors.New("unknown error")
}
