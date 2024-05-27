package brand_entity

import (
	protocol_entity "github.com/mariojuniortrab/hauling-api/internal/domain/entity/protocol"
)

type CreateBrandInputDto struct {
	Name *string `json:"name"`
}

func (u CreateBrandInputDto) IsEmpty() bool {
	return u == CreateBrandInputDto{}
}

func (u *CreateBrandInputDto) New() protocol_entity.Emptyable {
	return &CreateBrandInputDto{}
}

type CreateBrandOutputDto struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func NewCreateBrandOutputDto(brand *Brand) *CreateBrandOutputDto {
	return &CreateBrandOutputDto{
		Name: brand.Name,
		ID:   brand.ID,
	}
}
