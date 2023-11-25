package app

import (
	"fmt"
	"medical-vitals-management-system/models"
	"sort"
	"time"

	"github.com/jinzhu/gorm"
)

// AggregateVitals fetches vitals from the database and calculates the mean value for each requested vital
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

// CalculatePopulationInsights calculates population insights for each requested vital
func CalculatePopulationInsights(db *gorm.DB, request models.AggregateRequest) (map[string]string, error) {
	populationInsights := make(map[string]string)

	// Calculate population insight for each requested vital
	for _, vitalID := range request.VitalIDs {
		// Calculate population insight using aggregated values
		populationInsight, err := GetPopulationInsight(db, request.Username, vitalID, request.StartTimestamp, request.EndTimestamp)
		if err != nil {
			return nil, fmt.Errorf("failed to calculate population insight: %v", err)
		}
		populationInsights[vitalID] = populationInsight
	}

	return populationInsights, nil
}

// GetPopulationInsight compares a user's vitals against the population and provides percentile standings
func GetPopulationInsight(db *gorm.DB, username, vitalID string, startTimestamp, endTimestamp time.Time) (string, error) {
	// Calculate the aggregated value of the user's vital over the specified period
	var userAggregatedValue float64
	err := db.Table("vitals").
		Select("AVG(value)").
		Where("username = ? AND vital_id = ? AND timestamp BETWEEN ? AND ?", username, vitalID, startTimestamp, endTimestamp).
		Row().Scan(&userAggregatedValue)
	if err != nil {
		return "", fmt.Errorf("failed to get user's aggregated vital value: %v", err)
	}

	// Fetch aggregated values of all users for the specified vital
	var populationValues []float64
	err = db.Table("vitals").
		Select("AVG(value)").
		Where("vital_id = ? AND timestamp BETWEEN ? AND ?", vitalID, startTimestamp, endTimestamp).
		Group("username").
		Order("AVG(value)").
		Pluck("AVG(value)", &populationValues).
		Error
	if err != nil {
		return "", fmt.Errorf("failed to get population values: %v", err)
	}

	// Sort the population values in ascending order
	sort.Float64s(populationValues)

	// Find the user's aggregated value position in the sorted population dataset
	var position int
	for i, value := range populationValues {
		if userAggregatedValue == value {
			position = i + 1 // 1-based index
			break
		}
	}

	// Calculate the total count of aggregated values in the population dataset
	totalCount := len(populationValues)

	// Calculate the percentile rank
	percentileRank := (float64(position) / float64(totalCount)) * 100

	// Generate insight message
	insight := fmt.Sprintf("Your %s is in the %.2fth percentile.", vitalID, percentileRank)

	return insight, nil
}
