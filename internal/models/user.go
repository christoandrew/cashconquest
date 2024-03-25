package models

import (
	utils "github.com/christo-andrew/maybe-go/pkg/utils"
)
type User struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password" gorm:"type:varchar(256)"`
}

func (user *User) GenerateToken() string {
	return utils.GenerateMD5Hash(user.Email) + utils.GenerateMD5Hash(user.Password)
}

func (user *User) GetFullName() string {
	return user.FirstName + " " + user.LastName
}
