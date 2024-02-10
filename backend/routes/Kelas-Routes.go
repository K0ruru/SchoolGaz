package routes

import (
	"server/handlers"

	"github.com/gin-gonic/gin"
)

func KelasRoutes(router *gin.Engine){
  KelasGroup := router.Group("/kelas")
  {

    KelasGroup.GET("/",handlers.GetAllKelas)

  } 
}
