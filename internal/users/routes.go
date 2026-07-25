package users

import (
	"dbApi/internal/cache"
	db "dbApi/internal/db/sqlc"
	"dbApi/internal/validator"

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
