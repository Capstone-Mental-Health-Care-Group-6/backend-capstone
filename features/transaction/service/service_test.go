package service

import (
	"FinalProject/features/transaction"
	"FinalProject/features/transaction/mocks"
	mockUtil "FinalProject/utils/mocks"
	"errors"
	"mime/multipart"
	"testing"

	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/stretchr/testify/assert"
)

func TestGetTransactions(t *testing.T) {
	data := mocks.NewTransactionDataInterface(t)
	cld := mockUtil.NewCloudinaryInterface(t)
	mt := mockUtil.NewMidtransService(t)
	service := New(data, cld, mt)
	trans := []transaction.TransactionInfo{}

	t.Run("Server Error", func(t *testing.T) {
		data.On("GetAll", "asc").Return(nil, errors.New("Get All Transactions Failed")).Once()

		res, err := service.GetTransactions("asc")

		assert.Error(t, err)
		assert.Nil(t, res)
		assert.EqualError(t, err, "Get All Transactions Failed")
	})

	t.Run("Success Get", func(t *testing.T) {
		data.On("GetAll", "asc").Return(trans, nil).Once()

		res, err := service.GetTransactions("asc")

		assert.Nil(t, err)
		assert.NotNil(t, res)
		data.AssertExpectations(t)
	})
}

func TestGetTransaction(t *testing.T) {
	data := mocks.NewTransactionDataInterface(t)
	cld := mockUtil.NewCloudinaryInterface(t)
	mt := mockUtil.NewMidtransService(t)
	service := New(data, cld, mt)
	trans := []transaction.TransactionInfo{}

	t.Run("Server Error", func(t *testing.T) {
		data.On("GetByID", 1, "asc").Return(nil, errors.New("Get By User ID Process Failed")).Once()

		res, err := service.GetTransaction(1, "asc")

		assert.Error(t, err)
		assert.Nil(t, res)
		assert.EqualError(t, err, "Get By User ID Process Failed")
	})

	t.Run("Success Get", func(t *testing.T) {
		data.On("GetByID", 1, "asc").Return(trans, nil).Once()

		res, err := service.GetTransaction(1, "asc")

		assert.Nil(t, err)
		assert.NotNil(t, res)
		data.AssertExpectations(t)
	})
}

func TestGetTransactionByPatientID(t *testing.T) {
	data := mocks.NewTransactionDataInterface(t)
	cld := mockUtil.NewCloudinaryInterface(t)
	mt := mockUtil.NewMidtransService(t)
	service := New(data, cld, mt)
	trans := []transaction.TransactionInfo{}

	t.Run("Server Error", func(t *testing.T) {
		data.On("GetByPatientID", 1, "asc").Return(nil, errors.New("Get By Patient ID Process Failed")).Once()

		res, err := service.GetTransactionByPatientID(1, "asc")

		assert.Error(t, err)
		assert.Nil(t, res)
		assert.EqualError(t, err, "Get By Patient ID Process Failed")
	})

	t.Run("Success Get", func(t *testing.T) {
		data.On("GetByPatientID", 1, "asc").Return(trans, nil).Once()

		res, err := service.GetTransactionByPatientID(1, "asc")

		assert.Nil(t, err)
		assert.NotNil(t, res)
		data.AssertExpectations(t)
	})
}

func TestGetByIDMidtrans(t *testing.T) {
	data := mocks.NewTransactionDataInterface(t)
	cld := mockUtil.NewCloudinaryInterface(t)
	mt := mockUtil.NewMidtransService(t)
	service := New(data, cld, mt)
	trans := []transaction.TransactionInfo{}

	t.Run("Server Error", func(t *testing.T) {
		data.On("GetByIDMidtrans", "asc").Return(nil, errors.New("Get By ID Midtrans Process Failed")).Once()

		res, err := service.GetByIDMidtrans("asc")

		assert.Error(t, err)
		assert.Nil(t, res)
		assert.EqualError(t, err, "Get By ID Midtrans Process Failed")
	})

	t.Run("Success Get", func(t *testing.T) {
		data.On("GetByIDMidtrans", "asc").Return(trans, nil).Once()

		res, err := service.GetByIDMidtrans("asc")

		assert.Nil(t, err)
		assert.NotNil(t, res)
		data.AssertExpectations(t)
	})
}

