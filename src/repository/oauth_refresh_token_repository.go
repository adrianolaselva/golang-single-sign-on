package repository

import (
	"github.com/jinzhu/gorm"
	"oauth2/src/models"
)

type RefreshTokenRepository interface {
	Create(refreshToken *models.RefreshToken) error
	Update(refreshToken *models.RefreshToken) error
	FindById(id string) (*models.RefreshToken, error)
	FindByAccessToken(token string) (*models.RefreshToken, error)
	FindByRefreshToken(token string) (*models.RefreshToken, error)
}

type refreshTokenRepository struct {
	conn *gorm.DB
}

func NewRefreshTokenRepository(conn *gorm.DB) *refreshTokenRepository {
	return &refreshTokenRepository{conn}
}

func (r refreshTokenRepository) Create(refreshToken *models.RefreshToken) error {
	result := r.conn.Create(&refreshToken)
	if result != nil {
		return result.Error
	}

	return nil
}

func (r refreshTokenRepository) Update(refreshToken *models.RefreshToken) error {
	result := r.conn.Model(&refreshToken).Update(&refreshToken)
	if result != nil {
		return result.Error
	}

	return nil
}

func (r refreshTokenRepository) FindById(id string) (*models.RefreshToken, error) {
	refreshToken := models.RefreshToken{}
	result := r.conn.Find(&refreshToken, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &refreshToken, nil
}

func (r refreshTokenRepository) FindByAccessToken(accessToken string) (*models.RefreshToken, error) {
	refreshToken := models.RefreshToken{}
	result := r.conn.Where("access_token_id = ?", accessToken).First(&refreshToken)
	if result != nil {
		return nil, result.Error
	}

	return &refreshToken, nil
}

func (r refreshTokenRepository) FindByRefreshToken(token string) (*models.RefreshToken, error) {
	refreshToken := models.RefreshToken{}
	result := r.conn.Where("refresh_token = ?", token).Find(&refreshToken)
	if result.Error != nil {
		return nil, result.Error
	}

	return &refreshToken, nil
}
