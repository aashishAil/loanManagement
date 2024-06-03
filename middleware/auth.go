package middleware

import (
	"net/http"

	"loanManagement/constant"
	"loanManagement/logger"
	"loanManagement/util"

	"github.com/gin-gonic/gin"
)

type Auth interface {
	Authenticate() gin.HandlerFunc
}

type auth struct {
	contextUtil util.Context
	jwtUtil     util.Jwt
}

func (middleware *auth) Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		authToken := c.Request.Header.Get(constant.AuthHeader)
		requestPath := c.Request.URL.Path
		requestMethod := c.Request.Method

		message := constant.UnableToAuthenticateResponse

		if authToken == "" {
			logger.Log.Info("unauthorized access", logger.String("method", requestMethod),
				logger.String("path", requestPath))
			message = constant.MissingAuthTokenResponse
			c.AbortWithStatusJSON(http.StatusUnauthorized, message)
			return
		}

		userI, err := middleware.jwtUtil.ValidateToken(authToken)
		if err != nil {
			logger.Log.Error("error validating token", logger.Error(err),
				logger.String("method", requestMethod), logger.String("path", requestPath))
			c.AbortWithStatusJSON(http.StatusUnauthorized, message)
			return
		}

		if userI == nil {
			logger.Log.Error("nil user obtained", logger.Error(err),
				logger.String("method", requestMethod), logger.String("path", requestPath))
			c.AbortWithStatusJSON(http.StatusUnauthorized, message)
			return
		}

		ctx := c.Request.Context()
		ctx = middleware.contextUtil.StoreLoggedInUser(ctx, *userI)

		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func NewAuth(
	contextUtil util.Context,
	jwtUtl util.Jwt,
) Auth {
	return &auth{
		contextUtil: contextUtil,
		jwtUtil:     jwtUtl,
	}
}
