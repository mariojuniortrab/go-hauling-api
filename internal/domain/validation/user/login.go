package user_validation

import (
	"fmt"

	protocol_usecase "github.com/mariojuniortrab/hauling-api/internal/domain/usecase/protocol"
	user_usecase "github.com/mariojuniortrab/hauling-api/internal/domain/usecase/user"
	errors_validation "github.com/mariojuniortrab/hauling-api/internal/domain/validation/errors"
	protocol_validation "github.com/mariojuniortrab/hauling-api/internal/domain/validation/protocol"
)

type LoginValidation interface {
	Validate(input *user_usecase.LoginInputDto) []*errors_validation.CustomErrorMessage
	IsCredentialInvalid(input *user_usecase.UserDto, password string) bool
}

type loginValidation struct {
	validator protocol_validation.Validator
	encrypter protocol_usecase.Encrypter
}

func NewLoginValidation(validator protocol_validation.Validator,
	encrypter protocol_usecase.Encrypter) *loginValidation {
	return &loginValidation{validator, encrypter}
}

func (v *loginValidation) Validate(input *user_usecase.LoginInputDto) []*errors_validation.CustomErrorMessage {
	fmt.Println("[user_validation > loginValidation > Validate] input:", input)

	v.validateEmail(input.Email)
	v.validatePassword(input.Password)

	if v.validator.HasErrors() {
		return v.validator.GetErrorsAndClean()
	}

	fmt.Println("[user_validation > loginValidation > Validate] success")
	return nil
}

func (v *loginValidation) IsCredentialInvalid(input *user_usecase.UserDto, password string) bool {
	inactive := !input.Active
	return !v.encrypter.CheckPasswordHash(input.Password, password) || inactive
}

func (v *loginValidation) validateEmail(input string) {
	const fieldName = "email"

	v.validator.
		ValidateRequiredField(input, fieldName).
		ValidateFieldMaxLength(input, fieldName, 50).
		ValidateEmailField(input, fieldName)
}

func (v *loginValidation) validatePassword(input string) {
	const fieldName = "password"

	v.validator.
		ValidateRequiredField(input, fieldName).
		ValidateStringField(input, fieldName).
		ValidateFieldMaxLength(input, fieldName, 50).
		ValidateFieldMinLength(input, fieldName, 8)
}