func TestCreateTransaction(t *testing.T) {
	data := mocks.NewTransactionDataInterface(t)
	cld := mockUtil.NewCloudinaryInterface(t)
	mt := mockUtil.NewMidtransService(t)
	service := New(data, cld, mt)
	trans := transaction.Transaction{
		TopicID:           1,
		PatientID:         1,
		DoctorID:          1,
		MethodID:          1,
		DurationID:        1,
		CounselingID:      1,
		UserID:            1,
		MidtransID:        "",
		CounselingSession: 1,
		CounselingType:    "type",
		PriceMethod:       10000,
		PriceDuration:     10000,
		PriceCounseling:   10000,
		PaymentType:       "BCA",
	}
	var totalPrice = trans.PriceMethod + trans.PriceDuration + trans.PriceCounseling
	var coreApi coreapi.ChargeResponse

	t.Run("Generate Error", func(t *testing.T) {
		mt.On("GenerateTransaction", int(totalPrice), trans.PaymentType).Return(nil, nil, errors.New("Generate Transaction Failed")).Once()

		res1, res2, err := service.CreateTransaction(trans)

		assert.Error(t, err)
		assert.Nil(t, res1)
		assert.Nil(t, res2)
		assert.EqualError(t, err, "Generate Transaction Failed")
	})

	t.Run("Server Error", func(t *testing.T) {
		var response map[string]interface{}
		mt.On("GenerateTransaction", int(totalPrice), trans.PaymentType).Return(&coreApi, response, nil).Once()
		data.On("Insert", trans).Return(nil, errors.New("Insert Process Failed")).Once()

		res1, res2, err := service.CreateTransaction(trans)

		assert.Error(t, err)
		assert.Nil(t, res1)
		assert.Nil(t, res2)
		assert.EqualError(t, err, "Insert Process Failed")
	})

	t.Run("Success Insert", func(t *testing.T) {
		response := make(map[string]interface{})

		mt.On("GenerateTransaction", int(totalPrice), trans.PaymentType).Return(&coreApi, response, nil).Once()
		data.On("Insert", trans).Return(&trans, nil).Once()

		res1, res2, err := service.CreateTransaction(trans)

		assert.Nil(t, err)
		assert.NotNil(t, res1)
		assert.NotNil(t, res2)
		mt.AssertExpectations(t)
		data.AssertExpectations(t)
	})
}

func TestCreateManualTransaction(t *testing.T) {
	data := mocks.NewTransactionDataInterface(t)
	cld := mockUtil.NewCloudinaryInterface(t)
	mt := mockUtil.NewMidtransService(t)
	service := New(data, cld, mt)
	trans := transaction.Transaction{
		TopicID:           1,
		PatientID:         1,
		DoctorID:          1,
		MethodID:          1,
		DurationID:        1,
		CounselingID:      1,
		UserID:            1,
		CounselingSession: 1,
		CounselingType:    "type",
		PriceMethod:       10000,
		PriceDuration:     10000,
		PriceCounseling:   10000,
		PaymentType:       "BCA",
	}

	t.Run("Server Error", func(t *testing.T) {
		data.On("Insert", trans).Return(nil, errors.New("Insert Process Failed")).Once()

		res, err := service.CreateManualTransaction(trans)

		assert.Error(t, err)
		assert.Nil(t, res)
		assert.EqualError(t, err, "Insert Process Failed")
	})

	t.Run("Success Insert", func(t *testing.T) {
		data.On("Insert", trans).Return(&trans, nil).Once()

		res, err := service.CreateManualTransaction(trans)

		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, res.TopicID, trans.TopicID)
		assert.Equal(t, res.PatientID, trans.PatientID)
		assert.Equal(t, res.DoctorID, trans.DoctorID)
		assert.Equal(t, res.MethodID, trans.MethodID)
		assert.Equal(t, res.DurationID, trans.DurationID)
		assert.Equal(t, res.CounselingID, trans.CounselingID)
		data.AssertExpectations(t)
	})
}

func TestUpdateTransaction(t *testing.T) {
	data := mocks.NewTransactionDataInterface(t)
	cld := mockUtil.NewCloudinaryInterface(t)
	mt := mockUtil.NewMidtransService(t)
	service := New(data, cld, mt)
	var notificationPayload map[string]interface{}
	update := transaction.UpdateTransaction{
		PaymentStatus: 5,
	}

	t.Run("Transaction Status Failed", func(t *testing.T) {
		mt.On("TransactionStatus", notificationPayload).Return(0, "", errors.New("Transaction Status Failed")).Once()

		res, err := service.UpdateTransaction(notificationPayload, update)

		assert.Error(t, err)
		assert.Equal(t, res, false)
		assert.EqualError(t, err, "Transaction Status Failed")
	})

	t.Run("Server Error", func(t *testing.T) {
		mt.On("TransactionStatus", notificationPayload).Return(5, "order-id", nil).Once()
		data.On("GetAndUpdate", update, "order-id").Return(false, errors.New("Update Process Failed")).Once()

		res, err := service.UpdateTransaction(notificationPayload, update)

		assert.Error(t, err)
		assert.Equal(t, res, false)
		assert.EqualError(t, err, "Update Process Failed")
	})

	t.Run("Success Update", func(t *testing.T) {
		mt.On("TransactionStatus", notificationPayload).Return(5, "order-id", nil).Once()
		data.On("GetAndUpdate", update, "order-id").Return(true, nil).Once()

		res, err := service.UpdateTransaction(notificationPayload, update)

		assert.Nil(t, err)
		assert.Equal(t, res, true)
		mt.AssertExpectations(t)
		data.AssertExpectations(t)
	})
}

