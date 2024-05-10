package model_entity

import (
	"errors"

	"github.com/google/uuid"
	brand_entity "github.com/mariojuniortrab/hauling-api/internal/domain/entity/brand"
)

type Model struct {
	ID    string
	Name  string
	Brand *brand_entity.Brand
}

func NewModel(name string, brand *brand_entity.Brand) (*Model, error) {
	model := Model{
		ID:   uuid.New().String(),
		Name: name,
	}

	if brand == nil {
		return nil, errors.New("brand is required")
	}

	model.Brand = brand

	return &model, nil
}
