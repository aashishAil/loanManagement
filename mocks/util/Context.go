// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	gin "github.com/gin-gonic/gin"
	mock "github.com/stretchr/testify/mock"

	model "loanManagement/model"
)

// Context is an autogenerated mock type for the Context type
type Context struct {
	mock.Mock
}

type Context_Expecter struct {
	mock *mock.Mock
}

func (_m *Context) EXPECT() *Context_Expecter {
	return &Context_Expecter{mock: &_m.Mock}
}

// CreateContextFromGinContext provides a mock function with given fields: gCtx
func (_m *Context) CreateContextFromGinContext(gCtx *gin.Context) context.Context {
	ret := _m.Called(gCtx)

	if len(ret) == 0 {
		panic("no return value specified for CreateContextFromGinContext")
	}

	var r0 context.Context
	if rf, ok := ret.Get(0).(func(*gin.Context) context.Context); ok {
		r0 = rf(gCtx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(context.Context)
		}
	}

	return r0
}

// Context_CreateContextFromGinContext_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateContextFromGinContext'
type Context_CreateContextFromGinContext_Call struct {
	*mock.Call
}

// CreateContextFromGinContext is a helper method to define mock.On call
//   - gCtx *gin.Context
func (_e *Context_Expecter) CreateContextFromGinContext(gCtx interface{}) *Context_CreateContextFromGinContext_Call {
	return &Context_CreateContextFromGinContext_Call{Call: _e.mock.On("CreateContextFromGinContext", gCtx)}
}

func (_c *Context_CreateContextFromGinContext_Call) Run(run func(gCtx *gin.Context)) *Context_CreateContextFromGinContext_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gin.Context))
	})
	return _c
}

func (_c *Context_CreateContextFromGinContext_Call) Return(_a0 context.Context) *Context_CreateContextFromGinContext_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Context_CreateContextFromGinContext_Call) RunAndReturn(run func(*gin.Context) context.Context) *Context_CreateContextFromGinContext_Call {
	_c.Call.Return(run)
	return _c
}

// GetLoggedInUser provides a mock function with given fields: ctx
func (_m *Context) GetLoggedInUser(ctx context.Context) *model.LoggedInUser {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetLoggedInUser")
	}

	var r0 *model.LoggedInUser
	if rf, ok := ret.Get(0).(func(context.Context) *model.LoggedInUser); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.LoggedInUser)
		}
	}

	return r0
}

// Context_GetLoggedInUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLoggedInUser'
type Context_GetLoggedInUser_Call struct {
	*mock.Call
}

// GetLoggedInUser is a helper method to define mock.On call
//   - ctx context.Context
func (_e *Context_Expecter) GetLoggedInUser(ctx interface{}) *Context_GetLoggedInUser_Call {
	return &Context_GetLoggedInUser_Call{Call: _e.mock.On("GetLoggedInUser", ctx)}
}

func (_c *Context_GetLoggedInUser_Call) Run(run func(ctx context.Context)) *Context_GetLoggedInUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *Context_GetLoggedInUser_Call) Return(_a0 *model.LoggedInUser) *Context_GetLoggedInUser_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Context_GetLoggedInUser_Call) RunAndReturn(run func(context.Context) *model.LoggedInUser) *Context_GetLoggedInUser_Call {
	_c.Call.Return(run)
	return _c
}

// StoreLoggedInUser provides a mock function with given fields: ctx, user
func (_m *Context) StoreLoggedInUser(ctx context.Context, user model.LoggedInUser) context.Context {
	ret := _m.Called(ctx, user)

	if len(ret) == 0 {
		panic("no return value specified for StoreLoggedInUser")
	}

	var r0 context.Context
	if rf, ok := ret.Get(0).(func(context.Context, model.LoggedInUser) context.Context); ok {
		r0 = rf(ctx, user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(context.Context)
		}
	}

	return r0
}

// Context_StoreLoggedInUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'StoreLoggedInUser'
type Context_StoreLoggedInUser_Call struct {
	*mock.Call
}

// StoreLoggedInUser is a helper method to define mock.On call
//   - ctx context.Context
//   - user model.LoggedInUser
func (_e *Context_Expecter) StoreLoggedInUser(ctx interface{}, user interface{}) *Context_StoreLoggedInUser_Call {
	return &Context_StoreLoggedInUser_Call{Call: _e.mock.On("StoreLoggedInUser", ctx, user)}
}

func (_c *Context_StoreLoggedInUser_Call) Run(run func(ctx context.Context, user model.LoggedInUser)) *Context_StoreLoggedInUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.LoggedInUser))
	})
	return _c
}

func (_c *Context_StoreLoggedInUser_Call) Return(_a0 context.Context) *Context_StoreLoggedInUser_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Context_StoreLoggedInUser_Call) RunAndReturn(run func(context.Context, model.LoggedInUser) context.Context) *Context_StoreLoggedInUser_Call {
	_c.Call.Return(run)
	return _c
}

// NewContext creates a new instance of Context. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewContext(t interface {
	mock.TestingT
	Cleanup(func())
}) *Context {
	mock := &Context{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
