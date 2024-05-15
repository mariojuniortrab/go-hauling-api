package web_response_manager

import (
	"encoding/json"
	"log"
	"net/http"

	errors_validation "github.com/mariojuniortrab/hauling-api/internal/domain/validation/errors"
)

type ResponseManager interface {
	SetStatusCreated() ResponseManager
	SetStatusOk() ResponseManager
	SetBadRequestStatus() ResponseManager
	SetConflictStatus() ResponseManager
	SetInternalServerErrorStatus() ResponseManager
	SetUnauthorizedStatus() ResponseManager
	AddError(err *errors_validation.CustomErrorMessage) ResponseManager
	AddErrors(errs []*errors_validation.CustomErrorMessage) ResponseManager
	SetMessage(message string) ResponseManager
	SetData(data interface{}) ResponseManager
	Respond()
	RespondInternalServerError(err error)
	RespondLoginInvalid()
	RespondUnauthorized()
	RawRespond(statusCode int, data interface{})
	RespondUiidInvalid()
	RespondUiidIsRequired()
}

type messageSucessful struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type messageFieldError struct {
	Errors []*errors_validation.CustomErrorMessage
}

type responseManager struct {
	w          http.ResponseWriter
	statusCode int
	errors     []*errors_validation.CustomErrorMessage
	message    string
	data       interface{}
}

func NewResponseManager(w http.ResponseWriter) *responseManager {
	w.Header().Set("Content-Type", "application/json")
	return &responseManager{
		w:          w,
		statusCode: 200,
	}
}

func (r *responseManager) SetStatusCreated() ResponseManager {
	return r.setStatusCode(http.StatusCreated)
}

func (r *responseManager) SetStatusOk() ResponseManager {
	return r.setStatusCode(http.StatusOK)
}

func (r *responseManager) SetBadRequestStatus() ResponseManager {
	return r.setStatusCode(http.StatusBadRequest)
}

func (r *responseManager) SetConflictStatus() ResponseManager {
	return r.setStatusCode(http.StatusConflict)
}

func (r *responseManager) SetInternalServerErrorStatus() ResponseManager {
	return r.setStatusCode(http.StatusInternalServerError)
}

func (r *responseManager) SetUnauthorizedStatus() ResponseManager {
	return r.setStatusCode(http.StatusUnauthorized)
}

func (r *responseManager) setStatusCode(statusCode int) ResponseManager {
	r.statusCode = statusCode
	return r
}

func (r *responseManager) AddError(err *errors_validation.CustomErrorMessage) ResponseManager {
	r.errors = append(r.errors, err)
	return r
}

func (r *responseManager) AddErrors(errs []*errors_validation.CustomErrorMessage) ResponseManager {
	r.errors = append(r.errors, errs...)
	return r
}

func (r *responseManager) SetMessage(message string) ResponseManager {
	r.message = message
	return r
}

func (r *responseManager) SetData(data interface{}) ResponseManager {
	r.data = data
	return r
}

func (r *responseManager) Respond() {
	r.w.WriteHeader(r.statusCode)

	switch {
	case r.statusCode == http.StatusCreated:
		json.NewEncoder(r.w).Encode(&messageSucessful{Message: r.message, Data: r.data})
		return
	case r.statusCode == http.StatusOK:
		json.NewEncoder(r.w).Encode(&messageSucessful{Message: r.message, Data: r.data})
		return
	default:
		json.NewEncoder(r.w).Encode(&messageFieldError{Errors: r.errors})
		return
	}
}

func (r *responseManager) RespondInternalServerError(err error) {
	errorMessage := errors_validation.NewCustomErrorMessage(errors_validation.InternalServerError(), "")
	r.SetInternalServerErrorStatus().AddError(errorMessage)
	log.Println(err)
	r.Respond()
}

func (r *responseManager) RespondLoginInvalid() {
	errorMessage := errors_validation.NewCustomErrorMessage(errors_validation.UserNotFound(), "")
	r.SetBadRequestStatus().AddError(errorMessage)
	r.Respond()
}

func (r *responseManager) RespondUnauthorized() {
	errorMessage := errors_validation.NewCustomErrorMessage(errors_validation.Unauthorized(), "")
	r.SetUnauthorizedStatus().AddError(errorMessage)
	r.Respond()
}

func (r *responseManager) RawRespond(statusCode int, data interface{}) {
	r.w.WriteHeader(r.statusCode)
	json.NewEncoder(r.w).Encode(data)
}

func (r *responseManager) RespondUiidInvalid() {
	errorMessage := errors_validation.NewCustomErrorMessage(errors_validation.UiidFromPathInvalid(), "")
	r.SetBadRequestStatus().AddError(errorMessage)
	r.Respond()
}

func (r *responseManager) RespondUiidIsRequired() {
	errorMessage := errors_validation.NewCustomErrorMessage(errors_validation.UiidFromPathIsRequired(), "")
	r.SetBadRequestStatus().AddError(errorMessage)
	r.Respond()
}
