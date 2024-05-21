package user_validation

import (
	auth_entity "github.com/mariojuniortrab/hauling-api/internal/domain/entity/auth"
	util_entity "github.com/mariojuniortrab/hauling-api/internal/domain/entity/util"
	protocol_usecase "github.com/mariojuniortrab/hauling-api/internal/domain/usecase/protocol"
	errors_validation "github.com/mariojuniortrab/hauling-api/internal/domain/validation/errors"
	protocol_validation "github.com/mariojuniortrab/hauling-api/internal/domain/validation/protocol"
)

type SignupValidation interface {
	Validate(input *auth_entity.SignupInputDto) []*errors_validation.CustomErrorMessage
	AlreadyExists(email, id string) (*errors_validation.CustomErrorMessage, error)
}
type signUpValidation struct {
	validator      protocol_validation.Validator
	userRepository protocol_usecase.UserRepository
}

func NewSignUpValidation(validator protocol_validation.Validator, signupRepository protocol_usecase.UserRepository) *signUpValidation {
	return &signUpValidation{
		validator,
		signupRepository,
	}
}

func (v *signUpValidation) Validate(input *auth_entity.SignupInputDto) []*errors_validation.CustomErrorMessage {
	v.validateEmail(input.Email)
	v.validatePassword(input.Password)
	v.validateName(input.Name)
	v.validateBirth(input.Birth)
	v.validatePasswordConfirmation(input.PasswordConfirmation, input.Password)

	if v.validator.HasErrors() {
		return v.validator.GetErrorsAndClean()
	}

	return nil
}

func (v *signUpValidation) AlreadyExists(email, id string) (*errors_validation.CustomErrorMessage, error) {
	exists, err := v.userRepository.GetByEmail(email, id)

	if err != nil {
		return nil, err
	}

	if exists != nil {
		return errors_validation.NewCustomErrorMessage(errors_validation.AlreadyExists("user"), "email"), nil
	}

	return nil, nil
}

func (v *signUpValidation) validateEmail(input string) {
	const fieldName = "email"

	v.validator.
		ValidateRequiredField(input, fieldName).
		ValidateFieldMaxLength(input, fieldName, 50).
		ValidateEmailField(input, fieldName)
}

func (v *signUpValidation) validatePassword(input string) {
	const fieldName = "password"

	v.validator.
		ValidateRequiredField(input, fieldName).
		ValidateStringField(input, fieldName).
		ValidateFieldMaxLength(input, fieldName, 50).
		ValidateFieldMinLength(input, fieldName, 8)
}

func (v *signUpValidation) validateName(input string) {
	const fieldName = "name"

	v.validator.
		ValidateRequiredField(input, fieldName).
		ValidateFieldMaxLength(input, fieldName, 255)
}

func (v *signUpValidation) validateBirth(input string) {
	const fieldName = "birth"

	v.validator.ValidateRequiredField(input, fieldName)

	_, err := util_entity.GetDateFromString(input)
	if err != nil {
		v.validator.AddError(errors_validation.MustBeDateFormat(fieldName), fieldName)
	}
}

func (v *signUpValidation) validatePasswordConfirmation(input, password string) {
	const fieldName = "passwordConfirmation"

	v.validator.
		ValidateRequiredField(input, fieldName).
		ValidateStringField(input, fieldName).
		ValidateFieldMaxLength(input, fieldName, 50).
		ValidatePasswordConfirmationEquals(input, password)
}
