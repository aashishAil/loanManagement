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
	"github.com/pkg/errors"
)

type Admin interface {
	ViewLoan(c *gin.Context)
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
		Status:                  status,
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

func NewAdmin(
	adminHandler handler.Admin,

	contextUtil util.Context,
) Admin {
	return &admin{
		adminHandler: adminHandler,

		contextUtil: contextUtil,
	}
}
