package repository

import (
	"github.com/jinzhu/gorm"
	"oauth2/src/models"
)

type AccessTokenRepository interface {
	Create(user *models.AccessToken) error
	Update(user *models.AccessToken) error
	FindById(id string) (*models.AccessToken, error)
	FindByAccessToken(accessToken string) (*models.AccessToken, error)
	FindByUserId(id string) (*models.AccessToken, error)
}

type accessTokenRepository struct {
	conn *gorm.DB
}

func NewAccessTokenRepository(conn *gorm.DB) *accessTokenRepository {
	return &accessTokenRepository{conn}
}

func (a accessTokenRepository) Create(user *models.AccessToken) error {
	panic("implement me")
}

func (a accessTokenRepository) Update(user *models.AccessToken) error {
	panic("implement me")
}

func (a accessTokenRepository) FindById(id string) (*models.AccessToken, error) {
	panic("implement me")
}

func (a accessTokenRepository) FindByAccessToken(accessToken string) (*models.AccessToken, error) {
	panic("implement me")
}

func (a accessTokenRepository) FindByUserId(id string) (*models.AccessToken, error) {
	panic("implement me")
}