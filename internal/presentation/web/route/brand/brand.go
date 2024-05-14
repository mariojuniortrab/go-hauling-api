package brand_routes

import (
	"fmt"

	brand_entity "github.com/mariojuniortrab/hauling-api/internal/domain/entity/brand"
	brand_usecase "github.com/mariojuniortrab/hauling-api/internal/domain/usecase/brand"
	brand_validation "github.com/mariojuniortrab/hauling-api/internal/domain/validation/brand"
	brand_handler "github.com/mariojuniortrab/hauling-api/internal/presentation/web/handler/brand"
	web_protocol "github.com/mariojuniortrab/hauling-api/internal/presentation/web/protocol"
)

type router struct {
	brandRepository brand_entity.BrandRepository
}

func NewRouter(brandRepository brand_entity.BrandRepository) *router {
	return &router{
		brandRepository: brandRepository,
	}
}

func (r *router) Route(route web_protocol.Router) web_protocol.Router {
	createBrandUseCase := brand_usecase.NewCreateBrandUseCase(r.brandRepository)
	listBrandUseCase := brand_usecase.NewListBrandUseCase(r.brandRepository)
	createBrandValidation := brand_validation.NewCreateBrandValidation(r.brandRepository)

	createBrandHandler := brand_handler.NewCreateBrandHandler(createBrandUseCase, createBrandValidation)
	listBrandHandler := brand_handler.NewListBrandHandler(listBrandUseCase)

	route.Post("/brands", createBrandHandler.Handle)
	route.Get("/brands", listBrandHandler.Handle)

	fmt.Println("[brand_routes > router > Route]")

	return route
}
