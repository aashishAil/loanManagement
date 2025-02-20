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

type Payment interface {
	Create(ctx context.Context, data repoModel.CreatePaymentInput) (*databaseModel.Payment, error)
}

type payment struct {
	dbInstance dbInstance.PostgresDB
}

func (repo *payment) Create(ctx context.Context, data repoModel.CreatePaymentInput) (*databaseModel.Payment, error) {
	amountInLowestCurrency := int64(data.Amount * constant.MinCurrencyConversionFactor)
	paymentI := databaseModel.Payment{
		UserID:   data.UserID,
		LoanID:   data.LoanID,
		Amount:   amountInLowestCurrency,
		Currency: data.Currency,
	}

	db := repo.dbInstance.GetWritableDb()
	if data.TxDb != nil {
		db = (*data.TxDb).Get()
	}
	result := db.WithContext(ctx).Create(&paymentI)
	if result.Error != nil {
		logger.Log.Error("unable to create payment", logger.Error(result.Error))
		return nil, errors.Wrap(result.Error, "unable to create payment")

	}

	return &paymentI, nil
}

func NewPayment(
	dbInstance dbInstance.PostgresDB,
) Payment {
	return &payment{
		dbInstance: dbInstance,
	}
}
