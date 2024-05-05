package model_entity_test

import (
	"testing"

	brand_entity "github.com/mariojuniortrab/hauling-api/internal/entity/brand"
	model_entity "github.com/mariojuniortrab/hauling-api/internal/entity/model"
	"github.com/stretchr/testify/assert"
)

func Test_If_NewModel_Return_An_ID(t *testing.T) {
	name := "any_name"
	brand := brand_entity.NewBrand("any_brand")
	model, _ := model_entity.NewModel(name, brand)

	assert.NotNil(t, model.ID)
}

func Test_If_NewModel_Return_An_Error_If_Brand_Is_Nil(t *testing.T) {
	name := "any_name"
	errMessage := "brand is required"

	_, err := model_entity.NewModel(name, nil)
	assert.EqualError(t, err, errMessage)

}

func Test_if_NewModel_Return_Valid_Brand(t *testing.T) {
	name := "any_name"
	brand := brand_entity.NewBrand("any_brand")
	model, _ := model_entity.NewModel(name, brand)

	assert.Equal(t, model.Brand.ID, brand.ID)
}
