package user_handler

import (
	"encoding/json"
	"fmt"
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
	fmt.Println("[user_handler > signupHandler > Handle] input:", input)
	if err != nil {
		fmt.Println("[user_handler > signupHandler > Handle] err:", err)
		responseManager.RespondInternalServerError(err)
		return
	}

	validationErrs := h.signUpValidation.Validate(&input)
	if validationErrs != nil {
		fmt.Println("[user_handler > signupHandler > Handle] validationErrs")
		responseManager.SetBadRequestStatus().AddErrors(validationErrs).Respond()
		return
	}

	alreadyExistsErr, err := h.signUpValidation.AlreadyExists(input.Email, "")
	if err != nil {
		fmt.Println("[user_handler > signupHandler > Handle] err:", err)
		responseManager.RespondInternalServerError(err)
		return
	}

	if alreadyExistsErr != nil {
		fmt.Println("[user_handler > signupHandler > Handle] alreadyExists")
		responseManager.SetConflictStatus().AddError(alreadyExistsErr).Respond()
		return
	}

	output, err := h.signUp.Execute(&input)
	if err != nil {
		fmt.Println("[user_handler > signupHandler > Handle] err:", err)
		responseManager.RespondInternalServerError(err)
		return
	}

	fmt.Println("[user_handler > signupHandler > Handle] successful")
	responseManager.SetStatusCreated().SetMessage("user created").SetData(output).Respond()
}
