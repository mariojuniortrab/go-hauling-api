package infra_errors

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

func GetErrorMessage(v validator.FieldError) error {
	if v.Tag() == "required" {
		return errors.New(isRequired(v.Field()))
	}

	return errors.New("unknown error")
}

func isRequired(field string) string {
	return fmt.Sprintf("%s is required", strings.ToLower(field))
}
