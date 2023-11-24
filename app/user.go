package app

import (
	"fmt"
	"medical-vitals-management-system/models"

	"github.com/jinzhu/gorm"
)

func CreateUser(db *gorm.DB, user models.User) error {
	err := db.Create(&user).Error
	if err != nil {
		return fmt.Errorf("failed to create user: %v", err)
	}
	return nil
}

func GetUser(db *gorm.DB, username string) (models.User, error) {
	var user models.User
	err := db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return models.User{}, fmt.Errorf("user not found: %v", err)
	}
	return user, nil
}

func UpdateUser(db *gorm.DB, user models.User) error {
	err := db.Save(&user).Error
	if err != nil {
		return fmt.Errorf("failed to update user: %v", err)
	}
	return nil
}

// DeleteUser deletes a user record from the database
func DeleteUser(db *gorm.DB, username string) error {
	err := db.Where("username = ?", username).Delete(&models.User{}).Error
	if err != nil {
		return fmt.Errorf("failed to delete user: %v", err)
	}
	return nil
}

func UserExists(db *gorm.DB, username string) (bool, error) {
	var count int64
	if err := db.Model(&models.User{}).Where("username = ?", username).Count(&count).Error; err != nil {
		return false, fmt.Errorf("failed to check user existence: %v", err)
	}
	return count > 0, nil
}
