package database

import (
	"gin-blog-app/models"

	"gorm.io/gorm"
)

func CreatePost(post *models.Post) error {
	result := db.Create(post)
	return result.Error
}

func GetPosts() ([]models.Post, error) {
	var posts []models.Post
	result := db.Find(&posts)
	return posts, result.Error
}

func GetPostByID(id uint) (*models.Post, error) {
	var post models.Post
	result := db.First(&post, id)
	if result.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &post, result.Error
}

func UpdatePost(updatedPost *models.Post) error {
	result := db.Save(updatedPost)
	return result.Error
}

func DeletePostByID(id uint) error {
	result := db.Delete(&models.Post{}, id)
	return result.Error
}
