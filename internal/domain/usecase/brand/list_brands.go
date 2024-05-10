package brand_usecase

import brand_entity "github.com/mariojuniortrab/hauling-api/internal/domain/entity/brand"

type ListBrandsOutputDto struct {
	ID   string
	Name string
}

type ListBrandsSearchDto struct {
	ID    string `validate:"uuid4"`
	Name  string
	Page  int
	Limit int
	Term  string
}

type ListBrandUseCase struct {
	BrandRepository brand_entity.BrandRepository
}

func NewListBrandUseCase(brandRepository brand_entity.BrandRepository) *ListBrandUseCase {
	return &ListBrandUseCase{BrandRepository: brandRepository}
}

func (u *ListBrandUseCase) Execute() ([]*ListBrandsOutputDto, error) {
	brands, err := u.BrandRepository.ListAll()
	if err != nil {
		return nil, err
	}

	var brandsOutput []*ListBrandsOutputDto

	for _, brand := range brands {
		brandsOutput = append(brandsOutput, &ListBrandsOutputDto{
			ID:   brand.ID,
			Name: brand.Name,
		})
	}

	return brandsOutput, nil

}
