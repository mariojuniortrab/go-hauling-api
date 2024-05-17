package user_routes

import (
	"fmt"

	protocol_usecase "github.com/mariojuniortrab/hauling-api/internal/domain/usecase/protocol"
	user_usecase "github.com/mariojuniortrab/hauling-api/internal/domain/usecase/user"
	protocol_validation "github.com/mariojuniortrab/hauling-api/internal/domain/validation/protocol"
	user_validation "github.com/mariojuniortrab/hauling-api/internal/domain/validation/user"
	infra_adapters "github.com/mariojuniortrab/hauling-api/internal/infra/adapters"
	user_handler "github.com/mariojuniortrab/hauling-api/internal/presentation/web/handler/user"
	web_middleware "github.com/mariojuniortrab/hauling-api/internal/presentation/web/middleware"
	web_protocol "github.com/mariojuniortrab/hauling-api/internal/presentation/web/protocol"
)

type router struct {
	userRepository protocol_usecase.UserRepository
	validator      protocol_validation.Validator
	encrypter      protocol_usecase.Encrypter
	tokenizer      protocol_usecase.Tokenizer
	urlParser      web_protocol.URLParser
}

func NewRouter(userRepository protocol_usecase.UserRepository,
	validator protocol_validation.Validator,
	encrypter protocol_usecase.Encrypter,
	tokenizer protocol_usecase.Tokenizer,
	urlParser web_protocol.URLParser) *router {
	return &router{
		userRepository,
		validator,
		encrypter,
		tokenizer,
		urlParser,
	}
}

func (r *router) Route(route web_protocol.Router) web_protocol.Router {
	authUseCase := user_usecase.NewAuthorization(r.tokenizer)
	protected := web_middleware.NewProtectedMiddleware(r.tokenizer, authUseCase)
	list := web_middleware.NewListMiddleware(r.validator, infra_adapters.NewChiUrlParserAdapter())
	uuidUrlParser := web_middleware.NewUuidParser(r.urlParser)

	route.Group(func(rr web_protocol.Router) {
		rr.Use(protected.GetMiddleware())
		rr.Use(list.GetMiddleware())

		rr.Get("/user", r.getListHandler().Handle)
	})

	route.Group(func(rr web_protocol.Router) {
		rr.Use(protected.GetMiddleware())
		rr.Use(uuidUrlParser.GetMiddleware())

		rr.Get("/user/{id}", r.getDetailHandler().Handle)
		rr.Delete("/user/{id}", r.getRemoveHandler().Handle)
		rr.Patch("/user/{id}", r.getUpdateHandler().Handle)
	})

	route.Post("/signup", r.getSignupHandler().Handle)
	route.Post("/login", r.getLoginHandler().Handle)

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
	listHandler := user_handler.NewListHandler(listUseCase, listValidation, r.urlParser)

	return listHandler
}

func (r *router) getDetailHandler() web_protocol.Handle {
	detailUseCase := user_usecase.NewDetailuserUsecase(r.userRepository)
	detailHandler := user_handler.NewDetailHandler(r.urlParser, detailUseCase)

	return detailHandler
}

func (r *router) getRemoveHandler() web_protocol.Handle {
	removeUseCase := user_usecase.NewRemoveUserUsecase(r.userRepository)
	removeHandler := user_handler.NewRemoveHandler(r.urlParser, removeUseCase)

	return removeHandler
}

func (r *router) getUpdateHandler() web_protocol.Handle {
	updateValidation := user_validation.NewUpdateValidation(r.validator)
	updateUseCase := user_usecase.NewUpdateUserUsecase(r.userRepository)
	userHandler := user_handler.NewUpdateHandler(r.urlParser, updateUseCase, updateValidation)

	return userHandler
}
