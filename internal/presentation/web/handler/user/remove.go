package user_handler

import (
	"fmt"
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
	var input user_usecase.UserRemoveInputDto

	responseManager := web_response_manager.NewResponseManager(w)

	input.ID = h.urlParser.GetPathParamFromURL(r, "id")

	fmt.Println("[user_handler > removeHandler > Handle] uuid:", input.ID)
	if input.ID == "" {
		responseManager.RespondUiidIsRequired()
	}

	err, errNotFound := h.removeUseCase.Execute(&input)
	if err != nil {
		fmt.Println("[user_handler > removeHandler > Handle] err:", err)
		responseManager.RespondInternalServerError(err)
		return
	}

	if errNotFound == nil {
		responseManager.RespondNotFound("user")
		return
	}

	responseManager.SetStatusOk().SetMessage("removed").Respond()
}
