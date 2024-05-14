package user_usecase

import (
	"fmt"

	user_entity "github.com/mariojuniortrab/hauling-api/internal/domain/entity/user"
	protocol_usecase "github.com/mariojuniortrab/hauling-api/internal/domain/usecase/protocol"
)

type filter struct {
	ID     string `json:"id"`
	Email  string `json:"email"`
	Name   string `json:"name"`
	Active string `json:"active"`
}

type ListUserInputDto struct {
	protocol_usecase.List
	filter
}

type listItemOutputDto struct {
	ID     string `json:"id"`
	Email  string `json:"email"`
	Name   string `json:"name"`
	Birth  string `json:"birth"`
	Active bool   `json:"active"`
}

type listOutputDto struct {
	Items []*listItemOutputDto `json:"items"`
}

type List struct {
	userRepository protocol_usecase.UserRepository
}

func NewListUseCase(userRepository protocol_usecase.UserRepository) *List {
	return &List{userRepository}
}

func newListUserParams(input *ListUserInputDto) *user_entity.ListUserParams {
	willFilterActives := false
	active := false

	if input.Active != "" {
		willFilterActives = true
	}

	if input.Active == "true" {
		active = true
	}

	listUserParams := &user_entity.ListUserParams{
		Active:            active,
		WillFilterActives: willFilterActives,
		ID:                input.ID,
		Name:              input.Name,
		Email:             input.Email,
	}

	listUserParams.Page = int(input.Page)
	listUserParams.Limit = int(input.Limit)
	listUserParams.OrderBy = input.OrderBy
	listUserParams.OrderType = input.OrderType
	listUserParams.Q = input.Q

	return listUserParams
}

func NewListItemOutputDto(user *user_entity.User) *listItemOutputDto {
	const shortForm = "2006-01-02"
	birth := user.Birth.Format(shortForm)

	return &listItemOutputDto{
		ID:     user.ID,
		Name:   user.Name,
		Email:  user.Email,
		Birth:  birth,
		Active: user.Active,
	}
}

func (u *List) Execute(input *ListUserInputDto) (*listOutputDto, error) {
	fmt.Println("[user_usecase > List > Execute] input:", input)
	listUserParams := newListUserParams(input)

	users, err := u.userRepository.List(listUserParams)
	if err != nil {
		fmt.Println("[user_usecase > List > Execute] err:", err)
		return nil, err
	}

	result := &listOutputDto{}

	for _, user := range users {
		result.Items = append(result.Items, NewListItemOutputDto(user))
	}

	fmt.Println("[user_usecase > List > Execute] success")
	return result, nil
}
