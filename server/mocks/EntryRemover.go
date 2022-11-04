// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// EntryRemover is an autogenerated mock type for the EntryRemover type
type EntryRemover struct {
	mock.Mock
}

type EntryRemover_Expecter struct {
	mock *mock.Mock
}

func (_m *EntryRemover) EXPECT() *EntryRemover_Expecter {
	return &EntryRemover_Expecter{mock: &_m.Mock}
}

// Remove provides a mock function with given fields: ctx, entryID
func (_m *EntryRemover) Remove(ctx context.Context, entryID uuid.UUID) error {
	ret := _m.Called(ctx, entryID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) error); ok {
		r0 = rf(ctx, entryID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EntryRemover_Remove_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Remove'
type EntryRemover_Remove_Call struct {
	*mock.Call
}

// Remove is a helper method to define mock.On call
//  - ctx context.Context
//  - entryID uuid.UUID
func (_e *EntryRemover_Expecter) Remove(ctx interface{}, entryID interface{}) *EntryRemover_Remove_Call {
	return &EntryRemover_Remove_Call{Call: _e.mock.On("Remove", ctx, entryID)}
}

func (_c *EntryRemover_Remove_Call) Run(run func(ctx context.Context, entryID uuid.UUID)) *EntryRemover_Remove_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID))
	})
	return _c
}

func (_c *EntryRemover_Remove_Call) Return(_a0 error) *EntryRemover_Remove_Call {
	_c.Call.Return(_a0)
	return _c
}

type mockConstructorTestingTNewEntryRemover interface {
	mock.TestingT
	Cleanup(func())
}

// NewEntryRemover creates a new instance of EntryRemover. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewEntryRemover(t mockConstructorTestingTNewEntryRemover) *EntryRemover {
	mock := &EntryRemover{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
