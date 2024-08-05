package cache

import "time"

type Cache interface {
	Get(key string) ([]byte, error)
	Set(key string, value []byte, duration time.Duration) error
	Delete(key string) error
	Ping() error
	Close() error
}
