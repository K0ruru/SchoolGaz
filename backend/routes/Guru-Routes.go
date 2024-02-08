package routes

import(
 "github.com/gin-gonic/gin"
 "server/handlers"
)

func GuruRoutes(R *gin.Engine){
  GuruGroup := R.Group("/guru")
  {
   GuruGroup.POST("/add", handlers.CreateGuru)
    GuruGroup.PUT("/update/:NIK", handlers.UpdateGuru)
  } 

}
