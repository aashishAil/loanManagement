package handler

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
	"loanManagement/appError"
	"loanManagement/constant"
	databaseModel "loanManagement/database/model"
	"loanManagement/handler"
	handlerModel "loanManagement/handler/model"
	"loanManagement/logger"
	databaseMocks "loanManagement/mocks/database/instance"
	repoMocks "loanManagement/mocks/repo"
	repoModel "loanManagement/repo/model"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type TestAdminHandlerFetchLoansSuite struct {
	suite.Suite

	LoanI              databaseModel.Loan
	LoanStatus         constant.LoanStatus
	ScheduledRepayment *databaseModel.ScheduledRepayment

	MockLoanRepo               *repoMocks.Loan
	MockPostgresDBInstance     *databaseMocks.PostgresDB
	MockScheduledRepaymentRepo *repoMocks.ScheduledRepayment
}

func (suite *TestAdminHandlerFetchLoansSuite) SetupTest() {
	suite.LoanI = databaseModel.Loan{
		BaseWithUpdatedAt: databaseModel.BaseWithUpdatedAt{
			ID: uuid.New(),
		},
	}
	suite.LoanStatus = constant.LoanStatusPending
	suite.ScheduledRepayment = &databaseModel.ScheduledRepayment{
		LoanID: suite.LoanI.ID,
	}

	suite.MockLoanRepo = new(repoMocks.Loan)
	suite.MockPostgresDBInstance = new(databaseMocks.PostgresDB)
	suite.MockScheduledRepaymentRepo = new(repoMocks.ScheduledRepayment)
}

func (suite *TestAdminHandlerFetchLoansSuite) TestHandlesLoanFetchError() {
	testData := handlerModel.FetchAdminLoansInput{
		LoanIDs: []uuid.UUID{suite.LoanI.ID},
		Status:  nil,
	}
	testError := errors.New("test error")
	suite.MockLoanRepo.On("FindAll", mock.Anything, mock.Anything).Return(nil, testError)
	suite.MockLoanRepo.EXPECT().FindAll(mock.Anything, repoModel.FindAllLoanInput{
		IDs:    testData.LoanIDs,
		Status: testData.Status,
	})
	suite.MockScheduledRepaymentRepo.AssertNumberOfCalls(suite.T(), "FindAll", 0)

	admin := handler.NewAdmin(suite.MockLoanRepo, suite.MockScheduledRepaymentRepo, suite.MockPostgresDBInstance)
	result, err := admin.FetchLoans(context.Background(), testData)
	assert.Nil(suite.T(), result)
	assert.EqualError(suite.T(), err, testError.Error())
}

func (suite *TestAdminHandlerFetchLoansSuite) TestQueriesByValidLoanIDs() {
	testData := handlerModel.FetchAdminLoansInput{
		LoanIDs: []uuid.UUID{suite.LoanI.ID},
		Status:  nil,
	}
	suite.MockLoanRepo.On("FindAll", mock.Anything, mock.Anything).Return([]*databaseModel.Loan{&suite.LoanI}, nil)
	suite.MockLoanRepo.EXPECT().FindAll(mock.Anything, repoModel.FindAllLoanInput{
		IDs:    testData.LoanIDs,
		Status: testData.Status,
	})
	suite.MockScheduledRepaymentRepo.AssertNumberOfCalls(suite.T(), "FindAll", 0)

	admin := handler.NewAdmin(suite.MockLoanRepo, suite.MockScheduledRepaymentRepo, suite.MockPostgresDBInstance)
	result, err := admin.FetchLoans(context.Background(), testData)
	assert.NotNil(suite.T(), result)
	assert.Nil(suite.T(), err)
}

