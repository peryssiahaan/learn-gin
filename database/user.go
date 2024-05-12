package database

import (
	"gin-blog-app/models"

	"gorm.io/gorm"
)

func CreateUser(user *models.User) error {
	result := db.Create(user)
	return result.Error
}

func GetUsers() ([]models.User, error) {
	var users []models.User
	result := db.Find(&users)
	return users, result.Error
}

func GetUserByID(id uint) (*models.User, error) {
	var user models.User
	result := db.Preload("Posts").First(&user, id)
	if result.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &user, result.Error
}

func UpdateUser(updatedUser *models.User) error {
	result := db.Save(updatedUser)
	return result.Error
}

func DeleteUserByID(id uint) error {
	result := db.Delete(&models.User{}, id)
	return result.Error
}
