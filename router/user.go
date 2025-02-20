package router

import (
	"errors"
	"github.com/google/uuid"
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
	ViewLoan(c *gin.Context)
	RecordPayment(c *gin.Context)
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
		logger.Log.Info("invalid input", logger.Error(err))
		c.JSON(http.StatusBadRequest, constant.InvalidInputResponse)
		return
	}

	if input.Email == "" {
		logger.Log.Info("email is empty")
		c.JSON(http.StatusBadRequest, constant.EmptyEmailResponse)
		return
	}
	if input.Password == "" {
		logger.Log.Info("password is empty")
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
			statusCode = customErr.Code
			errResp.Error = customErr.Err.Error()
		} else {
			logger.Log.Error("unexpected error",
				logger.Error(err),
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
		logger.Log.Info("invalid input", logger.Error(err))
		c.JSON(http.StatusBadRequest, constant.InvalidInputResponse)
		return
	}

	if input.Amount <= 0 {
		logger.Log.Info("invalid amount", logger.Int64("amount", input.Amount))
		c.JSON(http.StatusBadRequest, constant.InvalidAmountResponse)
		return
	}

	if input.Term <= 0 {
		logger.Log.Info("invalid term", logger.Int64("term", input.Term))
		c.JSON(http.StatusBadRequest, constant.InvalidTermResponse)
		return
	}

	if input.Currency == "" {
		logger.Log.Info("currency is empty")
		c.JSON(http.StatusBadRequest, constant.InvalidCurrencyResponse)
		return
	}

	if input.DisbursalDate.IsZero() {
		logger.Log.Info("disbursal date is empty")
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
			statusCode = customErr.Code
			errResp.Error = customErr.Err.Error()
		} else {
			logger.Log.Error("unexpected error",
				logger.Error(err),
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

func (router *user) ViewLoan(c *gin.Context) {
	ctx := router.contextUtil.CreateContextFromGinContext(c)

	userI := router.contextUtil.GetLoggedInUser(ctx)
	// skipping check for userI being null as that should be already validated in the authentication middleware
	if userI.Type != constant.UserTypeCustomer {
		c.JSON(http.StatusForbidden, constant.CustomerOnlyRouteResponse)
		return
	}

	loanData, err := router.userHandler.FetchLoans(ctx, handlerModel.FetchUserLoanInput{
		UserID: userI.ID,
	})
	if err != nil {
		statusCode := http.StatusInternalServerError
		errResp := constant.DefaultErrorResponse

		customErr := appError.Custom{}
		ok := errors.As(err, &customErr)
		if ok {
			statusCode = customErr.Code
			errResp.Error = customErr.Err.Error()
		} else {
			logger.Log.Error("unexpected error",
				logger.Error(err),
				logger.String("api", "ViewLoan"),
			)
		}
		c.JSON(statusCode, errResp)
		return
	}

	resp := routerModel.GetUserLoansOutput{
		Loans: make([]routerModel.UserLoan, len(loanData.Loans)),
	}

	for i := range loanData.Loans {
		loanI := loanData.Loans[i]
		resp.Loans[i] = loanI.TransformForRouter(loanData.LoanScheduledRepayments[loanI.ID])
	}

	c.JSON(http.StatusOK, resp)
}

func (router *user) RecordPayment(c *gin.Context) {
	ctx := router.contextUtil.CreateContextFromGinContext(c)

	userI := router.contextUtil.GetLoggedInUser(ctx)
	// skipping check for userI being null as that should be already validated in the authentication middleware
	if userI.Type != constant.UserTypeCustomer {
		c.JSON(http.StatusForbidden, constant.CustomerOnlyRouteResponse)
		return
	}

	loanID, err := uuid.Parse(c.Param("ID"))
	if err != nil {
		logger.Log.Error("failed to parse loanID",
			logger.Error(err),
			logger.String("api", "AdminUpdateLoan"),
		)
		c.JSON(http.StatusBadRequest, constant.InvalidLoanIDResponse)
		return
	}

	input := routerModel.UserRecordPaymentInput{}
	if err := c.BindJSON(&input); err != nil {
		logger.Log.Info("invalid input", logger.Error(err))
		c.JSON(http.StatusBadRequest, constant.InvalidInputResponse)
		return
	}

	if input.Amount <= 0 {
		logger.Log.Info("invalid amount", logger.Float64("amount", input.Amount))
		c.JSON(http.StatusBadRequest, constant.InvalidAmountResponse)
		return
	}

	loanData, err := router.userHandler.AddLoanPayment(ctx, handlerModel.AddUserLoanPaymentInput{
		LoanID: loanID,
		UserID: userI.ID,
		Amount: input.Amount,
	})
	if err != nil {
		statusCode := http.StatusInternalServerError
		errResp := constant.DefaultErrorResponse

		customErr := appError.Custom{}
		ok := errors.As(err, &customErr)
		if ok {
			statusCode = customErr.Code
			errResp.Error = customErr.Err.Error()
		} else {
			logger.Log.Error("unexpected error",
				logger.Error(err),
				logger.String("api", "ViewLoan"),
			)
		}
		c.JSON(statusCode, errResp)
		return
	}

	resp := routerModel.UserRecordPaymentOutput{
		IsLoanClosed:      loanData.IsLoanClosed,
		PendingAmount:     loanData.PendingAmount,
		NextDueDate:       loanData.NextDueDate,
		NextPaymentAmount: loanData.NextPaymentAmount,
	}

	c.JSON(http.StatusCreated, resp)
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
