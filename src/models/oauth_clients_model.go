package models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"time"
)

type Client struct {
	ID      	string 		`gorm:"column:id;type:varchar(36);primary_key;"`
	Name		string		`gorm:"column:name;type:varchar(120);not null;"`
	Scopes		string		`gorm:"column:scopes;type:varchar(1024);null;"`
	Redirect	string		`gorm:"column:redirect;type:varchar(255);null;"`
	Revoked		bool    	`gorm:"column:activated;not null;type:boolean;default:false"`
	CreatedAt 	*time.Time	`gorm:"column:created_at;not null;"`
	UpdatedAt   *time.Time	`gorm:"column:updated_at;not null;"`
	User 		*User		`gorm:"foreignkey:user_id"`
}

func (Client) TableName() string {
	return "oauth_clients"
}

func (Client) BeforeCreate(scope *gorm.Scope) error {
	_ = scope.SetColumn("ID", uuid.New().String())
	return nil
}