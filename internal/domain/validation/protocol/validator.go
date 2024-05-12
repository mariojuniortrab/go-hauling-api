package protocol_validation

import (
	errors_validation "github.com/mariojuniortrab/hauling-api/internal/domain/validation/errors"
)

type Validator interface {
	ValidateRequiredField(interface{}, string) Validator
	ValidateEmailField(interface{}, string) Validator
	ValidatePasswordConfirmationEquals(string, string) Validator
	ValidateFieldString(interface{}, string) Validator
	ValidateFieldLength(interface{}, string, int) Validator
	ValidateFieldMaxLength(interface{}, string, int) Validator
	ValidateFieldMinLength(interface{}, string, int) Validator
	HasErrors() bool
	AddError(error, string) Validator
	GetErrors() []*errors_validation.CustomErrorMessage
	GetErrorsAndClean() []*errors_validation.CustomErrorMessage
}
