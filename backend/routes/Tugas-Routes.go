package routes

import (
	"server/handlers"

	"github.com/gin-gonic/gin"
)

func TugasRoutes(r *gin.Engine) {
	tugasGroup := r.Group("/tugas")
	{
		tugasGroup.GET("/", handlers.IndexTugas)
		tugasGroup.GET("/:id", handlers.ShowTugas)
		tugasGroup.POST("/create", handlers.CreateTugas)
		tugasGroup.PUT("/update/:id", handlers.UpdateTugas)
		tugasGroup.DELETE("/delete/:id", handlers.DeleteTugas)
	}
}
