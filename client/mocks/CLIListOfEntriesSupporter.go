// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// CLIListOfEntriesSupporter is an autogenerated mock type for the CLIListOfEntriesSupporter type
type CLIListOfEntriesSupporter struct {
	mock.Mock
}

type CLIListOfEntriesSupporter_Expecter struct {
	mock *mock.Mock
}

func (_m *CLIListOfEntriesSupporter) EXPECT() *CLIListOfEntriesSupporter_Expecter {
	return &CLIListOfEntriesSupporter_Expecter{mock: &_m.Mock}
}

// ListOfEntries provides a mock function with given fields: ctx, entries, data
func (_m *CLIListOfEntriesSupporter) ListOfEntries(ctx context.Context, entries []string, data interface{}) error {
	ret := _m.Called(ctx, entries, data)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []string, interface{}) error); ok {
		r0 = rf(ctx, entries, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CLIListOfEntriesSupporter_ListOfEntries_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListOfEntries'
type CLIListOfEntriesSupporter_ListOfEntries_Call struct {
	*mock.Call
}

// ListOfEntries is a helper method to define mock.On call
//  - ctx context.Context
//  - entries []string
//  - data interface{}
func (_e *CLIListOfEntriesSupporter_Expecter) ListOfEntries(ctx interface{}, entries interface{}, data interface{}) *CLIListOfEntriesSupporter_ListOfEntries_Call {
	return &CLIListOfEntriesSupporter_ListOfEntries_Call{Call: _e.mock.On("ListOfEntries", ctx, entries, data)}
}

func (_c *CLIListOfEntriesSupporter_ListOfEntries_Call) Run(run func(ctx context.Context, entries []string, data interface{})) *CLIListOfEntriesSupporter_ListOfEntries_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]string), args[2].(interface{}))
	})
	return _c
}

func (_c *CLIListOfEntriesSupporter_ListOfEntries_Call) Return(_a0 error) *CLIListOfEntriesSupporter_ListOfEntries_Call {
	_c.Call.Return(_a0)
	return _c
}

type mockConstructorTestingTNewCLIListOfEntriesSupporter interface {
	mock.TestingT
	Cleanup(func())
}

// NewCLIListOfEntriesSupporter creates a new instance of CLIListOfEntriesSupporter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCLIListOfEntriesSupporter(t mockConstructorTestingTNewCLIListOfEntriesSupporter) *CLIListOfEntriesSupporter {
	mock := &CLIListOfEntriesSupporter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}