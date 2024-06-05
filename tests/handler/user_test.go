package handler

import (
	"context"
	handlerModel "loanManagement/handler/model"
	"math/rand"
	"net/http"
	"testing"
	"time"

	"loanManagement/appError"
	"loanManagement/constant"
	databaseModel "loanManagement/database/model"
	"loanManagement/handler"
	"loanManagement/logger"
	databaseMocks "loanManagement/mocks/database/instance"
	repoMocks "loanManagement/mocks/repo"
	utilMocks "loanManagement/mocks/util"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type TestUserHandlerCheckValidCredentialsSuite struct {
	suite.Suite

	Email    string
	Password string
	Token    string
	UserI    *databaseModel.User

	MockLoanRepo               *repoMocks.Loan
	MockPaymentRepo            *repoMocks.Payment
	MockPostgresDBInstance     *databaseMocks.PostgresDB
	MockScheduledRepaymentRepo *repoMocks.ScheduledRepayment
	MockUserRepo               *repoMocks.User
	MockJwtUtil                *utilMocks.Jwt
	MockTimeUtil               *utilMocks.Time
}

func (suite *TestUserHandlerCheckValidCredentialsSuite) SetupTest() {
	suite.Email = "test@test.com"
	suite.Password = "test1234"
	suite.Token = "testToken"
	suite.UserI = &databaseModel.User{
		Base: databaseModel.Base{
			ID: uuid.New(),
		},
		Type: constant.UserTypeCustomer,
	}

	suite.MockLoanRepo = new(repoMocks.Loan)
	suite.MockPaymentRepo = new(repoMocks.Payment)
	suite.MockPostgresDBInstance = new(databaseMocks.PostgresDB)
	suite.MockScheduledRepaymentRepo = new(repoMocks.ScheduledRepayment)
	suite.MockUserRepo = new(repoMocks.User)
	suite.MockJwtUtil = new(utilMocks.Jwt)
	suite.MockTimeUtil = new(utilMocks.Time)
}

func (suite *TestUserHandlerCheckValidCredentialsSuite) TestCheckValidCredentialsUserNotFoundUnhandledError() {
	testError := errors.New("test error")
	suite.MockUserRepo.On("FindOne", mock.Anything, mock.Anything).Return(nil, testError)

	userHandler := handler.NewUser(suite.MockLoanRepo, suite.MockPaymentRepo, suite.MockScheduledRepaymentRepo, suite.MockUserRepo, suite.MockPostgresDBInstance, suite.MockJwtUtil, suite.MockTimeUtil)
	_, err := userHandler.CheckValidCredentials(context.Background(), suite.Email, suite.Password)

	assert.NotNil(suite.T(), err)
	customErr := appError.Custom{}
	ok := errors.As(err, &customErr)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), customErr.Code, http.StatusInternalServerError)
}

func (suite *TestUserHandlerCheckValidCredentialsSuite) TestCheckValidCredentialsUserNotFoundHandledError() {
	testError := appError.Custom{
		Err: errors.New("test error"),
	}
	suite.MockUserRepo.On("FindOne", mock.Anything, mock.Anything).Return(nil, testError)

	userHandler := handler.NewUser(suite.MockLoanRepo, suite.MockPaymentRepo, suite.MockScheduledRepaymentRepo, suite.MockUserRepo, suite.MockPostgresDBInstance, suite.MockJwtUtil, suite.MockTimeUtil)
	_, err := userHandler.CheckValidCredentials(context.Background(), suite.Email, suite.Password)

	assert.NotNil(suite.T(), err)
	customErr := appError.Custom{}
	ok := errors.As(err, &customErr)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), customErr.Code, http.StatusUnauthorized)
}

