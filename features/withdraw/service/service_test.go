package service

import (
	"FinalProject/features/withdraw"
	"FinalProject/features/withdraw/mocks"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetAllWithdraw(t *testing.T) {
	data := mocks.NewWithdrawDataInterface(t)
	service := New(data)
	wd := []withdraw.WithdrawInfo{}

	t.Run("Server Error", func(t *testing.T) {
		data.On("GetAll").Return(nil, errors.New("Get All Process Failed")).Once()

		res, err := service.GetAllWithdraw()

		assert.Error(t, err)
		assert.Nil(t, res)
		assert.EqualError(t, err, "Get All Process Failed")
	})

	t.Run("Success Get", func(t *testing.T) {
		data.On("GetAll").Return(wd, nil).Once()

		res, err := service.GetAllWithdraw()

		assert.Nil(t, err)
		assert.NotNil(t, res)
		data.AssertExpectations(t)
	})
}

func TestGetAllWithdrawDokter(t *testing.T) {
	data := mocks.NewWithdrawDataInterface(t)
	service := New(data)
	wd := []withdraw.WithdrawInfo{}

	t.Run("Server Error", func(t *testing.T) {
		data.On("GetAllDoctor", uint(1)).Return(nil, errors.New("Get All Withdraw Doctor Failed")).Once()

		res, err := service.GetAllWithdrawDokter(uint(1))

		assert.Error(t, err)
		assert.Nil(t, res)
		assert.EqualError(t, err, "Get All Withdraw Doctor Failed")
	})

	t.Run("Success Get", func(t *testing.T) {
		data.On("GetAllDoctor", uint(1)).Return(wd, nil).Once()

		res, err := service.GetAllWithdrawDokter(uint(1))

		assert.Nil(t, err)
		assert.NotNil(t, res)
		data.AssertExpectations(t)
	})
}

func TestCreateWithdraw(t *testing.T) {
	data := mocks.NewWithdrawDataInterface(t)
	service := New(data)
	wd := withdraw.Withdraw{
		DoctorID:      1,
		ConfirmByID:   1,
		BalanceBefore: 50000,
		BalanceAfter:  20000,
		BalanceReq:    20000,
		PaymentMethod: "BCA",
		PaymentNumber: "1234",
		PaymentName:   "Hau",
		DateConfirmed: time.Now(),
		Status:        "Pending",
	}

	t.Run("Server Error", func(t *testing.T) {
		data.On("Insert", wd).Return(nil, errors.New("Insert Process Failed")).Once()

		res, err := service.CreateWithdraw(wd)

		assert.Error(t, err)
		assert.Nil(t, res)
		assert.EqualError(t, err, "Insert Process Failed")
	})

	t.Run("Less Balance Process Error", func(t *testing.T) {
		data.On("Insert", wd).Return(&wd, nil).Once()
		data.On("LessBalance", wd.DoctorID, wd.BalanceReq).Return(false, errors.New("Less Balance Process Failed")).Once()

		res, err := service.CreateWithdraw(wd)

		assert.Error(t, err)
		assert.Nil(t, res)
		assert.EqualError(t, err, "Less Balance Process Failed")
	})

	t.Run("Less Balance", func(t *testing.T) {
		data.On("Insert", wd).Return(&wd, nil).Once()
		data.On("LessBalance", wd.DoctorID, wd.BalanceReq).Return(false, nil).Once()

		res, err := service.CreateWithdraw(wd)

		assert.Error(t, err)
		assert.Nil(t, res)
		assert.EqualError(t, err, "Less Balance")
	})

	t.Run("Success Create", func(t *testing.T) {
		data.On("Insert", wd).Return(&wd, nil).Once()
		data.On("LessBalance", wd.DoctorID, wd.BalanceReq).Return(true, nil).Once()

		res, err := service.CreateWithdraw(wd)

		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, res.DoctorID, wd.DoctorID)
		assert.Equal(t, res.ConfirmByID, wd.ConfirmByID)
		assert.Equal(t, res.BalanceBefore, wd.BalanceBefore)
		assert.Equal(t, res.BalanceBefore, wd.BalanceBefore)
		data.AssertExpectations(t)
	})
}

