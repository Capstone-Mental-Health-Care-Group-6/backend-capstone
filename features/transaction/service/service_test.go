package service_test

import (
	"FinalProject/features/transaction"
	"FinalProject/features/transaction/mocks"
	"FinalProject/features/transaction/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTransactions(t *testing.T) {
	mockData := new(mocks.MockTransactionDataInterface)
	mockMidtrans := new(mocks.MockMidtransServiceInterface)

	expectedTransactions := []transaction.TransactionInfo{{UserID: 1, MidtransID: "abc123", PriceResult: 100, PaymentStatus: 1, PaymentType: "CreditCard"}}
	mockData.On("GetAll").Return(expectedTransactions, nil)

	service := service.New(mockData, mockMidtrans)

	result, err := service.GetTransactions()

	assert.NoError(t, err)
	assert.Equal(t, expectedTransactions, result)

	mockData.AssertExpectations(t)
}

func TestGetTransaction(t *testing.T) {
	mockData := new(mocks.MockTransactionDataInterface)
	mockMidtrans := new(mocks.MockMidtransServiceInterface)

	expectedTransaction := []transaction.Transaction{
		{
			TopicID:           1,
			PatientID:         1,
			DoctorID:          1,
			MethodID:          1,
			DurationID:        1,
			CounselingID:      1,
			UserID:            1,
			MidtransID:        "abc123",
			CounselingSession: 1,
			CounselingType:    "type",
			PriceMethod:       100,
			PriceDuration:     100,
			PriceCounseling:   100,
			PriceResult:       100,
			PaymentStatus:     1,
			PaymentType:       "gopay",
		},
	}

	mockData.On("GetByID", 1).Return(expectedTransaction, nil)

	service := service.New(mockData, mockMidtrans)
	result, err := service.GetTransaction(1)

	assert.NoError(t, err)
	assert.Equal(t, expectedTransaction, result)

	mockData.AssertExpectations(t)
}

// func TestCreateTransaction(t *testing.T) {

// 	mockData := new(mocks.MockTransactionDataInterface)
// 	mockMidtrans := new(mocks.MockMidtransServiceInterface)

// 	newTransaction := transaction.Transaction{UserID: 1, MidtransID: "abc123", PriceResult: 1000, PaymentStatus: 1, PaymentType: "gopay"}
// 	expectedResponse := map[string]interface{}{"key": "value"}
// 	mockData.On("Insert", newTransaction).Return(&newTransaction, nil)
// 	mockMidtrans.On("GenerateTransaction", 100, "CreditCard").Return(&midtrans.ChargeResponse{}, expectedResponse, nil)

// 	service := service.New(mockData, mockMidtrans)
// 	result, response, err := service.CreateTransaction(newTransaction)

// 	assert.NoError(t, err)
// 	assert.Equal(t, &newTransaction, result)
// 	assert.Equal(t, expectedResponse, response)

// 	mockData.AssertExpectations(t)
// 	mockMidtrans.AssertExpectations(t)
// }
