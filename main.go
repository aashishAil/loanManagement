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
		logger.Log.Errorf("failed to start server: %v", err)
	} else {
		logger.Log.Info("server started successfully")
	}
}
