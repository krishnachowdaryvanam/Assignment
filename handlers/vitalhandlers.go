package handlers

import (
	"fmt"
	"medical-vitals-management-system/app"
	"medical-vitals-management-system/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func CreateVitalHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request struct {
			Username  string  `json:"username"`
			VitalID   string  `json:"vital_id"`
			Value     float64 `json:"value"`
			Timestamp string  `json:"timestamp"`
		}

		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request format", "status": "error"})
			return
		}

		exists, err := app.UserExists(db, request.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": fmt.Sprintf("Failed to check user existence: %v", err)})
			return
		} else if !exists {
			c.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "User not found"})
			return
		}

		timestamp, err := time.Parse(time.RFC3339, request.Timestamp)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid timestamp format"})
			return
		}

		if err := app.CreateVital(db, models.Vital{
			Username:  request.Username,
			VitalID:   request.VitalID,
			Value:     request.Value,
			Timestamp: timestamp,
		}); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": fmt.Sprintf("Failed to insert vital: %v", err)})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"status": "success", "message": fmt.Sprintf("Vital inserted for %s.", request.Username)})
	}
}

func GetVitalsHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var getData struct {
			Username string   `json:"username"`
			Period   []string `json:"period"`
		}
		if err := c.ShouldBindJSON(&getData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": fmt.Sprintf("Invalid request: %v", err)})
			return
		}

		username := getData.Username
		period := getData.Period

		if len(period) != 2 {
			c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid period"})
			return
		}

		exists, err := app.UserExists(db, username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": fmt.Sprintf("Failed to check user existence: %v", err)})
			return
		} else if !exists {
			c.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "User not found"})
			return
		}

		vitals, err := app.GetVitals(db, username, period)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": fmt.Sprintf("Failed to get vitals: %v", err)})
			return
		}

		var transformedVitals []map[string]interface{}
		for _, vital := range vitals {
			transformedVitals = append(transformedVitals, map[string]interface{}{
				"vitalID":   vital.VitalID,
				"value":     vital.Value,
				"timestamp": vital.Timestamp,
			})
		}

		c.JSON(http.StatusOK, gin.H{"status": "success", "data": transformedVitals})
	}
}

func UpdateVitalHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var updateData struct {
			Username  string  `json:"username"`
			VitalID   string  `json:"vitalID"`
			Timestamp string  `json:"timestamp"`
			NewValue  float64 `json:"newValue"`
		}

		if err := c.ShouldBindJSON(&updateData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": fmt.Sprintf("Invalid request: %v", err)})
			return
		}

		err := app.UpdateVital(db, updateData.Username, updateData.VitalID, updateData.Timestamp, updateData.NewValue)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": fmt.Sprintf("Failed to edit vital: %v", err)})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "success", "message": fmt.Sprintf("Vital value updated for %s.", updateData.Username)})
	}
}

func DeleteVitalHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var deleteRequest models.DeleteVitalRequest
		if err := c.ShouldBindJSON(&deleteRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": fmt.Sprintf("Invalid request: %v", err)})
			return
		}

		if err := app.DeleteVital(db, deleteRequest); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": fmt.Sprintf("Failed to delete vital: %v", err)})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "success", "message": fmt.Sprintf("Vital deleted for %s.", deleteRequest.Username)})
	}
}
