package dto

import (
	"gopkg.in/go-playground/validator.v9"
	"oauth2/src/common"
	"oauth2/src/models"
	"strings"
	"time"
)

type UserDto struct {
	Name     	string  		`json:"name" validate:"required,min=2,max=100"`
	LastName    string  		`json:"last_name" validate:"required,min=2,max=100"`
	Email       string  		`json:"email" validate:"required,min=2,max=100"`
	Username    string  		`json:"username" validate:"required,min=2,max=20"`
	Password    *string  		`json:"password,omitempty" validate:"required,min=2,max=255"`
	Birthday    *Birthday  		`json:"birthday" validate:"required"`
	Activated 	bool    		`json:"activated" validate:"required"`
	Roles    	[]*models.Role 	`json:"roles,omitempty"`
}

func (u *UserDto) ToUser() (*models.User, error) {
	v := validator.New()
	err := v.Struct(u)
	if err != nil {
		return nil, err
	}

	var birthday time.Time
	if u.Birthday != nil {
		birthday = u.Birthday.Time
	}

	return &models.User{
		Name:      u.Name,
		LastName:  u.LastName,
		Email:     u.Email,
		Username:  u.Username,
		Password:  u.Password,
		Birthday:  &birthday,
		Activated: u.Activated,
		Roles: 	   u.Roles,
	}, nil
}

type Birthday struct {
	time.Time
}

func (b *Birthday) UnmarshalJSON(input []byte) error {
	strInput := string(input)
	strInput = strings.Trim(strInput, `"`)
	newTime, err := time.Parse(common.YYYY_MM_DD, strInput)
	if err != nil {
		return err
	}

	b.Time = newTime
	return nil
}
