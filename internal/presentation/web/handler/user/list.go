package user_handler

import (
	"fmt"
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

	h.parseUrlParams(r, &input)
	fmt.Println("[user_handler > listHandler > Handle] input:", input)

	validationErrs := h.listValidation.Validate(&input)
	if validationErrs != nil {
		fmt.Println("[user_handler > listHandler > Handle] validationErrs")
		responseManager.SetBadRequestStatus().AddErrors(validationErrs).Respond()
		return
	}

	result, err := h.listUseCase.Execute(&input)
	if err != nil {
		fmt.Println("[user_handler > listHandler > Handle] err:", err)
		responseManager.RespondInternalServerError(err)
		return
	}

	fmt.Println("[user_handler > listHandler > Handle] successful")
	responseManager.SetStatusOk().SetMessage("success").SetData(result).Respond()
}

func (h *listHandler) parseUrlParams(r *http.Request, input *user_usecase.ListUserInputDto) {
	input.Page = r.URL.Query().Get("page")
	input.Limit = r.URL.Query().Get("limit")
	input.OrderBy = r.URL.Query().Get("orderBy")
	input.OrderType = r.URL.Query().Get("orderType")
	input.Q = r.URL.Query().Get("q")

	input.ID = r.URL.Query().Get("id")
	input.Email = r.URL.Query().Get("email")
	input.Name = r.URL.Query().Get("name")
	input.Active = r.URL.Query().Get("active")
}
