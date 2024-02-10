package main

import (
	"net/http"
	"server/db"
	"server/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:7070"}
	router.Use(cors.New(config))

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})

	routes.AuthRoutes(router)
	routes.KelasRoutes(router)
	routes.MapelRoutes(router)

	router.Run(":8080")
}
