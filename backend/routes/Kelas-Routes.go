package routes

import (
	"server/handlers"

	"github.com/gin-gonic/gin"
)

func KelasRoutes(router *gin.Engine){
  KelasGroup := router.Group("/kelas")
  {

    KelasGroup.GET("/",handlers.GetAllKelas)
    KelasGroup.GET("/:id", handlers.GetKelas)
    KelasGroup.POST("/create", handlers.CreateKelas)
    KelasGroup.PUT("/:id", handlers.EditKelas)
    KelasGroup.DELETE("/:id", handlers.DeleteKelas)
  } 
}
