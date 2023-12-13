// Code generated by mockery v2.35.2. DO NOT EDIT.

package mocks

import (
	chatbot "FinalProject/features/chatbot"

	mock "github.com/stretchr/testify/mock"
)

// ChatbotDataInterface is an autogenerated mock type for the ChatbotDataInterface type
type ChatbotDataInterface struct {
	mock.Mock
}

// GetAllChatBot provides a mock function with given fields: user_id
func (_m *ChatbotDataInterface) GetAllChatBot(user_id int) ([]chatbot.Chatbot, error) {
	ret := _m.Called(user_id)

	var r0 []chatbot.Chatbot
	var r1 error
	if rf, ok := ret.Get(0).(func(int) ([]chatbot.Chatbot, error)); ok {
		return rf(user_id)
	}
	if rf, ok := ret.Get(0).(func(int) []chatbot.Chatbot); ok {
		r0 = rf(user_id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]chatbot.Chatbot)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(user_id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertChatBot provides a mock function with given fields: input
func (_m *ChatbotDataInterface) InsertChatBot(input chatbot.Chatbot) (chatbot.Chatbot, error) {
	ret := _m.Called(input)

	var r0 chatbot.Chatbot
	var r1 error
	if rf, ok := ret.Get(0).(func(chatbot.Chatbot) (chatbot.Chatbot, error)); ok {
		return rf(input)
	}
	if rf, ok := ret.Get(0).(func(chatbot.Chatbot) chatbot.Chatbot); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Get(0).(chatbot.Chatbot)
	}

	if rf, ok := ret.Get(1).(func(chatbot.Chatbot) error); ok {
		r1 = rf(input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewChatbotDataInterface creates a new instance of ChatbotDataInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewChatbotDataInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *ChatbotDataInterface {
	mock := &ChatbotDataInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
