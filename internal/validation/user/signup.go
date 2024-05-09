package user_validation

import (
	"net/http"

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
	err := s.validator.Validate(input)
	if err != nil {
		return err
	}

	err = s.passwordConfirmationIsInvalid(input.Password, input.PasswordConfirmation)
	if err != nil {
		return err
	}

	err = s.alreadyExists(input.Name)
	if err != nil {
		return err
	}

	return nil
}

func (s *signUpValidation) alreadyExists(email string) *infra_errors.CustomError {
	exists, err := s.userRepository.GetByEmail(email)

	if err != nil {
		return infra_errors.NewCustomError(err, http.StatusInternalServerError, "")
	}

	if exists != nil {
		return infra_errors.NewCustomError(infra_errors.AlreadyExists("user"), http.StatusBadRequest, "email")
	}

	return nil
}

func (s *signUpValidation) passwordConfirmationIsInvalid(password, passwordConfirmation string) *infra_errors.CustomError {
	if password != passwordConfirmation {
		return infra_errors.NewCustomError(infra_errors.MustMatch("password_confirmation", "password"), http.StatusBadRequest, "name")
	}

	return nil
}
