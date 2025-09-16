package main

import (
	"video-platform-backend/config"
	"video-platform-backend/router"
	//"video-platform-backend/utils"
)

func main() {
	config.InitConfig()
	//utils.InitMinio()


	r := router.SetupRouter()
	r.Run(config.AppConfig.App.Port)

}
