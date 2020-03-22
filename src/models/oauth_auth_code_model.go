package models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"time"
)

type AuthCode struct {
	ID      	string 		`gorm:"column:id;type:varchar(36);primary_key;"`
	Code		string		`gorm:"column:code;type:varchar(2048);not null;"`
	Scopes		string		`gorm:"column:scopes;type:varchar(1024);not null;"`
	Revoked		bool    	`gorm:"column:revoked;not null;type:boolean;default:false"`
	CreatedAt 	*time.Time	`gorm:"column:created_at;not null;"`
	ExpiresAt 	*time.Time	`gorm:"column:expires_at;not null;"`
	UserID 		string		`gorm:"column:user_id;not null"`
	User 		*User		`gorm:"foreignkey:user_id"`
	ClientID	string		`gorm:"column:client_id;not null"`
	Client 		*Client		`gorm:"foreignkey:client_id"`
}

func (AuthCode) TableName() string {
	return "oauth_auth_codes"
}

func (AuthCode) BeforeCreate(scope *gorm.Scope) error {
	_ = scope.SetColumn("ID", uuid.New().String())
	return nil
}