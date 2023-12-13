package service_test

// func TestGetTransactions(t *testing.T) {
// 	mockData := new(mocks.MockTransactionDataInterface)
// 	mockMidtrans := new(mocks.MockMidtransServiceInterface)

// 	expectedTransactions := []transaction.TransactionInfo{}
// 	mockData.On("GetAll").Return(expectedTransactions, nil).Once()

// 	service := service.New(mockData, nil, mockMidtrans)

// 	blank := ""
// 	result, err := service.GetTransactions(blank)

// 	assert.Nil(t, err)
// 	assert.NotNil(t, expectedTransactions, result)

// 	mockData.AssertExpectations(t)
// }

// func TestGetTransaction(t *testing.T) {
// 	mockData := new(mocks.MockTransactionDataInterface)
// 	mockMidtrans := new(mocks.MockMidtransServiceInterface)

// 	expectedTransaction := []transaction.Transaction{}

// 	mockData.On("GetByID", 1).Return(expectedTransaction, nil).Once()

// 	service := service.New(mockData, nil, mockMidtrans)
// 	blank := ""
// 	result, err := service.GetTransaction(1, blank)

// 	assert.Nil(t, err)
// 	assert.NotNil(t, expectedTransaction, result)

// 	mockData.AssertExpectations(t)
// }

// func TestCreateTransaction(t *testing.T) {

// 	mockData := new(mocks.MockTransactionDataInterface)
// 	mockMidtrans := new(mocks.MockMidtransServiceInterface)

// 	newTransaction := transaction.Transaction{UserID: 1, MidtransID: "abc123", PriceResult: 1000, PaymentStatus: 1, PaymentType: "gopay"}
// 	expectedResponse := map[string]interface{}{"key": "value"}
// 	mockData.On("Insert", newTransaction).Return(&newTransaction, nil)
// 	mockMidtrans.On("GenerateTransaction", 0, "gopay").Return(&midtrans.ChargeResponse{}, expectedResponse, nil)

// 	service := service.New(mockData, mockMidtrans)
// 	result, response, err := service.CreateTransaction(newTransaction)

// 	assert.NoError(t, err)
// 	assert.Equal(t, &newTransaction, result)
// 	assert.Equal(t, expectedResponse, response)

// 	mockData.AssertExpectations(t)
// 	mockMidtrans.AssertExpectations(t)
// }
