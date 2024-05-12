package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"gin-blog-app/models"
)

var db *gorm.DB

func Init() error {
	var err error
	connectionString := "postgres://postgres.tsbxgkpukeamvfmhzkkr:ITB10215033@aws-0-ap-southeast-1.pooler.supabase.com:5432/postgres"
	db, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		return err
	}

	db.AutoMigrate(&models.User{}, &models.Post{})

	return nil
}

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
