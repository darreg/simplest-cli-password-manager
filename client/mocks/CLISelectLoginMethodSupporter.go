// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// CLISelectLoginMethodSupporter is an autogenerated mock type for the CLISelectLoginMethodSupporter type
type CLISelectLoginMethodSupporter struct {
	mock.Mock
}

type CLISelectLoginMethodSupporter_Expecter struct {
	mock *mock.Mock
}

func (_m *CLISelectLoginMethodSupporter) EXPECT() *CLISelectLoginMethodSupporter_Expecter {
	return &CLISelectLoginMethodSupporter_Expecter{mock: &_m.Mock}
}

// SelectLoginMethod provides a mock function with given fields: ctx, options, data
func (_m *CLISelectLoginMethodSupporter) SelectLoginMethod(ctx context.Context, options []string, data interface{}) error {
	ret := _m.Called(ctx, options, data)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []string, interface{}) error); ok {
		r0 = rf(ctx, options, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CLISelectLoginMethodSupporter_SelectLoginMethod_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SelectLoginMethod'
type CLISelectLoginMethodSupporter_SelectLoginMethod_Call struct {
	*mock.Call
}

// SelectLoginMethod is a helper method to define mock.On call
//  - ctx context.Context
//  - options []string
//  - data interface{}
func (_e *CLISelectLoginMethodSupporter_Expecter) SelectLoginMethod(ctx interface{}, options interface{}, data interface{}) *CLISelectLoginMethodSupporter_SelectLoginMethod_Call {
	return &CLISelectLoginMethodSupporter_SelectLoginMethod_Call{Call: _e.mock.On("SelectLoginMethod", ctx, options, data)}
}

func (_c *CLISelectLoginMethodSupporter_SelectLoginMethod_Call) Run(run func(ctx context.Context, options []string, data interface{})) *CLISelectLoginMethodSupporter_SelectLoginMethod_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]string), args[2].(interface{}))
	})
	return _c
}

func (_c *CLISelectLoginMethodSupporter_SelectLoginMethod_Call) Return(_a0 error) *CLISelectLoginMethodSupporter_SelectLoginMethod_Call {
	_c.Call.Return(_a0)
	return _c
}

type mockConstructorTestingTNewCLISelectLoginMethodSupporter interface {
	mock.TestingT
	Cleanup(func())
}

// NewCLISelectLoginMethodSupporter creates a new instance of CLISelectLoginMethodSupporter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCLISelectLoginMethodSupporter(t mockConstructorTestingTNewCLISelectLoginMethodSupporter) *CLISelectLoginMethodSupporter {
	mock := &CLISelectLoginMethodSupporter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}