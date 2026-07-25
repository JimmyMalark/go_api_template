package db

import (
	"context"
	"github.com/jimmymalark/go_api_template/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgres(ctx context.Context, cfg config.DBConfig) (*pgxpool.Pool, error) {
	return pgxpool.New(ctx, cfg.Addr())
}
