package app

import (
	"fmt"
	"medical-vitals-management-system/models"

	"github.com/jinzhu/gorm"
)

func AggregateVitals(db *gorm.DB, request models.AggregateRequest) (map[string]float64, error) {
	aggregatedValues := make(map[string]float64)

	// Fetch vitals from the database based on the provided criteria
	var vitals []models.Vital
	err := db.Where("username = ? AND timestamp BETWEEN ? AND ?", request.Username, request.StartTimestamp, request.EndTimestamp).
		Find(&vitals).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get vitals: %v", err)
	}

	// Check if vitals are found for the specified user and time range
	if len(vitals) == 0 {
		return nil, fmt.Errorf("no vitals found for the specified user and time range")
	}

	// Calculate the mean value for each requested vital
	for _, vitalID := range request.VitalIDs {
		var sum float64
		count := 0

		for _, vital := range vitals {
			if vital.VitalID == vitalID {
				sum += vital.Value
				count++
			}
		}

		if count > 0 {
			meanValue := sum / float64(count)
			aggregatedValues[vitalID] = meanValue
		}
	}

	return aggregatedValues, nil
}
