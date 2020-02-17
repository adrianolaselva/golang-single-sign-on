package repository

import (
	"github.com/jinzhu/gorm"
	"oauth2/src/models"
)

type UserRepository interface {
	Create(user *models.User) error
	Update(user *models.User) error
	Delete(uuid string) error
	FindById(uuid string) (*models.User, error)
	FindByUsername(username string) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
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
	result := u.conn.Delete(&models.User{}).Where("uuid = ?", uuid)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *userRepositoryImpl) FindById(uuid string) (*models.User, error) {
	var user *models.User
	result := u.conn.First(&user).Where("uuid = ?", uuid)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (u *userRepositoryImpl) FindByUsername(username string) (*models.User, error) {
	var user *models.User
	result := u.conn.First(&user).Where("username = ?", username)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (u *userRepositoryImpl) FindByEmail(email string) (*models.User, error) {
	var user *models.User
	result := u.conn.First(&user).Where("email = ?", email)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}