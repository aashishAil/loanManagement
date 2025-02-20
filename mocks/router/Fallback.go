// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	gin "github.com/gin-gonic/gin"
	mock "github.com/stretchr/testify/mock"
)

// Fallback is an autogenerated mock type for the Fallback type
type Fallback struct {
	mock.Mock
}

type Fallback_Expecter struct {
	mock *mock.Mock
}

func (_m *Fallback) EXPECT() *Fallback_Expecter {
	return &Fallback_Expecter{mock: &_m.Mock}
}

// NoRouteForGinHandler provides a mock function with given fields:
func (_m *Fallback) NoRouteForGinHandler() gin.HandlerFunc {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for NoRouteForGinHandler")
	}

	var r0 gin.HandlerFunc
	if rf, ok := ret.Get(0).(func() gin.HandlerFunc); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(gin.HandlerFunc)
		}
	}

	return r0
}

// Fallback_NoRouteForGinHandler_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'NoRouteForGinHandler'
type Fallback_NoRouteForGinHandler_Call struct {
	*mock.Call
}

// NoRouteForGinHandler is a helper method to define mock.On call
func (_e *Fallback_Expecter) NoRouteForGinHandler() *Fallback_NoRouteForGinHandler_Call {
	return &Fallback_NoRouteForGinHandler_Call{Call: _e.mock.On("NoRouteForGinHandler")}
}

func (_c *Fallback_NoRouteForGinHandler_Call) Run(run func()) *Fallback_NoRouteForGinHandler_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Fallback_NoRouteForGinHandler_Call) Return(_a0 gin.HandlerFunc) *Fallback_NoRouteForGinHandler_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Fallback_NoRouteForGinHandler_Call) RunAndReturn(run func() gin.HandlerFunc) *Fallback_NoRouteForGinHandler_Call {
	_c.Call.Return(run)
	return _c
}

// PingForGinRoute provides a mock function with given fields: c
func (_m *Fallback) PingForGinRoute(c *gin.Context) {
	_m.Called(c)
}

// Fallback_PingForGinRoute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PingForGinRoute'
type Fallback_PingForGinRoute_Call struct {
	*mock.Call
}

// PingForGinRoute is a helper method to define mock.On call
//   - c *gin.Context
func (_e *Fallback_Expecter) PingForGinRoute(c interface{}) *Fallback_PingForGinRoute_Call {
	return &Fallback_PingForGinRoute_Call{Call: _e.mock.On("PingForGinRoute", c)}
}

func (_c *Fallback_PingForGinRoute_Call) Run(run func(c *gin.Context)) *Fallback_PingForGinRoute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gin.Context))
	})
	return _c
}

func (_c *Fallback_PingForGinRoute_Call) Return() *Fallback_PingForGinRoute_Call {
	_c.Call.Return()
	return _c
}

func (_c *Fallback_PingForGinRoute_Call) RunAndReturn(run func(*gin.Context)) *Fallback_PingForGinRoute_Call {
	_c.Call.Return(run)
	return _c
}

// NewFallback creates a new instance of Fallback. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewFallback(t interface {
	mock.TestingT
	Cleanup(func())
}) *Fallback {
	mock := &Fallback{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
