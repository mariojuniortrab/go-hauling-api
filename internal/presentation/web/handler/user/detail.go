package user_handler

import (
	"fmt"
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
	var input user_usecase.UserDetailInputDto

	responseManager := web_response_manager.NewResponseManager(w)

	input.ID = h.urlParser.GetPathParamFromURL(r, "id")

	fmt.Println("[user_handler > detailHandler > Handle] uuid:", input.ID)
	if input.ID == "" {
		responseManager.RespondUiidIsRequired()
	}

	result, err := h.detailUseCase.Execute(&input)
	if err != nil {
		fmt.Println("[user_handler > detailHandler > Handle] err:", err)
		responseManager.RespondInternalServerError(err)
		return
	}

	if result == nil {
		responseManager.RespondNotFound("user")
		return
	}

	fmt.Println("[user_handler > detailHandler > Handle] successful")
	responseManager.SetStatusOk().SetMessage("success").SetData(result).Respond()
}
