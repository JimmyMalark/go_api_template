package auth

import (
	"github.com/jimmymalark/go_api_template/internal/config"
	db "github.com/jimmymalark/go_api_template/internal/db"
	"github.com/jimmymalark/go_api_template/internal/validator"

	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(r chi.Router, store *db.Store, cookieConfig config.CookieConfig) {
	authService := NewService(store)
	authHandler := NewHandler(authService, validator.New(), cookieConfig)

	authenticator := NewAuthenticator(store)

	r.Route("/auth", func(r chi.Router) {
		r.Post("/register", authHandler.Register)
		r.Post("/login", authHandler.Login)

		r.With(Middleware(authenticator, cookieConfig)).Group(func(r chi.Router) {
			r.Get("/me", authHandler.Me)
			r.Post("/logout", authHandler.Logout)
		})
	})
}
