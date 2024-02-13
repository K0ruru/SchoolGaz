package routes

import (
	"server/handlers"

	"github.com/gin-gonic/gin"
)

func MapelRoutes(router *gin.Engine) {
	MapelGroup := router.Group("/mapel")
	{
		MapelGroup.POST("/add", handlers.CreateMapel)
		MapelGroup.PUT("/:id_mapel", handlers.UpdateMapel)
    MapelGroup.DELETE("/:id_mapel",handlers.DeleteMapel)
    MapelGroup.GET("/",handlers.GetALLMapel)
    MapelGroup.GET("/:id_mapel", handlers.GetMapel)
	}

}
