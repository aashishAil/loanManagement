package router

import (
	"errors"
	"net/http"

	"loanManagement/appError"
	"loanManagement/constant"
	"loanManagement/handler"
	handlerModel "loanManagement/handler/model"
	"loanManagement/logger"
	routerModel "loanManagement/router/model"
	"loanManagement/util"

	"github.com/gin-gonic/gin"
)

type User interface {
	Login(c *gin.Context)
	CreateLoan(c *gin.Context)
}

type user struct {
	userHandler handler.User

	contextUtil  util.Context
	jwtUtil      util.Jwt
	passwordUtil util.Password
}

func (router *user) Login(c *gin.Context) {
	ctx := router.contextUtil.CreateContextFromGinContext(c)
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

	token, err := router.userHandler.CheckValidCredentials(ctx, input.Email, input.Password)
	if err != nil {
		statusCode := http.StatusInternalServerError
		errResp := constant.DefaultErrorResponse

		customErr := appError.Custom{}
		ok := errors.As(err, &customErr)
		if ok {
			logger.Log.Info("handled error",
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

func (router *user) CreateLoan(c *gin.Context) {
	ctx := router.contextUtil.CreateContextFromGinContext(c)

	userI := router.contextUtil.GetLoggedInUser(ctx)
	// skipping check for userI being null as that should be already validated in the authentication middleware
	if userI.Type != constant.UserTypeCustomer {
		c.JSON(http.StatusForbidden, constant.CustomerOnlyRouteResponse)
		return
	}

	input := routerModel.UserCreateLoanInput{}
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, constant.InvalidInputResponse)
		return
	}

	if input.Amount <= 0 {
		c.JSON(http.StatusBadRequest, constant.InvalidAmountResponse)
		return
	}

	if input.Term <= 0 {
		c.JSON(http.StatusBadRequest, constant.InvalidTermResponse)
		return
	}

	if input.Currency == "" {
		c.JSON(http.StatusBadRequest, constant.InvalidCurrencyResponse)
		return
	}

	if input.DisbursalDate.IsZero() {
		c.JSON(http.StatusBadRequest, constant.InvalidDisbursalDateResponse)
		return
	}

	loanID, err := router.userHandler.CreateLoan(ctx, handlerModel.CreateUserLoanInput{
		UserID:        userI.ID,
		Amount:        input.Amount,
		Currency:      input.Currency,
		Term:          input.Term,
		DisbursalDate: input.DisbursalDate,
	})
	if err != nil {
		statusCode := http.StatusInternalServerError
		errResp := constant.DefaultErrorResponse

		customErr := appError.Custom{}
		ok := errors.As(err, &customErr)
		if ok {
			logger.Log.Info("handled error",
				logger.Any("err", customErr),
				logger.String("api", "CreateLoan"),
			)
			statusCode = customErr.Code
			errResp.Error = customErr.Err.Error()
		} else {
			logger.Log.Error("unexpected error",
				logger.Any("err", err),
				logger.String("api", "CreateLoan"),
			)
		}
		c.JSON(statusCode, errResp)
		return
	}

	resp := routerModel.UserCreateLoanOutput{
		LoanID: loanID,
	}

	c.JSON(http.StatusOK, resp)
}

func NewUser(
	userHandler handler.User,

	contextUtil util.Context,
	jwtUtil util.Jwt,
	passwordUtil util.Password,
) User {
	return &user{
		userHandler: userHandler,

		contextUtil:  contextUtil,
		jwtUtil:      jwtUtil,
		passwordUtil: passwordUtil,
	}
}
