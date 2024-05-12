package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username"`
	Email    string `json:"email"`
	Posts    []Post `grom:"foreignKey:UserID"`
}
