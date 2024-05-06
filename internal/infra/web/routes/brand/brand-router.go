package brand_routes

import (
	brand_entity "github.com/mariojuniortrab/hauling-api/internal/entity/brand"
	brand_handler "github.com/mariojuniortrab/hauling-api/internal/infra/web/handlers/brand"
	"github.com/mariojuniortrab/hauling-api/internal/infra/web/routes"
	brand_usecase "github.com/mariojuniortrab/hauling-api/internal/usecase/brand"
	brand_validation "github.com/mariojuniortrab/hauling-api/internal/validation/brand"
)

type router struct {
	brandRepository brand_entity.BrandRepository
}

func NewRouter(brandRepository brand_entity.BrandRepository) *router {
	return &router{
		brandRepository: brandRepository,
	}
}

func (r *router) Route(route routes.Router) routes.Router {
	createBrandUseCase := brand_usecase.NewCreateBrandUseCase(r.brandRepository)
	listBrandUseCase := brand_usecase.NewListBrandUseCase(r.brandRepository)
	createBrandValidation := brand_validation.NewCreateBrandValidation(r.brandRepository)

	createBrandHandler := brand_handler.NewCreateBrandHandler(createBrandUseCase, createBrandValidation)
	listBrandHandler := brand_handler.NewListBrandHandler(listBrandUseCase)

	route.Post("/brands", createBrandHandler.Handle)
	route.Get("/brands", listBrandHandler.Handle)

	return route
}
