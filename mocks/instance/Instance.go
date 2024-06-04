// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	databaseinstance "loanManagement/database/instance"

	mock "github.com/stretchr/testify/mock"

	util "loanManagement/util"
)

// Instance is an autogenerated mock type for the Instance type
type Instance struct {
	mock.Mock
}

type Instance_Expecter struct {
	mock *mock.Mock
}

func (_m *Instance) EXPECT() *Instance_Expecter {
	return &Instance_Expecter{mock: &_m.Mock}
}

// ContextUtil provides a mock function with given fields:
func (_m *Instance) ContextUtil() util.Context {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ContextUtil")
	}

	var r0 util.Context
	if rf, ok := ret.Get(0).(func() util.Context); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(util.Context)
		}
	}

	return r0
}

// Instance_ContextUtil_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ContextUtil'
type Instance_ContextUtil_Call struct {
	*mock.Call
}

// ContextUtil is a helper method to define mock.On call
func (_e *Instance_Expecter) ContextUtil() *Instance_ContextUtil_Call {
	return &Instance_ContextUtil_Call{Call: _e.mock.On("ContextUtil")}
}

func (_c *Instance_ContextUtil_Call) Run(run func()) *Instance_ContextUtil_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Instance_ContextUtil_Call) Return(_a0 util.Context) *Instance_ContextUtil_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Instance_ContextUtil_Call) RunAndReturn(run func() util.Context) *Instance_ContextUtil_Call {
	_c.Call.Return(run)
	return _c
}

// DatabaseInstance provides a mock function with given fields:
func (_m *Instance) DatabaseInstance() databaseinstance.PostgresDB {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for DatabaseInstance")
	}

	var r0 databaseinstance.PostgresDB
	if rf, ok := ret.Get(0).(func() databaseinstance.PostgresDB); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(databaseinstance.PostgresDB)
		}
	}

	return r0
}

// Instance_DatabaseInstance_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DatabaseInstance'
type Instance_DatabaseInstance_Call struct {
	*mock.Call
}

// DatabaseInstance is a helper method to define mock.On call
func (_e *Instance_Expecter) DatabaseInstance() *Instance_DatabaseInstance_Call {
	return &Instance_DatabaseInstance_Call{Call: _e.mock.On("DatabaseInstance")}
}

func (_c *Instance_DatabaseInstance_Call) Run(run func()) *Instance_DatabaseInstance_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Instance_DatabaseInstance_Call) Return(_a0 databaseinstance.PostgresDB) *Instance_DatabaseInstance_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Instance_DatabaseInstance_Call) RunAndReturn(run func() databaseinstance.PostgresDB) *Instance_DatabaseInstance_Call {
	_c.Call.Return(run)
	return _c
}

// JwtUtil provides a mock function with given fields:
func (_m *Instance) JwtUtil() util.Jwt {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for JwtUtil")
	}

	var r0 util.Jwt
	if rf, ok := ret.Get(0).(func() util.Jwt); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(util.Jwt)
		}
	}

	return r0
}

// Instance_JwtUtil_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'JwtUtil'
type Instance_JwtUtil_Call struct {
	*mock.Call
}

// JwtUtil is a helper method to define mock.On call
func (_e *Instance_Expecter) JwtUtil() *Instance_JwtUtil_Call {
	return &Instance_JwtUtil_Call{Call: _e.mock.On("JwtUtil")}
}

func (_c *Instance_JwtUtil_Call) Run(run func()) *Instance_JwtUtil_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Instance_JwtUtil_Call) Return(_a0 util.Jwt) *Instance_JwtUtil_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Instance_JwtUtil_Call) RunAndReturn(run func() util.Jwt) *Instance_JwtUtil_Call {
	_c.Call.Return(run)
	return _c
}

// PasswordUtil provides a mock function with given fields:
func (_m *Instance) PasswordUtil() util.Password {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for PasswordUtil")
	}

	var r0 util.Password
	if rf, ok := ret.Get(0).(func() util.Password); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(util.Password)
		}
	}

	return r0
}

// Instance_PasswordUtil_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PasswordUtil'
type Instance_PasswordUtil_Call struct {
	*mock.Call
}

// PasswordUtil is a helper method to define mock.On call
func (_e *Instance_Expecter) PasswordUtil() *Instance_PasswordUtil_Call {
	return &Instance_PasswordUtil_Call{Call: _e.mock.On("PasswordUtil")}
}

func (_c *Instance_PasswordUtil_Call) Run(run func()) *Instance_PasswordUtil_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Instance_PasswordUtil_Call) Return(_a0 util.Password) *Instance_PasswordUtil_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Instance_PasswordUtil_Call) RunAndReturn(run func() util.Password) *Instance_PasswordUtil_Call {
	_c.Call.Return(run)
	return _c
}

// TimeUtil provides a mock function with given fields:
func (_m *Instance) TimeUtil() util.Time {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for TimeUtil")
	}

	var r0 util.Time
	if rf, ok := ret.Get(0).(func() util.Time); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(util.Time)
		}
	}

	return r0
}

// Instance_TimeUtil_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'TimeUtil'
type Instance_TimeUtil_Call struct {
	*mock.Call
}

// TimeUtil is a helper method to define mock.On call
func (_e *Instance_Expecter) TimeUtil() *Instance_TimeUtil_Call {
	return &Instance_TimeUtil_Call{Call: _e.mock.On("TimeUtil")}
}

func (_c *Instance_TimeUtil_Call) Run(run func()) *Instance_TimeUtil_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Instance_TimeUtil_Call) Return(_a0 util.Time) *Instance_TimeUtil_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Instance_TimeUtil_Call) RunAndReturn(run func() util.Time) *Instance_TimeUtil_Call {
	_c.Call.Return(run)
	return _c
}

// NewInstance creates a new instance of Instance. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewInstance(t interface {
	mock.TestingT
	Cleanup(func())
}) *Instance {
	mock := &Instance{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