func (suite *TestUserHandlerCheckValidCredentialsSuite) TestCheckValidCredentialsUserNotFoundError() {
	suite.MockUserRepo.On("FindOne", mock.Anything, mock.Anything).Return(nil, nil)

	userHandler := handler.NewUser(suite.MockLoanRepo, suite.MockPaymentRepo, suite.MockScheduledRepaymentRepo, suite.MockUserRepo, suite.MockPostgresDBInstance, suite.MockJwtUtil, suite.MockTimeUtil)
	_, err := userHandler.CheckValidCredentials(context.Background(), suite.Email, suite.Password)

	assert.NotNil(suite.T(), err)
	customErr := appError.Custom{}
	ok := errors.As(err, &customErr)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), customErr.Code, http.StatusNotFound)
}

func (suite *TestUserHandlerCheckValidCredentialsSuite) TestCheckValidCredentialsTokenGenerationError() {
	suite.MockUserRepo.On("FindOne", mock.Anything, mock.Anything).Return(suite.UserI, nil)
	suite.MockJwtUtil.On("GenerateToken", mock.Anything).Return("", errors.New("failed to generate token"))

	userHandler := handler.NewUser(suite.MockLoanRepo, suite.MockPaymentRepo, suite.MockScheduledRepaymentRepo, suite.MockUserRepo, suite.MockPostgresDBInstance, suite.MockJwtUtil, suite.MockTimeUtil)
	_, err := userHandler.CheckValidCredentials(context.Background(), suite.Email, suite.Password)

	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), "failed to generate token", err.Error())
}

func (suite *TestUserHandlerCheckValidCredentialsSuite) TestCheckValidCredentialsSuccess() {
	suite.MockUserRepo.On("FindOne", mock.Anything, mock.Anything).Return(suite.UserI, nil)
	suite.MockJwtUtil.On("GenerateToken", mock.Anything).Return(suite.Token, nil)

	userHandler := handler.NewUser(suite.MockLoanRepo, suite.MockPaymentRepo, suite.MockScheduledRepaymentRepo, suite.MockUserRepo, suite.MockPostgresDBInstance, suite.MockJwtUtil, suite.MockTimeUtil)
	result, err := userHandler.CheckValidCredentials(context.Background(), suite.Email, suite.Password)

	assert.NotNil(suite.T(), result)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), suite.Token, result)
}

type TestUserHandlerCreateLoanSuite struct {
	suite.Suite

	UserID        uuid.UUID
	Amount        int64
	Currency      constant.Currency
	Term          int64
	DisbursalDate time.Time
	LoanI         *databaseModel.Loan

	MockLoanRepo               *repoMocks.Loan
	MockPaymentRepo            *repoMocks.Payment
	MockPostgresDBInstance     *databaseMocks.PostgresDB
	MockScheduledRepaymentRepo *repoMocks.ScheduledRepayment
	MockTransactionDB          *databaseMocks.PostgresTransactionDB

	MockUserRepo *repoMocks.User
	MockJwtUtil  *utilMocks.Jwt
	MockTimeUtil *utilMocks.Time
}

func (suite *TestUserHandlerCreateLoanSuite) SetupTest() {
	suite.MockLoanRepo = new(repoMocks.Loan)
	suite.MockPaymentRepo = new(repoMocks.Payment)
	suite.MockPostgresDBInstance = new(databaseMocks.PostgresDB)
	suite.MockScheduledRepaymentRepo = new(repoMocks.ScheduledRepayment)
	suite.MockTransactionDB = new(databaseMocks.PostgresTransactionDB)
	suite.MockUserRepo = new(repoMocks.User)
	suite.MockJwtUtil = new(utilMocks.Jwt)
	suite.MockTimeUtil = new(utilMocks.Time)

	suite.UserID = uuid.New()
	suite.LoanI = &databaseModel.Loan{
		BaseWithUpdatedAt: databaseModel.BaseWithUpdatedAt{
			ID: uuid.New(),
		},
	}
	suite.Amount = constant.MinDisbursalAmount + rand.Int63()
	if suite.Amount > constant.MaxDisbursalAmount {
		suite.Amount = constant.MaxDisbursalAmount
	}
	suite.Currency = constant.CurrencyINR
	suite.Term = constant.MinLoanRepaymentTenure + rand.Int63()
	if suite.Term > constant.MaxLoanRepaymentTenure {
		suite.Term = constant.MaxLoanRepaymentTenure
	}
	suite.DisbursalDate = time.Now().AddDate(0, 1, 0)
}

