// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// TypeRemover is an autogenerated mock type for the TypeRemover type
type TypeRemover struct {
	mock.Mock
}

type TypeRemover_Expecter struct {
	mock *mock.Mock
}

func (_m *TypeRemover) EXPECT() *TypeRemover_Expecter {
	return &TypeRemover_Expecter{mock: &_m.Mock}
}

// Remove provides a mock function with given fields: ctx, tpID
func (_m *TypeRemover) Remove(ctx context.Context, tpID uuid.UUID) error {
	ret := _m.Called(ctx, tpID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) error); ok {
		r0 = rf(ctx, tpID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TypeRemover_Remove_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Remove'
type TypeRemover_Remove_Call struct {
	*mock.Call
}

// Remove is a helper method to define mock.On call
//  - ctx context.Context
//  - tpID uuid.UUID
func (_e *TypeRemover_Expecter) Remove(ctx interface{}, tpID interface{}) *TypeRemover_Remove_Call {
	return &TypeRemover_Remove_Call{Call: _e.mock.On("Remove", ctx, tpID)}
}

func (_c *TypeRemover_Remove_Call) Run(run func(ctx context.Context, tpID uuid.UUID)) *TypeRemover_Remove_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID))
	})
	return _c
}

func (_c *TypeRemover_Remove_Call) Return(_a0 error) *TypeRemover_Remove_Call {
	_c.Call.Return(_a0)
	return _c
}

type mockConstructorTestingTNewTypeRemover interface {
	mock.TestingT
	Cleanup(func())
}

// NewTypeRemover creates a new instance of TypeRemover. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTypeRemover(t mockConstructorTestingTNewTypeRemover) *TypeRemover {
	mock := &TypeRemover{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
