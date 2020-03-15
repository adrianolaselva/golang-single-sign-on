package repository

import (
	"github.com/jinzhu/gorm"
	"oauth2/src/common"
	"oauth2/src/models"
)

type AccessTokenRepository interface {
	Create(accessToken *models.AccessToken) error
	Update(accessToken *models.AccessToken) error
	FindById(id string) (*models.AccessToken, error)
	FindByAccessToken(token string) (*models.AccessToken, error)
	FindByUserId(userId string) (*models.AccessToken, error)
	Paginate(filters *map[string]interface{}, orderBy *string, orderDir *string, limit *int, page *int)  (*common.PaginationCommon, error)
}

type accessTokenRepository struct {
	conn *gorm.DB
}

func NewAccessTokenRepository(conn *gorm.DB) *accessTokenRepository {
	return &accessTokenRepository{conn}
}

func (a accessTokenRepository) Create(accessToken *models.AccessToken) error {
	result := a.conn.Create(&accessToken)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (a accessTokenRepository) Update(accessToken *models.AccessToken) error {
	result := a.conn.Model(&accessToken).Update(&accessToken)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (a accessTokenRepository) FindById(id string) (*models.AccessToken, error) {
	accessToken := models.AccessToken{}
	result := a.conn.Find(&accessToken, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &accessToken, nil
}

func (a accessTokenRepository) FindByAccessToken(token string) (*models.AccessToken, error) {
	accessToken := models.AccessToken{}
	result := a.conn.Where("access_token_id = ? ", token).First(&accessToken)
	if result != nil {
		return nil, result.Error
	}
	
	return &accessToken, nil
}

func (a accessTokenRepository) FindByUserId(userId string) (*models.AccessToken, error) {
	accessToken := models.AccessToken{}
	result := a.conn.Where("user_id = ? ", userId).Find(&accessToken)
	if result.Error != nil {
		return nil, result.Error
	}

	return &accessToken, nil
}

func (a accessTokenRepository) Paginate(filters *map[string]interface{}, orderBy *string, orderDir *string, limit *int, page *int)  (*common.PaginationCommon, error) {
	var databaseCommon common.Database

	rows, total, pages, err := databaseCommon.InitializePaginate(
		a.conn,
		&[]*models.AccessToken{},
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

	var accessTokens []*models.AccessToken
	for rows.Next() {
		var accessToken models.AccessToken
		err := a.conn.ScanRows(rows, &accessToken)
		if err != nil {
			return nil, err
		}
		accessTokens = append(accessTokens, &accessToken)
	}

	return &common.PaginationCommon{
		Current:      *page,
		PerPage:      *limit,
		TotalPages:   pages,
		TotalRecords: total,
		Data:         accessTokens,
	}, nil
}