func (suite *TestUserHandlerCreateLoanSuite) TestDisbursalDateLowerThanLimit() {
	suite.DisbursalDate = time.Now()
	suite.MockTimeUtil.On("GetCurrent").Return(time.Now())

	userHandler := handler.NewUser(suite.MockLoanRepo, suite.MockPaymentRepo, suite.MockScheduledRepaymentRepo, suite.MockUserRepo, suite.MockPostgresDBInstance, suite.MockJwtUtil, suite.MockTimeUtil)
	loanID, err := userHandler.CreateLoan(context.Background(), handlerModel.CreateUserLoanInput{
		UserID:        suite.UserID,
		Amount:        suite.Amount,
		Currency:      suite.Currency,
		Term:          suite.Term,
		DisbursalDate: suite.DisbursalDate,
	})

	assert.Equal(suite.T(), loanID, uuid.Nil)
	assert.NotNil(suite.T(), err)
	customErr := appError.Custom{}
	ok := errors.As(err, &customErr)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), customErr.Code, http.StatusBadRequest)
	assert.Equal(suite.T(), customErr.Error(), constant.InvalidDisbursalDateGapError.Error())
}

func (suite *TestUserHandlerCreateLoanSuite) TestTermLowerThanLimit() {
	suite.Term = constant.MinLoanRepaymentTenure - 1
	suite.MockTimeUtil.On("GetCurrent").Return(time.Now())

	userHandler := handler.NewUser(suite.MockLoanRepo, suite.MockPaymentRepo, suite.MockScheduledRepaymentRepo, suite.MockUserRepo, suite.MockPostgresDBInstance, suite.MockJwtUtil, suite.MockTimeUtil)
	loanID, err := userHandler.CreateLoan(context.Background(), handlerModel.CreateUserLoanInput{
		UserID:        suite.UserID,
		Amount:        suite.Amount,
		Currency:      suite.Currency,
		Term:          suite.Term,
		DisbursalDate: suite.DisbursalDate,
	})

	assert.Equal(suite.T(), loanID, uuid.Nil)
	assert.NotNil(suite.T(), err)
	customErr := appError.Custom{}
	ok := errors.As(err, &customErr)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), customErr.Code, http.StatusBadRequest)
	assert.Equal(suite.T(), customErr.Error(), constant.MinRepaymentTenureError.Error())
}

func (suite *TestUserHandlerCreateLoanSuite) TestTermHigherThanLimit() {
	suite.Term = constant.MaxLoanRepaymentTenure + 1
	suite.MockTimeUtil.On("GetCurrent").Return(time.Now())

	userHandler := handler.NewUser(suite.MockLoanRepo, suite.MockPaymentRepo, suite.MockScheduledRepaymentRepo, suite.MockUserRepo, suite.MockPostgresDBInstance, suite.MockJwtUtil, suite.MockTimeUtil)
	loanID, err := userHandler.CreateLoan(context.Background(), handlerModel.CreateUserLoanInput{
		UserID:        suite.UserID,
		Amount:        suite.Amount,
		Currency:      suite.Currency,
		Term:          suite.Term,
		DisbursalDate: suite.DisbursalDate,
	})

	assert.Equal(suite.T(), loanID, uuid.Nil)
	assert.NotNil(suite.T(), err)
	customErr := appError.Custom{}
	ok := errors.As(err, &customErr)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), customErr.Code, http.StatusBadRequest)
	assert.Equal(suite.T(), customErr.Error(), constant.MaxRepaymentTenureError.Error())
}

