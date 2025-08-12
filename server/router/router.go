package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Initialize() {
	router := gin.Default()

	router.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": "Health",
		})
	})

	initializeRoutes(router)
	router.Run()
}
