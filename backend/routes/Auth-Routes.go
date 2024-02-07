package routes

import (
    "github.com/gin-gonic/gin"
    "server/handlers"
)

func AuthRoutes(router *gin.Engine) {
    AuthGroup := router.Group("/Auth")
    {
        AuthGroup.GET("/",handlers.GetUsers)
    AuthGroup.POST("/add", handlers.CreateUser)
        AuthGroup.DELETE("/delete/:nis", handlers.DeleteUser)
        AuthGroup.PUT("/update/:nis", handlers.UpdateUser)
        AuthGroup.POST("/login", handlers.LoginHandler)
        
    }
}
