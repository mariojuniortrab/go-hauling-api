package user_entity

import util_entity "github.com/mariojuniortrab/hauling-api/internal/domain/entity/util"

type UserDetailOutputDto struct {
	ID     string `json:"id"`
	Email  string `json:"email"`
	Name   string `json:"name"`
	Birth  string `json:"birth"`
	Active bool   `json:"active"`
}

func NewUserDetailOutputDto(user *User) *UserDetailOutputDto {
	birth := util_entity.GetStringFromDate(user.Birth)

	return &UserDetailOutputDto{
		ID:     user.ID,
		Name:   user.Name,
		Email:  user.Email,
		Birth:  birth,
		Active: user.Active,
	}
}
