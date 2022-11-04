// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/alrund/yp-2-project/server/internal/domain/entity"
	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// TypeRepository is an autogenerated mock type for the TypeRepository type
type TypeRepository struct {
	mock.Mock
}

type TypeRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *TypeRepository) EXPECT() *TypeRepository_Expecter {
	return &TypeRepository_Expecter{mock: &_m.Mock}
}

// Add provides a mock function with given fields: ctx, tp
func (_m *TypeRepository) Add(ctx context.Context, tp *entity.Type) error {
	ret := _m.Called(ctx, tp)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Type) error); ok {
		r0 = rf(ctx, tp)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TypeRepository_Add_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Add'
type TypeRepository_Add_Call struct {
	*mock.Call
}

// Add is a helper method to define mock.On call
//  - ctx context.Context
//  - tp *entity.Type
func (_e *TypeRepository_Expecter) Add(ctx interface{}, tp interface{}) *TypeRepository_Add_Call {
	return &TypeRepository_Add_Call{Call: _e.mock.On("Add", ctx, tp)}
}

func (_c *TypeRepository_Add_Call) Run(run func(ctx context.Context, tp *entity.Type)) *TypeRepository_Add_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*entity.Type))
	})
	return _c
}

func (_c *TypeRepository_Add_Call) Return(_a0 error) *TypeRepository_Add_Call {
	_c.Call.Return(_a0)
	return _c
}

// Change provides a mock function with given fields: ctx, tp
func (_m *TypeRepository) Change(ctx context.Context, tp *entity.Type) error {
	ret := _m.Called(ctx, tp)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Type) error); ok {
		r0 = rf(ctx, tp)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TypeRepository_Change_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Change'
type TypeRepository_Change_Call struct {
	*mock.Call
}

// Change is a helper method to define mock.On call
//  - ctx context.Context
//  - tp *entity.Type
func (_e *TypeRepository_Expecter) Change(ctx interface{}, tp interface{}) *TypeRepository_Change_Call {
	return &TypeRepository_Change_Call{Call: _e.mock.On("Change", ctx, tp)}
}

func (_c *TypeRepository_Change_Call) Run(run func(ctx context.Context, tp *entity.Type)) *TypeRepository_Change_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*entity.Type))
	})
	return _c
}

func (_c *TypeRepository_Change_Call) Return(_a0 error) *TypeRepository_Change_Call {
	_c.Call.Return(_a0)
	return _c
}

// Get provides a mock function with given fields: ctx, tpID
func (_m *TypeRepository) Get(ctx context.Context, tpID uuid.UUID) (*entity.Type, error) {
	ret := _m.Called(ctx, tpID)

	var r0 *entity.Type
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) *entity.Type); ok {
		r0 = rf(ctx, tpID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Type)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, tpID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TypeRepository_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type TypeRepository_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//  - ctx context.Context
//  - tpID uuid.UUID
func (_e *TypeRepository_Expecter) Get(ctx interface{}, tpID interface{}) *TypeRepository_Get_Call {
	return &TypeRepository_Get_Call{Call: _e.mock.On("Get", ctx, tpID)}
}

func (_c *TypeRepository_Get_Call) Run(run func(ctx context.Context, tpID uuid.UUID)) *TypeRepository_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID))
	})
	return _c
}

func (_c *TypeRepository_Get_Call) Return(_a0 *entity.Type, _a1 error) *TypeRepository_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// Remove provides a mock function with given fields: ctx, tpID
func (_m *TypeRepository) Remove(ctx context.Context, tpID uuid.UUID) error {
	ret := _m.Called(ctx, tpID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) error); ok {
		r0 = rf(ctx, tpID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TypeRepository_Remove_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Remove'
type TypeRepository_Remove_Call struct {
	*mock.Call
}

// Remove is a helper method to define mock.On call
//  - ctx context.Context
//  - tpID uuid.UUID
func (_e *TypeRepository_Expecter) Remove(ctx interface{}, tpID interface{}) *TypeRepository_Remove_Call {
	return &TypeRepository_Remove_Call{Call: _e.mock.On("Remove", ctx, tpID)}
}

func (_c *TypeRepository_Remove_Call) Run(run func(ctx context.Context, tpID uuid.UUID)) *TypeRepository_Remove_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID))
	})
	return _c
}

func (_c *TypeRepository_Remove_Call) Return(_a0 error) *TypeRepository_Remove_Call {
	_c.Call.Return(_a0)
	return _c
}

type mockConstructorTestingTNewTypeRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewTypeRepository creates a new instance of TypeRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTypeRepository(t mockConstructorTestingTNewTypeRepository) *TypeRepository {
	mock := &TypeRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}