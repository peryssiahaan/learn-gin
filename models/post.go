package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	UserID  uint   `json:"-"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
