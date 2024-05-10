package web_response_manager

import (
	"encoding/json"
	"log"
	"net/http"

	errors_validation "github.com/mariojuniortrab/hauling-api/internal/domain/validation/errors"
)

type ResponseManager interface {
	SetStatusCreated() ResponseManager
	SetBadRequestStatus() ResponseManager
	SetConflictStatus() ResponseManager
	AddError(err *errors_validation.CustomErrorMessage) ResponseManager
	AddErrors(errs []*errors_validation.CustomErrorMessage) ResponseManager
	SetMessage(message string) ResponseManager
	SetData(data interface{}) ResponseManager
	Respond()
	RespondInternalServerError(err error)
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
	return &responseManager{
		w:          w,
		statusCode: 200,
	}
}

func (r *responseManager) SetStatusCreated() ResponseManager {
	return r.setStatusCode(http.StatusCreated)
}

func (r *responseManager) SetBadRequestStatus() ResponseManager {
	return r.setStatusCode(http.StatusBadRequest)
}

func (r *responseManager) SetConflictStatus() ResponseManager {
	return r.setStatusCode(http.StatusBadRequest)
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
	case r.statusCode == http.StatusOK:
		json.NewEncoder(r.w).Encode(&messageSucessful{Message: r.message, Data: r.data})
		return
	default:
		json.NewEncoder(r.w).Encode(&messageFieldError{Errors: r.errors})
	}
}

func (r *responseManager) RespondInternalServerError(err error) {
	r.w.WriteHeader(http.StatusInternalServerError)
	log.Println(err)
	json.NewEncoder(r.w).Encode(errors_validation.InternalServerError())
}
