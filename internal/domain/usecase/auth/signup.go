package auth_usecase

import (
	auth_entity "github.com/mariojuniortrab/hauling-api/internal/domain/entity/auth"
	user_entity "github.com/mariojuniortrab/hauling-api/internal/domain/entity/user"
	util_entity "github.com/mariojuniortrab/hauling-api/internal/domain/entity/util"
	protocol_usecase "github.com/mariojuniortrab/hauling-api/internal/domain/usecase/protocol"
)

type Signup struct {
	userRepository protocol_usecase.UserRepository
	encrypter      protocol_usecase.Encrypter
}

func NewSignupUseCase(userRepository protocol_usecase.UserRepository,
	encrypter protocol_usecase.Encrypter) *Signup {
	return &Signup{userRepository, encrypter}
}

func (s *Signup) Execute(input *auth_entity.SignupInputDto) (*auth_entity.SignupOutputDto, error) {

	formattedDate, err := util_entity.GetDateFromString(input.Birth)
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

	return auth_entity.NewSignupOutputDto(user), nil
}
