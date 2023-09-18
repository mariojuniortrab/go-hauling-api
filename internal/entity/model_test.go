package entity_test

import (
	"testing"

	"github.com/mariojuniortrab/hauling-api/internal/entity"
	"github.com/stretchr/testify/assert"
)

func Test_If_Model_Gets_AN_Error_If_Name_Is_Blank(t *testing.T) {
	model := entity.Model{}

	assert.ErrorContains(t, model.Validate(), "name is required")
}
func Test_If_Model_Gets_AN_Error_If_Brand_Is_Blank(t *testing.T) {
	model := entity.Model{Name: "any_name"}

	assert.ErrorContains(t, model.Validate(), "brand is required")
}

func Test_If_NewModel_Return_An_Error_If_Name_Is_Blank(t *testing.T) {
	_, error := entity.NewBrand("")

	assert.ErrorContains(t, error, "name is required")
}

func Test_If_NewModel_Return_An_Error_If_Brand_Is_Blank(t *testing.T) {
	_, error := entity.NewBrand("")

	assert.ErrorContains(t, error, "name is required")
}

func Test_If_NewModel_Return_An_ID(t *testing.T) {
	name := "any_name"
	brand, _ := entity.NewBrand(name)

	assert.NotNil(t, brand.ID)
}
func Test_If_NewModel_Return_An_Correct_Mode(t *testing.T) {
	name := "any_name"
	brand, _ := entity.NewBrand(name)

	assert.NotNil(t, brand.ID)
	assert.Equal(t, name, brand.Name)
}