func TestGetBalance(t *testing.T) {
	data := mocks.NewWithdrawDataInterface(t)
	service := New(data)

	t.Run("Success Get", func(t *testing.T) {
		data.On("GetBalance", uint(1)).Return(uint(20000), nil).Once()

		res, err := service.GetBalance(uint(1))

		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, res, uint(20000))
		data.AssertExpectations(t)
	})

	t.Run("Server Error", func(t *testing.T) {
		data.On("GetBalance", uint(1)).Return(uint(0), errors.New("Get Balance Error")).Once()

		res, err := service.GetBalance(uint(1))

		assert.NotNil(t, err)
		assert.Equal(t, uint(0), res)
		assert.EqualError(t, err, "Get Balance Error")
	})
}

func TestGetUserDoctor(t *testing.T) {
	data := mocks.NewWithdrawDataInterface(t)
	service := New(data)

	t.Run("Success Get", func(t *testing.T) {
		data.On("GetUserDoctor", uint(1)).Return(uint(1), nil).Once()

		res, err := service.GetUserDoctor(uint(1))

		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, res, uint(1))
		data.AssertExpectations(t)
	})

	t.Run("Server Error", func(t *testing.T) {
		data.On("GetUserDoctor", uint(1)).Return(uint(0), errors.New("Get User Doctor Error")).Once()

		res, err := service.GetUserDoctor(uint(1))

		assert.NotNil(t, err)
		assert.Equal(t, uint(0), res)
		assert.EqualError(t, err, "Get User Doctor Error")
	})
}

func TestGetByID(t *testing.T) {
	data := mocks.NewWithdrawDataInterface(t)
	service := New(data)
	var wd withdraw.WithdrawInfo

	t.Run("Server Error", func(t *testing.T) {
		data.On("GetByID", 1).Return(nil, errors.New("Get By ID Process Failed")).Once()

		res, err := service.GetByID(1)

		assert.Error(t, err)
		assert.Nil(t, res)
		assert.EqualError(t, err, "Get By ID Process Failed")
	})

	t.Run("Success Get", func(t *testing.T) {
		data.On("GetByID", 1).Return(&wd, nil).Once()

		res, err := service.GetByID(1)

		assert.Nil(t, err)
		assert.NotNil(t, res)
		data.AssertExpectations(t)
	})
}

func TestUpdateStatus(t *testing.T) {
	data := mocks.NewWithdrawDataInterface(t)
	service := New(data)
	wd := withdraw.Withdraw{
		DoctorID:      1,
		ConfirmByID:   1,
		BalanceBefore: 20000,
		BalanceAfter:  10000,
		BalanceReq:    10000,
	}
	wdInfo := withdraw.WithdrawInfo{
		ID: 1,
	}

	t.Run("Get By ID Error", func(t *testing.T) {
		data.On("GetByID", 1).Return(nil, errors.New("Get By ID Process Failed")).Once()

		res, err := service.UpdateStatus(1, wd)

		assert.Error(t, err)
		assert.Equal(t, res, false)
		assert.EqualError(t, err, "Get By ID Process Failed")
	})

	t.Run("Data Not Found", func(t *testing.T) {
		wdFail := withdraw.WithdrawInfo{
			ID: 0,
		}
		data.On("GetByID", 1).Return(&wdFail, nil).Once()

		res, err := service.UpdateStatus(1, wd)

		assert.Error(t, err)
		assert.Equal(t, res, false)
		assert.EqualError(t, err, "Data Not Found")
	})

	t.Run("Server Error", func(t *testing.T) {
		data.On("GetByID", 1).Return(&wdInfo, nil).Once()
		data.On("UpdateStatus", 1, wd).Return(false, errors.New("Update Status Process Failed")).Once()

		res, err := service.UpdateStatus(1, wd)

		assert.Error(t, err)
		assert.Equal(t, res, false)
		assert.EqualError(t, err, "Update Status Process Failed")
	})

	t.Run("Success Update", func(t *testing.T) {
		data.On("GetByID", 1).Return(&wdInfo, nil).Once()
		data.On("UpdateStatus", 1, wd).Return(true, nil).Once()

		res, err := service.UpdateStatus(1, wd)

		assert.Nil(t, err)
		assert.Equal(t, res, true)
		data.AssertExpectations(t)
	})
}
