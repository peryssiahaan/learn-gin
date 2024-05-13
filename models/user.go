package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"-"`
	Token    string `json:"-"`
	Posts    []Post `grom:"foreignKey:UserID"`
}