func (suite *TestAdminHandlerFetchLoansSuite) TestQueriesByValidStatus() {
	testData := handlerModel.FetchAdminLoansInput{
		Status: &suite.LoanStatus,
	}
	suite.MockLoanRepo.On("FindAll", mock.Anything, mock.Anything).Return([]*databaseModel.Loan{&suite.LoanI}, nil)
	suite.MockLoanRepo.EXPECT().FindAll(mock.Anything, repoModel.FindAllLoanInput{
		IDs:    testData.LoanIDs,
		Status: testData.Status,
	})

	admin := handler.NewAdmin(suite.MockLoanRepo, suite.MockScheduledRepaymentRepo, suite.MockPostgresDBInstance)
	result, err := admin.FetchLoans(context.Background(), testData)

	assert.NotNil(suite.T(), result)
	assert.Nil(suite.T(), err)
}

func (suite *TestAdminHandlerFetchLoansSuite) TestHandlesScheduledRepaymentsFetchError() {
	testData := handlerModel.FetchAdminLoansInput{
		LoanIDs:                 []uuid.UUID{suite.LoanI.ID},
		Status:                  nil,
		FetchScheduledRepayment: true,
	}
	testError := errors.New("test error")
	suite.MockLoanRepo.On("FindAll", mock.Anything, mock.Anything).Return([]*databaseModel.Loan{&suite.LoanI}, nil)
	suite.MockLoanRepo.EXPECT().FindAll(mock.Anything, repoModel.FindAllLoanInput{
		IDs:    testData.LoanIDs,
		Status: testData.Status,
	})
	suite.MockScheduledRepaymentRepo.On("FindAll", mock.Anything, mock.Anything).Return(nil, testError)

	admin := handler.NewAdmin(suite.MockLoanRepo, suite.MockScheduledRepaymentRepo, suite.MockPostgresDBInstance)
	result, err := admin.FetchLoans(context.Background(), testData)
	assert.Nil(suite.T(), result)
	assert.EqualError(suite.T(), err, testError.Error())
}

func (suite *TestAdminHandlerFetchLoansSuite) TestQueriesForScheduledRepayments() {
	testData := handlerModel.FetchAdminLoansInput{
		LoanIDs:                 []uuid.UUID{suite.LoanI.ID},
		Status:                  nil,
		FetchScheduledRepayment: true,
	}
	suite.MockLoanRepo.On("FindAll", mock.Anything, mock.Anything).Return([]*databaseModel.Loan{&suite.LoanI}, nil)
	suite.MockLoanRepo.EXPECT().FindAll(mock.Anything, repoModel.FindAllLoanInput{
		IDs:    testData.LoanIDs,
		Status: testData.Status,
	})
	suite.MockScheduledRepaymentRepo.On("FindAll", mock.Anything, mock.Anything).Return([]*databaseModel.ScheduledRepayment{suite.ScheduledRepayment}, nil)

	admin := handler.NewAdmin(suite.MockLoanRepo, suite.MockScheduledRepaymentRepo, suite.MockPostgresDBInstance)
	result, err := admin.FetchLoans(context.Background(), testData)

	assert.NotNil(suite.T(), result)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), &suite.LoanI, result.Loans[0])
	assert.Len(suite.T(), result.LoanScheduledRepayments, 1)
	assert.Equal(suite.T(), []*databaseModel.ScheduledRepayment{suite.ScheduledRepayment}, result.LoanScheduledRepayments[suite.LoanI.ID])
}

type TestAdminHandlerUpdateLoanAndScheduledRepaymentSuite struct {
	suite.Suite

	LoanI              databaseModel.Loan
	LoanStatus         constant.LoanStatus
	ScheduledRepayment *databaseModel.ScheduledRepayment

	MockLoanRepo               *repoMocks.Loan
	MockPostgresDBInstance     *databaseMocks.PostgresDB
	MockScheduledRepaymentRepo *repoMocks.ScheduledRepayment
	MockTransactionDB          *databaseMocks.PostgresTransactionDB
}

