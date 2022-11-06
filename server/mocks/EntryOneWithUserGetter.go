// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/alrund/yp-2-project/server/internal/domain/entity"
	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// EntryOneWithUserGetter is an autogenerated mock type for the EntryOneWithUserGetter type
type EntryOneWithUserGetter struct {
	mock.Mock
}

type EntryOneWithUserGetter_Expecter struct {
	mock *mock.Mock
}

func (_m *EntryOneWithUserGetter) EXPECT() *EntryOneWithUserGetter_Expecter {
	return &EntryOneWithUserGetter_Expecter{mock: &_m.Mock}
}

// GetOneWithUser provides a mock function with given fields: ctx, entryID, user
func (_m *EntryOneWithUserGetter) GetOneWithUser(ctx context.Context, entryID uuid.UUID, user *entity.User) (*entity.Entry, error) {
	ret := _m.Called(ctx, entryID, user)

	var r0 *entity.Entry
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, *entity.User) *entity.Entry); ok {
		r0 = rf(ctx, entryID, user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Entry)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID, *entity.User) error); ok {
		r1 = rf(ctx, entryID, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EntryOneWithUserGetter_GetOneWithUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetOneWithUser'
type EntryOneWithUserGetter_GetOneWithUser_Call struct {
	*mock.Call
}

// GetOneWithUser is a helper method to define mock.On call
//  - ctx context.Context
//  - entryID uuid.UUID
//  - user *entity.User
func (_e *EntryOneWithUserGetter_Expecter) GetOneWithUser(ctx interface{}, entryID interface{}, user interface{}) *EntryOneWithUserGetter_GetOneWithUser_Call {
	return &EntryOneWithUserGetter_GetOneWithUser_Call{Call: _e.mock.On("GetOneWithUser", ctx, entryID, user)}
}

func (_c *EntryOneWithUserGetter_GetOneWithUser_Call) Run(run func(ctx context.Context, entryID uuid.UUID, user *entity.User)) *EntryOneWithUserGetter_GetOneWithUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID), args[2].(*entity.User))
	})
	return _c
}

func (_c *EntryOneWithUserGetter_GetOneWithUser_Call) Return(_a0 *entity.Entry, _a1 error) *EntryOneWithUserGetter_GetOneWithUser_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

type mockConstructorTestingTNewEntryOneWithUserGetter interface {
	mock.TestingT
	Cleanup(func())
}

// NewEntryOneWithUserGetter creates a new instance of EntryOneWithUserGetter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewEntryOneWithUserGetter(t mockConstructorTestingTNewEntryOneWithUserGetter) *EntryOneWithUserGetter {
	mock := &EntryOneWithUserGetter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
