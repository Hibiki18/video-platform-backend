package controllers

import (
	"net/http"
	"video-platform-backend/internal/models"
	"video-platform-backend/utils"

	"github.com/gin-gonic/gin"
)

func CurrentUser(c *gin.Context) {
	// 从token中解析出username
	username, err := utils.ExtractUsername(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 根据username从数据库查询数据
	u, err := models.GetUserByName(username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data": u,
	})
}