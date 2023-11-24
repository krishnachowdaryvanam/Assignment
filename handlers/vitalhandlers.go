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

// CreateVitalHandler handles the insertion of a new vital record into the database
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

		// Check if the user exists before inserting the vital
		exists, err := app.UserExists(db, request.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": fmt.Sprintf("Failed to check user existence: %v", err)})
			return
		} else if !exists {
			c.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "User not found"})
			return
		}

		// Parse the timestamp from the request
		timestamp, err := time.Parse(time.RFC3339, request.Timestamp)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid timestamp format"})
			return
		}

		// Insert the new vital
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

// GetVitalsHandler retrieves vital records for a user over a specified period
func GetVitalsHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Bind the JSON request body to the GetData struct
		var getData struct {
			Username string   `json:"username"`
			Period   []string `json:"period"`
		}
		if err := c.ShouldBindJSON(&getData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": fmt.Sprintf("Invalid request: %v", err)})
			return
		}

		// Extract values from the GetData struct
		username := getData.Username
		period := getData.Period

		if len(period) != 2 {
			c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid period"})
			return
		}

		// Check if the user exists before getting vitals
		exists, err := app.UserExists(db, username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": fmt.Sprintf("Failed to check user existence: %v", err)})
			return
		} else if !exists {
			c.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "User not found"})
			return
		}

		// Get vitals for the user over the specified period
		vitals, err := app.GetVitals(db, username, period)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": fmt.Sprintf("Failed to get vitals: %v", err)})
			return
		}

		// Transform the result to the desired format
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

// UpdateVitalHandler updates an existing vital record in the database
func UpdateVitalHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.Param("username")
		var updateData struct {
			VitalID   string  `json:"vitalID"`
			Timestamp string  `json:"timestamp"`
			NewValue  float64 `json:"newValue"`
		}
		if err := c.ShouldBindJSON(&updateData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": fmt.Sprintf("Invalid request: %v", err)})
			return
		}

		// Check if the user exists before updating the vital
		exists, err := app.UserExists(db, username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": fmt.Sprintf("Failed to check user existence: %v", err)})
			return
		} else if !exists {
			c.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "User not found"})
			return
		}

		// Update the vital
		if err := app.UpdateVital(db, username, updateData.VitalID, updateData.Timestamp, updateData.NewValue); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": fmt.Sprintf("Failed to edit vital: %v", err)})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "success", "message": fmt.Sprintf("Vital updated for %s.", username)})
	}
}

// DeleteVitalHandler deletes a vital record for a user using the vital ID and timestamp
func DeleteVitalHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.Param("username")
		var deleteData struct {
			VitalID   string `json:"vitalID"`
			Timestamp string `json:"timestamp"`
		}
		if err := c.ShouldBindJSON(&deleteData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": fmt.Sprintf("Invalid request: %v", err)})
			return
		}

		// Check if the user exists before deleting the vital
		exists, err := app.UserExists(db, username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": fmt.Sprintf("Failed to check user existence: %v", err)})
			return
		} else if !exists {
			c.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "User not found"})
			return
		}

		// Delete the vital
		if err := app.DeleteVital(db, username, deleteData.VitalID, deleteData.Timestamp); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": fmt.Sprintf("Failed to delete vital: %v", err)})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "success", "message": fmt.Sprintf("Vital deleted for %s.", username)})
	}
}
