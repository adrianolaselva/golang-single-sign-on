package dto

import (
	"gopkg.in/go-playground/validator.v9"
	"oauth2/src/models"
)

type RoleDto struct {
	Name	string	`json:"name" validate:"required,min=2,max=100"`
}

func (u *RoleDto) ToRole() (*models.Role, error) {
	v := validator.New()
	err := v.Struct(u)

	if err != nil {
		return nil, err
	}

	return &models.Role{
		Name:	u.Name,
	}, nil
}

