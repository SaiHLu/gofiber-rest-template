package user_repository

import (
	"errors"
	"fmt"

	user_model "github.com/SaiHLu/rest-template/internal/users/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	Get() (*[]user_model.UserModel, error)
	Create(data *user_model.UserModel) (*user_model.UserModel, error)
}

type repository struct {
	DB *gorm.DB
}

func NewUserRepository(dbInstance *gorm.DB) UserRepository {
	return &repository{
		DB: dbInstance,
	}
}

func (r *repository) Get() (*[]user_model.UserModel, error) {
	users := []user_model.UserModel{}

	result := r.DB.Table("users").Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	return &users, nil
}

func (r *repository) Create(data *user_model.UserModel) (*user_model.UserModel, error) {
	result := r.DB.Table("users").Save(&data)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, fmt.Errorf("%s already exists", data.Email)
		}

		return nil, result.Error
	}

	return data, nil
}
