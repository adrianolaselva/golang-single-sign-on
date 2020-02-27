package repository

import (
	"github.com/jinzhu/gorm"
	"oauth2/src/common"
	"oauth2/src/models"
)

type RoleRepository interface {
	Create(role *models.Role) error
	Update(role *models.Role) error
	Delete(uuid string) error
	FindById(uuid string) (*models.Role, error)
	FindByName(username string) (*models.Role, error)
	Paginate(filters *map[string]interface{}, orderBy *string, orderDir *string, limit *int, page *int)  (*common.PaginationCommon, error)
}

type roleRepositoryImpl struct {
	conn *gorm.DB
}

func NewRoleRepository(conn *gorm.DB) *roleRepositoryImpl {
	return &roleRepositoryImpl{conn}
}

func (u *roleRepositoryImpl) Create(role *models.Role) error {
	result := u.conn.Create(&role)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *roleRepositoryImpl) Update(role *models.Role) error {
	result := u.conn.Model(&role).Updates(&role)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *roleRepositoryImpl) Delete(uuid string) error {
	result := u.conn.Where("id = ?", uuid).Delete(&models.Role{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *roleRepositoryImpl) FindById(uuid string) (*models.Role, error) {
	var role models.Role
	result := u.conn.Where("id = ?", uuid).First(&role)
	if result.Error != nil {
		return nil, result.Error
	}
	return &role, nil
}

func (u *roleRepositoryImpl) FindByName(username string) (*models.Role, error) {
	var role models.Role
	result := u.conn.Where("name = ?", username).First(&role)
	if result.Error != nil {
		return nil, result.Error
	}
	return &role, nil
}

func (u *roleRepositoryImpl) Paginate(filters *map[string]interface{}, orderBy *string, orderDir *string, limit *int, page *int)  (*common.PaginationCommon, error) {
	var databaseCommon common.Database

	rows, total, pages, err := databaseCommon.InitializePaginate(
		u.conn,
		&[]*models.Role{},
		filters,
		orderBy,
		orderDir,
		*limit,
		*page,
		"id",
		"ASC")

	if err != nil {
		return nil, err
	}

	var roles []*models.Role
	for rows.Next() {
		var user models.Role
		err := u.conn.ScanRows(rows, &user)
		if err != nil {
			return nil, err
		}
		roles = append(roles, &user)
	}

	return &common.PaginationCommon{
		Current:      *page,
		PerPage:      *limit,
		TotalPages:   pages,
		TotalRecords: total,
		Data:         roles,
	}, nil
}