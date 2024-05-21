package user_validation

import (
	user_entity "github.com/mariojuniortrab/hauling-api/internal/domain/entity/user"
	errors_validation "github.com/mariojuniortrab/hauling-api/internal/domain/validation/errors"
	protocol_validation "github.com/mariojuniortrab/hauling-api/internal/domain/validation/protocol"
)

type ListValidation interface {
	Validate(input *user_entity.ListUserInputDto) []*errors_validation.CustomErrorMessage
}

type listValidation struct {
	validator protocol_validation.Validator
}

func NewListValidation(validator protocol_validation.Validator) *listValidation {
	return &listValidation{validator}
}

func (v *listValidation) Validate(input *user_entity.ListUserInputDto) []*errors_validation.CustomErrorMessage {

	v.validateEmail(input.Email)
	v.validateName(input.Name)
	v.validateActive(input.Active)

	if v.validator.HasErrors() {
		return v.validator.GetErrorsAndClean()
	}

	return nil
}

func (v *listValidation) validateEmail(input string) {
	const fieldName = "email"

	v.validator.
		ValidateFieldMaxLength(input, fieldName, 50).
		ValidateEmailField(input, fieldName)
}

func (v *listValidation) validateName(input string) {
	const fieldName = "name"

	v.validator.
		ValidateFieldMaxLength(input, fieldName, 255)
}

func (v *listValidation) validateActive(input string) {
	const fieldName = "active"

	v.validator.ValidateStringBooleanField(input, fieldName)
}
