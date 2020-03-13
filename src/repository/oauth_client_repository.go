package repository

import (
	"github.com/jinzhu/gorm"
	"oauth2/src/models"
)

type ClientRepository interface {
	Create(user *models.Client) error
	Update(user *models.Client) error
	FindById(id string) (*models.Client, error)
}

type clientRepository struct {
	conn *gorm.DB
}

func NewClientRepository(conn *gorm.DB) *clientRepository {
	return &clientRepository{conn}
}

func (c clientRepository) Create(user *models.Client) error {
	panic("implement me")
}

func (c clientRepository) Update(user *models.Client) error {
	panic("implement me")
}

func (c clientRepository) FindById(id string) (*models.Client, error) {
	panic("implement me")
}
