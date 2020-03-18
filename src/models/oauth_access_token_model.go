package models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"time"
)

type AccessToken struct {
	ID      		string 		`gorm:"column:id;type:varchar(255);primary_key;"`
	AccessToken		string		`gorm:"column:access_token;type:varchar(5000);not null;"`
	Scopes			string		`gorm:"column:scopes;type:varchar(1024);not null;"`
	Revoked			bool    	`gorm:"column:revoked;not null;type:boolean;default:false"`
	ExpiresAt 		*time.Time	`gorm:"column:expires_at;not null;"`
	UserID 			string		`gorm:"column:user_id;not null;"`
	User 			*User
	ClientID 		string		`gorm:"column:client_id;not null;"`
	Client 			*Client
}

func (AccessToken) TableName() string {
	return "oauth_access_tokens"
}

func (AccessToken) BeforeCreate(scope *gorm.Scope) error {
	_ = scope.SetColumn("ID", uuid.New().String())
	return nil
}