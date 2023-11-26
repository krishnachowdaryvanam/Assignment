package app

import (
	"fmt"
	"medical-vitals-management-system/models"

	"github.com/jinzhu/gorm"
)

// CreateVital inserts a new vital record into the database
func CreateVital(db *gorm.DB, vital models.Vital) error {
	validVitalIDs := map[string]bool{
		"HeartRate":   true,
		"Temperature": true,
	}

	if !validVitalIDs[vital.VitalID] {

		return fmt.Errorf("invalid vitalID: %s", vital.VitalID)

	}

	err := db.Create(&vital).Error
	if err != nil {
		return fmt.Errorf("failed to insert vital: %v", err)
	}
	return nil
}

// GetVitals retrieves vital records for a user over a specified period
func GetVitals(db *gorm.DB, username string, period []string) ([]models.Vital, error) {
	var vitals []models.Vital
	if len(period) != 2 {
		return nil, fmt.Errorf("period should have exactly 2 elements")
	}

	err := db.Where("username = ? AND timestamp BETWEEN ? AND ?", username, period[0], period[1]).Find(&vitals).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get vitals: %v", err)
	}

	return vitals, nil
}

// UpdateVital updates an existing vital record in the database
func UpdateVital(db *gorm.DB, username, vitalID, timestamp string, newValue float64) error {
	err := db.Model(&models.Vital{}).
		Where("username = ? AND vital_id = ? AND timestamp = ?", username, vitalID, timestamp).
		Update("value", newValue).Error
	if err != nil {
		return fmt.Errorf("failed to edit vital: %v", err)
	}
	return nil
}

// DeleteVital deletes a vital record for a user using the vital ID and timestamp
func DeleteVital(db *gorm.DB, request models.DeleteVitalRequest) error {
	fmt.Printf("Deleting vital: %+v\n", request)

	var vital models.Vital
	err := db.Where("username = ? AND vital_id = ? AND timestamp = ?", request.Username, request.VitalID, request.Timestamp).
		First(&vital).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil
		}
		return fmt.Errorf("failed to check vital existence: %v", err)
	}

	err = db.Delete(&vital).Error
	if err != nil {
		return fmt.Errorf("failed to delete vital: %v", err)
	}

	return nil
}
