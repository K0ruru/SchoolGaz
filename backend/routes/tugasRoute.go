package routes

import (
	"server/handlers/tugasHandlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	tugasGroup := r.Group("/tugas")
	{
		tugasGroup.GET("/", tugasHandlers.Index)
		tugasGroup.GET("/:id", tugasHandlers.Show)
		tugasGroup.POST("/create", tugasHandlers.Create)
		tugasGroup.PUT("/update/:id", tugasHandlers.Update)
		tugasGroup.DELETE("/delete", tugasHandlers.Delete)
	}
}
