package dto

import (
	"oauth2/src/common"
	"oauth2/src/models"
	"strings"
	"time"
)

type UserDto struct {
	Name     	string  		`json:"name"`
	LastName    string  		`json:"last_name"`
	Email       string  		`json:"email"`
	Username    string  		`json:"username"`
	Password    *string  		`json:"password,omitempty"`
	Birthday    *Birthday  		`json:"birthday"`
	Activated 	bool    		`json:"activated"`
	Roles    	[]*models.Role 	`json:"roles,omitempty"`
}

func (u *UserDto) ToUser() *models.User {
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
	}
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
