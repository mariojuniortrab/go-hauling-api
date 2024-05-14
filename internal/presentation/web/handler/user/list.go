package user_handler

import (
	"encoding/json"
	"net/http"

	user_usecase "github.com/mariojuniortrab/hauling-api/internal/domain/usecase/user"
	user_validation "github.com/mariojuniortrab/hauling-api/internal/domain/validation/user"
	web_response_manager "github.com/mariojuniortrab/hauling-api/internal/presentation/web/response-manager"
)

type listHandler struct {
	listUseCase    *user_usecase.List
	listValidation user_validation.ListValidation
}

func NewListHandler(listUseCase *user_usecase.List,
	listValidation user_validation.ListValidation) *listHandler {
	return &listHandler{
		listUseCase,
		listValidation,
	}
}

func (h *listHandler) Handle(w http.ResponseWriter, r *http.Request) {
	var input user_usecase.ListUserInputDto

	responseManager := web_response_manager.NewResponseManager(w)

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		responseManager.RespondInternalServerError(err)
		return
	}

	validationErrs := h.listValidation.Validate(&input)
	if validationErrs != nil {
		responseManager.SetBadRequestStatus().AddErrors(validationErrs).Respond()
		return
	}

	result, err := h.listUseCase.Execute(&input)
	if err != nil {
		responseManager.RespondInternalServerError(err)
		return
	}

	responseManager.SetStatusOk().SetMessage("login successful").SetData(result).Respond()
}
