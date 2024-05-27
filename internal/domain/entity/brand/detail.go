package brand_entity

type DetailBrandOutputDto struct {
	ID     string `json:"id"`
	Email  string `json:"email"`
	Name   string `json:"name"`
	Birth  string `json:"birth"`
	Active bool   `json:"active"`
}

func NewUserDetailOutputDto(brand *Brand) *DetailBrandOutputDto {
	return &DetailBrandOutputDto{
		ID:   brand.ID,
		Name: brand.Name,
	}
}