func (suite *TestAdminHandlerUpdateLoanAndScheduledRepaymentSuite) SetupTest() {
	suite.LoanI = databaseModel.Loan{
		BaseWithUpdatedAt: databaseModel.BaseWithUpdatedAt{
			ID: uuid.New(),
		},
		Status: constant.LoanStatusPending,
	}
	suite.LoanStatus = constant.LoanStatusPending
	suite.ScheduledRepayment = &databaseModel.ScheduledRepayment{
		LoanID: suite.LoanI.ID,
		Status: constant.ScheduleRepaymentStatusPending,
	}

	suite.MockLoanRepo = new(repoMocks.Loan)
	suite.MockPostgresDBInstance = new(databaseMocks.PostgresDB)
	suite.MockScheduledRepaymentRepo = new(repoMocks.ScheduledRepayment)
	suite.MockTransactionDB = new(databaseMocks.PostgresTransactionDB)
}

func (suite *TestAdminHandlerUpdateLoanAndScheduledRepaymentSuite) TestHandlesLoanTerminalStatus() {
	suite.LoanI.Status = constant.LoanStatusPaid
	testData := handlerModel.UpdateLoanAndScheduledRepaymentInput{
		LoanI:              &suite.LoanI,
		ScheduleRepayments: []*databaseModel.ScheduledRepayment{suite.ScheduledRepayment},
		Status:             suite.LoanStatus,
	}
	suite.MockPostgresDBInstance.AssertNumberOfCalls(suite.T(), "GetTransactionDb", 0)
	suite.MockLoanRepo.AssertNumberOfCalls(suite.T(), "Update", 0)
	suite.MockScheduledRepaymentRepo.AssertNumberOfCalls(suite.T(), "Update", 0)
	suite.MockTransactionDB.AssertNumberOfCalls(suite.T(), "Rollback", 0)
	suite.MockTransactionDB.AssertNumberOfCalls(suite.T(), "Commit", 0)

	admin := handler.NewAdmin(suite.MockLoanRepo, suite.MockScheduledRepaymentRepo, suite.MockPostgresDBInstance)
	result, err := admin.UpdateLoanAndScheduledRepayment(context.Background(), testData)
	assert.Nil(suite.T(), result)
	customErr := appError.Custom{}
	ok := errors.As(err, &customErr)
	assert.True(suite.T(), ok)
	assert.EqualError(suite.T(), err, constant.LoanInTerminalStatusError.Error())
}

func (suite *TestAdminHandlerUpdateLoanAndScheduledRepaymentSuite) TestHandlesTransactionDatabaseError() {
	testData := handlerModel.UpdateLoanAndScheduledRepaymentInput{
		LoanI:              &suite.LoanI,
		ScheduleRepayments: []*databaseModel.ScheduledRepayment{suite.ScheduledRepayment},
		Status:             suite.LoanStatus,
	}

	testError := errors.New("test error")
	suite.MockPostgresDBInstance.On("GetTransactionDb").Return(nil, testError)
	suite.MockLoanRepo.AssertNumberOfCalls(suite.T(), "Update", 0)
	suite.MockScheduledRepaymentRepo.AssertNumberOfCalls(suite.T(), "Update", 0)
	suite.MockTransactionDB.AssertNumberOfCalls(suite.T(), "Rollback", 0)
	suite.MockTransactionDB.AssertNumberOfCalls(suite.T(), "Commit", 0)

	admin := handler.NewAdmin(suite.MockLoanRepo, suite.MockScheduledRepaymentRepo, suite.MockPostgresDBInstance)
	result, err := admin.UpdateLoanAndScheduledRepayment(context.Background(), testData)
	assert.Nil(suite.T(), result)
	assert.EqualError(suite.T(), err, testError.Error())
}

