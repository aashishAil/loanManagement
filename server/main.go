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

	attachRoutes(r)

	err := r.Run(envConfig.ServerPort())
	if err != nil {
		return errors.Wrap(err, "failed to start server")
	}
	return nil
}

func attachRoutes(r *gin.Engine) {
	// recover in case of panic between any endpoint calls
	appRouter := router.Init()
	r.Use(gin.CustomRecovery(middleware.RecoverGinError()))

	r.GET("/ping", appRouter.Default().PingForGinRoute)

	r.NoRoute(appRouter.Default().NoRouteForGinHandler())
}
