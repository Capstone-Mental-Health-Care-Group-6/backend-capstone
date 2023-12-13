// mocks/mocks.go

package mocks

import (
	"FinalProject/features/transaction"
	"FinalProject/utils/midtrans"

	"github.com/stretchr/testify/mock"
)

type MockTransactionDataInterface struct {
	mock.Mock
}

// GetByPatientID implements transaction.TransactionDataInterface.
func (*MockTransactionDataInterface) GetByPatientID(id int, sort string) ([]transaction.TransactionInfo, error) {
	panic("unimplemented")
}

type MockMidtransServiceInterface struct {
	mock.Mock
}

func (m *MockTransactionDataInterface) GetAll(sort string) ([]transaction.TransactionInfo, error) {
	args := m.Called()
	return args.Get(0).([]transaction.TransactionInfo), args.Error(1)
}

func (m *MockTransactionDataInterface) GetByID(id int, sort string) ([]transaction.TransactionInfo, error) {
	args := m.Called(id)
	return args.Get(0).([]transaction.TransactionInfo), args.Error(1)
}

func (m *MockTransactionDataInterface) GetByIDMidtrans(id string) ([]transaction.TransactionInfo, error) {
	args := m.Called(id)
	return args.Get(0).([]transaction.TransactionInfo), args.Error(1)
}

func (m *MockTransactionDataInterface) Insert(newData transaction.Transaction) (*transaction.Transaction, error) {
	args := m.Called(newData)
	return args.Get(0).(*transaction.Transaction), args.Error(1)
}

func (m *MockTransactionDataInterface) GetAndUpdate(newData transaction.UpdateTransaction, id string) (bool, error) {
	args := m.Called(newData, id)
	return args.Bool(0), args.Error(1)
}

func (m *MockTransactionDataInterface) Update(newData transaction.UpdateTransactionManual, id int) (bool, error) {
	args := m.Called(newData, id)
	return args.Bool(0), args.Error(1)
}

func (m *MockTransactionDataInterface) UpdateWithTrxID(newData transaction.UpdateTransactionManual, id string) (bool, error) {
	args := m.Called(newData, id)
	return args.Bool(0), args.Error(1)
}

func (m *MockTransactionDataInterface) Delete(id int) (bool, error) {
	args := m.Called(id)
	return args.Bool(0), args.Error(1)
}

func (m *MockMidtransServiceInterface) GenerateTransaction(amount int, paymentType string) (*midtrans.ChargeResponse, map[string]interface{}, error) {
	args := m.Called(amount, paymentType)
	return args.Get(0).(*midtrans.ChargeResponse), args.Get(1).(map[string]interface{}), args.Error(2)
}

func (m *MockMidtransServiceInterface) TransactionStatus(notificationPayload map[string]interface{}) (int, string, error) {
	args := m.Called(notificationPayload)
	return args.Int(0), args.String(1), args.Error(2)
}