func (suite *TestAdminHandlerUpdateLoanAndScheduledRepaymentSuite) TestHandlesLoanRepoUpdateRollbackError() {
	testData := handlerModel.UpdateLoanAndScheduledRepaymentInput{
		LoanI:              &suite.LoanI,
		ScheduleRepayments: []*databaseModel.ScheduledRepayment{suite.ScheduledRepayment},
		Status:             suite.LoanStatus,
	}

	testError := errors.New("test error")
	rollbackError := errors.New("test error 2")
	suite.MockPostgresDBInstance.On("GetTransactionDb").Return(suite.MockTransactionDB, nil)
	suite.MockLoanRepo.On("Update", mock.Anything, mock.Anything).Return(testError)
	suite.MockTransactionDB.On("Rollback").Return(rollbackError)
	suite.MockScheduledRepaymentRepo.AssertNumberOfCalls(suite.T(), "Update", 0)
	suite.MockTransactionDB.AssertNumberOfCalls(suite.T(), "Commit", 0)

	admin := handler.NewAdmin(suite.MockLoanRepo, suite.MockScheduledRepaymentRepo, suite.MockPostgresDBInstance)
	result, err := admin.UpdateLoanAndScheduledRepayment(context.Background(), testData)
	assert.Nil(suite.T(), result)
	assert.EqualError(suite.T(), err, rollbackError.Error())
}

func (suite *TestAdminHandlerUpdateLoanAndScheduledRepaymentSuite) TestHandlesLoanRepoUpdateError() {
	testData := handlerModel.UpdateLoanAndScheduledRepaymentInput{
		LoanI:              &suite.LoanI,
		ScheduleRepayments: []*databaseModel.ScheduledRepayment{suite.ScheduledRepayment},
		Status:             suite.LoanStatus,
	}

	testError := errors.New("test error")
	suite.MockPostgresDBInstance.On("GetTransactionDb").Return(suite.MockTransactionDB, nil)
	suite.MockLoanRepo.On("Update", mock.Anything, mock.Anything).Return(testError)
	suite.MockTransactionDB.On("Rollback").Return(nil)
	suite.MockScheduledRepaymentRepo.AssertNumberOfCalls(suite.T(), "Update", 0)
	suite.MockTransactionDB.AssertNumberOfCalls(suite.T(), "Commit", 0)

	admin := handler.NewAdmin(suite.MockLoanRepo, suite.MockScheduledRepaymentRepo, suite.MockPostgresDBInstance)
	result, err := admin.UpdateLoanAndScheduledRepayment(context.Background(), testData)
	assert.Nil(suite.T(), result)
	assert.EqualError(suite.T(), err, testError.Error())
}

func (suite *TestAdminHandlerUpdateLoanAndScheduledRepaymentSuite) TestHandlesScheduledRepaymentRepoUpdateRollbackError() {
	testData := handlerModel.UpdateLoanAndScheduledRepaymentInput{
		LoanI:              &suite.LoanI,
		ScheduleRepayments: []*databaseModel.ScheduledRepayment{suite.ScheduledRepayment},
		Status:             suite.LoanStatus,
	}

	testError := errors.New("test error")
	rollbackError := errors.New("test error 2")
	suite.MockPostgresDBInstance.On("GetTransactionDb").Return(suite.MockTransactionDB, nil)
	suite.MockLoanRepo.On("Update", mock.Anything, mock.Anything).Return(nil)
	suite.MockScheduledRepaymentRepo.On("Update", mock.Anything, mock.Anything).Return(testError)
	suite.MockTransactionDB.On("Rollback").Return(rollbackError)
	suite.MockTransactionDB.AssertNumberOfCalls(suite.T(), "Commit", 0)

	admin := handler.NewAdmin(suite.MockLoanRepo, suite.MockScheduledRepaymentRepo, suite.MockPostgresDBInstance)
	result, err := admin.UpdateLoanAndScheduledRepayment(context.Background(), testData)
	assert.Nil(suite.T(), result)
	assert.EqualError(suite.T(), err, rollbackError.Error())
}

