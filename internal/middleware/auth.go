package middleware

import (
	"net/http"
	"video-platform-backend/utils"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 解析并验证 token
		err := utils.ParseToken(c)
		if err != nil {
			c.String(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}
		username, err := utils.ExtractUsername(c)
		if err != nil {
			c.String(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}
		c.Set("username", username)
		c.Next()
	}
}
