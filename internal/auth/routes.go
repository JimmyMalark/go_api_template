package auth

import (
	db "github.com/jimmymalark/go_api_template/internal/db"
	"github.com/jimmymalark/go_api_template/internal/validator"

	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(r chi.Router, store *db.Store) {
	authService := NewService(store)
	authHandler := NewHandler(authService, validator.New())

	r.Route("/auth", func(r chi.Router) {
		r.Post("/register", authHandler.Register)
		r.Post("/login", authHandler.Login)
		r.Post("/logout", authHandler.Logout)
		r.Get("/me", authHandler.Me)
	})
}
