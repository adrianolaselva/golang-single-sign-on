package models

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"oauth2/src/common"
	"time"
)

type User struct {
	ID          string 		`json:"id"gorm:"column:id;primary_key;type:varchar(36);not null;"`
	Name     	string  	`json:"name"gorm:"column:name;type:varchar(120);not null;"`
	LastName    string  	`json:"last_name"gorm:"column:last_name;type:varchar(120);not null;"`
	Email       string  	`json:"email"gorm:"column:email;type:varchar(120);not null;unique_index"`
	Username    string  	`json:"username"gorm:"column:username;type:varchar(120);not null;unique_index"`
	Password    *string  	`json:"password,omitempty"gorm:"column:password;type:varchar(120);not null;"`
	Birthday    *time.Time  `json:"birthday"gorm:"column:birthday;type:date;null;"`
	Activated 	bool    	`json:"activated"gorm:"column:activated;not null;type:boolean;default:true"`
	CreatedAt 	*time.Time	`json:"created_at"gorm:"column:created_at;type:datetime;not null;"`
	UpdatedAt   *time.Time	`json:"updated_at"gorm:"column:updated_at;type:datetime;not null;"`
	ExpiresAt   *time.Time	`json:"expires_at"gorm:"column:expires_at;type:datetime;null;"`
	DeletedAt   *time.Time	`json:"deleted_at,omit"gorm:"column:deleted_at;type:datetime;null;"sql:"index"`
	Roles    	[]*Role 	`json:"roles,omitempty"gorm:"many2many:oauth_user_roles"`
}

func (User) TableName() string {
	return "oauth_users"
}

func (User) BeforeCreate(scope *gorm.Scope) error {
	_ = scope.SetColumn("ID", uuid.New().String())
	return nil
}

func (u *User) MarshalJSON() ([]byte, error) {
	type Alias User
	u.Password = nil

	var birthday, createdAt, updatedAt, expiresAt, deletedAt string
	if u.Birthday != nil {
		birthday = u.Birthday.Format(common.YYYY_MM_DD)
	}
	if u.CreatedAt != nil {
		createdAt = u.CreatedAt.Format(common.YYYY_MM_DD_HH_MM_SS)
	}
	if u.UpdatedAt != nil {
		updatedAt = u.UpdatedAt.Format(common.YYYY_MM_DD_HH_MM_SS)
	}
	if u.ExpiresAt != nil {
		expiresAt = u.ExpiresAt.Format(common.YYYY_MM_DD_HH_MM_SS)
	}
	if u.DeletedAt != nil {
		deletedAt = u.DeletedAt.Format(common.YYYY_MM_DD_HH_MM_SS)
	}
	return json.Marshal(&struct {
		*Alias
		Birthday string `json:"birthday"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
		ExpiresAt string `json:"expires_at"`
		DeletedAt string `json:"deleted_at"`
	}{
		Alias: (*Alias)(u),
		Birthday: birthday,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		ExpiresAt: expiresAt,
		DeletedAt: deletedAt,
	})
}
