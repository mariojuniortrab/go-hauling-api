package brand_usecase

import (
	brand_entity "github.com/mariojuniortrab/hauling-api/internal/domain/entity/brand"
	errors_validation "github.com/mariojuniortrab/hauling-api/internal/domain/validation/errors"
)

type CreateBrandValidation interface {
	Validate(createBrandInputDto CreateBrandInputDto) []*errors_validation.CustomErrorMessage
}

type CreateBrandInputDto struct {
	Name string `validate:"required" json:"name"`
}

type CreateBrandOutputDto struct {
	ID   string
	Name string
}

type CreateBrandUseCase struct {
	BrandRepository brand_entity.BrandRepository
}

func NewCreateBrandUseCase(brandRepository brand_entity.BrandRepository) *CreateBrandUseCase {
	return &CreateBrandUseCase{BrandRepository: brandRepository}
}

func (u *CreateBrandUseCase) Execute(input CreateBrandInputDto) (*CreateBrandOutputDto, error) {
	brand := brand_entity.NewBrand(input.Name)

	err := u.BrandRepository.Create(brand)
	if err != nil {
		return nil, err
	}

	return &CreateBrandOutputDto{
		ID:   brand.ID,
		Name: brand.Name,
	}, nil

}
