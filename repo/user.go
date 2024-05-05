package repo

import (
	dbInstance "loanManagement/database/instance"
	databaseModel "loanManagement/database/model"
	"loanManagement/logger"

	"github.com/pkg/errors"
)

type User interface {
	FindOne(email, encryptedPassword string) (databaseModel.User, error)
}

type user struct {
	dbInstance dbInstance.PostgresDB
}

func (u *user) FindOne(email, encryptedPassword string) (databaseModel.User, error) {
	var userI databaseModel.User

	err := u.dbInstance.GetReadableDb().Where(&databaseModel.User{
		Email:             email,
		EncryptedPassword: encryptedPassword,
	}).First(&userI).Error
	if err != nil {
		if errors.Is(err, dbInstance.ErrNoRecordFound) {
			logger.Log.Info("user not found", logger.String("email", email))
			return databaseModel.User{}, nil
		}
		logger.Log.Errorf("failed to find user: %v", err)
		return databaseModel.User{}, errors.Wrap(err, "failed to find user")
	}
	logger.Log.Info("user found", logger.String("email", email))

	return userI, nil
}

func NewUser(
	dbInstance dbInstance.PostgresDB,
) User {
	return &user{
		dbInstance: dbInstance,
	}
}
