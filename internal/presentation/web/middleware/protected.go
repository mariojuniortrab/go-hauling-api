package web_middleware

import (
	"net/http"

	protocol_usecase "github.com/mariojuniortrab/hauling-api/internal/domain/usecase/protocol"
	user_usecase "github.com/mariojuniortrab/hauling-api/internal/domain/usecase/user"
	web_response_manager "github.com/mariojuniortrab/hauling-api/internal/presentation/web/response-menager"
)

type Protected struct {
	tokenizer protocol_usecase.Tokenizer
	auth      *user_usecase.Authorization
}

func NewProtectedMiddleware(tokenizer protocol_usecase.Tokenizer,
	auth *user_usecase.Authorization) *Protected {
	return &Protected{
		tokenizer,
		auth,
	}
}

func (p *Protected) GetMiddleware() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			responseManager := web_response_manager.NewResponseManager(w)
			token := r.Header.Get("Authorization")

			if token == "" {
				responseManager.RespondUnauthorized()
				return
			}

			output, err := p.auth.Execute(&user_usecase.AuthInputDto{Token: token})

			if err != nil {
				responseManager.RespondUnauthorized()
				return
			}

			responseManager.RawRespond(200, output)
		})
	}
}
