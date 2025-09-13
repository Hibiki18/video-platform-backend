package router

import (
	"video-platform-backend/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	auth := r.Group("/auth")
	{
		auth.POST("/login", func(c *gin.Context) {
			// Handle login
			c.JSON(200, gin.H{"message": "login"})
		})
		auth.POST("/register", func(c *gin.Context) {
			// Handle registration
			controllers.Register(c)
		})
	}

	return r
}