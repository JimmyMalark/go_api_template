package db

import (
	"context"
	"dbApi/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgres(ctx context.Context, cfg config.DBConfig) (*pgxpool.Pool, error) {
	return pgxpool.New(ctx, cfg.Addr())
}
