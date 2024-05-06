package brand_validation

import (
	"net/http"

	"github.com/go-playground/validator/v10"
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
	validate := validator.New()

	err := validate.Struct(input)
	if err != nil {
		var messages []*infra_errors.CustomErrorMessage

		for _, validationErr := range err.(validator.ValidationErrors) {
			adaptedError := infra_errors.ValidationErrorAdapter(validationErr)
			messages = append(messages, infra_errors.NewCustomErrorMessage(adaptedError, validationErr.Field()))
		}

		return infra_errors.NewCustomErrorArray(messages, http.StatusBadRequest)
	}

	exists, err := v.BrandRepository.GetByName(input.Name)

	if err != nil {
		return infra_errors.NewCustomError(err, http.StatusInternalServerError, "")
	}

	if exists != nil {
		return infra_errors.NewCustomError(infra_errors.AlreadyExists("brand"), http.StatusInternalServerError, "name")
	}

	return nil
}
