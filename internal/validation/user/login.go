package user_validation

import (
	"net/http"

	infra_errors "github.com/mariojuniortrab/hauling-api/internal/infra/errors"
	user_usecase "github.com/mariojuniortrab/hauling-api/internal/usecase/user"
	util_usecase "github.com/mariojuniortrab/hauling-api/internal/usecase/util"
	"github.com/mariojuniortrab/hauling-api/internal/validation"
)

type loginValidation struct {
	validator validation.Validator
	encrypter util_usecase.Encrypter
}

func NewLoginValidation(validator validation.Validator,
	encrypter util_usecase.Encrypter) *loginValidation {
	return &loginValidation{validator, encrypter}
}

func (v *loginValidation) Validate(input *user_usecase.LoginInputDto) *infra_errors.CustomError {

	return nil
}

func (v *loginValidation) ValidateCredentials(input *user_usecase.UserDto, password string) *infra_errors.CustomError {
	passwordIsValid := v.encrypter.CheckPasswordHash(input.Password, password)

	if input.ID == "" || !passwordIsValid {
		return infra_errors.NewCustomError(infra_errors.WrongPassword(), http.StatusBadRequest, "password")
	}

	return nil
}
