package infra_adapters

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	errors_validation "github.com/mariojuniortrab/hauling-api/internal/domain/validation/errors"
	protocol_validation "github.com/mariojuniortrab/hauling-api/internal/domain/validation/protocol"
)

type validatorAdapter struct {
	validator *validator.Validate
	errors    []*errors_validation.CustomErrorMessage
}

func NewValidator() *validatorAdapter {
	return &validatorAdapter{
		validator: validator.New(validator.WithRequiredStructEnabled()),
	}
}

func (v *validatorAdapter) HasErrors() bool {
	return len(v.errors) > 0
}

func (v *validatorAdapter) AddError(err error, fieldName string) protocol_validation.Validator {
	v.errors = append(v.errors, errors_validation.NewCustomErrorMessage(err, fieldName))
	return v
}

func (v *validatorAdapter) GetErrors() []*errors_validation.CustomErrorMessage {
	return v.errors
}

func (v *validatorAdapter) GetErrorsAndClean() []*errors_validation.CustomErrorMessage {
	errors := v.errors
	v.errors = []*errors_validation.CustomErrorMessage{}

	return errors
}

func (v *validatorAdapter) ValidateRequiredField(f interface{}, fieldName string) protocol_validation.Validator {
	return v.defaultValidation(f, fieldName, "required", errors_validation.IsRequired)
}

func (v *validatorAdapter) ValidateEmailField(f interface{}, fieldName string) protocol_validation.Validator {
	return v.defaultValidation(f, fieldName, "omitempty,email", errors_validation.IsInvalid)
}

func (v *validatorAdapter) ValidatePasswordConfirmationEquals(password, passwordConfirmation string) protocol_validation.Validator {
	fn := errors_validation.IsPasswordConfirmationInvalid
	return v.defaultFieldCompareValidation(password, passwordConfirmation, "passwordConfirmation", fn)
}

func (v *validatorAdapter) ValidateFieldString(f interface{}, fieldName string) protocol_validation.Validator {
	return v.defaultValidation(f, fieldName, "omitempty,alphaunicode", errors_validation.MustBeString)
}

func (v *validatorAdapter) ValidateFieldLength(f interface{}, fieldName string, length int) protocol_validation.Validator {
	return v.defaultLenghValidation(f, fieldName, fmt.Sprintf("omitempty,len=%d", length), errors_validation.LengthMustBe(fieldName, length))
}

func (v *validatorAdapter) ValidateFieldMaxLength(f interface{}, fieldName string, length int) protocol_validation.Validator {
	return v.defaultLenghValidation(f, fieldName, fmt.Sprintf("omitempty,max=%d", length), errors_validation.LengthMustBeOrLess(fieldName, length))
}

func (v *validatorAdapter) defaultLenghValidation(f interface{}, fieldName, flag string, errMessage error) protocol_validation.Validator {
	err := v.validator.Var(f, flag)

	if err != nil {
		v.errors = append(v.errors, errors_validation.NewCustomErrorMessage(errMessage, fieldName))
	}

	return v
}

func (v *validatorAdapter) defaultValidation(f interface{}, fieldName, flag string, fn func(string) error) protocol_validation.Validator {
	err := v.validator.Var(f, flag)

	if err != nil {
		v.errors = append(v.errors, errors_validation.NewCustomErrorMessage(fn(fieldName), fieldName))
	}

	return v
}

func (v *validatorAdapter) defaultFieldCompareValidation(f1, f2, fieldName string, fn func() error) protocol_validation.Validator {
	err := v.validator.VarWithValue(f1, f2, "eqfield")

	if err != nil {
		v.errors = append(v.errors, errors_validation.NewCustomErrorMessage(fn(), fieldName))
	}

	return v
}
