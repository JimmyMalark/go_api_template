package users

import (
	"context"

	"github.com/jimmymalark/go_api_template/internal/apperrors"
	db "github.com/jimmymalark/go_api_template/internal/db/sqlc"
	"github.com/jimmymalark/go_api_template/internal/pagination"
)

type Repository struct {
	queries *db.Queries
}

type UserRepository interface {
	ListUsers(ctx context.Context, p pagination.Params) ([]db.User, error)
	CreateUser(ctx context.Context, params db.CreateUserParams) (db.User, error)

	GetUserByEmail(ctx context.Context, email string) (db.User, error)
	GetUserByUsername(ctx context.Context, username string) (db.User, error)
	GetUserByID(ctx context.Context, id int64) (db.User, error)

	CountUsers(ctx context.Context) (int64, error)
}

func NewRepository(q *db.Queries) *Repository {
	return &Repository{
		queries: q,
	}
}

func (r *Repository) ListUsers(
	ctx context.Context,
	p pagination.Params,
) ([]db.User, error) {
	offset := (p.Page - 1) * p.Limit

	users, err := r.queries.ListUsers(ctx, db.ListUsersParams{
		Limit:  int32(p.Limit),
		Offset: int32(offset),
	})
	if err != nil {
		return nil, apperrors.FromPostgres(err)
	}

	return users, nil
}

func (r *Repository) CreateUser(
	ctx context.Context,
	params db.CreateUserParams,
) (db.User, error) {
	user, err := r.queries.CreateUser(ctx, params)
	if err != nil {
		return db.User{}, apperrors.FromPostgres(err)
	}

	return user, nil
}

func (r *Repository) CountUsers(
	ctx context.Context,
) (int64, error) {
	return r.queries.CountUsers(ctx)
}

func (r *Repository) GetUserByEmail(
	ctx context.Context,
	email string,
) (db.User, error) {
	user, err := r.queries.GetUserByEmail(ctx, email)
	if err != nil {
		return db.User{}, apperrors.FromPostgres(err)
	}

	return user, nil
}

func (r *Repository) GetUserByUsername(
	ctx context.Context,
	username string,
) (db.User, error) {
	user, err := r.queries.GetUserByUsername(ctx, username)
	if err != nil {
		return db.User{}, apperrors.FromPostgres(err)
	}

	return user, nil
}

func (r *Repository) GetUserByID(
	ctx context.Context,
	id int64,
) (db.User, error) {
	user, err := r.queries.GetUserByID(ctx, id)
	if err != nil {
		return db.User{}, apperrors.FromPostgres(err)
	}

	return user, nil
}
