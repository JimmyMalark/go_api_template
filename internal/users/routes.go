package users

import (
	"github.com/jimmymalark/go_api_template/internal/cache"
	db "github.com/jimmymalark/go_api_template/internal/db/sqlc"
	"github.com/jimmymalark/go_api_template/internal/validator"

	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(r chi.Router, queries *db.Queries, cache cache.Cache) {
	repository := NewRepository(queries)
	service := NewService(repository, cache)
	handler := NewHandler(service, validator.New())

	r.Route("/admin/users", func(r chi.Router) {
		r.Get("/", handler.ListUsers)
		r.Post("/", handler.CreateUser)
	})
}
