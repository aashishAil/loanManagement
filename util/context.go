package util

import (
	goContext "context"

	"loanManagement/constant"
	"loanManagement/model"

	"github.com/gin-gonic/gin"
)

type Context interface {
	CreateContextFromGinContext(gCtx *gin.Context) goContext.Context
	StoreLoggedInUser(ctx goContext.Context, user model.LoggedInUser) goContext.Context
	GetLoggedInUser(ctx goContext.Context) *model.LoggedInUser
}

type context struct {
}

func (util *context) CreateContextFromGinContext(gCtx *gin.Context) goContext.Context {
	return gCtx.Request.Context()
}

func (util *context) StoreLoggedInUser(ctx goContext.Context, user model.LoggedInUser) goContext.Context {
	return goContext.WithValue(ctx, constant.LoggedInUserContextKey, user)
}

func (util *context) GetLoggedInUser(ctx goContext.Context) *model.LoggedInUser {
	user, convertable := ctx.Value(constant.LoggedInUserContextKey).(model.LoggedInUser)
	if !convertable {
		return nil
	}

	return &user
}

func NewContext() Context {
	return &context{}
}
