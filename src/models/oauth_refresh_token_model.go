package models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"time"
)

type RefreshToken struct {
	ID      		string 		`gorm:"column:id;type:varchar(36);primary_key;"`
	RefreshToken	string		`gorm:"column:refresh_token;type:varchar(1024);not null;"`
	Revoked			bool    	`gorm:"column:activated;not null;type:boolean;default:false"`
	CreatedAt 		*time.Time	`gorm:"column:created_at;not null;"`
	ExpiresAt   	*time.Time	`gorm:"column:expires_at;not null;"`
	AccessToken 	*User		`gorm:"foreignkey:access_token_id"`
}

func (RefreshToken) TableName() string {
	return "oauth_refresh_tokens"
}

func (RefreshToken) BeforeCreate(scope *gorm.Scope) error {
	_ = scope.SetColumn("ID", uuid.New().String())
	return nil
}