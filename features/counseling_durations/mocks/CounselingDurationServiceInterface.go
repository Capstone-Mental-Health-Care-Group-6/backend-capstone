// Code generated by mockery v2.35.2. DO NOT EDIT.

package mocks

import (
	counselingdurations "FinalProject/features/counseling_durations"

	mock "github.com/stretchr/testify/mock"
)

// CounselingDurationServiceInterface is an autogenerated mock type for the CounselingDurationServiceInterface type
type CounselingDurationServiceInterface struct {
	mock.Mock
}

// GetAll provides a mock function with given fields:
func (_m *CounselingDurationServiceInterface) GetAll() ([]counselingdurations.CounselingDurationInfo, error) {
	ret := _m.Called()

	var r0 []counselingdurations.CounselingDurationInfo
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]counselingdurations.CounselingDurationInfo, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []counselingdurations.CounselingDurationInfo); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]counselingdurations.CounselingDurationInfo)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: id
func (_m *CounselingDurationServiceInterface) GetByID(id int) ([]counselingdurations.CounselingDurationInfo, error) {
	ret := _m.Called(id)

	var r0 []counselingdurations.CounselingDurationInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(int) ([]counselingdurations.CounselingDurationInfo, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int) []counselingdurations.CounselingDurationInfo); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]counselingdurations.CounselingDurationInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewCounselingDurationServiceInterface creates a new instance of CounselingDurationServiceInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCounselingDurationServiceInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *CounselingDurationServiceInterface {
	mock := &CounselingDurationServiceInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}