package repo

import (
	"context"

	"loanManagement/constant"
	dbInstance "loanManagement/database/instance"
	databaseModel "loanManagement/database/model"
	"loanManagement/logger"
	repoModel "loanManagement/repo/model"
)

type ScheduledRepayment interface {
	FindAll(ctx context.Context, data repoModel.FindAllScheduledRepaymentInput) ([]*databaseModel.ScheduledRepayment, error)
	BulkCreate(ctx context.Context, data repoModel.BulkCreateScheduledRepaymentInput) error
	Update(ctx context.Context, data repoModel.UpdateScheduledRepaymentInput) error
}

type scheduledRepayment struct {
	dbInstance dbInstance.PostgresDB
}

func (repo *scheduledRepayment) FindAll(ctx context.Context, data repoModel.FindAllScheduledRepaymentInput) ([]*databaseModel.ScheduledRepayment, error) {
	var scheduledRepayments []*databaseModel.ScheduledRepayment

	err := repo.dbInstance.GetReadableDb().WithContext(ctx).Where(&databaseModel.ScheduledRepayment{
		LoanID: data.LoanID,
	}).Find(&scheduledRepayments).Error
	if err != nil {
		logger.Log.Error("failed to find scheduledRepayment", logger.Error(err))
		return nil, err
	}

	return scheduledRepayments, nil
}

func (repo *scheduledRepayment) BulkCreate(ctx context.Context, data repoModel.BulkCreateScheduledRepaymentInput) error {
	scheduledRepayments := make([]*databaseModel.ScheduledRepayment, len(data.ScheduledDates))
	amountInLowestCurrency := data.LoanAmount * 100
	totalDate := len(data.ScheduledDates)
	scheduledAmount := amountInLowestCurrency / int64(totalDate)
	diffAmount := amountInLowestCurrency - (scheduledAmount * int64(totalDate))

	for i := range data.ScheduledDates {
		repaymentAmount := scheduledAmount
		if i == 0 {
			repaymentAmount += diffAmount
		}
		scheduledDate := data.ScheduledDates[i]
		scheduledRepayments[i] = &databaseModel.ScheduledRepayment{
			LoanID:          data.LoanID,
			ScheduledAmount: repaymentAmount,
			PendingAmount:   repaymentAmount,
			Currency:        data.Currency,
			Status:          constant.SchedulePaymentStatusPending,
			ScheduledDate:   scheduledDate,
		}
	}

	db := repo.dbInstance.GetWritableDb()
	if data.TxDb != nil {
		db = (*data.TxDb).Get()
	}

	result := db.WithContext(ctx).Create(scheduledRepayments)
	if result.Error != nil {
		logger.Log.Error("unable to bulk create scheduledRepayment", logger.Error(result.Error))
		return result.Error
	}

	return nil
}

func (repo *scheduledRepayment) Update(ctx context.Context, data repoModel.UpdateScheduledRepaymentInput) error {
	updateModel := databaseModel.ScheduledRepayment{}

	if data.PendingAmount != nil {
		updateModel.PendingAmount = *data.PendingAmount
	}

	if data.Status != nil {
		updateModel.Status = *data.Status
	}

	db := repo.dbInstance.GetWritableDb()
	if data.TxDb != nil {
		db = (*data.TxDb).Get()
	}

	result := db.WithContext(ctx).Model(&databaseModel.ScheduledRepayment{
		BaseWithUpdatedAt: databaseModel.BaseWithUpdatedAt{
			ID: data.ID,
		},
	}).Updates(updateModel)
	if result.Error != nil {
		logger.Log.Error("unable to update scheduledRepayment", logger.Error(result.Error))
		return result.Error
	}

	return nil

}

func NewScheduledRepayment(
	dbInstance dbInstance.PostgresDB,
) ScheduledRepayment {
	return &scheduledRepayment{
		dbInstance: dbInstance,
	}
}
