package brand_usecase

import (
	brand_entity "github.com/mariojuniortrab/hauling-api/internal/entity/brand"
	infra_errors "github.com/mariojuniortrab/hauling-api/internal/infra/errors"
)

type CreateBrandValidation interface {
	Validate(createBrandInputDto CreateBrandInputDto) *infra_errors.CustomError
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
