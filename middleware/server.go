package middleware

import (
	"github.com/gin-gonic/gin"
	"loanManagement/logger"
	"net/http"
)

func RecoverGinError() func(c *gin.Context, err interface{}) {
	handler := func(c *gin.Context, err interface{}) {
		message := gin.H{
			"error": "unknown error",
		}
		var e error
		if err == nil {
			e = nil
			logger.Logger.Info("nil error in gin recovery handler")
		} else {
			e = err.(error)
			message = gin.H{
				"error": e.Error(),
			}
			logger.Logger.Error("gin recovery handler", logger.Any("error", e))
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, message)
	}
	return handler
}
