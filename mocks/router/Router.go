// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	router "loanManagement/router"

	mock "github.com/stretchr/testify/mock"
)

// Router is an autogenerated mock type for the Router type
type Router struct {
	mock.Mock
}

type Router_Expecter struct {
	mock *mock.Mock
}

func (_m *Router) EXPECT() *Router_Expecter {
	return &Router_Expecter{mock: &_m.Mock}
}

// Admin provides a mock function with given fields:
func (_m *Router) Admin() router.Admin {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Admin")
	}

	var r0 router.Admin
	if rf, ok := ret.Get(0).(func() router.Admin); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(router.Admin)
		}
	}

	return r0
}

// Router_Admin_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Admin'
type Router_Admin_Call struct {
	*mock.Call
}

// Admin is a helper method to define mock.On call
func (_e *Router_Expecter) Admin() *Router_Admin_Call {
	return &Router_Admin_Call{Call: _e.mock.On("Admin")}
}

func (_c *Router_Admin_Call) Run(run func()) *Router_Admin_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Router_Admin_Call) Return(_a0 router.Admin) *Router_Admin_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Router_Admin_Call) RunAndReturn(run func() router.Admin) *Router_Admin_Call {
	_c.Call.Return(run)
	return _c
}

// Fallback provides a mock function with given fields:
func (_m *Router) Fallback() router.Fallback {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Fallback")
	}

	var r0 router.Fallback
	if rf, ok := ret.Get(0).(func() router.Fallback); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(router.Fallback)
		}
	}

	return r0
}

// Router_Fallback_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Fallback'
type Router_Fallback_Call struct {
	*mock.Call
}

// Fallback is a helper method to define mock.On call
func (_e *Router_Expecter) Fallback() *Router_Fallback_Call {
	return &Router_Fallback_Call{Call: _e.mock.On("Fallback")}
}

func (_c *Router_Fallback_Call) Run(run func()) *Router_Fallback_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Router_Fallback_Call) Return(_a0 router.Fallback) *Router_Fallback_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Router_Fallback_Call) RunAndReturn(run func() router.Fallback) *Router_Fallback_Call {
	_c.Call.Return(run)
	return _c
}

// User provides a mock function with given fields:
func (_m *Router) User() router.User {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for User")
	}

	var r0 router.User
	if rf, ok := ret.Get(0).(func() router.User); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(router.User)
		}
	}

	return r0
}

// Router_User_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'User'
type Router_User_Call struct {
	*mock.Call
}

// User is a helper method to define mock.On call
func (_e *Router_Expecter) User() *Router_User_Call {
	return &Router_User_Call{Call: _e.mock.On("User")}
}

func (_c *Router_User_Call) Run(run func()) *Router_User_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Router_User_Call) Return(_a0 router.User) *Router_User_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Router_User_Call) RunAndReturn(run func() router.User) *Router_User_Call {
	_c.Call.Return(run)
	return _c
}

// NewRouter creates a new instance of Router. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRouter(t interface {
	mock.TestingT
	Cleanup(func())
}) *Router {
	mock := &Router{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
