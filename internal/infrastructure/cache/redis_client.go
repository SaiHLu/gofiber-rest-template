package cache

import (
	"context"

	"github.com/gofiber/storage/redis/v3"
)

type RedisClient struct {
	Client *redis.Storage
}

func NewRedisClient(config redis.Config) *RedisClient {
	store := redis.New(config)

	return &RedisClient{Client: store}
}

func (r *RedisClient) Ping(ctx context.Context) error {
	pingStatus := r.Client.Conn().Ping(ctx)
	if pingStatus.Err() != nil {
		return pingStatus.Err()
	}

	return nil
}
