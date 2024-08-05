package service

import (
	"github.com/SaiHLu/rest-template/internal/core/dto"
	"github.com/SaiHLu/rest-template/internal/core/entity"
	gormscope "github.com/SaiHLu/rest-template/internal/core/gorm_scope"
	"github.com/SaiHLu/rest-template/internal/core/repository/user"
	"github.com/google/uuid"
)

type UserService interface {
	GetAll(dto.QueryUserDto) ([]entity.User, error)
	GetOneByEmail(string) (entity.User, error)
	GetOneById(uuid.UUID) (entity.User, error)
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

func (u *service) GetOneByEmail(email string) (entity.User, error) {
	conditions := []gormscope.Condition{
		{Type: gormscope.AndCondition, Query: "email = ? ", Args: email},
	}

	user, err := u.repo.GetOne(conditions...)
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (u *service) GetOneById(id uuid.UUID) (entity.User, error) {
	conditions := []gormscope.Condition{
		{Type: gormscope.AndCondition, Query: "id = ? ", Args: id},
	}

	user, err := u.repo.GetOne(conditions...)
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
