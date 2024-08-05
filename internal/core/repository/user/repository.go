package user

import (
	"github.com/SaiHLu/rest-template/internal/core/dto"
	"github.com/SaiHLu/rest-template/internal/core/entity"
	gormscope "github.com/SaiHLu/rest-template/internal/core/gorm_scope"
)

type UserRepository interface {
	GetAll(dto.QueryUserDto) ([]entity.User, error)
	GetOne(...gormscope.Condition) (entity.User, error)
	Create(dto.CreateUserDto) (entity.User, error)
	Delete(uint) (entity.User, error)
	Update(uint, dto.UpdateUserDto) (entity.User, error)
}
