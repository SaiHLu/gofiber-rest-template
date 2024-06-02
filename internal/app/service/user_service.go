package service

import (
	"github.com/SaiHLu/rest-template/internal/app/dto"
	"github.com/SaiHLu/rest-template/internal/app/entity"
	"github.com/SaiHLu/rest-template/internal/app/repository/user"
)

type UserService interface {
	GetAll(dto.QueryUserDto) ([]entity.User, error)
	GetOne(map[string]interface{}) (entity.User, error)
	Create(dto.CreateUserDto) (entity.User, error)
	Delete(uint) (entity.User, error)
	Update(uint, dto.UpdateUserDto) (entity.User, error)
}

type service struct {
	repo user.UserRepository
}

func NewUserService(repo user.UserRepository) UserService {
	return &service{
		repo: repo,
	}
}

func (u *service) GetAll(query dto.QueryUserDto) ([]entity.User, error) {
	users, err := u.repo.GetAll(query)
	if err != nil {
		return []entity.User{}, err
	}

	return users, nil
}

func (u *service) GetOne(conditions map[string]interface{}) (entity.User, error) {
	user, err := u.repo.GetOne(conditions)
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (u *service) Create(body dto.CreateUserDto) (entity.User, error) {
	user, err := u.repo.Create(body)
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (u *service) Update(id uint, body dto.UpdateUserDto) (entity.User, error) {
	user, err := u.repo.Update(id, body)
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (u *service) Delete(id uint) (entity.User, error) {
	user, err := u.repo.Delete(id)
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}