func (suite *TestUserHandlerCreateLoanSuite) TestDisbursalAmountLowerThanLimit() {
	suite.Amount = constant.MinDisbursalAmount - 1
	suite.MockTimeUtil.On("GetCurrent").Return(time.Now())

	userHandler := handler.NewUser(suite.MockLoanRepo, suite.MockPaymentRepo, suite.MockScheduledRepaymentRepo, suite.MockUserRepo, suite.MockPostgresDBInstance, suite.MockJwtUtil, suite.MockTimeUtil)
	loanID, err := userHandler.CreateLoan(context.Background(), handlerModel.CreateUserLoanInput{
		UserID:        suite.UserID,
		Amount:        suite.Amount,
		Currency:      suite.Currency,
		Term:          suite.Term,
		DisbursalDate: suite.DisbursalDate,
	})

	assert.Equal(suite.T(), loanID, uuid.Nil)
	assert.NotNil(suite.T(), err)
	customErr := appError.Custom{}
	ok := errors.As(err, &customErr)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), customErr.Code, http.StatusBadRequest)
	assert.Equal(suite.T(), customErr.Error(), constant.InvalidMinAmountError.Error())
}

func (suite *TestUserHandlerCreateLoanSuite) TestDisbursalAmountHigherThanLimit() {
	suite.Amount = constant.MaxDisbursalAmount + 1
	suite.MockTimeUtil.On("GetCurrent").Return(time.Now())

	userHandler := handler.NewUser(suite.MockLoanRepo, suite.MockPaymentRepo, suite.MockScheduledRepaymentRepo, suite.MockUserRepo, suite.MockPostgresDBInstance, suite.MockJwtUtil, suite.MockTimeUtil)
	loanID, err := userHandler.CreateLoan(context.Background(), handlerModel.CreateUserLoanInput{
		UserID:        suite.UserID,
		Amount:        suite.Amount,
		Currency:      suite.Currency,
		Term:          suite.Term,
		DisbursalDate: suite.DisbursalDate,
	})

	assert.Equal(suite.T(), loanID, uuid.Nil)
	assert.NotNil(suite.T(), err)
	customErr := appError.Custom{}
	ok := errors.As(err, &customErr)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), customErr.Code, http.StatusBadRequest)
	assert.Equal(suite.T(), customErr.Error(), constant.InvalidMaxAmountError.Error())
}

func (suite *TestUserHandlerCreateLoanSuite) TestInvalidCurrency() {
	suite.Currency = ""
	suite.MockTimeUtil.On("GetCurrent").Return(time.Now())

	userHandler := handler.NewUser(suite.MockLoanRepo, suite.MockPaymentRepo, suite.MockScheduledRepaymentRepo, suite.MockUserRepo, suite.MockPostgresDBInstance, suite.MockJwtUtil, suite.MockTimeUtil)
	loanID, err := userHandler.CreateLoan(context.Background(), handlerModel.CreateUserLoanInput{
		UserID:        suite.UserID,
		Amount:        suite.Amount,
		Currency:      suite.Currency,
		Term:          suite.Term,
		DisbursalDate: suite.DisbursalDate,
	})

	assert.Equal(suite.T(), loanID, uuid.Nil)
	assert.NotNil(suite.T(), err)
	customErr := appError.Custom{}
	ok := errors.As(err, &customErr)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), customErr.Code, http.StatusBadRequest)
	assert.Equal(suite.T(), customErr.Error(), constant.InvalidCurrencyError.Error())
}

func (suite *TestUserHandlerCreateLoanSuite) TestTransactionCreateError() {
	testError := errors.New("test error")
	suite.MockTimeUtil.On("GetCurrent").Return(time.Now())
	suite.MockPostgresDBInstance.On("GetTransactionDb").Return(nil, testError)

	userHandler := handler.NewUser(suite.MockLoanRepo, suite.MockPaymentRepo, suite.MockScheduledRepaymentRepo, suite.MockUserRepo, suite.MockPostgresDBInstance, suite.MockJwtUtil, suite.MockTimeUtil)
	loanID, err := userHandler.CreateLoan(context.Background(), handlerModel.CreateUserLoanInput{
		UserID:        suite.UserID,
		Amount:        suite.Amount,
		Currency:      suite.Currency,
		Term:          suite.Term,
		DisbursalDate: suite.DisbursalDate,
	})

	assert.Equal(suite.T(), loanID, uuid.Nil)
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), err.Error(), testError.Error())
}

