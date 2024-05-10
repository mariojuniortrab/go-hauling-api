package validation

import infra_errors "github.com/mariojuniortrab/hauling-api/internal/infra/errors"

type Validator interface {
	ValidateRequiredField(interface{}, string) Validator
	ValidateEmailField(interface{}, string) Validator
	ValidatePasswordConfirmationEquals(string, string) Validator
	ValidateFieldString(interface{}, string) Validator
	ValidateFieldLength(interface{}, string, int) Validator
	HasErrors() bool
	AddError(error, string) Validator
	GetErrors() *infra_errors.CustomError
}
