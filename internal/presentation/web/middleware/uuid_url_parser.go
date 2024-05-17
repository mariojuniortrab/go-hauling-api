package web_middleware

import (
	"net/http"

	util_validation "github.com/mariojuniortrab/hauling-api/internal/domain/validation/util"
	web_protocol "github.com/mariojuniortrab/hauling-api/internal/presentation/web/protocol"
	web_response_manager "github.com/mariojuniortrab/hauling-api/internal/presentation/web/response-manager"
)

type uuidParser struct {
	urlParser web_protocol.URLParser
}

func NewUuidParser(urlParser web_protocol.URLParser) *uuidParser {
	return &uuidParser{urlParser}
}

func (m *uuidParser) GetMiddleware() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			responseManager := web_response_manager.NewResponseManager(w)

			uuid := m.urlParser.GetPathParamFromURL(r, "id")
			if !util_validation.IsUIID(uuid) {
				responseManager.RespondUiidInvalid()
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
