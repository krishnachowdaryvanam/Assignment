package handlers

import (
	"fmt"
	"medical-vitals-management-system/app"
	"medical-vitals-management-system/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func AggregateVitalsHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var aggregateRequest models.AggregateRequest
		if err := c.ShouldBindJSON(&aggregateRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": fmt.Sprintf("Invalid request: %v", err)})
			return
		}

		userExists, err := app.UserExists(db, aggregateRequest.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": fmt.Sprintf("Failed to check user existence: %v", err)})
			return
		}

		if !userExists {
			c.JSON(http.StatusNotFound, gin.H{"status": "error", "message": fmt.Sprintf("User not found: %s", aggregateRequest.Username)})
			return
		}

		aggregatedValues, err := app.AggregateVitals(db, aggregateRequest)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": fmt.Sprintf("Failed to calculate aggregate values: %v", err)})
			return
		}

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

		c.JSON(http.StatusOK, response)
	}
}

func PopulationInsightHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var insightRequest models.PopulationInsightRequest
		if err := c.ShouldBindJSON(&insightRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": fmt.Sprintf("Invalid request: %v", err)})
			return
		}

		userExists, err := app.UserExists(db, insightRequest.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": fmt.Sprintf("Failed to check user existence: %v", err)})
			return
		}

		if !userExists {
			c.JSON(http.StatusNotFound, gin.H{"status": "error", "message": fmt.Sprintf("User not found: %s", insightRequest.Username)})
			return
		}

		insight, err := app.GetPopulationInsight(db, insightRequest.Username, insightRequest.VitalID, insightRequest.StartTimestamp, insightRequest.EndTimestamp)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": fmt.Sprintf("Failed to calculate population insight: %v", err)})
			return
		}

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

		c.JSON(http.StatusOK, response)
	}
}
