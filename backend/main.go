package main

import (
	"fmt"
	"net/http"
	"server/db"
	"server/handlers"
	"server/model"
	"server/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the database connection
	dbInstance, err := db.InitDB()
	if err != nil {
		fmt.Println("Failed to initialize database:", err)
		return
	}
	defer dbInstance.Close()

	// Perform auto migration for all models
	if err := model.AutoMigrate(dbInstance); err != nil {
		fmt.Println("Error migrating models:", err)
		return
	}

	// Initialize the Gin router
	router := gin.Default()

	// Configure CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:7070"}
	router.Use(cors.New(config))

	router.Static("/profile_pictures", "./profile_pictures")

	// Define routes
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})
	routes.AuthRoutes(router)
	routes.KelasRoutes(router)
	routes.MapelRoutes(router)
	routes.GuruRoutes(router)
	routes.TugasRoutes(router)

	handlers.InitCloudinary()

	// Start the HTTP server
	router.Run(":8080")
}
