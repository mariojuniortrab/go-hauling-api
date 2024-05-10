package infra_validation

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	infra_errors "github.com/mariojuniortrab/hauling-api/internal/infra/errors"
	"github.com/mariojuniortrab/hauling-api/internal/validation"
)

type validatorAdapter struct {
	validator *validator.Validate
	errors    []*infra_errors.CustomErrorMessage
}

func NewValidator() *validatorAdapter {
	return &validatorAdapter{
		validator: validator.New(validator.WithRequiredStructEnabled()),
	}
}

func (v *validatorAdapter) HasErrors() bool {
	return len(v.errors) > 0
}

func (v *validatorAdapter) AddError(err error, fieldName string) validation.Validator {
	v.errors = append(v.errors, infra_errors.NewCustomErrorMessage(err, fieldName))
	return v
}

func (v *validatorAdapter) GetErrors() *infra_errors.CustomError {
	return infra_errors.NewCustomErrorArray(v.errors, http.StatusBadRequest)
}

func (v *validatorAdapter) GetErrorsAndClean() *infra_errors.CustomError {
	errors := v.errors

	v.errors = []*infra_errors.CustomErrorMessage{}

	return infra_errors.NewCustomErrorArray(errors, http.StatusBadRequest)
}

func (v *validatorAdapter) ValidateRequiredField(f interface{}, fieldName string) validation.Validator {
	return v.defaultValidation(f, fieldName, "required", infra_errors.IsRequired)
}

func (v *validatorAdapter) ValidateEmailField(f interface{}, fieldName string) validation.Validator {
	return v.defaultValidation(f, fieldName, "omitempty,email", infra_errors.IsInvalid)
}

func (v *validatorAdapter) ValidatePasswordConfirmationEquals(password, passwordConfirmation string) validation.Validator {
	fn := infra_errors.IsPasswordConfirmationInvalid
	return v.defaultFieldCompareValidation(password, passwordConfirmation, "passwordConfirmation", fn)
}

func (v *validatorAdapter) ValidateFieldString(f interface{}, fieldName string) validation.Validator {
	return v.defaultValidation(f, fieldName, "omitempty,alphaunicode", infra_errors.MustBeString)
}

func (v *validatorAdapter) ValidateFieldLength(f interface{}, fieldName string, length int) validation.Validator {
	return v.defaultLenghValidation(f, fieldName, fmt.Sprintf("omitempty,len=%d", length), infra_errors.LengthMustBe(fieldName, length))
}

func (v *validatorAdapter) ValidateFieldMaxLength(f interface{}, fieldName string, length int) validation.Validator {
	return v.defaultLenghValidation(f, fieldName, fmt.Sprintf("omitempty,max=%d", length), infra_errors.LengthMustBeOrLess(fieldName, length))
}

func (v *validatorAdapter) defaultLenghValidation(f interface{}, fieldName, flag string, errMessage error) validation.Validator {
	err := v.validator.Var(f, flag)

	if err != nil {
		v.errors = append(v.errors, infra_errors.NewCustomErrorMessage(errMessage, fieldName))
	}

	return v
}

func (v *validatorAdapter) defaultValidation(f interface{}, fieldName, flag string, fn func(string) error) validation.Validator {
	err := v.validator.Var(f, flag)

	if err != nil {
		v.errors = append(v.errors, infra_errors.NewCustomErrorMessage(fn(fieldName), fieldName))
	}

	return v
}

func (v *validatorAdapter) defaultFieldCompareValidation(f1, f2, fieldName string, fn func() error) validation.Validator {
	err := v.validator.VarWithValue(f1, f2, "eqfield")

	if err != nil {
		v.errors = append(v.errors, infra_errors.NewCustomErrorMessage(fn(), fieldName))
	}

	return v
}
