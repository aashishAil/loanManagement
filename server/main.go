package server

import (
	"loanManagement/config"
	"loanManagement/instance"
	"loanManagement/middleware"
	"loanManagement/router"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func Start(envConfig config.Config) error {
	r := gin.Default()

	err := attachRoutes(r)
	if err != nil {
		return errors.Wrap(err, "failed to attach routes")
	}

	err = r.Run(envConfig.ServerPort())
	if err != nil {
		return errors.Wrap(err, "failed to start server")
	}
	return nil
}

func attachRoutes(r *gin.Engine) error {
	// recover in case of panic between any endpoint calls
	appInstance, err := instance.Init()
	if err != nil {
		return errors.Wrap(err, "failed to initialize app instances")
	}

	appRouter := router.Init(appInstance)
	r.Use(gin.CustomRecovery(middleware.RecoverGinError()))

	r.GET("/ping", appRouter.Default().PingForGinRoute)
	r.POST("/login", appRouter.User().Login)

	r.NoRoute(appRouter.Default().NoRouteForGinHandler())

	return nil
}
