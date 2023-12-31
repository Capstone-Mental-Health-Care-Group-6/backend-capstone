// Code generated by mockery v2.35.2. DO NOT EDIT.

package mocks

import (
	transaction "FinalProject/features/transaction"

	mock "github.com/stretchr/testify/mock"
)

// TransactionServiceInterface is an autogenerated mock type for the TransactionServiceInterface type
type TransactionServiceInterface struct {
	mock.Mock
}

// CreateManualTransaction provides a mock function with given fields: newData
func (_m *TransactionServiceInterface) CreateManualTransaction(newData transaction.Transaction) (*transaction.Transaction, error) {
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

// CreateTransaction provides a mock function with given fields: newData
func (_m *TransactionServiceInterface) CreateTransaction(newData transaction.Transaction) (*transaction.Transaction, map[string]interface{}, error) {
	ret := _m.Called(newData)

	var r0 *transaction.Transaction
	var r1 map[string]interface{}
	var r2 error
	if rf, ok := ret.Get(0).(func(transaction.Transaction) (*transaction.Transaction, map[string]interface{}, error)); ok {
		return rf(newData)
	}
	if rf, ok := ret.Get(0).(func(transaction.Transaction) *transaction.Transaction); ok {
		r0 = rf(newData)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*transaction.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(transaction.Transaction) map[string]interface{}); ok {
		r1 = rf(newData)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(map[string]interface{})
		}
	}

	if rf, ok := ret.Get(2).(func(transaction.Transaction) error); ok {
		r2 = rf(newData)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// DeleteTransaction provides a mock function with given fields: id
func (_m *TransactionServiceInterface) DeleteTransaction(id int) (bool, error) {
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

// GetByIDMidtrans provides a mock function with given fields: id
func (_m *TransactionServiceInterface) GetByIDMidtrans(id string) ([]transaction.TransactionInfo, error) {
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

// GetTransaction provides a mock function with given fields: id, sort
func (_m *TransactionServiceInterface) GetTransaction(id int, sort string) ([]transaction.TransactionInfo, error) {
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

// GetTransactionByPatientID provides a mock function with given fields: id, sort
func (_m *TransactionServiceInterface) GetTransactionByPatientID(id int, sort string) ([]transaction.TransactionInfo, error) {
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

// GetTransactions provides a mock function with given fields: sort
func (_m *TransactionServiceInterface) GetTransactions(sort string) ([]transaction.TransactionInfo, error) {
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

// PaymentProofUpload provides a mock function with given fields: newData
func (_m *TransactionServiceInterface) PaymentProofUpload(newData transaction.PaymentProofDataModel) (string, error) {
	ret := _m.Called(newData)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(transaction.PaymentProofDataModel) (string, error)); ok {
		return rf(newData)
	}
	if rf, ok := ret.Get(0).(func(transaction.PaymentProofDataModel) string); ok {
		r0 = rf(newData)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(transaction.PaymentProofDataModel) error); ok {
		r1 = rf(newData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateTransaction provides a mock function with given fields: notificationPayload, newData
func (_m *TransactionServiceInterface) UpdateTransaction(notificationPayload map[string]interface{}, newData transaction.UpdateTransaction) (bool, error) {
	ret := _m.Called(notificationPayload, newData)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(map[string]interface{}, transaction.UpdateTransaction) (bool, error)); ok {
		return rf(notificationPayload, newData)
	}
	if rf, ok := ret.Get(0).(func(map[string]interface{}, transaction.UpdateTransaction) bool); ok {
		r0 = rf(notificationPayload, newData)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(map[string]interface{}, transaction.UpdateTransaction) error); ok {
		r1 = rf(notificationPayload, newData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateTransactionManual provides a mock function with given fields: newData, id
func (_m *TransactionServiceInterface) UpdateTransactionManual(newData transaction.UpdateTransactionManual, id string) (bool, error) {
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

// NewTransactionServiceInterface creates a new instance of TransactionServiceInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTransactionServiceInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *TransactionServiceInterface {
	mock := &TransactionServiceInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
