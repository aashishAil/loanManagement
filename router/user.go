package router

import (
	"errors"
	"net/http"

	"loanManagement/appError"
	"loanManagement/constant"
	"loanManagement/handler"
	"loanManagement/logger"
	routerModel "loanManagement/router/model"
	"loanManagement/util"

	"github.com/gin-gonic/gin"
)

type User interface {
	Login(c *gin.Context)
}

type user struct {
	userHandler handler.User

	contextUtl   util.Context
	jwtUtil      util.Jwt
	passwordUtil util.Password
}

func (u *user) Login(c *gin.Context) {
	ctx := u.contextUtl.CreateContextFromGinContext(c)
	input := routerModel.UserLoginInput{}
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, constant.InvalidInputResponse)
		return
	}

	if input.Email == "" {
		c.JSON(http.StatusBadRequest, constant.EmptyEmailResponse)
		return
	}
	if input.Password == "" {
		c.JSON(http.StatusBadRequest, constant.EmptyPasswordResponse)
		return
	}

	token, err := u.userHandler.CheckValidCredentials(ctx, input.Email, input.Password)
	if err != nil {
		statusCode := http.StatusInternalServerError
		errResp := constant.DefaultErrorResponse

		var customErr appError.Custom
		ok := errors.As(err, &customErr)
		if ok {
			logger.Log.Error("handled error",
				logger.Any("err", customErr),
				logger.String("api", "Login"),
			)
			statusCode = customErr.Code
			errResp.Error = customErr.Err.Error()
		} else {
			logger.Log.Error("unexpected error",
				logger.Any("err", err),
				logger.String("api", "Login"),
			)
		}
		c.JSON(statusCode, errResp)
		return
	}

	resp := routerModel.UserLoginOutput{
		Token: token,
	}

	c.JSON(http.StatusOK, resp)

}

func NewUser(
	userHandler handler.User,

	contextUtl util.Context,
	jwtUtil util.Jwt,
	passwordUtil util.Password,
) User {
	return &user{
		userHandler: userHandler,

		contextUtl:   contextUtl,
		jwtUtil:      jwtUtil,
		passwordUtil: passwordUtil,
	}
}
