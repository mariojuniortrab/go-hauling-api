package user_routes

import (
	user_entity "github.com/mariojuniortrab/hauling-api/internal/entity/user"
	user_handler "github.com/mariojuniortrab/hauling-api/internal/infra/web/handlers/user"
	"github.com/mariojuniortrab/hauling-api/internal/infra/web/routes"
	user_usecase "github.com/mariojuniortrab/hauling-api/internal/usecase/user"
	"github.com/mariojuniortrab/hauling-api/internal/validation"
	user_validation "github.com/mariojuniortrab/hauling-api/internal/validation/user"
)

type router struct {
	userRepository user_entity.UserRepository
	validator      validation.Validator
}

func NewRouter(UserRepository user_entity.UserRepository) *router {
	return &router{
		userRepository: UserRepository,
	}
}

func (r *router) Route(route routes.Router) routes.Router {
	signUpValidation := user_validation.NewSignUpValidator(r.validator, r.userRepository)
	signUpUseCase := user_usecase.NewSignUpUseCase(r.userRepository)
	signupHandler := user_handler.NewSignupHandler(signUpValidation, signUpUseCase)

	route.Post("signup", signupHandler.Handle)

	return route
}
