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
	UserID    uint      `gorm:"column:user_id;constraint:OnDelete:CASCADE"`
	Username  string    `gorm:"column:username;not null"`
	VitalID   string    `gorm:"column:vital_id;not null"`
	Value     float64   `gorm:"column:value"`
	Timestamp time.Time `gorm:"column:timestamp"`
}

type AggregateRequest struct {
	Username       string    `json:"username"`
	VitalIDs       []string  `json:"vital_ids"`
	StartTimestamp time.Time `json:"start_timestamp"`
	EndTimestamp   time.Time `json:"end_timestamp"`
}

type AggregateResponse struct {
	Status         string        `json:"status"`
	Message        string        `json:"message"`
	Data           AggregateData `json:"data"`
	StartTimestamp time.Time     `json:"start_timestamp"`
	EndTimestamp   time.Time     `json:"end_timestamp"`
}

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

type PopulationInsightData struct {
	Username       string    `json:"username"`
	VitalID        string    `json:"vital_id"`
	StartTimestamp time.Time `json:"start_timestamp"`
	EndTimestamp   time.Time `json:"end_timestamp"`
	Insight        string    `json:"insight"`
}
