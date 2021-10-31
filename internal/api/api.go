package api

import (
	"cem-taxi-api/config"
	"cem-taxi-api/internal/router"
	"fmt"
)

func setConfiguration(configPath string) {
	config.Setup(configPath)
}

func Run(configPath string) {
	if configPath == "" {
		configPath = "config/config.yaml"
	}
	setConfiguration(configPath)

	conf := config.GetConfig()
	web := router.Setup()
	port := conf.Server.Port
	fmt.Println("Server running on port: " + port)
	_ = web.Run(":" + port)
}
