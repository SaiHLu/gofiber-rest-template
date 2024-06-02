package service

import (
	"time"

	"github.com/SaiHLu/rest-template/internal/infrastructure/cache/repository"
)

type CacheService interface {
	Set(key string, val []byte, exp time.Duration) error
	Get(key string) ([]byte, error)
	Delete(key string) error
}

type service struct {
	repo repository.CacheRepository
}

func NewCacheService(repo repository.CacheRepository) CacheService {
	return &service{repo: repo}
}

func (r *service) Set(key string, val []byte, exp time.Duration) error {
	if err := r.repo.Set(key, val, exp); err != nil {
		return err
	}

	return nil
}

func (r *service) Get(key string) ([]byte, error) {
	data, err := r.repo.Get(key)
	if err != nil {
		return []byte{}, err
	}

	return data, nil
}

func (r *service) Delete(key string) error {
	if err := r.repo.Delete(key); err != nil {
		return err
	}

	return nil
}
