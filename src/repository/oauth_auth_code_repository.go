package repository

import (
	"github.com/jinzhu/gorm"
	"oauth2/src/models"
)

type AuthCodeRepository interface {
	Create(authCode *models.AuthCode) error
	Update(authCode *models.AuthCode) error
	FindById(id string) (*models.AuthCode, error)
	FindByCode(code string) (*models.AuthCode, error)
	FindByUserId(userId string) (*models.AuthCode, error)
}

type authCodeRepository struct {
	conn *gorm.DB
}

func NewAuthCodeRepository(conn *gorm.DB) *authCodeRepository {
	return &authCodeRepository{conn}
}

func (a authCodeRepository) Create(authCode *models.AuthCode) error {
	result := a.conn.Create(&authCode)
	if result.Error != nil {
		return result. Error
	}

	return nil
}

func (a authCodeRepository) Update(authCode *models.AuthCode) error {
	result := a.conn.Model(&authCode).Update(&authCode)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (a authCodeRepository) FindById(id string) (*models.AuthCode, error) {
	authCode := models.AuthCode{}
	result := a.conn.Find(&authCode, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &authCode, nil
}

func (a authCodeRepository) FindByCode(code string) (*models.AuthCode, error) {
	authCode := models.AuthCode{}
	result := a.conn.Where("code = ? ", code).First(&authCode)
	if result.Error != nil {
		return nil, result.Error
	}

	return &authCode, nil
}

func (a authCodeRepository) FindByUserId(userId string) (*models.AuthCode, error) {
	authCode := models.AuthCode{}
	result := a.conn.Where("user_id = ? ", userId).First(&authCode)
	if result.Error != nil {
		return nil, result.Error
	}

	return &authCode, nil
}
