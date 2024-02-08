package main

import (
	"server/db"
	"server/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
  
  router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:7070"} // Add the origins that are allowed to make requests
	router.Use(cors.New(config))

  routes.AuthRoutes(router)
  routes.GuruRoutes(router)
  router.Run(":8080")

}
