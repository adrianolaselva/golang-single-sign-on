package repository

import (
	"github.com/jinzhu/gorm"
	"oauth2/src/common"
	"oauth2/src/models"
)

type UserRepository interface {
	Create(user *models.User) error
	Update(user *models.User) error
	Delete(uuid string) error
	FindById(uuid string) (*models.User, error)
	FindByUsername(username string) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	Paginate(filters *map[string]interface{}, orderBy *string, orderDir *string, limit *int, page *int)  (*common.PaginationCommon, error)
	LoadRolesById(uuid string) ([]*models.Role, error)
	LoadClientsById(uuid string) ([]*models.Client, error)
}

type userRepositoryImpl struct {
	conn *gorm.DB
}

func NewUserRepository(conn *gorm.DB) *userRepositoryImpl {
	return &userRepositoryImpl{conn}
}

func (u *userRepositoryImpl) Create(user *models.User) error {
	result := u.conn.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *userRepositoryImpl) Update(user *models.User) error {
	result := u.conn.Model(&user).Updates(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *userRepositoryImpl) Delete(uuid string) error {
	result := u.conn.Where("id = ?", uuid).Delete(&models.User{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *userRepositoryImpl) FindById(uuid string) (*models.User, error) {
	var user models.User
	result := u.conn.Where("id = ?", uuid).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (u *userRepositoryImpl) FindByUsername(username string) (*models.User, error) {
	var user models.User
	result := u.conn.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (u *userRepositoryImpl) FindByEmail(email string) (*models.User, error) {
	var user models.User
	result := u.conn.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (u *userRepositoryImpl) Paginate(filters *map[string]interface{}, orderBy *string, orderDir *string, limit *int, page *int)  (*common.PaginationCommon, error) {
	var databaseCommon common.Database

	rows, total, pages, err := databaseCommon.InitializePaginate(
		u.conn,
		&[]*models.User{},
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

	var users []*models.User
	for rows.Next() {
		var user models.User
		err := u.conn.ScanRows(rows, &user)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	return &common.PaginationCommon{
		Current:      *page,
		PerPage:      *limit,
		TotalPages:   pages,
		TotalRecords: total,
		Data:         users,
	}, nil
}

func (u *userRepositoryImpl) LoadRolesById(uuid string) ([]*models.Role, error) {
	var user models.User
	result := u.conn.Preload("Roles").Where("id = ?", uuid).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user.Roles, nil
}

func (u *userRepositoryImpl) LoadClientsById(uuid string) ([]*models.Client, error) {
	var user models.User
	result := u.conn.Preload("Clients").Where("id = ?", uuid).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user.Clients, nil
}
