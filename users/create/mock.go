// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package create

import (
	user "users/user"

	mock "github.com/stretchr/testify/mock"
)

// MockCreateService is an autogenerated mock type for the CreateService type
type MockCreateService struct {
	mock.Mock
}

// Create provides a mock function with given fields: _a0
func (_m *MockCreateService) Create(_a0 user.User) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(user.User) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
