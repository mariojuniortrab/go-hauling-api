package user_usecase

import (
	"time"

	user_entity "github.com/mariojuniortrab/hauling-api/internal/entity/user"
	infra_errors "github.com/mariojuniortrab/hauling-api/internal/infra/errors"
	util_usecase "github.com/mariojuniortrab/hauling-api/internal/usecase/util"
)

type SignupValidation interface {
	Validate(*SignupInputDto) *infra_errors.CustomError
}

type SignupInputDto struct {
	Password             string `validate:"required,eqfield=PasswordConfirmation" json:"password"`
	Name                 string `validate:"required" json:"name"`
	PasswordConfirmation string `validate:"required" json:"password_confirmation"`
	Email                string `validate:"required,email" json:"email"`
	Birth                string `validate:"required" json:"birth"`
}

type signupOutputDto struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

type Signup struct {
	userRepository user_entity.UserRepository
	encrypter      util_usecase.Encrypter
}

func NewSignupUseCase(userRepository user_entity.UserRepository,
	encrypter util_usecase.Encrypter) *Signup {
	return &Signup{
		userRepository,
		encrypter,
	}
}

func (s *Signup) Execute(input SignupInputDto) (*signupOutputDto, error) {
	formattedDate, err := getFormattedDate(input.Birth)
	if err != nil {
		return nil, err
	}

	hashPassword, err := s.encrypter.Hash(input.Password)
	if err != nil {
		return nil, err
	}

	user := user_entity.NewUser(input.Name, hashPassword, input.Email, formattedDate)

	err = s.userRepository.Create(user)
	if err != nil {
		return nil, err
	}

	return &signupOutputDto{
		Email: user.Email,
		Name:  user.Name,
		ID:    user.ID,
	}, nil
}

func getFormattedDate(date string) (time.Time, error) {
	const shortForm = "2006-Jan-02"
	result, err := time.Parse(shortForm, date)

	return result, err
}
