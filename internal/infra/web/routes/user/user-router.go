package user_routes

import (
	user_entity "github.com/mariojuniortrab/hauling-api/internal/entity/user"
	handlers_protocols "github.com/mariojuniortrab/hauling-api/internal/infra/web/handlers/protocols"
	user_handler "github.com/mariojuniortrab/hauling-api/internal/infra/web/handlers/user"
	"github.com/mariojuniortrab/hauling-api/internal/infra/web/routes"
	user_usecase "github.com/mariojuniortrab/hauling-api/internal/usecase/user"
	util_usecase "github.com/mariojuniortrab/hauling-api/internal/usecase/util"
	"github.com/mariojuniortrab/hauling-api/internal/validation"
	user_validation "github.com/mariojuniortrab/hauling-api/internal/validation/user"
)

type router struct {
	userRepository user_entity.UserRepository
	validator      validation.Validator
	encrypter      util_usecase.Encrypter
	tokenizer      util_usecase.Tokenizer
}

func NewRouter(userRepository user_entity.UserRepository,
	validator validation.Validator,
	encrypter util_usecase.Encrypter,
	tokenizer util_usecase.Tokenizer) *router {
	return &router{
		userRepository,
		validator,
		encrypter,
		tokenizer,
	}
}

func (r *router) Route(route routes.Router) routes.Router {
	signupHandler := r.getSignupHandler()
	loginHandler := r.getLoginHandler()

	route.Post("signup", signupHandler.Handle)
	route.Post("login", loginHandler.Handle)

	return route
}

func (r *router) getSignupHandler() handlers_protocols.Handle {
	signUpValidation := user_validation.NewSignUpValidation(r.validator, r.userRepository)
	signUpUseCase := user_usecase.NewSignupUseCase(r.userRepository, r.encrypter)
	signupHandler := user_handler.NewSignupHandler(signUpValidation, signUpUseCase)

	return signupHandler
}

func (r *router) getLoginHandler() handlers_protocols.Handle {
	loginValidation := user_validation.NewLoginValidation(r.validator, r.encrypter)
	loginUseCase := user_usecase.NewLoginUseCase(r.userRepository, r.tokenizer)
	signupHandler := user_handler.NewLoginHandle(loginValidation, loginUseCase)

	return signupHandler
}
