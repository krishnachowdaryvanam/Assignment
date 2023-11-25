package app

import (
	"medical-vitals-management-system/models"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type MockVitalDB struct {
	CreateVitalFunc func(vital models.Vital) error
	GetVitalsFunc   func(username string, period []string) ([]models.Vital, error)
	UpdateVitalFunc func(username, vitalID, timestamp string, newValue float64) error
	DeleteVitalFunc func(username, vitalID, timestamp string) error
}

func (m *MockVitalDB) CreateVital(vital models.Vital) error {
	if m.CreateVitalFunc != nil {
		return m.CreateVitalFunc(vital)
	}
	return nil
}

func (m *MockVitalDB) GetVitals(username string, period []string) ([]models.Vital, error) {
	if m.GetVitalsFunc != nil {
		return m.GetVitalsFunc(username, period)
	}
	return nil, nil
}

func (m *MockVitalDB) UpdateVital(username, vitalID, timestamp string, newValue float64) error {
	if m.UpdateVitalFunc != nil {
		return m.UpdateVitalFunc(username, vitalID, timestamp, newValue)
	}
	return nil
}

func (m *MockVitalDB) DeleteVital(username, vitalID, timestamp string) error {
	if m.DeleteVitalFunc != nil {
		return m.DeleteVitalFunc(username, vitalID, timestamp)
	}
	return nil
}

func TestCreateVitalWithMock(t *testing.T) {
	mockDB := &MockVitalDB{}

	vital := models.Vital{
		Username:  "JohnDoe",
		VitalID:   "HeartRate",
		Value:     75.0,
		Timestamp: time.Now(),
	}

	mockDB.CreateVitalFunc = func(v models.Vital) error {
		return nil
	}

	err := mockDB.CreateVital(vital)
	assert.NoError(t, err, "Error creating vital record")

}

func TestGetVitalsWithMock(t *testing.T) {
	mockDB := &MockVitalDB{}

	username := "JohnDoe"
	period := []string{"2023-01-01T00:00:00Z", "2023-01-02T00:00:00Z"}
	expectedVitals := []models.Vital{
		{Username: username, VitalID: "HeartRate", Value: 75.0, Timestamp: time.Now()},
		{Username: username, VitalID: "Temperature", Value: 98.6, Timestamp: time.Now()},
	}

	mockDB.GetVitalsFunc = func(username string, period []string) ([]models.Vital, error) {
		return expectedVitals, nil
	}

	vitals, err := mockDB.GetVitals(username, period)
	assert.NoError(t, err, "Error getting vitals")
	assert.Equal(t, expectedVitals, vitals, "Vital data mismatch")
}

func TestUpdateVitalWithMock(t *testing.T) {
	mockDB := &MockVitalDB{}

	username := "JohnDoe"
	vitalID := "HeartRate"
	timestamp := time.Now().Format(time.RFC3339)
	newValue := 80.0

	mockDB.UpdateVitalFunc = func(username, vitalID, timestamp string, newValue float64) error {
		return nil
	}

	err := mockDB.UpdateVital(username, vitalID, timestamp, newValue)
	assert.NoError(t, err, "Error updating vital record")
}

func TestDeleteVitalWithMock(t *testing.T) {
	mockDB := &MockVitalDB{}

	username := "JohnDoe"
	vitalID := "HeartRate"
	timestamp := time.Now().Format(time.RFC3339)

	mockDB.DeleteVitalFunc = func(username, vitalID, timestamp string) error {
		return nil
	}

	err := mockDB.DeleteVital(username, vitalID, timestamp)
	assert.NoError(t, err, "Error deleting vital record")

}
