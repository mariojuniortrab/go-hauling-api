package user_handler

import (
	"net/http"

	user_entity "github.com/mariojuniortrab/hauling-api/internal/domain/entity/user"
	user_usecase "github.com/mariojuniortrab/hauling-api/internal/domain/usecase/user"
	user_validation "github.com/mariojuniortrab/hauling-api/internal/domain/validation/user"
	web_protocol "github.com/mariojuniortrab/hauling-api/internal/presentation/web/protocol"
	web_response_manager "github.com/mariojuniortrab/hauling-api/internal/presentation/web/response-manager"
)

type listHandler struct {
	listUseCase    *user_usecase.List
	listValidation user_validation.ListValidation
	urlParser      web_protocol.URLParser
}

func NewListHandler(listUseCase *user_usecase.List,
	listValidation user_validation.ListValidation,
	urlParser web_protocol.URLParser) *listHandler {
	return &listHandler{
		listUseCase,
		listValidation,
		urlParser,
	}
}

func (h *listHandler) Handle(w http.ResponseWriter, r *http.Request) {
	var input user_entity.ListUserInputDto

	responseManager := web_response_manager.NewResponseManager(w)

	h.parseUrlParams(r, &input)

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

	responseManager.SetStatusOk().SetMessage("success").SetData(result).Respond()
}

func (h *listHandler) parseUrlParams(r *http.Request, input *user_entity.ListUserInputDto) {
	input.Page = h.urlParser.GetQueryParamFromURL(r, "page")
	input.Limit = h.urlParser.GetQueryParamFromURL(r, "limit")
	input.OrderBy = h.urlParser.GetQueryParamFromURL(r, "orderBy")
	input.OrderType = h.urlParser.GetQueryParamFromURL(r, "orderType")
	input.Q = h.urlParser.GetQueryParamFromURL(r, "q")

	input.ID = h.urlParser.GetQueryParamFromURL(r, "id")
	input.Email = h.urlParser.GetQueryParamFromURL(r, "email")
	input.Name = h.urlParser.GetQueryParamFromURL(r, "name")
	input.Active = h.urlParser.GetQueryParamFromURL(r, "active")
}
