package handlers

import (
	"fmt"
	"medical-vitals-management-system/app"
	"medical-vitals-management-system/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// AggregateVitalsHandler retrieves average values of specific vitals for a user over a specified period
func AggregateVitalsHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var aggregateRequest models.AggregateRequest
		if err := c.ShouldBindJSON(&aggregateRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": fmt.Sprintf("Invalid request: %v", err)})
			return
		}

		// Check if the user exists
		userExists, err := app.UserExists(db, aggregateRequest.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": fmt.Sprintf("Failed to check user existence: %v", err)})
			return
		}

		if !userExists {
			c.JSON(http.StatusNotFound, gin.H{"status": "error", "message": fmt.Sprintf("User not found: %s", aggregateRequest.Username)})
			return
		}

		// Perform calculations in the code instead of fetching from the database
		aggregatedValues, err := app.AggregateVitals(db, aggregateRequest)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": fmt.Sprintf("Failed to calculate aggregate values: %v", err)})
			return
		}

		// Prepare the response
		response := models.AggregateResponse{
			Status:  "success",
			Message: "Aggregate fetched successfully",
			Data: models.AggregateData{
				Username:   aggregateRequest.Username,
				Aggregates: aggregatedValues,
			},
			StartTimestamp: aggregateRequest.StartTimestamp,
			EndTimestamp:   aggregateRequest.EndTimestamp,
		}

		// Return the response
		c.JSON(http.StatusOK, response)
	}
}

// PopulationInsightHandler compares a user's vitals against the population and provides percentile standings
func PopulationInsightHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var insightRequest models.PopulationInsightRequest
		if err := c.ShouldBindJSON(&insightRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": fmt.Sprintf("Invalid request: %v", err)})
			return
		}

		// Check if the user exists
		userExists, err := app.UserExists(db, insightRequest.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": fmt.Sprintf("Failed to check user existence: %v", err)})
			return
		}

		if !userExists {
			c.JSON(http.StatusNotFound, gin.H{"status": "error", "message": fmt.Sprintf("User not found: %s", insightRequest.Username)})
			return
		}

		// Calculate population insight
		insight, err := app.GetPopulationInsight(db, insightRequest.Username, insightRequest.VitalID, insightRequest.StartTimestamp, insightRequest.EndTimestamp)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": fmt.Sprintf("Failed to calculate population insight: %v", err)})
			return
		}

		// Prepare the response
		response := models.PopulationInsightResponse{
			Status:  "success",
			Message: "Population insight fetched successfully",
			Data: models.PopulationInsightData{
				Username:       insightRequest.Username,
				VitalID:        insightRequest.VitalID,
				StartTimestamp: insightRequest.StartTimestamp,
				EndTimestamp:   insightRequest.EndTimestamp,
				Insight:        insight,
			},
		}

		// Return the response
		c.JSON(http.StatusOK, response)
	}
}