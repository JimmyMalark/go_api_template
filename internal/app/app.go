package app

import (
	"context"
	"github.com/jimmymalark/go_api_template/internal/cache"
	"github.com/jimmymalark/go_api_template/internal/config"
	"github.com/jimmymalark/go_api_template/internal/db"
	sqlc "github.com/jimmymalark/go_api_template/internal/db/sqlc"
	"github.com/jimmymalark/go_api_template/internal/logger"
	"github.com/jimmymalark/go_api_template/internal/users"
	"github.com/jimmymalark/go_api_template/internal/validator"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httprate"

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

	pool, err := db.NewPostgres(ctx, cfg.DB)
	if err != nil {
		return nil, err
	}

	// redisClient := redis.NewClient(&redis.Options{
	// 	Addr: cfg.Cache.Addr(),
	// 	DB:   0,
	// })

	redis_client, err := cache.Client(ctx, cfg.Cache)
	if err != nil {
		return nil, err
	}
	cacheService := cache.New(redis_client)

	queries := sqlc.New(pool)

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Timeout(time.Second * 60))
	r.Use(httprate.LimitByIP(100, 1*time.Minute))
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

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
		Redis:  redis_client,
		Router: r,
	}, nil
}
