package user_handler

import (
	"encoding/json"
	"net/http"

	user_usecase "github.com/mariojuniortrab/hauling-api/internal/domain/usecase/user"
	user_validation "github.com/mariojuniortrab/hauling-api/internal/domain/validation/user"
	web_response_manager "github.com/mariojuniortrab/hauling-api/internal/presentation/web/response-manager"
)

type loginHandler struct {
	loginValidation user_validation.LoginValidation
	loginUseCase    *user_usecase.Login
}

func NewLoginHandle(loginValidation user_validation.LoginValidation,
	loginUseCase *user_usecase.Login) *loginHandler {
	return &loginHandler{
		loginValidation,
		loginUseCase,
	}
}

func (h *loginHandler) Handle(w http.ResponseWriter, r *http.Request) {
	var input user_usecase.LoginInputDto

	responseManager := web_response_manager.NewResponseManager(w)

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		responseManager.RespondInternalServerError(err)
		return
	}

	validationErrs := h.loginValidation.Validate(&input)
	if validationErrs != nil {
		responseManager.SetBadRequestStatus().AddErrors(validationErrs).Respond()
		return
	}

	user, err := h.loginUseCase.GetByEmail(&input)

	if err != nil {
		responseManager.RespondInternalServerError(err)
		return
	}
	if user == nil {
		responseManager.RespondLoginInvalid()
		return
	}

	if h.loginValidation.IsCredentialInvalid(user, input.Password) {
		responseManager.RespondLoginInvalid()
		return
	}

	output, err := h.loginUseCase.Execute(user)
	if err != nil {
		responseManager.RespondInternalServerError(err)
		return
	}

	responseManager.SetStatusOk().SetMessage("login successful").SetData(output).Respond()
}
