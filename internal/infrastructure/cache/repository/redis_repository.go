package repository

import (
	"time"

	"github.com/gofiber/storage/redis/v3"
)

type RedisCacheService struct {
	client *redis.Storage
}

func NewRedisCacheService(client *redis.Storage) CacheRepository {
	return &RedisCacheService{client: client}
}

func (r *RedisCacheService) Set(key string, val []byte, exp time.Duration) error {
	if err := r.client.Set(key, val, exp); err != nil {
		return err
	}

	return nil
}

func (r *RedisCacheService) Get(key string) ([]byte, error) {
	data, err := r.client.Get(key)
	if err != nil {
		return []byte{}, err
	}

	return data, nil
}

func (r *RedisCacheService) Delete(key string) error {
	if err := r.client.Delete(key); err != nil {
		return err
	}

	return nil
}
