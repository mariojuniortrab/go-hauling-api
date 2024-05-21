package user_handler

import (
	"net/http"

	user_usecase "github.com/mariojuniortrab/hauling-api/internal/domain/usecase/user"
	web_protocol "github.com/mariojuniortrab/hauling-api/internal/presentation/web/protocol"
	web_response_manager "github.com/mariojuniortrab/hauling-api/internal/presentation/web/response-manager"
)

type removeHandler struct {
	urlParser     web_protocol.URLParser
	removeUseCase *user_usecase.RemoveUserUseCase
}

func NewRemoveHandler(urlParser web_protocol.URLParser,
	removeUseCase *user_usecase.RemoveUserUseCase) *removeHandler {
	return &removeHandler{urlParser, removeUseCase}
}

func (h *removeHandler) Handle(w http.ResponseWriter, r *http.Request) {
	responseManager := web_response_manager.NewResponseManager(w)

	id := h.urlParser.GetPathParamFromURL(r, "id")

	if id == "" {
		responseManager.RespondUiidIsRequired()
	}

	err, errNotFound := h.removeUseCase.Execute(id)
	if err != nil {
		responseManager.RespondInternalServerError(err)
		return
	}

	if errNotFound != nil {
		responseManager.RespondNotFound("user")
		return
	}

	responseManager.SetStatusOk().SetMessage("removed").Respond()
}
