package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisCache struct {
	client *redis.Client
}

func New(client *redis.Client) Cache {
	return &RedisCache{
		client: client,
	}
}

func (r *RedisCache) Get(
	ctx context.Context,
	key string,
	dest any,
) bool {

	data, err := r.client.Get(ctx, key).Bytes()
	if err != nil {
		return false
	}

	if err := json.Unmarshal(data, dest); err != nil {
		return false
	}

	return true
}

func (r *RedisCache) Set(
	ctx context.Context,
	key string,
	value any,
	ttl time.Duration,
) error {

	data, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return r.client.Set(ctx, key, data, ttl).Err()
}

func (r *RedisCache) Delete(ctx context.Context, key string) error {
	return r.client.Del(ctx, key).Err()
}

func (r *RedisCache) DeletePattern(ctx context.Context, pattern string) error {
	var cursor uint64

	for {
		keys, nextCursor, err := r.client.Scan(ctx, cursor, pattern, 100).Result()
		if err != nil {
			return err
		}

		if len(keys) > 0 {
			if err := r.client.Del(ctx, keys...).Err(); err != nil {
				return err
			}
		}

		cursor = nextCursor
		if cursor == 0 {
			break
		}
	}

	return nil
}
