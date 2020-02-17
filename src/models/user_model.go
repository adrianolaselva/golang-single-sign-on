package models

import (
	"github.com/jinzhu/gorm"
	"github.com/google/uuid"
	"time"
)

type User struct {
	Uuid         string 	`gorm:"column:uuid;type:varchar(36);primary_key;"`
	Username     string  	`gorm:"column:username;type:varchar(100);not null;"`
	Email        string  	`gorm:"column:email;type:varchar(100);not null;unique_index"`
	Password     *string  	`gorm:"column:password;type:varchar(100);not null;"`
	Birthday     *time.Time  `gorm:"column:birthday;null;"`
	Activated 	 bool    	`gorm:"column:activated;not null;type:boolean;default:true"`
	Roles    	 []Role 	`gorm:"many2many:user_roles;foreignkey:user_uuid;association_foreignkey:name"`
	CreatedAt 	 *time.Time	`gorm:"column:created_at;not null;"`
	UpdatedAt    *time.Time	`gorm:"column:updated_at;not null;"`
	DeletedAt    *time.Time	`gorm:"column:deleted_at;null;"sql:"index"`
}

func (User) TableName() string {
	return "users"
}

func (d *User) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.New().String()
	scope.SetColumn("Uuid", uuid)
	return nil
}
