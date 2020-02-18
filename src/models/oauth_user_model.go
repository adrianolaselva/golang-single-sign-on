package models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	ID          string 		`gorm:"column:id;type:varchar(36);primary_key;"`
	Name     	string  	`gorm:"column:name;type:varchar(100);not null;"`
	LastName    string  	`gorm:"column:last_name;type:varchar(100);not null;"`
	Email       string  	`gorm:"column:email;type:varchar(100);not null;unique_index"`
	Username    string  	`gorm:"column:username;type:varchar(100);not null;unique_index"`
	Password     *string  	`gorm:"column:password;type:varchar(100);not null;"`
	Birthday     *time.Time `gorm:"column:birthday;null;"`
	Activated 	 bool    	`gorm:"column:activated;not null;type:boolean;default:true"`
	CreatedAt 	 *time.Time	`gorm:"column:created_at;not null;"`
	UpdatedAt    *time.Time	`gorm:"column:updated_at;not null;"`
	ExpiresAt    *time.Time	`gorm:"column:expires_at;null;"sql:"index"`
	DeletedAt    *time.Time	`gorm:"column:deleted_at;null;"sql:"index"`
	Roles    	 []Role 	`gorm:"many2many:oauth_user_roles;foreignkey:user_id;association_foreignkey:role_id"`
}

func (User) TableName() string {
	return "oauth_users"
}

func (User) BeforeCreate(scope *gorm.Scope) error {
	_ = scope.SetColumn("ID", uuid.New().String())
	return nil
}
