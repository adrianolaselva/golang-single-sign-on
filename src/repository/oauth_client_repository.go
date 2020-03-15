package repository

import (
	"github.com/jinzhu/gorm"
	"oauth2/src/models"
)

type ClientRepository interface {
	Create(client *models.Client) error
	Update(client *models.Client) error
	FindById(id string) (*models.Client, error)
}

type clientRepository struct {
	conn *gorm.DB
}

func NewClientRepository(conn *gorm.DB) *clientRepository {
	return &clientRepository{conn}
}

func (c clientRepository) Create(client *models.Client) error {
	result := c.conn.Create(&client)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (c clientRepository) Update(client *models.Client) error {
	result := c.conn.Model(&client).Update(&client)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (c clientRepository) FindById(id string) (*models.Client, error) {
	client := models.Client{}
	result := c.conn.Where("id = ?", id).Preload("User").First(&client)
	if result.Error != nil {
		return nil, result.Error
	}

	return &client, nil
}
