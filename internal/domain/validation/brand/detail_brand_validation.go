package brand_validation

import (
	brand_entity "github.com/mariojuniortrab/hauling-api/internal/domain/entity/brand"
	errors_validation "github.com/mariojuniortrab/hauling-api/internal/domain/validation/errors"
)

type DetailBrandValidation struct {
	BrandRepository *brand_entity.BrandRepository
}

func NewDetailBrandValidation(brandRepository *brand_entity.BrandRepository) *DetailBrandValidation {
	return &DetailBrandValidation{
		BrandRepository: brandRepository,
	}
}

func (v *DetailBrandValidation) Validate(id string) *errors_validation.CustomErrorMessage {
	if id == "" {
		return errors_validation.NewCustomErrorMessage(errors_validation.IsRequired("id"), "id")
	}

	return nil
}
