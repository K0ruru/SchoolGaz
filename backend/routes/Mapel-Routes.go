package routes

import (
	"server/handlers"

	"github.com/gin-gonic/gin"
)

func MapelRoutes(router *gin.Engine) {
	MapelGroup := router.Group("/mapel")
	{
		MapelGroup.POST("/add", handlers.CreateMapel)
    MapelGroup.PUT("/update/:Id_mapel",handlers.UpdateMapel) 
	}

}
