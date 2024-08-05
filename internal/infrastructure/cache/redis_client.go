package cache

import (
	"context"
	"time"

	"github.com/gofiber/storage/redis/v3"
)

type redisCache struct {
	client *redis.Storage
}

func NewRedisCache(config redis.Config) Cache {
	store := redis.New(config)

	return &redisCache{client: store}
}

func (r *redisCache) Set(key string, val []byte, exp time.Duration) error {
	if err := r.client.Set(key, val, exp); err != nil {
		return err
	}

	return nil
}

func (r *redisCache) Get(key string) ([]byte, error) {
	data, err := r.client.Get(key)
	if err != nil {
		return []byte{}, err
	}

	return data, nil
}

func (r *redisCache) Delete(key string) error {
	if err := r.client.Delete(key); err != nil {
		return err
	}

	return nil
}

func (r *redisCache) Ping() error {
	pingStatus := r.client.Conn().Ping(context.Background())
	if pingStatus.Err() != nil {
		return pingStatus.Err()
	}

	return nil
}

func (r *redisCache) Close() error {
	return r.client.Close()
}
