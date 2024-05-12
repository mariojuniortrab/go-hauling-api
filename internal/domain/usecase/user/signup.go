package user_usecase

import (
	"time"

	user_entity "github.com/mariojuniortrab/hauling-api/internal/domain/entity/user"
	protocol_usecase "github.com/mariojuniortrab/hauling-api/internal/domain/usecase/protocol"
)

type SignupInputDto struct {
	Password             string `json:"password"`
	Name                 string `json:"name"`
	PasswordConfirmation string `json:"passwordConfirmation"`
	Email                string `json:"email"`
	Birth                string `json:"birth"`
}

type signupOutputDto struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
	Birth string `json:"birth"`
}

type Signup struct {
	userRepository user_entity.UserRepository
	encrypter      protocol_usecase.Encrypter
}

func NewSignupUseCase(userRepository user_entity.UserRepository,
	encrypter protocol_usecase.Encrypter) *Signup {
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
		Birth: getStringDate(user.Birth),
	}, nil
}

func getFormattedDate(date string) (time.Time, error) {
	const shortForm = "2006-01-02"

	result, err := time.Parse(shortForm, date)

	return result, err
}

func getStringDate(date time.Time) string {
	const shortForm = "2006-01-02"
	return date.Format(shortForm)
}
