package user_routes

import (
	"encoding/json"
	"net/http"

	user_entity "github.com/mariojuniortrab/hauling-api/internal/domain/entity/user"
	protocol_usecase "github.com/mariojuniortrab/hauling-api/internal/domain/usecase/protocol"
	user_usecase "github.com/mariojuniortrab/hauling-api/internal/domain/usecase/user"
	protocol_validation "github.com/mariojuniortrab/hauling-api/internal/domain/validation/protocol"
	user_validation "github.com/mariojuniortrab/hauling-api/internal/domain/validation/user"
	user_handler "github.com/mariojuniortrab/hauling-api/internal/presentation/web/handler/user"
	web_middleware "github.com/mariojuniortrab/hauling-api/internal/presentation/web/middleware"
	web_protocol "github.com/mariojuniortrab/hauling-api/internal/presentation/web/protocol"
)

type router struct {
	userRepository user_entity.UserRepository
	validator      protocol_validation.Validator
	encrypter      protocol_usecase.Encrypter
	tokenizer      protocol_usecase.Tokenizer
}

func NewRouter(userRepository user_entity.UserRepository,
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
	signupHandler := r.getSignupHandler()
	loginHandler := r.getLoginHandler()
	authUseCase := user_usecase.NewAuthorization(r.tokenizer)
	protected := web_middleware.NewProtectedMiddleware(r.tokenizer, authUseCase)
	list := web_middleware.NewListMiddleware(r.validator)

	route.Group(func(rr web_protocol.Router) {

		rr.Use(protected.GetMiddleware())
		rr.Use(list.GetMiddleware())

		rr.Get("/user", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(200)
			json.NewEncoder(w).Encode("Response")
		})

	})

	route.Post("/signup", signupHandler.Handle)
	route.Post("/login", loginHandler.Handle)

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
