package user_service

import (
	user_model "github.com/SaiHLu/rest-template/internal/users/models"
	user_repository "github.com/SaiHLu/rest-template/internal/users/repository"
)

type UserService interface {
	user_repository.UserRepository
}

type service struct {
	repository user_repository.UserRepository
}

func NewUserService(r user_repository.UserRepository) UserService {
	return &service{
		repository: r,
	}
}

func (s *service) Get() (*[]user_model.UserModel, error) {
	return s.repository.Get()
}

func (s *service) Create(data *user_model.UserModel) (*user_model.UserModel, error) {
	return s.repository.Create(data)
}
