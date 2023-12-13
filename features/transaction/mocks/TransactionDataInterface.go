// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks

import (
	transaction "FinalProject/features/transaction"

	mock "github.com/stretchr/testify/mock"
)

// TransactionDataInterface is an autogenerated mock type for the TransactionDataInterface type
type TransactionDataInterface struct {
	mock.Mock
}

// Delete provides a mock function with given fields: id
func (_m *TransactionDataInterface) Delete(id int) (bool, error) {
	ret := _m.Called(id)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (bool, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int) bool); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields: sort
func (_m *TransactionDataInterface) GetAll(sort string) ([]transaction.TransactionInfo, error) {
	ret := _m.Called(sort)

	var r0 []transaction.TransactionInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]transaction.TransactionInfo, error)); ok {
		return rf(sort)
	}
	if rf, ok := ret.Get(0).(func(string) []transaction.TransactionInfo); ok {
		r0 = rf(sort)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]transaction.TransactionInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(sort)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAndUpdate provides a mock function with given fields: newData, id
func (_m *TransactionDataInterface) GetAndUpdate(newData transaction.UpdateTransaction, id string) (bool, error) {
	ret := _m.Called(newData, id)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(transaction.UpdateTransaction, string) (bool, error)); ok {
		return rf(newData, id)
	}
	if rf, ok := ret.Get(0).(func(transaction.UpdateTransaction, string) bool); ok {
		r0 = rf(newData, id)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(transaction.UpdateTransaction, string) error); ok {
		r1 = rf(newData, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: id, sort
func (_m *TransactionDataInterface) GetByID(id int, sort string) ([]transaction.TransactionInfo, error) {
	ret := _m.Called(id, sort)

	var r0 []transaction.TransactionInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(int, string) ([]transaction.TransactionInfo, error)); ok {
		return rf(id, sort)
	}
	if rf, ok := ret.Get(0).(func(int, string) []transaction.TransactionInfo); ok {
		r0 = rf(id, sort)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]transaction.TransactionInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(int, string) error); ok {
		r1 = rf(id, sort)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByIDMidtrans provides a mock function with given fields: id
func (_m *TransactionDataInterface) GetByIDMidtrans(id string) ([]transaction.TransactionInfo, error) {
	ret := _m.Called(id)

	var r0 []transaction.TransactionInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]transaction.TransactionInfo, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) []transaction.TransactionInfo); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]transaction.TransactionInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByPatientID provides a mock function with given fields: id, sort
func (_m *TransactionDataInterface) GetByPatientID(id int, sort string) ([]transaction.TransactionInfo, error) {
	ret := _m.Called(id, sort)

	var r0 []transaction.TransactionInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(int, string) ([]transaction.TransactionInfo, error)); ok {
		return rf(id, sort)
	}
	if rf, ok := ret.Get(0).(func(int, string) []transaction.TransactionInfo); ok {
		r0 = rf(id, sort)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]transaction.TransactionInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(int, string) error); ok {
		r1 = rf(id, sort)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Insert provides a mock function with given fields: newData
func (_m *TransactionDataInterface) Insert(newData transaction.Transaction) (*transaction.Transaction, error) {
	ret := _m.Called(newData)

	var r0 *transaction.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(transaction.Transaction) (*transaction.Transaction, error)); ok {
		return rf(newData)
	}
	if rf, ok := ret.Get(0).(func(transaction.Transaction) *transaction.Transaction); ok {
		r0 = rf(newData)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*transaction.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(transaction.Transaction) error); ok {
		r1 = rf(newData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: newData, id
func (_m *TransactionDataInterface) Update(newData transaction.UpdateTransactionManual, id int) (bool, error) {
	ret := _m.Called(newData, id)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(transaction.UpdateTransactionManual, int) (bool, error)); ok {
		return rf(newData, id)
	}
	if rf, ok := ret.Get(0).(func(transaction.UpdateTransactionManual, int) bool); ok {
		r0 = rf(newData, id)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(transaction.UpdateTransactionManual, int) error); ok {
		r1 = rf(newData, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateWithTrxID provides a mock function with given fields: newData, id
func (_m *TransactionDataInterface) UpdateWithTrxID(newData transaction.UpdateTransactionManual, id string) (bool, error) {
	ret := _m.Called(newData, id)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(transaction.UpdateTransactionManual, string) (bool, error)); ok {
		return rf(newData, id)
	}
	if rf, ok := ret.Get(0).(func(transaction.UpdateTransactionManual, string) bool); ok {
		r0 = rf(newData, id)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(transaction.UpdateTransactionManual, string) error); ok {
		r1 = rf(newData, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewTransactionDataInterface creates a new instance of TransactionDataInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTransactionDataInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *TransactionDataInterface {
	mock := &TransactionDataInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
