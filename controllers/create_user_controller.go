package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/pwnkcin/go-api/models"
)

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	// check if user already exists (this is a placeholder, implement your own logic)

	fmt.Println("User data:", user)
}
