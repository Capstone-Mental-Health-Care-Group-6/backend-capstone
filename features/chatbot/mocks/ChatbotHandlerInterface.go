// Code generated by mockery v2.35.2. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo/v4"
	mock "github.com/stretchr/testify/mock"
)

// ChatbotHandlerInterface is an autogenerated mock type for the ChatbotHandlerInterface type
type ChatbotHandlerInterface struct {
	mock.Mock
}

// CreateChatBot provides a mock function with given fields:
func (_m *ChatbotHandlerInterface) CreateChatBot() echo.HandlerFunc {
	ret := _m.Called()

	var r0 echo.HandlerFunc
	if rf, ok := ret.Get(0).(func() echo.HandlerFunc); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(echo.HandlerFunc)
		}
	}

	return r0
}

// GetAllChatBot provides a mock function with given fields:
func (_m *ChatbotHandlerInterface) GetAllChatBot() echo.HandlerFunc {
	ret := _m.Called()

	var r0 echo.HandlerFunc
	if rf, ok := ret.Get(0).(func() echo.HandlerFunc); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(echo.HandlerFunc)
		}
	}

	return r0
}

// NewChatbotHandlerInterface creates a new instance of ChatbotHandlerInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewChatbotHandlerInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *ChatbotHandlerInterface {
	mock := &ChatbotHandlerInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
