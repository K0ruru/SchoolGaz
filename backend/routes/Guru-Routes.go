package routes

import (
	"server/handlers"

	"github.com/gin-gonic/gin"
)


func GuruRoutes(router *gin.Engine){
  GuruGroups := router.Group("/guru")
  {
    GuruGroups.POST("/add", handlers.CreateGuru)   
    GuruGroups.POST("/login", handlers.Loginguru)
    GuruGroups.PUT("/update/:NIS", handlers.Updateguru)
    GuruGroups.DELETE("/del/:NIS",handlers.DelGuru)
    GuruGroups.GET("/",handlers.GetALLguru)
    GuruGroups.GET("/show/:NIS",handlers.GetGuru)
  }
}
