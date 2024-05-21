package user_handler

import (
	"net/http"

	user_usecase "github.com/mariojuniortrab/hauling-api/internal/domain/usecase/user"
	web_protocol "github.com/mariojuniortrab/hauling-api/internal/presentation/web/protocol"
	web_response_manager "github.com/mariojuniortrab/hauling-api/internal/presentation/web/response-manager"
)

type detailHandler struct {
	urlParser     web_protocol.URLParser
	detailUseCase *user_usecase.DetailUserUseCase
}

func NewDetailHandler(urlParser web_protocol.URLParser,
	detailUseCase *user_usecase.DetailUserUseCase) *detailHandler {
	return &detailHandler{urlParser, detailUseCase}
}

func (h *detailHandler) Handle(w http.ResponseWriter, r *http.Request) {
	responseManager := web_response_manager.NewResponseManager(w)

	id := h.urlParser.GetPathParamFromURL(r, "id")

	if id == "" {
		responseManager.RespondUiidIsRequired()
	}

	result, err := h.detailUseCase.Execute(id)
	if err != nil {
		responseManager.RespondInternalServerError(err)
		return
	}

	if result == nil {
		responseManager.RespondNotFound("user")
		return
	}

	responseManager.SetStatusOk().SetMessage("success").SetData(result).Respond()
}
