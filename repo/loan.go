package repo

import (
	"context"

	"loanManagement/appError"
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
		UserID:          data.UserID,
		DisbursalAmount: amountInLowestCurrency,
		PendingAmount:   amountInLowestCurrency,
		Currency:        data.Currency,
		Term:            data.Term,
		Status:          constant.LoanStatusPending,
		DisbursalDate:   data.DisbursalDate.UTC(), // for consistency all dates will be stored in UTC
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

	if data.ID == nil && data.UserID == nil {
		return nil, appError.Custom{
			Err: errors.New("loanID or userID is required"),
		}
	}

	queryModel := databaseModel.Loan{}

	if data.ID != nil {
		queryModel.ID = *data.ID
	}

	if data.UserID != nil {
		queryModel.UserID = *data.UserID
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

	if data.UserID == nil && data.Status == nil && len(data.IDs) == 0 {
		return nil, appError.Custom{
			Err: errors.New("userID, status or loanIDs is required"),
		}
	}

	db := repo.dbInstance.GetReadableDb().WithContext(ctx)
	if data.UserID != nil {
		db = db.Where("user_id = ?", data.UserID)
	}

	if data.Status != nil {
		db = db.Where("status = ?", data.Status)
	}

	if len(data.IDs) > 0 {
		if len(data.IDs) == 1 {
			db = db.Where("id = ?", data.IDs[0])
		} else {
			db = db.Where("id IN (?)", data.IDs)
		}
	}

	result := db.Find(&loans)
	if result.Error != nil {
		logger.Log.Error("unable to find loans: %v", logger.Error(result.Error))
		return nil, errors.Wrap(result.Error, "unable to find loans")
	}

	return loans, nil
}

func (repo *loan) Update(ctx context.Context, data repoModel.UpdateLoanInput) error {
	updatedModel := databaseModel.Loan{}

	if data.Status != nil {
		updatedModel.Status = *data.Status
	}

	if data.PendingAmount != nil {
		updatedModel.PendingAmount = *data.PendingAmount
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
