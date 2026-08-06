package sessions

import (
	"context"

	"github.com/jimmymalark/go_api_template/internal/apperrors"
	db "github.com/jimmymalark/go_api_template/internal/db/sqlc"
)

type Repository struct {
	queries *db.Queries
}

type SessionRepository interface {
	CreateSession(ctx context.Context, params db.CreateSessionParams) (db.Session, error)
	GetSessionByTokenHash(ctx context.Context, tokenHash string) (db.Session, error)
	DeleteSession(ctx context.Context, id int64) error
	DeleteSessionByTokenHash(ctx context.Context, tokenHash string) error
	DeleteSessionsByUserID(ctx context.Context, userID int64) error
}

func NewRepository(q *db.Queries) *Repository {
	return &Repository{
		queries: q,
	}
}

func (r *Repository) CreateSession(
	ctx context.Context,
	params db.CreateSessionParams,
) (db.Session, error) {
	session, err := r.queries.CreateSession(ctx, params)
	if err != nil {
		return db.Session{}, apperrors.FromPostgres(err)
	}

	return session, nil
}

func (r *Repository) GetSessionByTokenHash(
	ctx context.Context,
	tokenHash []byte,
) (db.Session, error) {
	session, err := r.queries.GetSessionByTokenHash(ctx, tokenHash)
	if err != nil {
		return db.Session{}, apperrors.FromPostgres(err)
	}

	return session, nil
}

func (r *Repository) DeleteSession(
	ctx context.Context,
	id int64,
) error {
	if err := r.queries.DeleteSession(ctx, id); err != nil {
		return apperrors.FromPostgres(err)
	}

	return nil
}

func (r *Repository) DeleteSessionsByUserID(
	ctx context.Context,
	userID int64,
) error {
	if err := r.queries.DeleteSessionsByUserID(ctx, userID); err != nil {
		return apperrors.FromPostgres(err)
	}

	return nil
}

func (r *Repository) DeleteSessionByTokenHash(
	ctx context.Context,
	tokenHash []byte,
) error {
	if err := r.queries.DeleteSessionByTokenHash(ctx, tokenHash); err != nil {
		return apperrors.FromPostgres(err)
	}

	return nil
}
