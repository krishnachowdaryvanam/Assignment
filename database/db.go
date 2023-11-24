package database

import (
	"fmt"
	"medical-vitals-management-system/models"
	"os"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"

	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

func InitDB() (*gorm.DB, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// Create database connection string
	conn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName,
	)

	db, err := gorm.Open("postgres", conn)
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.User{}, &models.Vital{})
	logrus.Info("Successfully connected to the database")
	return db, nil
}
