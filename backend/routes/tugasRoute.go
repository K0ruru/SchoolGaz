package routes

import (
	"server/handlers/tugasHandlers"

	"github.com/gin-gonic/gin"
)

func TugasRoutes(r *gin.Engine) {
	tugasGroup := r.Group("/tugas")
	{
		tugasGroup.GET("/", tugasHandlers.IndexTugas)
		tugasGroup.GET("/:id", tugasHandlers.ShowTugas)
		tugasGroup.POST("/create", tugasHandlers.CreateTugas)
		tugasGroup.PUT("/update/:id", tugasHandlers.UpdateTugas)
		tugasGroup.DELETE("/delete/:id", tugasHandlers.DeleteTugas)
	}
}
