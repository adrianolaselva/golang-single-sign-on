package models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Role struct {
	Name string	`gorm:"column:name;type:varchar(100);primary_key;"`
}

func (Role) TableName() string {
	return "roles"
}

func (Role) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.New()
	scope.SetColumn("Uuid", uuid)
	return nil
}