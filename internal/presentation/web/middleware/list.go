package web_middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	protocol_validation "github.com/mariojuniortrab/hauling-api/internal/domain/validation/protocol"
	web_response_manager "github.com/mariojuniortrab/hauling-api/internal/presentation/web/response-manager"
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
			orderBy := r.URL.Query().Get("orderBy")
			orderType := r.URL.Query().Get("orderType")

			fmt.Println("[web_middlewares > list > handlerFunc] page:", page)
			fmt.Println("[web_middlewares > list > handlerFunc] limit:", limit)
			fmt.Println("[web_middlewares > list > handlerFunc] orderBy:", orderBy)
			fmt.Println("[web_middlewares > list > handlerFunc] orderType:", orderType)

			p.validate(page, "page")
			p.validate(limit, "limit")
			p.validateOrderFields(orderBy, orderType)

			if p.validator.HasErrors() {
				fmt.Println("[web_middlewares > list > handlerFunc] hasErrors")
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

func (p *list) validate(input, fieldName string) {

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

func (p *list) validateOrderFields(orderBy, orderType string) {
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