func (suite *TestAdminHandlerUpdateLoanAndScheduledRepaymentSuite) TestHandlesScheduledRepaymentRepoUpdateError() {
	testData := handlerModel.UpdateLoanAndScheduledRepaymentInput{
		LoanI:              &suite.LoanI,
		ScheduleRepayments: []*databaseModel.ScheduledRepayment{suite.ScheduledRepayment},
		Status:             suite.LoanStatus,
	}

	testError := errors.New("test error")
	suite.MockPostgresDBInstance.On("GetTransactionDb").Return(suite.MockTransactionDB, nil)
	suite.MockLoanRepo.On("Update", mock.Anything, mock.Anything).Return(nil)
	suite.MockScheduledRepaymentRepo.On("Update", mock.Anything, mock.Anything).Return(testError)
	suite.MockTransactionDB.On("Rollback").Return(nil)
	suite.MockTransactionDB.AssertNumberOfCalls(suite.T(), "Commit", 0)

	admin := handler.NewAdmin(suite.MockLoanRepo, suite.MockScheduledRepaymentRepo, suite.MockPostgresDBInstance)
	result, err := admin.UpdateLoanAndScheduledRepayment(context.Background(), testData)
	assert.Nil(suite.T(), result)
	assert.EqualError(suite.T(), err, testError.Error())
}

func (suite *TestAdminHandlerUpdateLoanAndScheduledRepaymentSuite) TestHandlesTransactionCommitError() {
	testData := handlerModel.UpdateLoanAndScheduledRepaymentInput{
		LoanI:              &suite.LoanI,
		ScheduleRepayments: []*databaseModel.ScheduledRepayment{suite.ScheduledRepayment},
		Status:             suite.LoanStatus,
	}

	testError := errors.New("test error")
	suite.MockPostgresDBInstance.On("GetTransactionDb").Return(suite.MockTransactionDB, nil)
	suite.MockLoanRepo.On("Update", mock.Anything, mock.Anything).Return(nil)
	suite.MockScheduledRepaymentRepo.On("Update", mock.Anything, mock.Anything).Return(nil)
	suite.MockTransactionDB.On("Commit").Return(testError)

	admin := handler.NewAdmin(suite.MockLoanRepo, suite.MockScheduledRepaymentRepo, suite.MockPostgresDBInstance)
	result, err := admin.UpdateLoanAndScheduledRepayment(context.Background(), testData)
	assert.Nil(suite.T(), result)
	assert.EqualError(suite.T(), err, testError.Error())
}

func (suite *TestAdminHandlerUpdateLoanAndScheduledRepaymentSuite) TestReturnsSuccessOnUpdate() {
	testData := handlerModel.UpdateLoanAndScheduledRepaymentInput{
		LoanI:              &suite.LoanI,
		ScheduleRepayments: []*databaseModel.ScheduledRepayment{suite.ScheduledRepayment},
		Status:             suite.LoanStatus,
	}

	suite.MockPostgresDBInstance.On("GetTransactionDb").Return(suite.MockTransactionDB, nil)
	suite.MockLoanRepo.On("Update", mock.Anything, mock.Anything).Return(nil)
	suite.MockScheduledRepaymentRepo.On("Update", mock.Anything, mock.Anything).Return(nil)
	suite.MockTransactionDB.On("Commit").Return(nil)

	admin := handler.NewAdmin(suite.MockLoanRepo, suite.MockScheduledRepaymentRepo, suite.MockPostgresDBInstance)
	result, err := admin.UpdateLoanAndScheduledRepayment(context.Background(), testData)
	assert.NotNil(suite.T(), result)
	assert.Equal(suite.T(), result.Success, true)
	assert.Nil(suite.T(), err)
}

func TestAdminHandler(t *testing.T) {
	logger.Init(false)
	suite.Run(t, new(TestAdminHandlerFetchLoansSuite))
	suite.Run(t, new(TestAdminHandlerUpdateLoanAndScheduledRepaymentSuite))
}
