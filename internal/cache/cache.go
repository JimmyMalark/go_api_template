package cache

import (
	"context"
	"time"
)

type Cache interface {
	Get(ctx context.Context, key string, dest any) bool
	Set(ctx context.Context, key string, value any, ttl time.Duration) error

	Delete(ctx context.Context, key string) error
	DeletePattern(ctx context.Context, pattern string) error
}
