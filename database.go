package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	connectionString := "postgres://postgres.tsbxgkpukeamvfmhzkkr:ITB10215033@aws-0-ap-southeast-1.pooler.supabase.com:5432/postgres"
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	return db, nil
}

func CreatePost(post *Post, db *gorm.DB) error {
	result := db.Create(post)
	return result.Error
}

func GetPosts(db *gorm.DB) ([]Post, error) {
	var posts []Post
	result := db.Find(&posts)
	return posts, result.Error
}

func GetPostByID(id uint, db *gorm.DB) (*Post, error) {
	var post Post
	result := db.First(&post, id)
	if result.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &post, result.Error
}

func UpdatePost(updatedPost *Post, db *gorm.DB) error {
	result := db.Save(updatedPost)
	return result.Error
}

func DeletePostByID(id uint, db *gorm.DB) error {
	result := db.Delete(&Post{}, id)
	return result.Error
}
