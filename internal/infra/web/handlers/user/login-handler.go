package user_handler

import (
	"encoding/json"
	"net/http"

	user_usecase "github.com/mariojuniortrab/hauling-api/internal/usecase/user"
)

type loginHandler struct {
	loginValidation user_usecase.LoginValidation
	loginUseCase    *user_usecase.Login
}

func NewLoginHandle(loginValidation user_usecase.LoginValidation,
	loginUseCase *user_usecase.Login) *loginHandler {
	return &loginHandler{
		loginValidation,
		loginUseCase,
	}
}

func (h *loginHandler) Handle(w http.ResponseWriter, r *http.Request) {
	var input user_usecase.LoginInputDto

	w.Header().Set("Content-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	validationErr := h.loginValidation.Validate(&input)
	if validationErr != nil {
		w.WriteHeader(validationErr.StatusCode)
		json.NewEncoder(w).Encode(validationErr)
		return
	}

	user, err := h.loginUseCase.GetByUsername(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	validationErr = h.loginValidation.ValidateCredentials(user, input.Password)
	if validationErr != nil {
		w.WriteHeader(validationErr.StatusCode)
		json.NewEncoder(w).Encode(validationErr)
		return
	}

	output, err := h.loginUseCase.Execute(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}
