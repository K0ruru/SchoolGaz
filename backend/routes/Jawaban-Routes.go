package routes

import (
	"server/handlers"

	"github.com/gin-gonic/gin"
)

func JawabanRoutes(router *gin.Engine) {
	JawabanGroups := router.Group("/jawaban")
	{
		JawabanGroups.GET("/", handlers.GetJawaban)
		JawabanGroups.GET("/:id_jawaban", handlers.GetJwabanByid)
		JawabanGroups.POST("/add", handlers.CreateJawaban)
		JawabanGroups.PUT("/update/:id_jawaban", handlers.UpdateJawaban)
		JawabanGroups.DELETE("/delete/:id_jawaban", handlers.DeleteJawaban)
	}
}
