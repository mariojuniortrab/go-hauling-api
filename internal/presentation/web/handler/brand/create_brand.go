package brand_handler

import (
	"encoding/json"
	"net/http"

	brand_usecase "github.com/mariojuniortrab/hauling-api/internal/domain/usecase/brand"
)

type CreateBrandHandler struct {
	CreateBrandUseCase    *brand_usecase.CreateBrandUseCase
	CreateBrandValidation brand_usecase.CreateBrandValidation
}

func NewCreateBrandHandler(createBrandUseCase *brand_usecase.CreateBrandUseCase, createBrandValidation brand_usecase.CreateBrandValidation) *CreateBrandHandler {
	return &CreateBrandHandler{
		CreateBrandUseCase:    createBrandUseCase,
		CreateBrandValidation: createBrandValidation,
	}
}

func (c *CreateBrandHandler) Handle(w http.ResponseWriter, r *http.Request) {
	var input brand_usecase.CreateBrandInputDto

	w.Header().Set("Content-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	validationErr := c.CreateBrandValidation.Validate(input)
	if validationErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(validationErr)
		return
	}

	output, err := c.CreateBrandUseCase.Execute(input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}
