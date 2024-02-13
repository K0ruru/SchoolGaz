package routes

import (
	"server/handlers"

	"github.com/gin-gonic/gin"
)

func TugasRoutes(r *gin.Engine) {
	tugasGroup := r.Group("/tugas")
	{
		tugasGroup.GET("/", handlers.IndexTugas)
		tugasGroup.GET("/:id", handlers.GetTugas)
		tugasGroup.POST("/add", handlers.CreateTugas)
		tugasGroup.PUT("/update/:id", handlers.UpdateTugas)
		tugasGroup.DELETE("/delete/:id", handlers.DeleteTugas)

		// Route Upload Tugas
		tugasGroup.POST("/file/upload", handlers.UploadFile)
		tugasGroup.GET("/file", handlers.IndexFile)
		tugasGroup.GET("/file/:id", handlers.GetFile)
	}
}
