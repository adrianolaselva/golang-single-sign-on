package repository

import (
	"github.com/jinzhu/gorm"
	"oauth2/src/models"
)

type RefreshTokenRepository interface {
	Create(user *models.RefreshToken) error
	Update(user *models.RefreshToken) error
	FindById(id string) (*models.RefreshToken, error)
	FindByAccessToken(accessToken string) (*models.RefreshToken, error)
	FindByRefreshToken(refreshToken string) (*models.RefreshToken, error)
	FindByCode(id string) (*models.RefreshToken, error)
}

type refreshTokenRepository struct {
	conn *gorm.DB
}

func NewRefreshTokenRepository(conn *gorm.DB) *refreshTokenRepository {
	return &refreshTokenRepository{conn}
}

func (r refreshTokenRepository) Create(user *models.RefreshToken) error {
	panic("implement me")
}

func (r refreshTokenRepository) Update(user *models.RefreshToken) error {
	panic("implement me")
}

func (r refreshTokenRepository) FindById(id string) (*models.RefreshToken, error) {
	panic("implement me")
}

func (r refreshTokenRepository) FindByAccessToken(accessToken string) (*models.RefreshToken, error) {
	panic("implement me")
}

func (r refreshTokenRepository) FindByRefreshToken(refreshToken string) (*models.RefreshToken, error) {
	panic("implement me")
}

func (r refreshTokenRepository) FindByCode(id string) (*models.RefreshToken, error) {
	panic("implement me")
}
