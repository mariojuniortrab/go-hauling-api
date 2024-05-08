package user_handler

import (
	"encoding/json"
	"net/http"

	user_usecase "github.com/mariojuniortrab/hauling-api/internal/usecase/user"
)

type signupHandler struct {
	signUpValidation user_usecase.SignUpValidation
	signUpUseCase    *user_usecase.SignUpUseCase
}

func NewSignupHandler(signUpValidation user_usecase.SignUpValidation, signUpUseCase *user_usecase.SignUpUseCase) *signupHandler {
	return &signupHandler{
		signUpValidation: signUpValidation,
		signUpUseCase:    signUpUseCase,
	}
}

func (s *signupHandler) Handle(w http.ResponseWriter, r *http.Request) {
	var input user_usecase.SignUpInputDto

	w.Header().Set("Content-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	validationErr := s.signUpValidation.Validate(input)
	if validationErr != nil {
		w.WriteHeader(validationErr.StatusCode)
		json.NewEncoder(w).Encode(validationErr)
		return
	}

	output, err := s.signUpUseCase.Execute(input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}
