package auth_entity

import (
	user_entity "github.com/mariojuniortrab/hauling-api/internal/domain/entity/user"
	util_entity "github.com/mariojuniortrab/hauling-api/internal/domain/entity/util"
)

type SignupInputDto struct {
	Password             string `json:"password"`
	Name                 string `json:"name"`
	PasswordConfirmation string `json:"passwordConfirmation"`
	Email                string `json:"email"`
	Birth                string `json:"birth"`
}

type SignupOutputDto struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
	Birth string `json:"birth"`
}

func NewSignupOutputDto(user *user_entity.User) *SignupOutputDto {
	return &SignupOutputDto{
		Email: user.Email,
		Name:  user.Name,
		ID:    user.ID,
		Birth: util_entity.GetStringFromDate(user.Birth),
	}
}
