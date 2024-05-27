package web_response_manager

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	errors_validation "github.com/mariojuniortrab/hauling-api/internal/domain/validation/errors"
)

type messageSucessful struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type messageFieldErrorArray struct {
	Errors []*errors_validation.CustomFieldErrorMessage
}

type messageFieldError struct {
	Errors *errors_validation.CustomFieldErrorMessage
}
type messageError struct {
	Error *errors_validation.CustomErrorMessage
}

func RespondOk(w http.ResponseWriter, message string, data any) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&messageSucessful{Message: message, Data: data})
}

func RespondUiidIsRequired(w http.ResponseWriter) {
	errorMessage := errors_validation.NewCustomErrorMessage(errors_validation.UiidFromPathIsRequired())

	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(&messageError{Error: errorMessage})
}

func RespondInternalServerError(w http.ResponseWriter, err error) {
	log.Print(err)
	errorMessage := errors_validation.NewCustomErrorMessage(errors_validation.InternalServerError())
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(&messageError{Error: errorMessage})
}

func RespondNotFound(w http.ResponseWriter, resource string) {
	errorMessage := errors_validation.NewCustomErrorMessage(errors_validation.NotFound(resource))
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(&messageError{Error: errorMessage})
}

func RespondFieldErrorValidation(w http.ResponseWriter, errs []*errors_validation.CustomFieldErrorMessage) {
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(&messageFieldErrorArray{Errors: errs})
}

func RespondLoginInvalid(w http.ResponseWriter) {
	errorMessage := errors_validation.NewCustomErrorMessage(errors_validation.UserNotFound())
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(&messageError{Error: errorMessage})
}

func RespondConflictError(w http.ResponseWriter, errs *errors_validation.CustomFieldErrorMessage) {
	w.WriteHeader(http.StatusConflict)
	json.NewEncoder(w).Encode(&messageFieldError{Errors: errs})
}

func RespondCreated(w http.ResponseWriter, message string, data any) {
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&messageSucessful{Message: message, Data: data})
}

func RespondGenericError(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	errorMessage := errors_validation.NewCustomErrorMessage(errors.New(message))
	json.NewEncoder(w).Encode(&messageError{Error: errorMessage})
}

func RespondUnsupportedMediaType(w http.ResponseWriter) {
	w.WriteHeader(http.StatusCreated)
	errorMessage := errors_validation.NewCustomErrorMessage(errors_validation.ContentTypeIsNotJSON())
	json.NewEncoder(w).Encode(&messageError{Error: errorMessage})
}

func RespondUnauthorized(w http.ResponseWriter) {
	w.WriteHeader(http.StatusCreated)
	errorMessage := errors_validation.NewCustomErrorMessage(errors_validation.Unauthorized())
	json.NewEncoder(w).Encode(&messageError{Error: errorMessage})
}

func RespondUiidInvalid(w http.ResponseWriter) {
	w.WriteHeader(http.StatusCreated)
	errorMessage := errors_validation.NewCustomErrorMessage(errors_validation.Unauthorized())
	json.NewEncoder(w).Encode(&messageError{Error: errorMessage})
}
