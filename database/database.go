package database

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"gin-blog-app/models"
)

var db *gorm.DB

func Init() error {
	var err error
	connectionString := os.Getenv("POSTGRES_CS")
	db, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		return err
	}

	db.AutoMigrate(&models.User{}, &models.Post{})

	return nil
}
