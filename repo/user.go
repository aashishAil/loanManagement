package repo

import (
	"context"

	dbInstance "loanManagement/database/instance"
	databaseModel "loanManagement/database/model"
	"loanManagement/logger"
	repoModel "loanManagement/repo/model"
	"loanManagement/util"

	"github.com/pkg/errors"
)

type User interface {
	FindOne(ctx context.Context, data repoModel.FindOneUserInput) (*databaseModel.User, error)
}

type user struct {
	dbInstance dbInstance.PostgresDB

	passwordUtil util.Password
}

func (repo *user) FindOne(ctx context.Context, data repoModel.FindOneUserInput) (*databaseModel.User, error) {
	var userI databaseModel.User

	encryptedPassword, err := repo.passwordUtil.Hash(data.Password)
	if err != nil {
		logger.Log.Error("failed to hash password", logger.Error(err))
		return nil, errors.Wrap(err, "failed to hash password")
	}

	err = repo.dbInstance.GetReadableDb().Where(&databaseModel.User{
		Email:             data.Email,
		EncryptedPassword: encryptedPassword,
	}).First(&userI).Error
	if err != nil {
		if errors.Is(err, dbInstance.ErrNoRecordFound) {
			logger.Log.Info("user not found", logger.String("email", data.Email))
			return nil, nil
		}
		logger.Log.Error("failed to find user", logger.Error(err))
		return nil, errors.Wrap(err, "failed to find user")
	}
	logger.Log.Info("user found", logger.String("email", data.Email))

	return &userI, nil
}

func NewUser(
	dbInstance dbInstance.PostgresDB,

	passwordUtil util.Password,
) User {
	return &user{
		dbInstance: dbInstance,

		passwordUtil: passwordUtil,
	}
}
