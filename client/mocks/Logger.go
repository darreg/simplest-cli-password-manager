// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

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

// Debug provides a mock function with given fields: msg, args
func (_m *Logger) Debug(msg string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, msg)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Logger_Debug_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Debug'
type Logger_Debug_Call struct {
	*mock.Call
}

// Debug is a helper method to define mock.On call
//  - msg string
//  - args ...interface{}
func (_e *Logger_Expecter) Debug(msg interface{}, args ...interface{}) *Logger_Debug_Call {
	return &Logger_Debug_Call{Call: _e.mock.On("Debug",
		append([]interface{}{msg}, args...)...)}
}

func (_c *Logger_Debug_Call) Run(run func(msg string, args ...interface{})) *Logger_Debug_Call {
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

func (_c *Logger_Debug_Call) Return() *Logger_Debug_Call {
	_c.Call.Return()
	return _c
}

// EnableDebug provides a mock function with given fields:
func (_m *Logger) EnableDebug() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Logger_EnableDebug_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'EnableDebug'
type Logger_EnableDebug_Call struct {
	*mock.Call
}

// EnableDebug is a helper method to define mock.On call
func (_e *Logger_Expecter) EnableDebug() *Logger_EnableDebug_Call {
	return &Logger_EnableDebug_Call{Call: _e.mock.On("EnableDebug")}
}

func (_c *Logger_EnableDebug_Call) Run(run func()) *Logger_EnableDebug_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Logger_EnableDebug_Call) Return(_a0 error) *Logger_EnableDebug_Call {
	_c.Call.Return(_a0)
	return _c
}

// Error provides a mock function with given fields: err, args
func (_m *Logger) Error(err error, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, err)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Logger_Error_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Error'
type Logger_Error_Call struct {
	*mock.Call
}

// Error is a helper method to define mock.On call
//  - err error
//  - args ...interface{}
func (_e *Logger_Expecter) Error(err interface{}, args ...interface{}) *Logger_Error_Call {
	return &Logger_Error_Call{Call: _e.mock.On("Error",
		append([]interface{}{err}, args...)...)}
}

func (_c *Logger_Error_Call) Run(run func(err error, args ...interface{})) *Logger_Error_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(error), variadicArgs...)
	})
	return _c
}

func (_c *Logger_Error_Call) Return() *Logger_Error_Call {
	_c.Call.Return()
	return _c
}

// Fatal provides a mock function with given fields: err, args
func (_m *Logger) Fatal(err error, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, err)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Logger_Fatal_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Fatal'
type Logger_Fatal_Call struct {
	*mock.Call
}

// Fatal is a helper method to define mock.On call
//  - err error
//  - args ...interface{}
func (_e *Logger_Expecter) Fatal(err interface{}, args ...interface{}) *Logger_Fatal_Call {
	return &Logger_Fatal_Call{Call: _e.mock.On("Fatal",
		append([]interface{}{err}, args...)...)}
}

func (_c *Logger_Fatal_Call) Run(run func(err error, args ...interface{})) *Logger_Fatal_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(error), variadicArgs...)
	})
	return _c
}

func (_c *Logger_Fatal_Call) Return() *Logger_Fatal_Call {
	_c.Call.Return()
	return _c
}

// Info provides a mock function with given fields: msg, args
func (_m *Logger) Info(msg string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, msg)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Logger_Info_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Info'
type Logger_Info_Call struct {
	*mock.Call
}

// Info is a helper method to define mock.On call
//  - msg string
//  - args ...interface{}
func (_e *Logger_Expecter) Info(msg interface{}, args ...interface{}) *Logger_Info_Call {
	return &Logger_Info_Call{Call: _e.mock.On("Info",
		append([]interface{}{msg}, args...)...)}
}

func (_c *Logger_Info_Call) Run(run func(msg string, args ...interface{})) *Logger_Info_Call {
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

func (_c *Logger_Info_Call) Return() *Logger_Info_Call {
	_c.Call.Return()
	return _c
}

// Warn provides a mock function with given fields: msg, args
func (_m *Logger) Warn(msg string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, msg)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Logger_Warn_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Warn'
type Logger_Warn_Call struct {
	*mock.Call
}

// Warn is a helper method to define mock.On call
//  - msg string
//  - args ...interface{}
func (_e *Logger_Expecter) Warn(msg interface{}, args ...interface{}) *Logger_Warn_Call {
	return &Logger_Warn_Call{Call: _e.mock.On("Warn",
		append([]interface{}{msg}, args...)...)}
}

func (_c *Logger_Warn_Call) Run(run func(msg string, args ...interface{})) *Logger_Warn_Call {
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

func (_c *Logger_Warn_Call) Return() *Logger_Warn_Call {
	_c.Call.Return()
	return _c
}

type mockConstructorTestingTNewLogger interface {
	mock.TestingT
	Cleanup(func())
}

// NewLogger creates a new instance of Logger. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewLogger(t mockConstructorTestingTNewLogger) *Logger {
	mock := &Logger{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
