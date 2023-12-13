package service

import (
	counselingsession "FinalProject/features/counseling_session"
	"FinalProject/features/counseling_session/mocks"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetAllCounseling(t *testing.T) {
	data := mocks.NewCounselingSessionDataInterface(t)
	service := New(data)
	dataCounseling := []counselingsession.CounselingSession{}

	t.Run("Server Error", func(t *testing.T) {
		data.On("GetAll").Return(nil, errors.New("Get All Process Failed")).Once()

		res, err := service.GetAllCounseling()

		assert.Error(t, err)
		assert.Nil(t, res)
		assert.EqualError(t, err, "Get All Process Failed")
	})

	t.Run("Success Get", func(t *testing.T) {
		data.On("GetAll").Return(dataCounseling, nil).Once()

		res, err := service.GetAllCounseling()

		assert.Nil(t, err)
		assert.NotNil(t, res)
		data.AssertExpectations(t)
	})
}

func TestGetByIDCounseling(t *testing.T) {
	data := mocks.NewCounselingSessionDataInterface(t)
	service := New(data)
	var dataCounseling counselingsession.CounselingSession

	t.Run("Server Error", func(t *testing.T) {
		data.On("GetById", 1).Return(nil, errors.New("Get By ID Process Failed")).Once()

		res, err := service.GetCounseling(1)

		assert.Error(t, err)
		assert.Nil(t, res)
		assert.EqualError(t, err, "Get By ID Process Failed")
	})

	t.Run("Success Get", func(t *testing.T) {
		data.On("GetById", 1).Return(&dataCounseling, nil).Once()

		res, err := service.GetCounseling(1)

		assert.Nil(t, err)
		assert.NotNil(t, res)
		data.AssertExpectations(t)
	})
}

func TestGetAllByUserIDCounseling(t *testing.T) {
	data := mocks.NewCounselingSessionDataInterface(t)
	service := New(data)
	dataCounseling := []counselingsession.CounselingSession{}

	t.Run("Server Error", func(t *testing.T) {
		data.On("GetAllCounselingByUserID", 1).Return(nil, errors.New("Get All By ID Process Failed")).Once()

		res, err := service.GetAllCounselingByUserID(1)

		assert.Error(t, err)
		assert.Nil(t, res)
		assert.EqualError(t, err, "Get All By ID Process Failed")
	})

	t.Run("Success Get", func(t *testing.T) {
		data.On("GetAllCounselingByUserID", 1).Return(dataCounseling, nil).Once()

		res, err := service.GetAllCounselingByUserID(1)

		assert.Nil(t, err)
		assert.NotNil(t, res)
		data.AssertExpectations(t)
	})
}

func TestCreateCounseling(t *testing.T) {
	data := mocks.NewCounselingSessionDataInterface(t)
	service := New(data)
	dataCounseling := counselingsession.CounselingSession{
		TransactionID: 1,
		Date:          time.Now(),
		Time:          time.Now(),
		Duration:      120,
		Status:        "Finished",
	}

	t.Run("Server Error", func(t *testing.T) {
		data.On("Create", dataCounseling).Return(nil, errors.New("Create Process Failed")).Once()

		res, err := service.CreateCounseling(dataCounseling)

		assert.Error(t, err)
		assert.Nil(t, res)
		assert.EqualError(t, err, "Create Process Failed")
	})

	t.Run("Success Create", func(t *testing.T) {
		data.On("Create", dataCounseling).Return(&dataCounseling, nil).Once()

		res, err := service.CreateCounseling(dataCounseling)

		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, res.TransactionID, dataCounseling.TransactionID)
		assert.Equal(t, res.Date, dataCounseling.Date)
		assert.Equal(t, res.Time, dataCounseling.Time)
		assert.Equal(t, res.Duration, dataCounseling.Duration)
		assert.Equal(t, res.Status, dataCounseling.Status)
		data.AssertExpectations(t)
	})
}

func TestUpdateCounseling(t *testing.T) {
	data := mocks.NewCounselingSessionDataInterface(t)
	service := New(data)
	dataCounseling := counselingsession.CounselingSession{
		TransactionID: 1,
		Date:          time.Now(),
		Time:          time.Now(),
		Duration:      120,
		Status:        "Finished",
	}

	t.Run("Server Error", func(t *testing.T) {
		data.On("Update", 1, dataCounseling).Return(false, errors.New("Update Process Failed")).Once()

		res, err := service.UpdateCounseling(1, dataCounseling)

		assert.Error(t, err)
		assert.Equal(t, false, res)
		assert.EqualError(t, err, "Update Process Failed")
	})

	t.Run("Success Update", func(t *testing.T) {
		data.On("Update", 1, dataCounseling).Return(true, nil).Once()

		res, err := service.UpdateCounseling(1, dataCounseling)

		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, true, res)
		data.AssertExpectations(t)
	})
}

func TestDeleteCounseling(t *testing.T) {
	data := mocks.NewCounselingSessionDataInterface(t)
	service := New(data)

	t.Run("Server Error", func(t *testing.T) {
		data.On("Delete", 1).Return(false, errors.New("Delete Process Failed")).Once()

		res, err := service.DeleteCounseling(1)

		assert.Error(t, err)
		assert.Equal(t, false, res)
		assert.EqualError(t, err, "Delete Process Failed")
	})

	t.Run("Success Get", func(t *testing.T) {
		data.On("Delete", 1).Return(true, nil).Once()

		res, err := service.DeleteCounseling(1)

		assert.Nil(t, err)
		assert.Equal(t, true, res)
		data.AssertExpectations(t)
	})
}
