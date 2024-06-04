package router

import (
	"net/http"
	"strconv"

	"loanManagement/appError"
	"loanManagement/constant"
	"loanManagement/handler"
	handlerModel "loanManagement/handler/model"
	"loanManagement/logger"
	routerModel "loanManagement/router/model"
	"loanManagement/util"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type Admin interface {
	ViewLoan(c *gin.Context)
	UpdateLoan(c *gin.Context)
}

type admin struct {
	adminHandler handler.Admin

	contextUtil util.Context
}

func (router *admin) ViewLoan(c *gin.Context) {
	ctx := router.contextUtil.CreateContextFromGinContext(c)

	userI := router.contextUtil.GetLoggedInUser(ctx)
	// skipping check for userI being null as that should be already validated in the authentication middleware
	if userI.Type != constant.UserTypeAdmin {
		c.JSON(http.StatusForbidden, constant.AdminOnlyRouteResponse)
		return
	}

	status := constant.LoanStatus(c.Param("status"))

	if status == "" {
		c.JSON(http.StatusBadRequest, constant.InvalidStatusResponse)
		return
	}

	if _, found := constant.LoanStatusMap[status]; !found {
		c.JSON(http.StatusBadRequest, constant.InvalidStatusResponse)
		return
	}

	repaymentFetch := c.Query("fetchScheduledRepayments")
	fetchRepayments := false
	if repaymentFetch != "" {
		var err error
		fetchRepayments, err = strconv.ParseBool(repaymentFetch)
		if err != nil {
			logger.Log.Error("failed to parse fetchScheduledRepayments",
				logger.Error(err),
				logger.String("api", "AdminViewLoan"),
			)
			c.JSON(http.StatusBadRequest, constant.InvalidFetchScheduledRepaymentsResponse)
			return
		}
	}

	loanData, err := router.adminHandler.FetchLoans(ctx, handlerModel.FetchAdminLoansInput{
		Status:                  &status,
		FetchScheduledRepayment: fetchRepayments,
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
				logger.String("api", "AdminViewLoan"),
			)
		}
		c.JSON(statusCode, errResp)
		return
	}

	resp := routerModel.GetAdminLoansOutput{
		Loans: make([]routerModel.UserLoan, len(loanData.Loans)),
	}

	for i := range loanData.Loans {
		loanI := loanData.Loans[i]
		resp.Loans[i] = loanI.TransformForRouter(loanData.LoanScheduledRepayments[loanI.ID])
	}

	c.JSON(http.StatusOK, resp)
}

func (router *admin) UpdateLoan(c *gin.Context) {
	ctx := router.contextUtil.CreateContextFromGinContext(c)

	userI := router.contextUtil.GetLoggedInUser(ctx)
	// skipping check for userI being null as that should be already validated in the authentication middleware
	if userI.Type != constant.UserTypeAdmin {
		c.JSON(http.StatusForbidden, constant.AdminOnlyRouteResponse)
		return
	}

	input := routerModel.UpdateAdminLoanInput{}
	if err := c.BindJSON(&input); err != nil {
		logger.Log.Info("invalid input", logger.Error(err))
		c.JSON(http.StatusBadRequest, constant.InvalidInputResponse)
		return
	}

	if input.Status == "" {
		logger.Log.Info("blank status provided")
		c.JSON(http.StatusBadRequest, constant.InvalidStatusResponse)
		return
	}

	if _, found := constant.LoanStatusMap[input.Status]; !found {
		logger.Log.Info("invalid status provided", logger.String("status", string(input.Status)))
		c.JSON(http.StatusBadRequest, constant.InvalidStatusResponse)
		return
	}

	if input.Status == constant.LoanStatusPending || input.Status == constant.LoanStatusPaid {
		logger.Log.Info("invalid status provided", logger.String("status", string(input.Status)))
		c.JSON(http.StatusBadRequest, constant.InvalidStatusResponse)
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

	loanData, err := router.adminHandler.FetchLoans(ctx, handlerModel.FetchAdminLoansInput{
		LoanIDs: []uuid.UUID{loanID},
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
				logger.String("api", "AdminUpdateLoan"),
			)
		}
		c.JSON(statusCode, errResp)
		return
	}

	if len(loanData.Loans) == 0 {
		c.JSON(http.StatusNotFound, constant.InvalidLoanIDResponse)
		return
	}

	loanI := loanData.Loans[0]

	_, err = router.adminHandler.UpdateLoanAndScheduledRepayment(ctx, handlerModel.UpdateLoanAndScheduledRepaymentInput{
		LoanI:  loanI,
		Status: input.Status,
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
				logger.String("api", "AdminUpdateLoan"),
			)
		}
		c.JSON(statusCode, errResp)
		return
	}

	c.Status(http.StatusNoContent)
}

func NewAdmin(
	adminHandler handler.Admin,

	contextUtil util.Context,
) Admin {
	return &admin{
		adminHandler: adminHandler,

		contextUtil: contextUtil,
	}
}
