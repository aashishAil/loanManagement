// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	zapcore "go.uber.org/zap/zapcore"
)

// Logger is an autogenerated mock type for the Logger type
type Logger struct {
	mock.Mock
}

type Logger_Expecter struct {
	mock *mock.Mock
}

func (_m *Logger) EXPECT() *Logger_Expecter {
	return &Logger_Expecter{mock: &_m.Mock}
}

// Debug provides a mock function with given fields: message, args
func (_m *Logger) Debug(message string, args ...zapcore.Field) {
	_va := make([]interface{}, len(args))
	for _i := range args {
		_va[_i] = args[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, message)
	_ca = append(_ca, _va...)
	_m.Called(_ca...)
}

// Logger_Debug_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Debug'
type Logger_Debug_Call struct {
	*mock.Call
}

// Debug is a helper method to define mock.On call
//   - message string
//   - args ...zapcore.Field
func (_e *Logger_Expecter) Debug(message interface{}, args ...interface{}) *Logger_Debug_Call {
	return &Logger_Debug_Call{Call: _e.mock.On("Debug",
		append([]interface{}{message}, args...)...)}
}

func (_c *Logger_Debug_Call) Run(run func(message string, args ...zapcore.Field)) *Logger_Debug_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]zapcore.Field, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(zapcore.Field)
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *Logger_Debug_Call) Return() *Logger_Debug_Call {
	_c.Call.Return()
	return _c
}

func (_c *Logger_Debug_Call) RunAndReturn(run func(string, ...zapcore.Field)) *Logger_Debug_Call {
	_c.Call.Return(run)
	return _c
}

// Error provides a mock function with given fields: message, args
func (_m *Logger) Error(message string, args ...zapcore.Field) {
	_va := make([]interface{}, len(args))
	for _i := range args {
		_va[_i] = args[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, message)
	_ca = append(_ca, _va...)
	_m.Called(_ca...)
}

// Logger_Error_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Error'
type Logger_Error_Call struct {
	*mock.Call
}

// Error is a helper method to define mock.On call
//   - message string
//   - args ...zapcore.Field
func (_e *Logger_Expecter) Error(message interface{}, args ...interface{}) *Logger_Error_Call {
	return &Logger_Error_Call{Call: _e.mock.On("Error",
		append([]interface{}{message}, args...)...)}
}

func (_c *Logger_Error_Call) Run(run func(message string, args ...zapcore.Field)) *Logger_Error_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]zapcore.Field, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(zapcore.Field)
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *Logger_Error_Call) Return() *Logger_Error_Call {
	_c.Call.Return()
	return _c
}

func (_c *Logger_Error_Call) RunAndReturn(run func(string, ...zapcore.Field)) *Logger_Error_Call {
	_c.Call.Return(run)
	return _c
}

// Errorf provides a mock function with given fields: message, args
func (_m *Logger) Errorf(message string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, message)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Logger_Errorf_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Errorf'
type Logger_Errorf_Call struct {
	*mock.Call
}

// Errorf is a helper method to define mock.On call
//   - message string
//   - args ...interface{}
func (_e *Logger_Expecter) Errorf(message interface{}, args ...interface{}) *Logger_Errorf_Call {
	return &Logger_Errorf_Call{Call: _e.mock.On("Errorf",
		append([]interface{}{message}, args...)...)}
}

func (_c *Logger_Errorf_Call) Run(run func(message string, args ...interface{})) *Logger_Errorf_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *Logger_Errorf_Call) Return() *Logger_Errorf_Call {
	_c.Call.Return()
	return _c
}

func (_c *Logger_Errorf_Call) RunAndReturn(run func(string, ...interface{})) *Logger_Errorf_Call {
	_c.Call.Return(run)
	return _c
}

// Info provides a mock function with given fields: message, args
func (_m *Logger) Info(message string, args ...zapcore.Field) {
	_va := make([]interface{}, len(args))
	for _i := range args {
		_va[_i] = args[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, message)
	_ca = append(_ca, _va...)
	_m.Called(_ca...)
}

// Logger_Info_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Info'
type Logger_Info_Call struct {
	*mock.Call
}

// Info is a helper method to define mock.On call
//   - message string
//   - args ...zapcore.Field
func (_e *Logger_Expecter) Info(message interface{}, args ...interface{}) *Logger_Info_Call {
	return &Logger_Info_Call{Call: _e.mock.On("Info",
		append([]interface{}{message}, args...)...)}
}

func (_c *Logger_Info_Call) Run(run func(message string, args ...zapcore.Field)) *Logger_Info_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]zapcore.Field, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(zapcore.Field)
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *Logger_Info_Call) Return() *Logger_Info_Call {
	_c.Call.Return()
	return _c
}

