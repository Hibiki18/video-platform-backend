package router

import (
	"fmt"
	"video-platform-backend/internal/controllers"
	"video-platform-backend/internal/middleware"
	"video-platform-backend/logger"
	"video-platform-backend/utils"

	"time"

	"github.com/gin-contrib/cors"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	// 初始化 zap
	logger.InitLogger()
	zapLogger := logger.Logger.Desugar() // 拿到 *zap.Logger，给 gin 使用

	// 替换 gin 默认中间件：日志 和 恢复
	r := gin.New()
	r.Use(ginzap.Ginzap(zapLogger, time.RFC3339, true)) // zap 格式的访问日志
	r.Use(ginzap.RecoveryWithZap(zapLogger, true))      // panic 时自动记录堆栈

	// CORS 中间件配置
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // 前端地址
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	//路由
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
			fmt.Println(c.GetString("username"))
			c.JSON(200, gin.H{"message": "you are logged in", "username": username})
		})
	}

	logger.Logger.Info("服务启动成功")
	logger.Logger.Infof("当前端口：%d", 8080)

	return r
}
