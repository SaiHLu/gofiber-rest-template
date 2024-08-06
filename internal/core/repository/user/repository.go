package user

import (
	"github.com/SaiHLu/rest-template/internal/core/dto"
	"github.com/SaiHLu/rest-template/internal/core/entity"
	gormscope "github.com/SaiHLu/rest-template/internal/core/gorm_scope"
	"github.com/google/uuid"
)

type UserRepository interface {
	GetAll(dto.QueryUserDto) ([]entity.User, error)
	GetOne(...gormscope.Condition) (entity.User, error)
	Create(dto.CreateUserDto) (entity.User, error)
	Delete(uuid.UUID) (entity.User, error)
	Update(uuid.UUID, dto.UpdateUserDto) (entity.User, error)
}
