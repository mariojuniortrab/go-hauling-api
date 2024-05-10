package brand_validation

import (
	brand_entity "github.com/mariojuniortrab/hauling-api/internal/entity/brand"
	infra_errors "github.com/mariojuniortrab/hauling-api/internal/infra/errors"
	brand_usecase "github.com/mariojuniortrab/hauling-api/internal/usecase/brand"
)

type CreateBrandValidation struct {
	BrandRepository brand_entity.BrandRepository
}

func NewCreateBrandValidation(brandRepository brand_entity.BrandRepository) *CreateBrandValidation {
	return &CreateBrandValidation{BrandRepository: brandRepository}
}

func (v *CreateBrandValidation) Validate(input brand_usecase.CreateBrandInputDto) *infra_errors.CustomError {

	return nil
}
