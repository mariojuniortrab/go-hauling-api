package user_handler

import (
	"encoding/json"
	"net/http"

	user_usecase "github.com/mariojuniortrab/hauling-api/internal/usecase/user"
)

type signupHandler struct {
	signUpValidation user_usecase.SignupValidation
	signUp           *user_usecase.Signup
}

func NewSignupHandler(signUpValidation user_usecase.SignupValidation,
	signUp *user_usecase.Signup) *signupHandler {
	return &signupHandler{
		signUpValidation: signUpValidation,
		signUp:           signUp,
	}
}

func (h *signupHandler) Handle(w http.ResponseWriter, r *http.Request) {
	var input user_usecase.SignupInputDto

	w.Header().Set("Content-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	validationErr := h.signUpValidation.Validate(&input)
	if validationErr != nil {
		w.WriteHeader(validationErr.StatusCode)
		json.NewEncoder(w).Encode(validationErr)
		return
	}

	output, err := h.signUp.Execute(input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}
