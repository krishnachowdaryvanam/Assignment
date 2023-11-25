package routes

import (
	"medical-vitals-management-system/handlers"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetUpRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(200, "Welcome To Medical Vital Management System")
	})

	// User routes
	userGroup := router.Group("/users")
	{
		userGroup.GET("/get_user", handlers.GetUserHandler(db))
		userGroup.POST("/create_user", handlers.CreateUserHandler(db))
		userGroup.PUT("/update_user", handlers.UpdateUserHandler(db))
		userGroup.DELETE("/delete_user", handlers.DeleteUserHandler(db))
	}

	// Vital routes
	vitalGroup := router.Group("/vitals")
	{
		vitalGroup.POST("/insert_vital", handlers.CreateVitalHandler(db))
		vitalGroup.GET("/get_vitals", handlers.GetVitalsHandler(db))
		vitalGroup.PUT("/edit_vital", handlers.UpdateVitalHandler(db))
		vitalGroup.DELETE("/delete_vital", handlers.DeleteVitalHandler(db))

	}

	//Insight routes
	insightGroup := router.Group("/insights")
	{
		insightGroup.POST("/aggregate", handlers.AggregateVitalsHandler(db))
		insightGroup.POST("/population_insight", handlers.PopulationInsightHandler(db))
	}

	return router
}
