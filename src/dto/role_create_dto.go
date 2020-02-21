package dto

import (
	"oauth2/src/models"
)

type RoleDto struct {
	Name	string	`json:"name"`
}

func (u *RoleDto) ToRole() *models.Role {
	return &models.Role{
		Name:	u.Name,
	}
}

