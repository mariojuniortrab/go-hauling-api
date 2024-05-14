package user_handler

import (
	"encoding/json"
	"fmt"
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
	fmt.Println("[user_handler > loginHandler > Handle] input:", input)
	if err != nil {
		fmt.Println("[user_handler > loginHandler > Handle] err:", err)
		responseManager.RespondInternalServerError(err)
		return
	}

	validationErrs := h.loginValidation.Validate(&input)
	if validationErrs != nil {
		fmt.Println("[user_handler > loginHandler > Handle] validationErrs")
		responseManager.SetBadRequestStatus().AddErrors(validationErrs).Respond()
		return
	}

	user, err := h.loginUseCase.GetByEmail(&input)
	fmt.Println("[user_handler > loginHandler > Handle] user:", user)
	if err != nil {
		fmt.Println("[user_handler > loginHandler > Handle] err:", err)
		responseManager.RespondInternalServerError(err)
		return
	}
	if user == nil {
		responseManager.RespondLoginInvalid()
		return
	}

	if h.loginValidation.IsCredentialInvalid(user, input.Password) {
		fmt.Println("[user_handler > loginHandler > Handle] IsCredentialInvalid")
		responseManager.RespondLoginInvalid()
		return
	}

	output, err := h.loginUseCase.Execute(user)
	if err != nil {
		fmt.Println("[user_handler > loginHandler > Handle] err:", err)
		responseManager.RespondInternalServerError(err)
		return
	}

	fmt.Println("[user_handler > loginHandler > Handle] successful")
	responseManager.SetStatusOk().SetMessage("login successful").SetData(output).Respond()
}
