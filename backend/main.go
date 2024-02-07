package main

import (
	"server/db"
	"server/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
  
  router := gin.Default()
  routes.AuthRoutes(router)
  
  router.Run(":8080")

}
