package routes

import (
	"server/handlers"

	"github.com/gin-gonic/gin"
)

func KelasRoutes(router *gin.Engine){
  KelasGroup := router.Group("/kelas")
  {
    KelasGroup.POST("/add", handlers.Createkelas)
    KelasGroup.PUT("/update/:id_kelas", handlers.Updatekelas)
    KelasGroup.GET("/show",handlers.Getkelas)
    KelasGroup.DELETE("/delete/:id_kelas", handlers.DeleteKelas)
  } 
}
