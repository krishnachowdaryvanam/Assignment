package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"column:username;not null"`
	Age      int    `gorm:"column:age"`
	Gender   string `gorm:"column:gender"`
	Vitals   []Vital
}

type Vital struct {
	gorm.Model
	UserID    uint      `gorm:"column:user_id;constraint:OnDelete:CASCADE"` // Foreign key to User
	Username  string    `gorm:"column:username;not null"`
	VitalID   string    `gorm:"column:vital_id;not null"`
	Value     float64   `gorm:"column:value"`
	Timestamp time.Time `gorm:"column:timestamp"`
}

// AggregateRequest represents the request for aggregate operations
type AggregateRequest struct {
	Username       string    `json:"username"`
	VitalIDs       []string  `json:"vital_ids"`
	StartTimestamp time.Time `json:"start_timestamp"`
	EndTimestamp   time.Time `json:"end_timestamp"`
}

// AggregateResponse represents the response format for aggregate operations
type AggregateResponse struct {
	Status         string        `json:"status"`
	Message        string        `json:"message"`
	Data           AggregateData `json:"data"`
	StartTimestamp time.Time     `json:"start_timestamp"`
	EndTimestamp   time.Time     `json:"end_timestamp"`
}

// AggregateData represents the data section of the AggregateResponse
type AggregateData struct {
	Username   string             `json:"username"`
	Aggregates map[string]float64 `json:"aggregates"`
}

type PopulationInsightRequest struct {
	Username       string    `json:"username"`
	VitalID        string    `json:"vital_id"`
	StartTimestamp time.Time `json:"start_timestamp"`
	EndTimestamp   time.Time `json:"end_timestamp"`
}

// PopulationInsightResponse represents the response structure for population insight.
type PopulationInsightResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Username       string    `json:"username"`
		VitalID        string    `json:"vital_id"`
		StartTimestamp time.Time `json:"start_timestamp"`
		EndTimestamp   time.Time `json:"end_timestamp"`
		Insight        string    `json:"insight"`
	} `json:"data"`
}

// PopulationInsightData represents the data structure for the response.
type PopulationInsightData struct {
	Username       string    `json:"username"`
	VitalID        string    `json:"vital_id"`
	StartTimestamp time.Time `json:"start_timestamp"`
	EndTimestamp   time.Time `json:"end_timestamp"`
	Insight        string    `json:"insight"`
}
