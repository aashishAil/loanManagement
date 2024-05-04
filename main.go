package main

import (
	"loanManagement/config"
	"loanManagement/logger"
	"loanManagement/server"
)

func main() {
	envConfig := config.Env
	err := server.Start(envConfig)
	if err != nil {
		logger.Logger.Errorf("failed to start server: %v", err)
	} else {
		logger.Logger.Info("server started successfully")
	}
}