func (suite *TestUserHandlerCreateLoanSuite) TestLoanCreateAndRollbackError() {
	testError := errors.New("test error")
	rollbackError := errors.New("rollback error")
	suite.MockTimeUtil.On("GetCurrent").Return(time.Now())
	suite.MockPostgresDBInstance.On("GetTransactionDb").Return(suite.MockTransactionDB, nil)
	suite.MockLoanRepo.On("Create", mock.Anything, mock.Anything, mock.Anything).Return(nil, testError)
	suite.MockTransactionDB.On("Rollback").Return(rollbackError)

	userHandler := handler.NewUser(suite.MockLoanRepo, suite.MockPaymentRepo, suite.MockScheduledRepaymentRepo, suite.MockUserRepo, suite.MockPostgresDBInstance, suite.MockJwtUtil, suite.MockTimeUtil)
	loanID, err := userHandler.CreateLoan(context.Background(), handlerModel.CreateUserLoanInput{
		UserID:        suite.UserID,
		Amount:        suite.Amount,
		Currency:      suite.Currency,
		Term:          suite.Term,
		DisbursalDate: suite.DisbursalDate,
	})

	assert.Equal(suite.T(), loanID, uuid.Nil)
	assert.NotNil(suite.T(), err)
}

func (suite *TestUserHandlerCreateLoanSuite) TestLoanCreateError() {
	testError := errors.New("test error")
	suite.MockTimeUtil.On("GetCurrent").Return(time.Now())
	suite.MockPostgresDBInstance.On("GetTransactionDb").Return(suite.MockTransactionDB, nil)
	suite.MockLoanRepo.On("Create", mock.Anything, mock.Anything, mock.Anything).Return(nil, testError)
	suite.MockTransactionDB.On("Rollback").Return(nil)

	userHandler := handler.NewUser(suite.MockLoanRepo, suite.MockPaymentRepo, suite.MockScheduledRepaymentRepo, suite.MockUserRepo, suite.MockPostgresDBInstance, suite.MockJwtUtil, suite.MockTimeUtil)
	loanID, err := userHandler.CreateLoan(context.Background(), handlerModel.CreateUserLoanInput{
		UserID:        suite.UserID,
		Amount:        suite.Amount,
		Currency:      suite.Currency,
		Term:          suite.Term,
		DisbursalDate: suite.DisbursalDate,
	})

	assert.Equal(suite.T(), loanID, uuid.Nil)
	assert.NotNil(suite.T(), err)
}

func (suite *TestUserHandlerCreateLoanSuite) TestScheduledRepaymentCreateAndRollbackError() {
	testError := errors.New("test error")
	suite.MockTimeUtil.On("GetCurrent").Return(time.Now())
	suite.MockPostgresDBInstance.On("GetTransactionDb").Return(suite.MockTransactionDB, nil)
	suite.MockLoanRepo.On("Create", mock.Anything, mock.Anything, mock.Anything).Return(suite.LoanI, nil)
	suite.MockScheduledRepaymentRepo.On("BulkCreate", mock.Anything, mock.Anything).Return(testError)
	suite.MockTransactionDB.On("Rollback").Return(testError)

	userHandler := handler.NewUser(suite.MockLoanRepo, suite.MockPaymentRepo, suite.MockScheduledRepaymentRepo, suite.MockUserRepo, suite.MockPostgresDBInstance, suite.MockJwtUtil, suite.MockTimeUtil)
	loanID, err := userHandler.CreateLoan(context.Background(), handlerModel.CreateUserLoanInput{
		UserID:        suite.UserID,
		Amount:        suite.Amount,
		Currency:      suite.Currency,
		Term:          suite.Term,
		DisbursalDate: suite.DisbursalDate,
	})

	assert.Equal(suite.T(), loanID, uuid.Nil)
	assert.NotNil(suite.T(), err)
}