func TestUpdateTransactionManual(t *testing.T) {
	data := mocks.NewTransactionDataInterface(t)
	cld := mockUtil.NewCloudinaryInterface(t)
	mt := mockUtil.NewMidtransService(t)
	service := New(data, cld, mt)
	trans := transaction.UpdateTransactionManual{
		UserID:          1,
		PriceMethod:     10000,
		PriceDuration:   10000,
		PriceCounseling: 10000,
		PriceResult:     30000,
		PaymentStatus:   5,
		PaymentType:     "Manual",
	}

	t.Run("Update Process Fail Text", func(t *testing.T) {
		data.On("UpdateWithTrxID", trans, "MASDASD").Return(false, errors.New("Update Process Failed")).Once()

		res, err := service.UpdateTransactionManual(trans, "MASDASD")

		assert.Error(t, err)
		assert.Equal(t, res, false)
		assert.EqualError(t, err, "Update Process Failed")
	})

	t.Run("Success Update Text", func(t *testing.T) {
		data.On("UpdateWithTrxID", trans, "MASDASD").Return(true, nil).Once()

		res, err := service.UpdateTransactionManual(trans, "MASDASD")

		assert.Nil(t, err)
		assert.Equal(t, res, true)
		data.AssertExpectations(t)
	})

	t.Run("Update Process Fail Number", func(t *testing.T) {
		data.On("Update", trans, 1).Return(false, errors.New("Update Process Failed")).Once()

		res, err := service.UpdateTransactionManual(trans, "1")

		assert.Error(t, err)
		assert.Equal(t, res, false)
		assert.EqualError(t, err, "Update Process Failed")
	})

	t.Run("Success Update Number", func(t *testing.T) {
		data.On("Update", trans, 1).Return(true, nil).Once()

		res, err := service.UpdateTransactionManual(trans, "1")

		assert.Nil(t, err)
		assert.Equal(t, res, true)
		data.AssertExpectations(t)
	})
}

func TestDeleteTransaction(t *testing.T) {
	data := mocks.NewTransactionDataInterface(t)
	cld := mockUtil.NewCloudinaryInterface(t)
	mt := mockUtil.NewMidtransService(t)
	service := New(data, cld, mt)

	t.Run("Server Error", func(t *testing.T) {
		data.On("Delete", 1).Return(false, errors.New("Delete Process Failed")).Once()

		res, err := service.DeleteTransaction(1)

		assert.Error(t, err)
		assert.Equal(t, res, false)
		assert.EqualError(t, err, "Delete Process Failed")
	})

	t.Run("Success Delete", func(t *testing.T) {
		data.On("Delete", 1).Return(true, nil).Once()

		res, err := service.DeleteTransaction(1)

		assert.Nil(t, err)
		assert.Equal(t, res, true)
		data.AssertExpectations(t)
	})
}

func TestPaymentProofUpload(t *testing.T) {
	data := mocks.NewTransactionDataInterface(t)
	cld := mockUtil.NewCloudinaryInterface(t)
	mt := mockUtil.NewMidtransService(t)
	service := New(data, cld, mt)
	var mockFile multipart.File
	trans := transaction.PaymentProofDataModel{
		PaymentProofPhoto: mockFile,
	}

	t.Run("Upload Failed", func(t *testing.T) {
		cld.On("UploadImageHelper", trans.PaymentProofPhoto).Return("", errors.New("Upload Payment Proof Failed")).Once()

		res, err := service.PaymentProofUpload(trans)

		assert.Error(t, err)
		assert.NotNil(t, res)
		assert.EqualError(t, err, "Upload Payment Proof Failed")
	})

	t.Run("Success Upload", func(t *testing.T) {
		cld.On("UploadImageHelper", trans.PaymentProofPhoto).Return("https://", nil).Once()

		res, err := service.PaymentProofUpload(trans)

		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, res, "https://")
		cld.AssertExpectations(t)
	})
}
