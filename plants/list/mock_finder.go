// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package list

import (
	plant "plants/plant"

	mock "github.com/stretchr/testify/mock"
)

// MockFinder is an autogenerated mock type for the Finder type
type MockFinder struct {
	mock.Mock
}

// FindAll provides a mock function with given fields: _a0
func (_m *MockFinder) FindAll(_a0 string) ([]plant.Plant, error) {
	ret := _m.Called(_a0)

	var r0 []plant.Plant
	if rf, ok := ret.Get(0).(func(string) []plant.Plant); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]plant.Plant)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindOne provides a mock function with given fields: _a0, _a1
func (_m *MockFinder) FindOne(_a0 string, _a1 string) (plant.Plant, error) {
	ret := _m.Called(_a0, _a1)

	var r0 plant.Plant
	if rf, ok := ret.Get(0).(func(string, string) plant.Plant); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(plant.Plant)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
