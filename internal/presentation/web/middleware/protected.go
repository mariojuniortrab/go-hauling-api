package web_middleware

import (
	"context"
	"net/http"

	auth_usecase "github.com/mariojuniortrab/hauling-api/internal/domain/usecase/auth"
	web_response_manager "github.com/mariojuniortrab/hauling-api/internal/presentation/web/response-manager"
)

type LoggedUser struct {
	User string
}

type Protected struct {
	auth *auth_usecase.Authorization
}

func NewProtectedMiddleware(auth *auth_usecase.Authorization) *Protected {
	return &Protected{
		auth,
	}
}

func (m *Protected) Middleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {

		token := r.Header.Get("Authorization")

		if token == "" {
			web_response_manager.RespondUnauthorized(w)
			return
		}

		output, err := m.auth.Execute(&auth_usecase.AuthInputDto{Token: token})
		if err != nil {
			web_response_manager.RespondUnauthorized(w)
			return
		}

		ctx := context.WithValue(r.Context(), LoggedUser{"loggedUser"}, output)

		newRequest := r.Clone(ctx)

		next.ServeHTTP(w, newRequest)
	}

	return http.HandlerFunc(fn)
}
