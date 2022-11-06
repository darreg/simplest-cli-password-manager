// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// EncryptDecryptor is an autogenerated mock type for the EncryptDecryptor type
type EncryptDecryptor struct {
	mock.Mock
}

type EncryptDecryptor_Expecter struct {
	mock *mock.Mock
}

func (_m *EncryptDecryptor) EXPECT() *EncryptDecryptor_Expecter {
	return &EncryptDecryptor_Expecter{mock: &_m.Mock}
}

// Decrypt provides a mock function with given fields: encrypted
func (_m *EncryptDecryptor) Decrypt(encrypted string) ([]byte, error) {
	ret := _m.Called(encrypted)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(string) []byte); ok {
		r0 = rf(encrypted)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(encrypted)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EncryptDecryptor_Decrypt_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Decrypt'
type EncryptDecryptor_Decrypt_Call struct {
	*mock.Call
}

// Decrypt is a helper method to define mock.On call
//  - encrypted string
func (_e *EncryptDecryptor_Expecter) Decrypt(encrypted interface{}) *EncryptDecryptor_Decrypt_Call {
	return &EncryptDecryptor_Decrypt_Call{Call: _e.mock.On("Decrypt", encrypted)}
}

func (_c *EncryptDecryptor_Decrypt_Call) Run(run func(encrypted string)) *EncryptDecryptor_Decrypt_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *EncryptDecryptor_Decrypt_Call) Return(_a0 []byte, _a1 error) *EncryptDecryptor_Decrypt_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// Encrypt provides a mock function with given fields: data
func (_m *EncryptDecryptor) Encrypt(data []byte) (string, error) {
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

// EncryptDecryptor_Encrypt_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Encrypt'
type EncryptDecryptor_Encrypt_Call struct {
	*mock.Call
}

// Encrypt is a helper method to define mock.On call
//  - data []byte
func (_e *EncryptDecryptor_Expecter) Encrypt(data interface{}) *EncryptDecryptor_Encrypt_Call {
	return &EncryptDecryptor_Encrypt_Call{Call: _e.mock.On("Encrypt", data)}
}

func (_c *EncryptDecryptor_Encrypt_Call) Run(run func(data []byte)) *EncryptDecryptor_Encrypt_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]byte))
	})
	return _c
}

func (_c *EncryptDecryptor_Encrypt_Call) Return(_a0 string, _a1 error) *EncryptDecryptor_Encrypt_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

type mockConstructorTestingTNewEncryptDecryptor interface {
	mock.TestingT
	Cleanup(func())
}

// NewEncryptDecryptor creates a new instance of EncryptDecryptor. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewEncryptDecryptor(t mockConstructorTestingTNewEncryptDecryptor) *EncryptDecryptor {
	mock := &EncryptDecryptor{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}