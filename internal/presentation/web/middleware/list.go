package web_middleware

import (
	"net/http"

	protocol_validation "github.com/mariojuniortrab/hauling-api/internal/domain/validation/protocol"
	web_response_manager "github.com/mariojuniortrab/hauling-api/internal/presentation/web/response-menager"
)

type list struct {
	validator protocol_validation.Validator
}

func NewListMiddleware(validator protocol_validation.Validator) *list {
	return &list{validator}
}

func (p *list) GetMiddleware() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			responseManager := web_response_manager.NewResponseManager(w)

			page := r.URL.Query().Get("page")
			limit := r.URL.Query().Get("limit")
			p.validatePage(page)
			p.validateLimit(limit)

			if p.validator.HasErrors() {
				responseManager.
					SetBadRequestStatus().
					AddErrors(p.validator.GetErrorsAndClean()).
					Respond()
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func (p *list) validatePage(page string) {
	p.validator.
		ValidateRequiredField(page, "page").
		ValidateNumberField(page, "page")
}

func (p *list) validateLimit(page string) {
	p.validator.
		ValidateRequiredField(page, "limit").
		ValidateNumberField(page, "limit")
}
