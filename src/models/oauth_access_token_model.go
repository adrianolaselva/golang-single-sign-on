package models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"time"
)

type AccessToken struct {
	ID      		string 		`gorm:"column:id;type:varchar(36);primary_key;"`
	AccessToken		string		`gorm:"column:access_token;type:varchar(1024);not null;"`
	Scopes			string		`gorm:"column:scopes;type:varchar(1024);not null;"`
	Revoked			bool    	`gorm:"column:revoked;not null;type:boolean;default:false"`
	ExpiresAt 		*time.Time	`gorm:"column:expires_at;not null;"`
	User 			*User		`gorm:"foreignkey:user_id"`
	Client 			*Client		`gorm:"foreignkey:client_id"`
}

func (AccessToken) TableName() string {
	return "oauth_access_tokens"
}

func (AccessToken) BeforeCreate(scope *gorm.Scope) error {
	_ = scope.SetColumn("ID", uuid.New().String())
	return nil
}