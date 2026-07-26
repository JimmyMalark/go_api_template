package db

import (
	"context"

	sqlc "github.com/jimmymalark/go_api_template/internal/db/sqlc"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Store struct {
	db *pgxpool.Pool
}

func NewStore(db *pgxpool.Pool) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) Transaction(
	ctx context.Context,
	fn func(r *Repositories) error,
) error {

	tx, err := s.db.Begin(ctx)
	if err != nil {
		return err
	}

	defer tx.Rollback(ctx)

	repos := NewRepositories(sqlc.New(tx))

	if err := fn(repos); err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func RunInTransaction[T any](
	ctx context.Context,
	store *Store,
	fn func(*Repositories) (T, error),
) (T, error) {
	tx, err := store.db.Begin(ctx)
	if err != nil {
		var zero T
		return zero, err
	}

	defer func() {
		_ = tx.Rollback(ctx)
	}()

	repos := NewRepositories(sqlc.New(tx))

	result, err := fn(repos)
	if err != nil {
		var zero T
		return zero, err
	}

	if err := tx.Commit(ctx); err != nil {
		var zero T
		return zero, err
	}

	return result, nil
}
