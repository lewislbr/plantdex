// Code generated by mockery (devel). DO NOT EDIT.

package server

import mock "github.com/stretchr/testify/mock"

// mockTokenService is an autogenerated mock type for the tokenService type
type mockTokenService struct {
	mock.Mock
}

// Generate provides a mock function with given fields: _a0
func (_m *mockTokenService) Generate(_a0 string) (string, error) {
	ret := _m.Called(_a0)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Revoke provides a mock function with given fields: _a0
func (_m *mockTokenService) Revoke(_a0 string) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Validate provides a mock function with given fields: _a0
func (_m *mockTokenService) Validate(_a0 string) (string, error) {
	ret := _m.Called(_a0)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}