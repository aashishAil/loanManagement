package router

import (
	"net/http"

	"loanManagement/logger"

	"github.com/gin-gonic/gin"
)

type Fallback interface {
	NoRouteForGinHandler() gin.HandlerFunc
	PingForGinRoute(c *gin.Context)
}

type fallback struct {
}

func (*fallback) NoRouteForGinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Log.Error("unknown route requested", logger.String("path", c.Request.URL.Path))
		c.JSON(http.StatusNotFound, gin.H{
			"error": "not found",
		})
	}
}

func (*fallback) PingForGinRoute(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func NewDefault() Fallback {
	return &fallback{}
}
