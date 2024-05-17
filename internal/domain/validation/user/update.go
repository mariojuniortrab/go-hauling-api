package user_validation

import (
	"fmt"

	user_usecase "github.com/mariojuniortrab/hauling-api/internal/domain/usecase/user"
	util_usecase "github.com/mariojuniortrab/hauling-api/internal/domain/usecase/util"
	errors_validation "github.com/mariojuniortrab/hauling-api/internal/domain/validation/errors"
	protocol_validation "github.com/mariojuniortrab/hauling-api/internal/domain/validation/protocol"
)

type UpdateValidation interface {
	Validate(input *user_usecase.UpdateFields) (*errors_validation.CustomErrorMessage, []*errors_validation.CustomErrorMessage)
}
type updateValidation struct {
	validator protocol_validation.Validator
}

func NewUpdateValidation(validator protocol_validation.Validator) *updateValidation {
	return &updateValidation{
		validator,
	}
}

func (v *updateValidation) Validate(input *user_usecase.UpdateFields) (*errors_validation.CustomErrorMessage, []*errors_validation.CustomErrorMessage) {
	fmt.Println("[user_validation > updateValidation > Validate] input:", input)

	if v.isEmptyRequest(input) {
		return errors_validation.NewCustomErrorMessage(errors_validation.EmptyRequest(), ""), nil
	}

	v.validatePassword(input.Password)
	v.validateName(input.Name)
	v.validateBirth(input.Birth)
	v.validatePasswordConfirmation(input.PasswordConfirmation, input.Password)

	if v.validator.HasErrors() {
		return nil, v.validator.GetErrorsAndClean()
	}

	fmt.Println("[user_validation > updateValidation > Validate] success")
	return nil, nil
}

func (v *updateValidation) validatePassword(input string) {
	const fieldName = "password"

	v.validator.
		ValidateStringField(input, fieldName).
		ValidateFieldMaxLength(input, fieldName, 50).
		ValidateFieldMinLength(input, fieldName, 8)
}

func (v *updateValidation) validateName(input string) {
	const fieldName = "name"

	v.validator.
		ValidateFieldMaxLength(input, fieldName, 255)
}

func (v *updateValidation) validateBirth(input string) {
	const fieldName = "birth"

	_, err := util_usecase.GetDateFromString(input)
	if err != nil {
		v.validator.AddError(errors_validation.MustBeDateFormat(fieldName), fieldName)
	}
}

func (v *updateValidation) validatePasswordConfirmation(input, password string) {
	const fieldName = "passwordConfirmation"

	if input == "" && password == "" {
		return
	}

	v.validator.
		ValidateStringField(input, fieldName).
		ValidateFieldMaxLength(input, fieldName, 50).
		ValidatePasswordConfirmationEquals(input, password)
}

func (v *updateValidation) isEmptyRequest(fields *user_usecase.UpdateFields) bool {
	return (&user_usecase.UpdateFields{}) == fields
}
