package user_validation

import (
	"net/http"
	"time"

	user_entity "github.com/mariojuniortrab/hauling-api/internal/entity/user"
	infra_errors "github.com/mariojuniortrab/hauling-api/internal/infra/errors"
	user_usecase "github.com/mariojuniortrab/hauling-api/internal/usecase/user"
	"github.com/mariojuniortrab/hauling-api/internal/validation"
)

type signUpValidation struct {
	validator      validation.Validator
	userRepository user_entity.UserRepository
}

func NewSignUpValidation(validator validation.Validator, userRepository user_entity.UserRepository) *signUpValidation {
	return &signUpValidation{
		validator,
		userRepository,
	}
}

func (s *signUpValidation) Validate(input *user_usecase.SignupInputDto) *infra_errors.CustomError {
	s.validateEmail(input.Email)
	s.validatePassword(input.Password)
	s.validateName(input.Name)
	s.validateBirth(input.Birth)
	s.validatePasswordConfirmation(input.PasswordConfirmation, input.Password)

	if s.validator.HasErrors() {
		return s.validator.GetErrorsAndClean()
	}

	err := s.alreadyExists(input.Email, "")
	if err != nil {
		return err
	}

	return nil
}

func (s *signUpValidation) alreadyExists(email, id string) *infra_errors.CustomError {
	exists, err := s.userRepository.GetByEmail(email, id)

	if err != nil {
		return infra_errors.NewCustomError(err, http.StatusInternalServerError, "")
	}

	if exists != nil {
		return infra_errors.NewCustomError(infra_errors.AlreadyExists("user"), http.StatusConflict, "email")
	}

	return nil
}

func (s *signUpValidation) validateEmail(input string) {
	const fieldName = "email"

	s.validator.
		ValidateRequiredField(input, fieldName).
		ValidateFieldMaxLength(input, fieldName, 50).
		ValidateEmailField(input, fieldName)
}

func (s *signUpValidation) validatePassword(input string) {
	const fieldName = "password"

	s.validator.
		ValidateRequiredField(input, fieldName).
		ValidateFieldString(input, fieldName).
		ValidateFieldMaxLength(input, fieldName, 50)
}

func (s *signUpValidation) validateName(input string) {
	const fieldName = "name"

	s.validator.
		ValidateRequiredField(input, fieldName).
		ValidateFieldMaxLength(input, fieldName, 255)
}

func (s *signUpValidation) validateBirth(input string) {
	const fieldName = "birth"
	const shortForm = "2006-01-02"

	s.validator.ValidateRequiredField(input, fieldName)

	_, err := time.Parse(shortForm, input)
	if err != nil {
		s.validator.AddError(infra_errors.MustBeDateFormat(fieldName), fieldName)
	}
}

func (s *signUpValidation) validatePasswordConfirmation(input, password string) {
	const fieldName = "passwordConfirmation"

	s.validator.
		ValidateRequiredField(input, fieldName).
		ValidateFieldString(input, fieldName).
		ValidateFieldMaxLength(input, fieldName, 50).
		ValidatePasswordConfirmationEquals(input, password)
}
