package repository

import (
	"github.com/SaiHLu/rest-template/internal/app/domain/user/dto"
	"github.com/SaiHLu/rest-template/internal/app/entity"
)

type UserRepository interface {
	GetAll(dto.QueryUserDto) ([]entity.User, error)
	GetOne(map[string]interface{}) (entity.User, error)
	Create(dto.CreateUserDto) (entity.User, error)
	Delete(uint) (entity.User, error)
	Update(uint, dto.UpdateUserDto) (entity.User, error)
}
