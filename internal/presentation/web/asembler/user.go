package web_assembler

import (
	"database/sql"
	"net/http"

	auth_usecase "github.com/mariojuniortrab/hauling-api/internal/domain/usecase/auth"
	protocol_application "github.com/mariojuniortrab/hauling-api/internal/domain/usecase/protocol/application"
	user_usecase "github.com/mariojuniortrab/hauling-api/internal/domain/usecase/user"
	user_validation "github.com/mariojuniortrab/hauling-api/internal/domain/validation/user"
	user_mysql_repository "github.com/mariojuniortrab/hauling-api/internal/infra/repository/mysql/user"
	user_handler "github.com/mariojuniortrab/hauling-api/internal/presentation/web/handler/user"
)

type UserAssembler struct {
	mysqlDB   *sql.DB
	validator protocol_application.Validator
	encrypter protocol_application.Encrypter
	tokenizer protocol_application.Tokenizer
	urlParser protocol_application.URLParser
}

func NewUserAssembler(
	mysqlDB *sql.DB,
	validator protocol_application.Validator,
	encrypter protocol_application.Encrypter,
	tokenizer protocol_application.Tokenizer,
	urlParser protocol_application.URLParser,
) *UserAssembler {
	return &UserAssembler{
		mysqlDB,
		validator,
		encrypter,
		tokenizer,
		urlParser,
	}
}

func (a *UserAssembler) GetAssembledSignupHandle() http.HandlerFunc {
	return a.assembleSignupHandler().Handle
}

func (a *UserAssembler) GetAssembledLoginHandle() http.HandlerFunc {
	return a.assembleLoginHandler().Handle
}

func (a *UserAssembler) GetAssembledListUserHandle() http.HandlerFunc {
	return a.assembleListUserHandler().Handle
}

func (a *UserAssembler) GetAssembledDetailUserHandle() http.HandlerFunc {
	return a.assembleDetailUserHandler().Handle
}

func (a *UserAssembler) GetAssembledRemoveUserHandle() http.HandlerFunc {
	return a.assembleRemoveUserHandler().Handle
}

func (a *UserAssembler) GetAssembledUpdateUserHandle() http.HandlerFunc {
	return a.assembleUpdateUserHandler().Handle
}

func (a *UserAssembler) assembleSignupHandler() protocol_application.Handler {
	signupRepository := user_mysql_repository.NewSignupRepository(a.mysqlDB)
	signUpValidation := user_validation.NewSignUpValidation(a.validator, signupRepository)
	signUpUseCase := auth_usecase.NewSignupUseCase(signupRepository, a.encrypter)
	signupHandler := user_handler.NewSignupHandler(signUpValidation, signUpUseCase)

	return signupHandler
}

func (a *UserAssembler) assembleLoginHandler() protocol_application.Handler {
	loginRepository := user_mysql_repository.NewLoginRepository(a.mysqlDB)
	loginValidation := user_validation.NewLoginValidation(a.validator, a.encrypter)
	loginUseCase := auth_usecase.NewLoginUseCase(loginRepository, a.tokenizer)
	loginHandler := user_handler.NewLoginHandle(loginValidation, loginUseCase)

	return loginHandler
}

func (a *UserAssembler) assembleListUserHandler() protocol_application.Handler {
	listRepository := user_mysql_repository.NewListUserRepository(a.mysqlDB)
	listValidation := user_validation.NewListValidation(a.validator)
	listUseCase := user_usecase.NewListUseCase(listRepository)
	listHandler := user_handler.NewListHandler(listUseCase, listValidation, a.urlParser)

	return listHandler
}

func (a *UserAssembler) assembleDetailUserHandler() protocol_application.Handler {
	detailRepository := user_mysql_repository.NewDetailUserRepository(a.mysqlDB)
	detailUseCase := user_usecase.NewDetailuserUsecase(detailRepository)
	detailHandler := user_handler.NewDetailHandler(a.urlParser, detailUseCase)

	return detailHandler
}

func (a *UserAssembler) assembleRemoveUserHandler() protocol_application.Handler {
	removeRepository := user_mysql_repository.NewRemoveUserRepository(a.mysqlDB)
	removeUseCase := user_usecase.NewRemoveUserUsecase(removeRepository)
	removeHandler := user_handler.NewRemoveHandler(a.urlParser, removeUseCase)

	return removeHandler
}

func (a *UserAssembler) assembleUpdateUserHandler() protocol_application.Handler {
	updateRepository := user_mysql_repository.NewUpdateUserRepository(a.mysqlDB)
	updateValidation := user_validation.NewUpdateValidation(a.validator)
	updateUseCase := user_usecase.NewUpdateUserUsecase(updateRepository, a.encrypter)
	updateHandler := user_handler.NewUpdateHandler(a.urlParser, updateUseCase, updateValidation)

	return updateHandler
}