func (suite *TestUserHandlerCreateLoanSuite) TestScheduledRepaymentCreateError() {
	testError := errors.New("test error")
	suite.MockTimeUtil.On("GetCurrent").Return(time.Now())
	suite.MockPostgresDBInstance.On("GetTransactionDb").Return(suite.MockTransactionDB, nil)
	suite.MockLoanRepo.On("Create", mock.Anything, mock.Anything, mock.Anything).Return(suite.LoanI, nil)
	suite.MockScheduledRepaymentRepo.On("BulkCreate", mock.Anything, mock.Anything).Return(testError)
	suite.MockTransactionDB.On("Rollback").Return(nil)

	userHandler := handler.NewUser(suite.MockLoanRepo, suite.MockPaymentRepo, suite.MockScheduledRepaymentRepo, suite.MockUserRepo, suite.MockPostgresDBInstance, suite.MockJwtUtil, suite.MockTimeUtil)
	loanID, err := userHandler.CreateLoan(context.Background(), handlerModel.CreateUserLoanInput{
		UserID:        suite.UserID,
		Amount:        suite.Amount,
		Currency:      suite.Currency,
		Term:          suite.Term,
		DisbursalDate: suite.DisbursalDate,
	})

	assert.Equal(suite.T(), loanID, uuid.Nil)
	assert.NotNil(suite.T(), err)
}

func (suite *TestUserHandlerCreateLoanSuite) TestTransactionCommitError() {
	testError := errors.New("test error")
	suite.MockTimeUtil.On("GetCurrent").Return(time.Now())
	suite.MockPostgresDBInstance.On("GetTransactionDb").Return(suite.MockTransactionDB, nil)
	suite.MockLoanRepo.On("Create", mock.Anything, mock.Anything, mock.Anything).Return(suite.LoanI, nil)
	suite.MockScheduledRepaymentRepo.On("BulkCreate", mock.Anything, mock.Anything).Return(nil)
	suite.MockTransactionDB.On("Commit").Return(testError)

	userHandler := handler.NewUser(suite.MockLoanRepo, suite.MockPaymentRepo, suite.MockScheduledRepaymentRepo, suite.MockUserRepo, suite.MockPostgresDBInstance, suite.MockJwtUtil, suite.MockTimeUtil)
	loanID, err := userHandler.CreateLoan(context.Background(), handlerModel.CreateUserLoanInput{
		UserID:        suite.UserID,
		Amount:        suite.Amount,
		Currency:      suite.Currency,
		Term:          suite.Term,
		DisbursalDate: suite.DisbursalDate,
	})

	assert.Equal(suite.T(), loanID, uuid.Nil)
	assert.NotNil(suite.T(), err)
}

func (suite *TestUserHandlerCreateLoanSuite) TestSuccess() {
	suite.MockTimeUtil.On("GetCurrent").Return(time.Now())
	suite.MockPostgresDBInstance.On("GetTransactionDb").Return(suite.MockTransactionDB, nil)
	suite.MockLoanRepo.On("Create", mock.Anything, mock.Anything, mock.Anything).Return(suite.LoanI, nil)
	suite.MockScheduledRepaymentRepo.On("BulkCreate", mock.Anything, mock.Anything).Return(nil)

	suite.MockTransactionDB.AssertNumberOfCalls(suite.T(), "Rollback", 0)
	suite.MockTransactionDB.On("Commit").Return(nil)

	userHandler := handler.NewUser(suite.MockLoanRepo, suite.MockPaymentRepo, suite.MockScheduledRepaymentRepo, suite.MockUserRepo, suite.MockPostgresDBInstance, suite.MockJwtUtil, suite.MockTimeUtil)
	loanID, err := userHandler.CreateLoan(context.Background(), handlerModel.CreateUserLoanInput{
		UserID:        suite.UserID,
		Amount:        suite.Amount,
		Currency:      suite.Currency,
		Term:          suite.Term,
		DisbursalDate: suite.DisbursalDate,
	})

	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), loanID)
	assert.Equal(suite.T(), suite.LoanI.ID, loanID)
}

func TestUserHandler(t *testing.T) {
	logger.Init(false)
	suite.Run(t, new(TestUserHandlerCheckValidCredentialsSuite))
	suite.Run(t, new(TestUserHandlerCreateLoanSuite))
}
