package db

import (
	sqlc "github.com/jimmymalark/go_api_template/internal/db/sqlc"
	"github.com/jimmymalark/go_api_template/internal/sessions"
	"github.com/jimmymalark/go_api_template/internal/users"
)

type Repositories struct {
	Users    *users.Repository
	Sessions *sessions.Repository
}

func NewRepositories(q *sqlc.Queries) *Repositories {
	return &Repositories{
		Users:    users.NewRepository(q),
		Sessions: sessions.NewRepository(q),
	}
}
