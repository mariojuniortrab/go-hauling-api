package web_middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	protocol_validation "github.com/mariojuniortrab/hauling-api/internal/domain/validation/protocol"
	web_protocol "github.com/mariojuniortrab/hauling-api/internal/presentation/web/protocol"
	web_response_manager "github.com/mariojuniortrab/hauling-api/internal/presentation/web/response-manager"
)

type paginate struct {
	validator protocol_validation.Validator
	urlParser web_protocol.URLParser
}

func NewPaginateMiddleware(validator protocol_validation.Validator,
	urlParser web_protocol.URLParser) *paginate {
	return &paginate{validator, urlParser}
}

func (p *paginate) GetMiddleware() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {

		fn := func(w http.ResponseWriter, r *http.Request) {
			responseManager := web_response_manager.NewResponseManager(w)

			page := p.urlParser.GetQueryParamFromURL(r, "page")
			limit := p.urlParser.GetQueryParamFromURL(r, "limit")
			orderBy := p.urlParser.GetQueryParamFromURL(r, "orderBy")
			orderType := p.urlParser.GetQueryParamFromURL(r, "orderType")

			p.validate(page, "page")
			p.validate(limit, "limit")
			p.validateOrderFields(orderBy, orderType)

			if p.validator.HasErrors() {
				responseManager.
					SetBadRequestStatus().
					AddErrors(p.validator.GetErrorsAndClean()).
					Respond()
				return
			}

			next.ServeHTTP(w, r)
		}

		return http.HandlerFunc(fn)
	}
}

func (p *paginate) validate(input, fieldName string) {

	p.validator.
		ValidateRequiredField(input, fieldName).
		ValidateNumberField(input, fieldName)

	convertedField, err := strconv.Atoi(input)
	if err != nil {
		return
	}

	if convertedField == 0 {
		p.validator.AddError(fmt.Errorf("%s must be higher than 0", fieldName), fieldName)
	}
}

func (p *paginate) validateOrderFields(orderBy, orderType string) {
	if orderBy == "" && orderType == "" {
		return
	}

	if orderBy != "" && orderType == "" {
		p.validator.AddError(errors.New("orderType is required when orderBy is informed"), "orderType")
		return
	}

	if orderBy == "" && orderType != "" {
		p.validator.AddError(errors.New("orderBy is required when orderType is informed"), "orderBy")
		return
	}

	orderByItems := strings.Split(orderBy, ",")
	if len(orderByItems) > 1 {
		p.validator.AddError(errors.New("you can only order by 1 field at time"), "orderBy")
	}

	if strings.ToLower(orderType) != "asc" && strings.ToLower(orderType) != "desc" {
		p.validator.AddError(errors.New("orderType must be 'ASC' or 'DESC'"), "orderType")
	}
}
