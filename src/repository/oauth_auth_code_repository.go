package repository

import (
	"github.com/jinzhu/gorm"
	"oauth2/src/models"
)

type AuthCodeRepository interface {
	Create(user *models.AuthCode) error
	Update(user *models.AuthCode) error
	FindById(id string) (*models.AuthCode, error)
	FindByCode(id string) (*models.AuthCode, error)
	FindByUserId(id string) (*models.AuthCode, error)
}

type authCodeRepository struct {
	conn *gorm.DB
}

func NewAuthCodeRepository(conn *gorm.DB) *authCodeRepository {
	return &authCodeRepository{conn}
}

func (a authCodeRepository) Create(user *models.AuthCode) error {
	panic("implement me")
}

func (a authCodeRepository) Update(user *models.AuthCode) error {
	panic("implement me")
}

func (a authCodeRepository) FindById(id string) (*models.AuthCode, error) {
	panic("implement me")
}

func (a authCodeRepository) FindByCode(id string) (*models.AuthCode, error) {
	panic("implement me")
}

func (a authCodeRepository) FindByUserId(id string) (*models.AuthCode, error) {
	panic("implement me")
}
