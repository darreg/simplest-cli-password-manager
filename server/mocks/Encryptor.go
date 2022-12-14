// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Encryptor is an autogenerated mock type for the Encryptor type
type Encryptor struct {
	mock.Mock
}

type Encryptor_Expecter struct {
	mock *mock.Mock
}

func (_m *Encryptor) EXPECT() *Encryptor_Expecter {
	return &Encryptor_Expecter{mock: &_m.Mock}
}

// Encrypt provides a mock function with given fields: data
func (_m *Encryptor) Encrypt(data []byte) (string, error) {
	ret := _m.Called(data)

	var r0 string
	if rf, ok := ret.Get(0).(func([]byte) string); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Encryptor_Encrypt_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Encrypt'
type Encryptor_Encrypt_Call struct {
	*mock.Call
}

// Encrypt is a helper method to define mock.On call
//  - data []byte
func (_e *Encryptor_Expecter) Encrypt(data interface{}) *Encryptor_Encrypt_Call {
	return &Encryptor_Encrypt_Call{Call: _e.mock.On("Encrypt", data)}
}

func (_c *Encryptor_Encrypt_Call) Run(run func(data []byte)) *Encryptor_Encrypt_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]byte))
	})
	return _c
}

func (_c *Encryptor_Encrypt_Call) Return(_a0 string, _a1 error) *Encryptor_Encrypt_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

type mockConstructorTestingTNewEncryptor interface {
	mock.TestingT
	Cleanup(func())
}

// NewEncryptor creates a new instance of Encryptor. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewEncryptor(t mockConstructorTestingTNewEncryptor) *Encryptor {
	mock := &Encryptor{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
