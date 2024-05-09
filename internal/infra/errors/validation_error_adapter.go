package infra_errors

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

func ValidationErrorAdapter(v validator.FieldError) error {
	switch {
	case v.Tag() == "required":
		return IsRequired(v.Field())
	case v.Tag() == "eqfield":
		return IsInvalid(v.Field())
	case v.Tag() == "alphanumunicode":
		return IsInvalid(v.Field())
	case v.Tag() == "email":
		return IsInvalid(v.Field())
	}

	return errors.New("unknown error")
}
