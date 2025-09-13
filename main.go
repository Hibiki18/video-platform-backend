package main

import (
	"video-platform-backend/config"
	"video-platform-backend/router"
)

func main() {
	config.InitConfig()
	r := router.SetupRouter()
	r.Run(config.AppConfig.App.Port)
}
