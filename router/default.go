package router

import (
	"net/http"

	"loanManagement/logger"

	"github.com/gin-gonic/gin"
)

func NoRouteForGinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Log.Error("unknown route requested", logger.String("path", c.Request.URL.Path))
		c.JSON(http.StatusNotFound, gin.H{
			"error": "not found",
		})
	}
}

func PingForGinRoute(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}
