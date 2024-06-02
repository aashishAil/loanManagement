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
	middlewareI := middleware.Init(appInstance.ContextUtil(), appInstance.JwtUtil())
	r.Use(gin.CustomRecovery(middlewareI.Server().RecoverGinError()))

	apiGroup := r.Group("/api")

	apiGroup.GET("/ping", appRouter.Fallback().PingForGinRoute)

	attachUserRoutes(apiGroup.Group("/user"), appRouter.User(), middlewareI)

	r.NoRoute(appRouter.Fallback().NoRouteForGinHandler())

	return nil
}

func attachUserRoutes(router *gin.RouterGroup, customerRouter router.User, middlewareI middleware.Middleware) {
	router.POST("/login", customerRouter.Login)
	// TODO: attach authentication middleware
}
