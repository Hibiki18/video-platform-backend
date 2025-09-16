package controllers

import (
	"context"
	"fmt"
	"net/http"
	"path/filepath"
	"time"

	"video-platform-backend/config"
	"video-platform-backend/internal/models"
	"video-platform-backend/utils"

	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
)

// UploadVideo 上传视频到 MinIO
func UploadVideo(c *gin.Context) {
	// 从 JWT 中获取 username（需要中间件提前设置）
	username, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// 获取表单参数
	title := c.PostForm("title")
	description := c.PostForm("description")

	// 获取文件
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File is required"})
		return
	}

	// 打开文件
	src, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open file"})
		return
	}
	defer src.Close()

	// 生成文件名
	fileName := fmt.Sprintf("%d_%s", time.Now().Unix(), filepath.Base(file.Filename))

	// 上传到 MinIO
	_, err = utils.MinioClient.PutObject(
		context.Background(),
		config.AppConfig.Minio.Bucket,
		fileName,
		src,
		file.Size,
		minio.PutObjectOptions{ContentType: file.Header.Get("Content-Type")},
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload to MinIO"})
		return
	}

	// 拼接视频访问 URL
	videoURL := fmt.Sprintf("http://%s/%s/%s",
		config.AppConfig.Minio.Endpoint,
		config.AppConfig.Minio.Bucket,
		fileName,
	)

	// 写入数据库
	video := models.Video{
		Username:    username.(string), // 这里假设 JWT 中间件存的是 string
		Title:       title,
		Description: description,
		FilePath:    videoURL,
		CreatedAt:   time.Now(),
	}
	if err := config.DB.Create(&video).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save video"})
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, gin.H{
		"message": "Upload successful",
		"video":   video,
	})
}
