package entity_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_If_Gets_AN_Error_If_Name_Is_Blank(t *testing.T) {
	brand := Brand{}

	assert.Error(t, brand.Validate(), "Name is Required")
}
