// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/alrund/yp-2-project/server/internal/domain/entity"
	mock "github.com/stretchr/testify/mock"
)

// TypeAllGetter is an autogenerated mock type for the TypeAllGetter type
type TypeAllGetter struct {
	mock.Mock
}

type TypeAllGetter_Expecter struct {
	mock *mock.Mock
}

func (_m *TypeAllGetter) EXPECT() *TypeAllGetter_Expecter {
	return &TypeAllGetter_Expecter{mock: &_m.Mock}
}

// GetAll provides a mock function with given fields: ctx
func (_m *TypeAllGetter) GetAll(ctx context.Context) ([]*entity.Type, error) {
	ret := _m.Called(ctx)

	var r0 []*entity.Type
	if rf, ok := ret.Get(0).(func(context.Context) []*entity.Type); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entity.Type)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TypeAllGetter_GetAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAll'
type TypeAllGetter_GetAll_Call struct {
	*mock.Call
}

// GetAll is a helper method to define mock.On call
//  - ctx context.Context
func (_e *TypeAllGetter_Expecter) GetAll(ctx interface{}) *TypeAllGetter_GetAll_Call {
	return &TypeAllGetter_GetAll_Call{Call: _e.mock.On("GetAll", ctx)}
}

func (_c *TypeAllGetter_GetAll_Call) Run(run func(ctx context.Context)) *TypeAllGetter_GetAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *TypeAllGetter_GetAll_Call) Return(_a0 []*entity.Type, _a1 error) *TypeAllGetter_GetAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

type mockConstructorTestingTNewTypeAllGetter interface {
	mock.TestingT
	Cleanup(func())
}

// NewTypeAllGetter creates a new instance of TypeAllGetter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTypeAllGetter(t mockConstructorTestingTNewTypeAllGetter) *TypeAllGetter {
	mock := &TypeAllGetter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}