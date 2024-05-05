package brand_handler

import (
	"encoding/json"
	"net/http"

	brand_usecase "github.com/mariojuniortrab/hauling-api/internal/usecase/brand"
)

type ListBrandHandler struct {
	ListBrandUseCase *brand_usecase.ListBrandUseCase
}

func NewListBrandHandler(listBrandUseCase *brand_usecase.ListBrandUseCase) *ListBrandHandler {
	return &ListBrandHandler{
		ListBrandUseCase: listBrandUseCase,
	}
}

func (l *ListBrandHandler) Handle(w http.ResponseWriter, r *http.Request) {
	output, err := l.ListBrandUseCase.Execute()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}
