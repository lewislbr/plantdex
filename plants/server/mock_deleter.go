// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package server

import mock "github.com/stretchr/testify/mock"

// MockDeleter is an autogenerated mock type for the Deleter type
type MockDeleter struct {
	mock.Mock
}

// Delete provides a mock function with given fields: _a0, _a1
func (_m *MockDeleter) Delete(_a0 string, _a1 string) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}