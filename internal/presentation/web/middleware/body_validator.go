package web_middleware

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	web_response_manager "github.com/mariojuniortrab/hauling-api/internal/presentation/web/response-manager"
)

type IsEmpty interface {
	IsEmpty() bool
}

type bodyValidator struct {
	fieldsToValidate IsEmpty
}

func NewBodyValidator(fieldsToValidate IsEmpty) *bodyValidator {
	return &bodyValidator{
		fieldsToValidate: fieldsToValidate,
	}
}

func (m *bodyValidator) GetMiddleware() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			responseManager := web_response_manager.NewResponseManager(w)

			fakebody, err := getFakeBody(r)

			if err != nil {
				responseManager.RespondInternalServerError(err)
				return
			}

			fakebody = http.MaxBytesReader(w, fakebody, 1048576)

			dec := json.NewDecoder(fakebody)
			dec.DisallowUnknownFields()

			err = dec.Decode(m.fieldsToValidate)
			if err != nil {
				var syntaxError *json.SyntaxError
				var unmarshalTypeError *json.UnmarshalTypeError

				switch {

				case errors.As(err, &syntaxError):
					msg := fmt.Sprintf("Request body contains badly-formed JSON (at position %d)", syntaxError.Offset)
					responseManager.
						SetBadRequestStatus().
						AddNewGenericError(msg).
						Respond()
					return

				case errors.Is(err, io.ErrUnexpectedEOF):
					msg := "Request body contains badly-formed JSON"
					responseManager.
						SetBadRequestStatus().
						AddNewGenericError(msg).
						Respond()
					return

				case errors.As(err, &unmarshalTypeError):
					msg := fmt.Sprintf("Request body contains an invalid value for the %q field (at position %d)", unmarshalTypeError.Field, unmarshalTypeError.Offset)
					responseManager.
						SetBadRequestStatus().
						AddNewGenericError(msg).
						Respond()
					return

				case strings.HasPrefix(err.Error(), "json: unknown field "):
					fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
					msg := fmt.Sprintf("Request body contains unknown field %s", fieldName)
					responseManager.
						SetBadRequestStatus().
						AddNewGenericError(msg).
						Respond()
					return

				case errors.Is(err, io.EOF):
					msg := "Request body must not be empty"
					responseManager.
						SetBadRequestStatus().
						AddNewGenericError(msg).
						Respond()
					return

				case err.Error() == "http: request body too large":
					msg := "Request body must not be larger than 1MB"
					responseManager.
						SetRequestEntityTooLargeStatus().
						AddNewGenericError(msg).
						Respond()
					return

				default:
					responseManager.RespondInternalServerError(err)
					return
				}
			}

			err = dec.Decode(&struct{}{})
			if !errors.Is(err, io.EOF) {
				msg := "Request body must only contain a single JSON object"
				responseManager.
					SetBadRequestStatus().
					AddNewGenericError(msg).
					Respond()
				return
			}

			if m.fieldsToValidate.IsEmpty() {
				msg := "Request body must not be empty"
				responseManager.
					SetBadRequestStatus().
					AddNewGenericError(msg).
					Respond()
				return
			}

			next.ServeHTTP(w, r)
		}

		return http.HandlerFunc(fn)
	}
}

func getFakeBody(r *http.Request) (io.ReadCloser, error) {
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	copyBody := io.NopCloser(bytes.NewBuffer(bodyBytes))

	return copyBody, nil
}
