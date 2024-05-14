package user_routes

import (
	"fmt"

	protocol_usecase "github.com/mariojuniortrab/hauling-api/internal/domain/usecase/protocol"
	user_usecase "github.com/mariojuniortrab/hauling-api/internal/domain/usecase/user"
	protocol_validation "github.com/mariojuniortrab/hauling-api/internal/domain/validation/protocol"
	user_validation "github.com/mariojuniortrab/hauling-api/internal/domain/validation/user"
	user_handler "github.com/mariojuniortrab/hauling-api/internal/presentation/web/handler/user"
	web_middleware "github.com/mariojuniortrab/hauling-api/internal/presentation/web/middleware"
	web_protocol "github.com/mariojuniortrab/hauling-api/internal/presentation/web/protocol"
)

type router struct {
	userRepository protocol_usecase.UserRepository
	validator      protocol_validation.Validator
	encrypter      protocol_usecase.Encrypter
	tokenizer      protocol_usecase.Tokenizer
}

func NewRouter(userRepository protocol_usecase.UserRepository,
	validator protocol_validation.Validator,
	encrypter protocol_usecase.Encrypter,
	tokenizer protocol_usecase.Tokenizer) *router {
	return &router{
		userRepository,
		validator,
		encrypter,
		tokenizer,
	}
}

func (r *router) Route(route web_protocol.Router) web_protocol.Router {
	authUseCase := user_usecase.NewAuthorization(r.tokenizer)
	protected := web_middleware.NewProtectedMiddleware(r.tokenizer, authUseCase)
	list := web_middleware.NewListMiddleware(r.validator)

	route.Group(func(rr web_protocol.Router) {
		rr.Use(protected.GetMiddleware())
		rr.Use(list.GetMiddleware())

		listHandler := r.getListHandler()

		rr.Get("/user", listHandler.Handle)

	})

	signupHandler := r.getSignupHandler()
	loginHandler := r.getLoginHandler()
	route.Post("/signup", signupHandler.Handle)
	route.Post("/login", loginHandler.Handle)

	fmt.Println("[user_routes > router > Route] routes up")
	return route
}

func (r *router) getSignupHandler() web_protocol.Handle {
	signUpValidation := user_validation.NewSignUpValidation(r.validator, r.userRepository)
	signUpUseCase := user_usecase.NewSignupUseCase(r.userRepository, r.encrypter)
	signupHandler := user_handler.NewSignupHandler(signUpValidation, signUpUseCase)

	return signupHandler
}

func (r *router) getLoginHandler() web_protocol.Handle {
	loginValidation := user_validation.NewLoginValidation(r.validator, r.encrypter)
	loginUseCase := user_usecase.NewLoginUseCase(r.userRepository, r.tokenizer)
	signupHandler := user_handler.NewLoginHandle(loginValidation, loginUseCase)

	return signupHandler
}

func (r *router) getListHandler() web_protocol.Handle {
	listValidation := user_validation.NewListValidation(r.validator)
	listUseCase := user_usecase.NewListUseCase(r.userRepository)
	listHandler := user_handler.NewListHandler(listUseCase, listValidation)

	return listHandler
}
