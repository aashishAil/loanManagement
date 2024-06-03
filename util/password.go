package util

import (
	"golang.org/x/crypto/bcrypt"

	"loanManagement/logger"

	"github.com/pkg/errors"
)

type Password interface {
	Hash(password string) (string, error)
	Compare(hashedPassword, password string) error
}

type password struct{}

func (util *password) Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		logger.Log.Error(err.Error())
		return "", errors.Wrap(err, "failed to hash password")
	}

	return string(bytes), nil
}

func (util *password) Compare(hashedPassword, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		logger.Log.Info("failed to compare password", logger.String("error", err.Error()))
		return errors.Wrap(err, "failed to compare password")
	}

	return nil
}

func NewPassword() Password {
	return &password{}
}
