package brand_entity_test

import (
	"testing"

	brand_entity "github.com/mariojuniortrab/hauling-api/internal/entity/brand"
	"github.com/stretchr/testify/assert"
)

func Test_If_NewBrand_Return_An_ID(t *testing.T) {
	name := "any_name"
	brand := brand_entity.NewBrand(name)

	assert.NotNil(t, brand.ID)
}
