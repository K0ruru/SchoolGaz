package routes

import (
	"server/handlers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {
    AuthGroup := router.Group("/Auth")
    {
       AuthGroup.POST("/add", handlers.CreateUser)
       AuthGroup.PUT("/update/:NIS", handlers.UpdateUser)
       AuthGroup.GET("/get/:NIS", handlers.GetUser);
       AuthGroup.GET("/show", handlers.GetALLuser) 
       AuthGroup.DELETE("/del/:NIS", handlers.DelUser)
       AuthGroup.POST("/login", handlers.LoginUser)
       AuthGroup.GET("/siswa/:Kelas", handlers.GetAllUserByKelas)
       AuthGroup.PUT("/pp/:NIS", handlers.UploadProfilePicture)
    }
}


