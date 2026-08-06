package auth

import (
	"net/http"

	"github.com/jimmymalark/go_api_template/internal/config"
	"github.com/jimmymalark/go_api_template/internal/httpx"
)

func Middleware(authenticator *Authenticator, cookieConfig config.CookieConfig) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			cookie, err := r.Cookie(cookieConfig.Name)
			if err != nil {
				httpx.WriteError(w, http.StatusUnauthorized, "not authenticated", r, nil)
				return
			}

			identity, err := authenticator.Authenticate(r.Context(), cookie.Value)
			if err != nil {
				httpx.WriteError(w, http.StatusUnauthorized, "not authenticated", r, nil)
				return
			}

			ctx := WithIdentity(r.Context(), identity)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
