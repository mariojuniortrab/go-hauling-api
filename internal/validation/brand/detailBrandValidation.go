package brand_validation

import (
	"net/http"

	brand_entity "github.com/mariojuniortrab/hauling-api/internal/entity/brand"
	infra_errors "github.com/mariojuniortrab/hauling-api/internal/infra/errors"
	util_validation "github.com/mariojuniortrab/hauling-api/internal/validation/util"
)

type DetailBrandValidation struct {
	BrandRepository *brand_entity.BrandRepository
}

func NewDetailBrandValidation(brandRepository *brand_entity.BrandRepository) *DetailBrandValidation {
	return &DetailBrandValidation{
		BrandRepository: brandRepository,
	}
}

func (v *DetailBrandValidation) Validate(id string) *infra_errors.CustomError {
	if id == "" {
		return infra_errors.NewCustomError(infra_errors.IsRequired("id"), http.StatusBadRequest, "id")
	}

	if !util_validation.IsUIID(id) {
		return infra_errors.NewCustomError(infra_errors.MustBeUUID("id"), http.StatusBadRequest, "id")
	}

	return nil
}
