package handler

import (
	"context"
	"net/http"

	"loanManagement/appError"
	"loanManagement/logger"
	"loanManagement/repo"
	repoModel "loanManagement/repo/model"
	"loanManagement/util"

	"github.com/pkg/errors"
)

type User interface {
	CheckValidCredentials(ctx context.Context, email, password string) (string, error)
}

type user struct {
	userRepo repo.User

	jwtUtil util.Jwt
}

func (h *user) CheckValidCredentials(ctx context.Context, email, password string) (string, error) {
	userI, err := h.userRepo.FindOne(ctx, repoModel.FindOneUserInput{
		Email:    email,
		Password: password,
	})
	if err != nil {
		logger.Log.Error("failed to find user", logger.Error(err))
		return "", appError.Custom{
			Err:  errors.Wrap(err, "failed to find user"),
			Code: http.StatusInternalServerError,
		}
	}

	if userI == nil {
		logger.Log.Info("user not found", logger.String("email", email))
		return "", appError.Custom{
			Err:  errors.New("user not found"),
			Code: http.StatusNotFound,
		}
	}

	token, err := h.jwtUtil.GenerateToken(*userI)
	if err != nil {
		logger.Log.Error("failed to generate token", logger.Error(err))
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
) User {
	return &user{
		userRepo: userRepo,

		jwtUtil: jwtUtil,
	}
}
