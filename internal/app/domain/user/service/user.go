package service

import (
	"github.com/SaiHLu/rest-template/internal/app/domain/user/dto"
	"github.com/SaiHLu/rest-template/internal/app/domain/user/repository"
	"github.com/SaiHLu/rest-template/internal/app/entity"
)

type Service interface {
	GetAll(dto.QueryUserDto) ([]entity.User, error)
	Create(dto.CreateUserDto) (entity.User, error)
	Delete(uint) (entity.User, error)
	Update(uint, dto.UpdateUserDto) (entity.User, error)
}

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) Service {
	return &UserService{
		repo: repo,
	}
}

func (u *UserService) GetAll(query dto.QueryUserDto) ([]entity.User, error) {
	users, err := u.repo.GetAll(query)
	if err != nil {
		return []entity.User{}, err
	}

	return users, nil
}

func (u *UserService) Create(body dto.CreateUserDto) (entity.User, error) {
	user, err := u.repo.Create(body)
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (u *UserService) Update(id uint, body dto.UpdateUserDto) (entity.User, error) {
	user, err := u.repo.Update(id, body)
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (u *UserService) Delete(id uint) (entity.User, error) {
	user, err := u.repo.Delete(id)
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}
