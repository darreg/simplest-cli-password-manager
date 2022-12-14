// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// ConfigLoader is an autogenerated mock type for the ConfigLoader type
type ConfigLoader struct {
	mock.Mock
}

type ConfigLoader_Expecter struct {
	mock *mock.Mock
}

func (_m *ConfigLoader) EXPECT() *ConfigLoader_Expecter {
	return &ConfigLoader_Expecter{mock: &_m.Mock}
}

// Load provides a mock function with given fields: cfg
func (_m *ConfigLoader) Load(cfg interface{}) error {
	ret := _m.Called(cfg)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(cfg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ConfigLoader_Load_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Load'
type ConfigLoader_Load_Call struct {
	*mock.Call
}

// Load is a helper method to define mock.On call
//  - cfg interface{}
func (_e *ConfigLoader_Expecter) Load(cfg interface{}) *ConfigLoader_Load_Call {
	return &ConfigLoader_Load_Call{Call: _e.mock.On("Load", cfg)}
}

func (_c *ConfigLoader_Load_Call) Run(run func(cfg interface{})) *ConfigLoader_Load_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}))
	})
	return _c
}

func (_c *ConfigLoader_Load_Call) Return(_a0 error) *ConfigLoader_Load_Call {
	_c.Call.Return(_a0)
	return _c
}

// LoadFile provides a mock function with given fields: path, cfg
func (_m *ConfigLoader) LoadFile(path string, cfg interface{}) error {
	ret := _m.Called(path, cfg)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, interface{}) error); ok {
		r0 = rf(path, cfg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ConfigLoader_LoadFile_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LoadFile'
type ConfigLoader_LoadFile_Call struct {
	*mock.Call
}

// LoadFile is a helper method to define mock.On call
//  - path string
//  - cfg interface{}
func (_e *ConfigLoader_Expecter) LoadFile(path interface{}, cfg interface{}) *ConfigLoader_LoadFile_Call {
	return &ConfigLoader_LoadFile_Call{Call: _e.mock.On("LoadFile", path, cfg)}
}

func (_c *ConfigLoader_LoadFile_Call) Run(run func(path string, cfg interface{})) *ConfigLoader_LoadFile_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(interface{}))
	})
	return _c
}

func (_c *ConfigLoader_LoadFile_Call) Return(_a0 error) *ConfigLoader_LoadFile_Call {
	_c.Call.Return(_a0)
	return _c
}

type mockConstructorTestingTNewConfigLoader interface {
	mock.TestingT
	Cleanup(func())
}

// NewConfigLoader creates a new instance of ConfigLoader. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewConfigLoader(t mockConstructorTestingTNewConfigLoader) *ConfigLoader {
	mock := &ConfigLoader{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
