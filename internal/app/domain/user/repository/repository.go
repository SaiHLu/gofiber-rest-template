package repository

import (
	"github.com/SaiHLu/rest-template/internal/app/domain/user/dto"
	"github.com/SaiHLu/rest-template/internal/app/entity"
)

type UserRepository interface {
	GetAll(dto.QueryUserDto) ([]entity.User, error)
	Create(dto.CreateUserDto) (entity.User, error)
}
