package util

import (
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"

	"loanManagement/logger"
)

type Password interface {
	Hash(password string) (string, error)
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

func NewPassword() Password {
	return &password{}
}
