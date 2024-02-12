// main.go
package main

import (
	"server/database"
	"server/models"
	"server/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	database.ConnectDatabase()

	models.AutoMigrateTugas()
	models.AutoMigrateJawaban()
	models.AutoMigrateUploadTugas()

	routes.TugasRoutes(r)

	r.Run()
}
