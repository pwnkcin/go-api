package router

import (
	"github.com/gin-gonic/gin"
	controller "github.com/pwnkcin/go-api/controllers"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/signup", controller.CreateUser)
		}
	}
}
