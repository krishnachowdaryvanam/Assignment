package main

import (
	"medical-vitals-management-system/database"
	"medical-vitals-management-system/routes"

	"github.com/sirupsen/logrus"
)

func main() {
	// Connect to the database
	db, err := database.InitDB()
	if err != nil {
		logrus.Fatalf("Failed to connect to the database: %v", err)
	}

	// Create a new Gin router using SetupRouter
	router := routes.SetUpRouter(db)
	router.Run(":8080")
}
