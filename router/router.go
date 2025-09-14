package router

import (
	"video-platform-backend/internal/controllers"
	"video-platform-backend/internal/middleware"
	"video-platform-backend/utils"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	auth := r.Group("/auth")
	{
		auth.POST("/login", func(c *gin.Context) {
			// Handle login
			controllers.Login(c)
		})
		auth.POST("/register", func(c *gin.Context) {
			// Handle registration
			controllers.Register(c)
		})
	}

	api := r.Group("/api")
	api.Use(middleware.AuthMiddleware())
	{
		api.GET("/me", func(c *gin.Context) {
			username, _ := utils.ExtractUsername(c)
			c.JSON(200, gin.H{"message": "you are logged in", "username": username})
		})
	}

	return r
}
