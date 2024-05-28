package user_entity

import protocol_entity "github.com/mariojuniortrab/hauling-api/internal/domain/entity/protocol"

type ListUserDto struct {
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
	protocol_entity.ListOutputDto
	Items []*UserDetailOutputDto `json:"items"`
}

func NewListUserDto(input *ListUserInputDto) (*ListUserDto, error) {
	willFilterActives := false
	active := false

	if input.Active != "" {
		willFilterActives = true
	}

	if input.Active == "true" {
		active = true
	}

	listUserDto := &ListUserDto{
		Active:            active,
		WillFilterActives: willFilterActives,
		ID:                input.ID,
		Name:              input.Name,
		Email:             input.Email,
	}

	err := protocol_entity.FillFromInput(&input.ListInputDto, &listUserDto.List)
	if err != nil {
		return nil, err
	}

	return listUserDto, nil
}
