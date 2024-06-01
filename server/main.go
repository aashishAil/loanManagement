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

	apiGroup := r.Group("/api")

	apiGroup.GET("/ping", appRouter.Default().PingForGinRoute)

	attachUserRoutes(apiGroup.Group("/user"), appRouter.User())

	r.NoRoute(appRouter.Default().NoRouteForGinHandler())

	return nil
}

func attachUserRoutes(router *gin.RouterGroup, customerRouter router.User) {
	router.POST("/login", customerRouter.Login)
	// TODO: attach authentication middleware
}
