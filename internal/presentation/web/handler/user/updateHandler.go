package user_handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	user_usecase "github.com/mariojuniortrab/hauling-api/internal/domain/usecase/user"
	user_validation "github.com/mariojuniortrab/hauling-api/internal/domain/validation/user"
	web_protocol "github.com/mariojuniortrab/hauling-api/internal/presentation/web/protocol"
	web_response_manager "github.com/mariojuniortrab/hauling-api/internal/presentation/web/response-manager"
)

type updateHandler struct {
	urlParser        web_protocol.URLParser
	updateUseCase    *user_usecase.UpdateUserUseCase
	updateValidation user_validation.UpdateValidation
}

func NewUpdateHandler(urlParser web_protocol.URLParser,
	updateUseCase *user_usecase.UpdateUserUseCase,
	updateValidation user_validation.UpdateValidation) *updateHandler {
	return &updateHandler{urlParser, updateUseCase, updateValidation}
}

func (h *updateHandler) Handle(w http.ResponseWriter, r *http.Request) {
	var input user_usecase.UserUpdateInputDto

	responseManager := web_response_manager.NewResponseManager(w)

	input.ID = h.urlParser.GetPathParamFromURL(r, "id")

	fmt.Println("[user_handler > updateHandler > Handle] uuid:", input.ID)
	if input.ID == "" {
		responseManager.RespondUiidIsRequired()
	}

	var updateValues user_usecase.UpdateFields

	err := json.NewDecoder(r.Body).Decode(&updateValues)
	fmt.Println("[user_handler > signupHandler > Handle] input:", input)
	if err != nil {
		fmt.Println("[user_handler > signupHandler > Handle] err:", err)
		responseManager.RespondInternalServerError(err)
		return
	}

	emptyRequestError, validationErrs := h.updateValidation.Validate(&updateValues)
	if validationErrs != nil {
		fmt.Println("[user_handler > signupHandler > Handle] validationErrs")
		responseManager.SetBadRequestStatus().AddErrors(validationErrs).Respond()
		return
	}
	if emptyRequestError != nil {
		fmt.Println("[user_handler > signupHandler > Handle] validationErrs")
		responseManager.SetBadRequestStatus().AddError(emptyRequestError).Respond()
		return
	}

	result, err := h.updateUseCase.Execute(&input, &updateValues)
	if err != nil {
		fmt.Println("[user_handler > updateHandler > Handle] err:", err)
		responseManager.RespondInternalServerError(err)
		return
	}

	if result == nil {
		responseManager.RespondNotFound("user")
		return
	}

	fmt.Println("[user_handler > updateHandler > Handle] successful")
	responseManager.SetStatusOk().SetMessage("success").SetData(result).Respond()
}
