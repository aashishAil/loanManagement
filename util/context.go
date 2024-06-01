package util

import (
	goContext "context"

	"github.com/gin-gonic/gin"
)

type Context interface {
	CreateContextFromGinContext(gCtx *gin.Context) goContext.Context
}

type context struct {
}

func (util *context) CreateContextFromGinContext(gCtx *gin.Context) goContext.Context {
	return gCtx.Request.Context()
}

func NewContext() Context {
	return &context{}
}
