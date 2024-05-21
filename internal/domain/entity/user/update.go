package user_entity

import (
	"strings"

	util_entity "github.com/mariojuniortrab/hauling-api/internal/domain/entity/util"
)

type UserUpdateInputDto struct {
	Name                 string `json:"name"`
	Birth                string `json:"birth"`
	Active               string `json:"active"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"passwordConfirmation"`
}

func (u UserUpdateInputDto) IsEmpty() bool {
	return u == UserUpdateInputDto{}
}

func NewUserUpdateInputDto(user *User) *UserUpdateInputDto {
	active := "true"

	if !user.Active {
		active = "false"
	}

	return &UserUpdateInputDto{
		Name:   user.Name,
		Birth:  util_entity.GetStringFromDate(user.Birth),
		Active: active,
	}
}

func NewUserFromUpdateInputDto(input *UserUpdateInputDto) (*User, error) {
	formattedDate, err := util_entity.GetDateFromString(input.Birth)
	if err != nil {
		return nil, err
	}

	active := strings.ToLower(input.Active) == "true"

	return &User{
		Name:     input.Name,
		Password: input.Password,
		Active:   active,
		Birth:    formattedDate,
	}, nil
}
