package brand_validation

import (
	brand_entity "github.com/mariojuniortrab/hauling-api/internal/domain/entity/brand"
	brand_usecase "github.com/mariojuniortrab/hauling-api/internal/domain/usecase/brand"
	errors_validation "github.com/mariojuniortrab/hauling-api/internal/domain/validation/errors"
)

type CreateBrandValidation struct {
	BrandRepository brand_entity.BrandRepository
}

func NewCreateBrandValidation(brandRepository brand_entity.BrandRepository) *CreateBrandValidation {
	return &CreateBrandValidation{BrandRepository: brandRepository}
}

func (v *CreateBrandValidation) Validate(input brand_usecase.CreateBrandInputDto) []*errors_validation.CustomErrorMessage {

	return nil
}
