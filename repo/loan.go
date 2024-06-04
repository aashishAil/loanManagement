package repo

import (
	"context"

	"loanManagement/constant"
	dbInstance "loanManagement/database/instance"
	databaseModel "loanManagement/database/model"
	"loanManagement/logger"
	repoModel "loanManagement/repo/model"

	"github.com/pkg/errors"
)

type Loan interface {
	Create(ctx context.Context, data repoModel.CreateLoanInput) (*databaseModel.Loan, error)
	FindOne(ctx context.Context, data repoModel.FindOneLoanInput) (*databaseModel.Loan, error)
	FindAll(ctx context.Context, data repoModel.FindAllLoanInput) ([]*databaseModel.Loan, error)
	Update(ctx context.Context, data repoModel.UpdateLoanInput) error
}

type loan struct {
	dbInstance dbInstance.PostgresDB
}

func (repo *loan) Create(ctx context.Context, data repoModel.CreateLoanInput) (*databaseModel.Loan, error) {
	amountInLowestCurrency := data.Amount * constant.MinCurrencyConversionFactor
	loanI := databaseModel.Loan{
		UserID:            data.UserID,
		DisbursalAmount:   amountInLowestCurrency,
		PendingAmount:     amountInLowestCurrency,
		WeeklyInstallment: amountInLowestCurrency / data.Term,
		Currency:          data.Currency,
		Term:              data.Term,
		Status:            constant.LoanStatusPending,
		DisbursalDate:     data.DisbursalDate.UTC(), // for consistency all dates will be stored in UTC
	}

	db := repo.dbInstance.GetWritableDb()
	if data.TxDb != nil {
		db = (*data.TxDb).Get()
	}
	result := db.WithContext(ctx).Create(&loanI)
	if result.Error != nil {
		logger.Log.Error("unable to create loan", logger.Error(result.Error))
		return nil, errors.Wrap(result.Error, "unable to create loan")

	}

	return &loanI, nil
}

func (repo *loan) FindOne(ctx context.Context, data repoModel.FindOneLoanInput) (*databaseModel.Loan, error) {
	var loanI databaseModel.Loan

	queryModel := databaseModel.Loan{
		UserID: data.UserID,
	}

	if data.DisbursalDate != nil {
		queryModel.DisbursalDate = *data.DisbursalDate
	}

	result := repo.dbInstance.GetReadableDb().WithContext(ctx).Where(queryModel).First(&loanI)
	if result.Error != nil {
		logger.Log.Error("unable to find loan", logger.Error(result.Error))
		return nil, errors.Wrap(result.Error, "unable to find loan")
	}

	return &loanI, nil
}

func (repo *loan) FindAll(ctx context.Context, data repoModel.FindAllLoanInput) ([]*databaseModel.Loan, error) {
	var loans []*databaseModel.Loan

	result := repo.dbInstance.GetReadableDb().WithContext(ctx).Where(databaseModel.Loan{
		UserID: data.UserID,
	}).Find(&loans)
	if result.Error != nil {
		logger.Log.Error("unable to find loans: %v", logger.Error(result.Error))
		return nil, errors.Wrap(result.Error, "unable to find loans")
	}

	return loans, nil
}

func (repo *loan) Update(ctx context.Context, data repoModel.UpdateLoanInput) error {
	updatedModel := databaseModel.Loan{
		Status: data.Status,
	}

	db := repo.dbInstance.GetWritableDb()
	if data.TxDb != nil {
		db = (*data.TxDb).Get()
	}

	result := db.WithContext(ctx).Model(&databaseModel.Loan{
		BaseWithUpdatedAt: databaseModel.BaseWithUpdatedAt{
			ID: data.ID,
		},
	}).Updates(updatedModel)
	if result.Error != nil {
		logger.Log.Error("unable to update loan: %v", logger.Error(result.Error))
		return errors.Wrap(result.Error, "unable to update loan")
	}

	return nil
}

func NewLoan(
	dbInstance dbInstance.PostgresDB,
) Loan {
	return &loan{
		dbInstance: dbInstance,
	}
}
