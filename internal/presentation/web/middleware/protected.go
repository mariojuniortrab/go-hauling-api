package web_middleware

import (
	"context"
	"fmt"
	"net/http"

	protocol_usecase "github.com/mariojuniortrab/hauling-api/internal/domain/usecase/protocol"
	user_usecase "github.com/mariojuniortrab/hauling-api/internal/domain/usecase/user"
	web_response_manager "github.com/mariojuniortrab/hauling-api/internal/presentation/web/response-manager"
)

type LoggedUser struct {
	User string
}

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

			fmt.Println("[web_middlewares > Protected > handlerFunc] token:", token)
			if token == "" {
				responseManager.RespondUnauthorized()
				return
			}

			output, err := p.auth.Execute(&user_usecase.AuthInputDto{Token: token})
			fmt.Println("[web_middlewares > Protected > handlerFunc] output:", output)
			if err != nil {
				fmt.Println("[web_middlewares > Protected > handlerFunc] err:", err)
				responseManager.RespondUnauthorized()
				return
			}

			ctx := context.WithValue(r.Context(), LoggedUser{"loggedUser"}, output)

			newRequest := r.Clone(ctx)

			next.ServeHTTP(w, newRequest)
		})
	}
}
