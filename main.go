package main

import (
	"fmt"
	"medical-vitals-management-system/database"
	"medical-vitals-management-system/routes"
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}

	db, err := database.InitDB()
	if err != nil {
		logrus.Fatalf("Failed to connect to the database: %v", err)
	}

	router := routes.SetUpRouter(db)
	router.Run(":8080")
}
