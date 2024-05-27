package user_entity

import protocol_entity "github.com/mariojuniortrab/hauling-api/internal/domain/entity/protocol"

type ListUserParams struct {
	protocol_entity.List
	ID                string
	Name              string
	WillFilterActives bool
	Active            bool
	Email             string
}

type filter struct {
	ID     string
	Email  string
	Name   string
	Active string
}

type ListUserInputDto struct {
	protocol_entity.ListInputDto
	filter
}

type ListOutputDto struct {
	Items []*UserDetailOutputDto `json:"items"`
}

func NewListUserParams(input *ListUserInputDto) (*ListUserParams, error) {
	willFilterActives := false
	active := false

	if input.Active != "" {
		willFilterActives = true
	}

	if input.Active == "true" {
		active = true
	}

	listUserParams := &ListUserParams{
		Active:            active,
		WillFilterActives: willFilterActives,
		ID:                input.ID,
		Name:              input.Name,
		Email:             input.Email,
	}

	err := protocol_entity.FillFromInput(&input.ListInputDto, &listUserParams.List)
	if err != nil {
		return nil, err
	}

	return listUserParams, nil
}
