package handler

import (
	"github.com/google/uuid"
	"loanManagement/appError"
	"loanManagement/logger"
	"loanManagement/repo"
	"loanManagement/util"
	"net/http"

	"github.com/pkg/errors"
)

type User interface {
	CheckValidCredentials(email, password string) (string, error)
}

type user struct {
	userRepo repo.User

	jwtUtil      util.Jwt
	passwordUtil util.Password
}

func (h *user) CheckValidCredentials(email, password string) (string, error) {
	encryptedPassword, err := h.passwordUtil.Hash(password)
	if err != nil {
		logger.Log.Error("failed to hash password")
		return "", appError.Custom{
			Err:  errors.Wrap(err, "failed to hash password"),
			Code: http.StatusInternalServerError,
		}
	}

	userI, err := h.userRepo.FindOne(email, encryptedPassword)
	if err != nil {
		logger.Log.Error("failed to find user")
		return "", appError.Custom{
			Err:  errors.Wrap(err, "failed to find user"),
			Code: http.StatusInternalServerError,
		}
	}

	if userI.ID == uuid.Nil {
		logger.Log.Info("user not found", logger.String("email", email))
		return "", appError.Custom{
			Err:  errors.New("user not found"),
			Code: http.StatusNotFound,
		}
	}

	token, err := h.jwtUtil.GenerateToken(userI)
	if err != nil {
		logger.Log.Error("failed to generate token")
		return "", appError.Custom{
			Err:  errors.Wrap(err, "failed to generate token"),
			Code: http.StatusInternalServerError,
		}
	}

	return token, nil
}

func NewUser(
	userRepo repo.User,

	jwtUtil util.Jwt,
	passwordUtil util.Password,
) User {
	return &user{
		userRepo: userRepo,

		jwtUtil:      jwtUtil,
		passwordUtil: passwordUtil,
	}
}
