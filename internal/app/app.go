package app

import (
	"context"
	"dbApi/internal/cache"
	"dbApi/internal/config"
	"dbApi/internal/db"
	sqlc "dbApi/internal/db/sqlc"
	"dbApi/internal/logger"
	"dbApi/internal/users"
	"dbApi/internal/validator"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	httpSwagger "github.com/swaggo/http-swagger"
)

type App struct {
	Config *config.Config

	DB *pgxpool.Pool

	Redis *redis.Client

	Validator *validator.Validator

	Router chi.Router
}

func New(ctx context.Context) (*App, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}

	pool, err := db.NewPostgres(ctx, cfg.Database)
	if err != nil {
		return nil, err
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr: cfg.Cache.CSN(),
		DB:   0,
	})

	cacheService := cache.New(redisClient)

	queries := sqlc.New(pool)

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Route("/v1", func(r chi.Router) {
		users.RegisterRoutes(r, queries, cacheService)
	})

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
	))

	logger.New(logger.Config{
		Level: cfg.App.LogLevel,
		JSON:  cfg.App.LogJSON,
	})

	return &App{
		Config: cfg,
		DB:     pool,
		Redis:  redisClient,
		Router: r,
	}, nil
}