func (_c *Logger_Info_Call) RunAndReturn(run func(string, ...zapcore.Field)) *Logger_Info_Call {
	_c.Call.Return(run)
	return _c
}

// Infof provides a mock function with given fields: message, args
func (_m *Logger) Infof(message string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, message)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Logger_Infof_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Infof'
type Logger_Infof_Call struct {
	*mock.Call
}

// Infof is a helper method to define mock.On call
//   - message string
//   - args ...interface{}
func (_e *Logger_Expecter) Infof(message interface{}, args ...interface{}) *Logger_Infof_Call {
	return &Logger_Infof_Call{Call: _e.mock.On("Infof",
		append([]interface{}{message}, args...)...)}
}

func (_c *Logger_Infof_Call) Run(run func(message string, args ...interface{})) *Logger_Infof_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *Logger_Infof_Call) Return() *Logger_Infof_Call {
	_c.Call.Return()
	return _c
}

func (_c *Logger_Infof_Call) RunAndReturn(run func(string, ...interface{})) *Logger_Infof_Call {
	_c.Call.Return(run)
	return _c
}

// ShutDownLogger provides a mock function with given fields:
func (_m *Logger) ShutDownLogger() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ShutDownLogger")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Logger_ShutDownLogger_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ShutDownLogger'
type Logger_ShutDownLogger_Call struct {
	*mock.Call
}

// ShutDownLogger is a helper method to define mock.On call
func (_e *Logger_Expecter) ShutDownLogger() *Logger_ShutDownLogger_Call {
	return &Logger_ShutDownLogger_Call{Call: _e.mock.On("ShutDownLogger")}
}

func (_c *Logger_ShutDownLogger_Call) Run(run func()) *Logger_ShutDownLogger_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Logger_ShutDownLogger_Call) Return(_a0 error) *Logger_ShutDownLogger_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Logger_ShutDownLogger_Call) RunAndReturn(run func() error) *Logger_ShutDownLogger_Call {
	_c.Call.Return(run)
	return _c
}

// Warn provides a mock function with given fields: message, args
func (_m *Logger) Warn(message string, args ...zapcore.Field) {
	_va := make([]interface{}, len(args))
	for _i := range args {
		_va[_i] = args[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, message)
	_ca = append(_ca, _va...)
	_m.Called(_ca...)
}

// Logger_Warn_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Warn'
type Logger_Warn_Call struct {
	*mock.Call
}

// Warn is a helper method to define mock.On call
//   - message string
//   - args ...zapcore.Field
func (_e *Logger_Expecter) Warn(message interface{}, args ...interface{}) *Logger_Warn_Call {
	return &Logger_Warn_Call{Call: _e.mock.On("Warn",
		append([]interface{}{message}, args...)...)}
}

func (_c *Logger_Warn_Call) Run(run func(message string, args ...zapcore.Field)) *Logger_Warn_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]zapcore.Field, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(zapcore.Field)
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *Logger_Warn_Call) Return() *Logger_Warn_Call {
	_c.Call.Return()
	return _c
}

func (_c *Logger_Warn_Call) RunAndReturn(run func(string, ...zapcore.Field)) *Logger_Warn_Call {
	_c.Call.Return(run)
	return _c
}

// Warnf provides a mock function with given fields: message, args
func (_m *Logger) Warnf(message string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, message)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Logger_Warnf_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Warnf'
type Logger_Warnf_Call struct {
	*mock.Call
}

// Warnf is a helper method to define mock.On call
//   - message string
//   - args ...interface{}
func (_e *Logger_Expecter) Warnf(message interface{}, args ...interface{}) *Logger_Warnf_Call {
	return &Logger_Warnf_Call{Call: _e.mock.On("Warnf",
		append([]interface{}{message}, args...)...)}
}

func (_c *Logger_Warnf_Call) Run(run func(message string, args ...interface{})) *Logger_Warnf_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *Logger_Warnf_Call) Return() *Logger_Warnf_Call {
	_c.Call.Return()
	return _c
}

func (_c *Logger_Warnf_Call) RunAndReturn(run func(string, ...interface{})) *Logger_Warnf_Call {
	_c.Call.Return(run)
	return _c
}

// NewLogger creates a new instance of Logger. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewLogger(t interface {
	mock.TestingT
	Cleanup(func())
}) *Logger {
	mock := &Logger{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
