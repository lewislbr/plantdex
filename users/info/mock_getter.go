// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package info

import (
	user "users/user"

	mock "github.com/stretchr/testify/mock"
)

// MockGetter is an autogenerated mock type for the Getter type
type MockGetter struct {
	mock.Mock
}

// GetUserInfo provides a mock function with given fields: _a0
func (_m *MockGetter) GetUserInfo(_a0 string) (user.Info, error) {
	ret := _m.Called(_a0)

	var r0 user.Info
	if rf, ok := ret.Get(0).(func(string) user.Info); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(user.Info)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
