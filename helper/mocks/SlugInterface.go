// Code generated by mockery v2.35.2. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// SlugInterface is an autogenerated mock type for the SlugInterface type
type SlugInterface struct {
	mock.Mock
}

// GenerateSlug provides a mock function with given fields: name
func (_m *SlugInterface) GenerateSlug(name string) string {
	ret := _m.Called(name)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// NewSlugInterface creates a new instance of SlugInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSlugInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *SlugInterface {
	mock := &SlugInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
