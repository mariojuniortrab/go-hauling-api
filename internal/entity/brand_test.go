package entity_test

import (
	"testing"

	"github.com/mariojuniortrab/hauling-api/internal/entity"
	"github.com/stretchr/testify/assert"
)

func Test_If_Gets_AN_Error_If_Name_Is_Blank(t *testing.T) {
	brand := entity.Brand{}

	assert.ErrorContains(t, brand.Validate(), "name is required")
}

func Test_If_NewBrand_Return_An_Error_If_Name_Is_Blank(t *testing.T) {
	_, error := entity.NewBrand("")

	assert.ErrorContains(t, error, "name is required")
}

func Test_If_NewBrand_Return_An_ID(t *testing.T) {
	name := "any_name"
	brand, _ := entity.NewBrand(name)

	assert.NotNil(t, brand.ID)
}
func Test_If_NewBrand_Return_An_Correct_Brand(t *testing.T) {
	name := "any_name"
	brand, _ := entity.NewBrand(name)

	assert.NotNil(t, brand.ID)
	assert.Equal(t, name, brand.Name)
}
