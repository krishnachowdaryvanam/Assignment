package handlers

import (
	"fmt"
	"medical-vitals-management-system/app"
	"medical-vitals-management-system/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func CreateUserHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newUser models.User
		if err := c.ShouldBindJSON(&newUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": fmt.Sprintf("Invalid request: %v", err)})
			return
		}

		exists, err := app.UserExists(db, newUser.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": fmt.Sprintf("Failed to check user existence: %v", err)})
			return
		} else if exists {
			c.JSON(http.StatusConflict, gin.H{"status": "error", "message": "User with this username already exists"})
			return
		}

		if err := app.CreateUser(db, newUser); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": fmt.Sprintf("Failed to create user: %v", err)})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"status": "success", "message": fmt.Sprintf("User %s created.", newUser.Username)})
	}
}

func GetUserHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.Param("username")

		exists, err := app.UserExists(db, username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": fmt.Sprintf("Failed to check user existence: %v", err)})
			return
		}

		if !exists {
			c.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "User not found"})
			return
		}

		user, err := app.GetUser(db, username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": fmt.Sprintf("Failed to get user details: %v", err)})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "success", "data": user})
	}
}

func UpdateUserHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var updatedUser models.User
		if err := c.ShouldBindJSON(&updatedUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": fmt.Sprintf("Invalid request: %v", err)})
			return
		}

		exists, err := app.UserExists(db, updatedUser.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": fmt.Sprintf("Failed to check user existence: %v", err)})
			return
		} else if !exists {
			c.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "User not found"})
			return
		}

		if err := app.UpdateUser(db, updatedUser); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": fmt.Sprintf("Failed to update user: %v", err)})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "success", "message": fmt.Sprintf("User %s updated.", updatedUser.Username)})
	}
}

func DeleteUserHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.Param("username")

		exists, err := app.UserExists(db, username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": fmt.Sprintf("Failed to check user existence: %v", err)})
			return
		} else if !exists {
			c.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "User not found"})
			return
		}

		if err := app.DeleteUser(db, username); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": fmt.Sprintf("Failed to delete user: %v", err)})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "success", "message": fmt.Sprintf("User %s deleted.", username)})
	}
}
