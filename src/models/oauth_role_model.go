package models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Role struct {
	ID      string 	`json:"id"gorm:"column:id;type:varchar(36);primary_key:true;"`
	Name 	string	`json:"name,omitempty"gorm:"column:name;type:varchar(100);not null;unique_index;"`
}

func (Role) TableName() string {
	return "oauth_roles"
}

func (Role) BeforeCreate(scope *gorm.Scope) error {
	_ = scope.SetColumn("ID", uuid.New().String())
	return nil
}