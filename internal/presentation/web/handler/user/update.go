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
	responseManager := web_response_manager.NewResponseManager(w)
	id := h.urlParser.GetPathParamFromURL(r, "id")

	editedUser, err := h.updateUseCase.GetForUpdate(id)
	if err != nil {
		responseManager.RespondInternalServerError(err)
	}
	if editedUser == nil {
		responseManager.RespondNotFound("user")
		return
	}

	fmt.Printf("r.Body: %+v\n", r.Body)
	err = json.NewDecoder(r.Body).Decode(editedUser)
	fmt.Println("Decodou:", editedUser)
	if err != nil {
		responseManager.RespondInternalServerError(err)
		return
	}

	emptyRequestError, validationErrs := h.updateValidation.Validate(editedUser)
	if validationErrs != nil {
		responseManager.SetBadRequestStatus().AddErrors(validationErrs).Respond()
		return
	}
	if emptyRequestError != nil {
		responseManager.SetBadRequestStatus().AddError(emptyRequestError).Respond()
		return
	}

	result, err := h.updateUseCase.Execute(id, editedUser)
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
