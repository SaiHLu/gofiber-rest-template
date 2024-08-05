package user

import (
	"errors"

	"github.com/SaiHLu/rest-template/common"
	"github.com/SaiHLu/rest-template/common/logger"
	"github.com/SaiHLu/rest-template/internal/core/dto"
	"github.com/SaiHLu/rest-template/internal/core/entity"
	gormscope "github.com/SaiHLu/rest-template/internal/core/gorm_scope"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type postgresRepository struct {
	db *gorm.DB
}

func NewPostgresRepository(db *gorm.DB) UserRepository {
	return &postgresRepository{
		db: db,
	}
}

func (r *postgresRepository) GetAll(query dto.QueryUserDto) ([]entity.User, error) {
	var (
		users []entity.User
		count int64
	)

	limit, _ := common.ConvertStringToInt(query.Limit)
	offset, _ := common.ConvertStringToInt(query.Offset)

	if err := r.db.Model(&entity.User{}).Scopes(gormscope.QueryFilter(query.Query), gormscope.OrderBy(query.Order), gormscope.Paginate(limit, offset)).Find(&users).Error; err != nil {
		logger.Error(err.Error())
		return users, err
	}

	if err := r.db.Model(&entity.User{}).Scopes(gormscope.QueryFilter(query.Query), gormscope.OrderBy(query.Order)).Count(&count).Error; err != nil {
		logger.Error(err.Error())
		return users, err
	}

	return users, nil
}

func (r *postgresRepository) GetOne(conditions ...gormscope.Condition) (entity.User, error) {
	var user entity.User

	if err := r.db.Model(&entity.User{}).Scopes(gormscope.WhereCondition(conditions...)).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *postgresRepository) Create(body dto.CreateUserDto) (entity.User, error) {
	var user = entity.User{
		Name:     &body.Name,
		Email:    body.Email,
		Password: body.Password,
	}

	result := r.db.Model(&entity.User{}).Create(&user)

	if result.Error != nil {
		return entity.User{}, result.Error
	}

	return user, nil
}

func (r *postgresRepository) Update(id uint, body dto.UpdateUserDto) (entity.User, error) {
	var user entity.User

	result := r.db.Model(&user).Clauses(clause.Returning{}).Where("id = ?", id).Updates(&body)

	if result.RowsAffected <= 0 {
		return entity.User{}, errors.New("no records found")
	}

	if result.Error != nil {
		return entity.User{}, result.Error
	}

	return user, nil
}

func (r *postgresRepository) Delete(id uint) (entity.User, error) {
	var user entity.User

	result := r.db.Clauses(clause.Returning{}).Where("id = ?", id).Delete(&user)

	if result.RowsAffected <= 0 {
		return entity.User{}, errors.New("no records found")
	}

	if result.Error != nil {
		return entity.User{}, result.Error
	}

	return user, nil
}
