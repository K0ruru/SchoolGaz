package routes

import (
	"server/handlers"

	"github.com/gin-gonic/gin"
)

func TugasRoutes(r *gin.Engine) {
	tugasGroup := r.Group("/tugas")
	{
		tugasGroup.GET("/", handlers.GetAllTugas)
		tugasGroup.GET("/:id_tugas", handlers.GetTugas)
		tugasGroup.POST("/add", handlers.CreateTugas)
		tugasGroup.PUT("/update/:id_tugas", handlers.UpdateTugas)
		tugasGroup.DELETE("/delete/:id_tugas", handlers.DeleteTugas)

		// Route Upload Tugas
		tugasGroup.POST("/file/upload", handlers.UploadFileTugas)
		tugasGroup.GET("/file", handlers.IndexFileTugas)
		tugasGroup.GET("/file/:id_UT", handlers.GetFileTugas)
	}
}
