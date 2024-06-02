package middleware

import (
	"net/http"

	"loanManagement/logger"

	"github.com/gin-gonic/gin"
)

type Server interface {
	RecoverGinError() func(c *gin.Context, err interface{})
}

type server struct {
}

func (middleware *server) RecoverGinError() func(c *gin.Context, err interface{}) {
	handler := func(c *gin.Context, err interface{}) {
		message := gin.H{
			"error": "unknown error",
		}
		var e error
		if err == nil {
			e = nil
			logger.Log.Info("nil error in gin recovery handler")
		} else {
			e = err.(error)
			message = gin.H{
				"error": e.Error(),
			}
			logger.Log.Error("gin recovery handler", logger.Error(e))
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, message)
	}
	return handler
}

func NewServer() Server {
	return &server{}
}
