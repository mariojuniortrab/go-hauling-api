package user_handler

import (
	"encoding/json"
	"net/http"

	user_usecase "github.com/mariojuniortrab/hauling-api/internal/domain/usecase/user"
	user_validation "github.com/mariojuniortrab/hauling-api/internal/domain/validation/user"
	web_response_manager "github.com/mariojuniortrab/hauling-api/internal/presentation/web/response-manager"
)

type signupHandler struct {
	signUpValidation user_validation.SignupValidation
	signUp           *user_usecase.Signup
}

func NewSignupHandler(signUpValidation user_validation.SignupValidation,
	signUp *user_usecase.Signup) *signupHandler {
	return &signupHandler{
		signUpValidation: signUpValidation,
		signUp:           signUp,
	}
}

func (h *signupHandler) Handle(w http.ResponseWriter, r *http.Request) {
	var input user_usecase.SignupInputDto

	responseManager := web_response_manager.NewResponseManager(w)

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		responseManager.RespondInternalServerError(err)
		return
	}

	validationErrs := h.signUpValidation.Validate(&input)
	if validationErrs != nil {
		responseManager.SetBadRequestStatus().AddErrors(validationErrs).Respond()
		return
	}

	alreadyExistsErr, err := h.signUpValidation.AlreadyExists(input.Email, "")
	if err != nil {
		responseManager.RespondInternalServerError(err)
		return
	}

	if alreadyExistsErr != nil {
		responseManager.SetConflictStatus().AddError(alreadyExistsErr).Respond()
		return
	}

	output, err := h.signUp.Execute(&input)
	if err != nil {
		responseManager.RespondInternalServerError(err)
		return
	}

	responseManager.SetStatusCreated().SetMessage("user created").SetData(output).Respond()
}
