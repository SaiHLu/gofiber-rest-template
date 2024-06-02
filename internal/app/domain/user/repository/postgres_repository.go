package repository

import (
	"errors"

	"github.com/SaiHLu/rest-template/common"
	"github.com/SaiHLu/rest-template/internal/app/domain/user/dto"
	"github.com/SaiHLu/rest-template/internal/app/entity"
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
	var users []entity.User

	page, _ := common.ConvertStringToInt(query.Page)
	pageSize, _ := common.ConvertStringToInt(query.PageSize)

	result := r.db.Model(&entity.User{}).Scopes(common.Paginate(page, pageSize)).Find(&users)

	if result.Error != nil {
		return []entity.User{}, result.Error
	}

	return users, nil
}

func (r *postgresRepository) GetOne(conditions map[string]interface{}) (entity.User, error) {
	var user entity.User

	if err := r.db.Model(&entity.User{}).Where(conditions).First(&user).Error; err != nil {
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
