package repository

import "time"

type CacheRepository interface {
	Set(key string, val []byte, exp time.Duration) error
	Get(key string) ([]byte, error)
	Delete(key string) error
}
