package server

import (
	"loanManagement/config"
	"loanManagement/middleware"
	"loanManagement/router"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func Start(envConfig config.Config) error {
	r := gin.Default()

	r.Use(gin.CustomRecovery(middleware.RecoverGinError()))

	r.GET("/ping", router.PingForGinRoute)
	r.NoRoute(router.NoRouteForGinHandler())

	err := r.Run(envConfig.ServerPort())
	if err != nil {
		return errors.Wrap(err, "failed to start server")
	}
	return nil
}
