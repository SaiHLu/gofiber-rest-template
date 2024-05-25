package postgresrepository

import (
	"github.com/SaiHLu/rest-template/common"
	"github.com/SaiHLu/rest-template/internal/app/domain/user/dto"
	"github.com/SaiHLu/rest-template/internal/app/domain/user/repository"
	"github.com/SaiHLu/rest-template/internal/app/entity"
	"gorm.io/gorm"
)

type postgresRepository struct {
	db *gorm.DB
}

func NewPostgresRepository(db *gorm.DB) repository.UserRepository {
	return &postgresRepository{
		db: db,
	}
}

func (r *postgresRepository) GetAll(query dto.QueryUserDto) ([]entity.User, error) {
	var users []entity.User

	page, _ := common.ConvertStringToInt(query.Page)
	pageSize, _ := common.ConvertStringToInt(query.PageSize)

	result := r.db.Model(&entity.User{}).Scopes(common.Paginate(page, pageSize)).Find(&users)

	if result.Error != nil {
		return []entity.User{}, result.Error
	}

	return users, nil
}

func (r *postgresRepository) Create(body dto.CreateUserDto) (entity.User, error) {
	var user = entity.User{
		Name:  &body.Name,
		Email: body.Email,
	}

	result := r.db.Model(&entity.User{}).Create(&user)

	if result.Error != nil {
		return entity.User{}, result.Error
	}

	return user, nil
}
