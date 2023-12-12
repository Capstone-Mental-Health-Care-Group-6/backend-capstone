package service

import (
	counselingtopics "FinalProject/features/counseling_topics"
	"FinalProject/features/counseling_topics/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	data := mocks.NewCounselingTopicDataInterface(t)
	service := New(data)
	counselingDuration := []counselingtopics.CounselingTopicInfo{}

	t.Run("Server Error", func(t *testing.T) {
		data.On("GetAll").Return(nil, errors.New("Get All Process Failed")).Once()

		res, err := service.GetAll()

		assert.Error(t, err)
		assert.Nil(t, res)
		assert.EqualError(t, err, "Get All Process Failed")
	})

	t.Run("Success Get", func(t *testing.T) {
		data.On("GetAll").Return(counselingDuration, nil).Once()

		res, err := service.GetAll()

		assert.Nil(t, err)
		assert.NotNil(t, res)
		data.AssertExpectations(t)
	})
}

func TestGetByID(t *testing.T) {
	data := mocks.NewCounselingTopicDataInterface(t)
	service := New(data)
	counselingDuration := []counselingtopics.CounselingTopicInfo{}

	t.Run("Server Error", func(t *testing.T) {
		data.On("GetByID", 1).Return(nil, errors.New("Get By ID Process Failed")).Once()

		res, err := service.GetByID(1)

		assert.Error(t, err)
		assert.Nil(t, res)
		assert.EqualError(t, err, "Get By ID Process Failed")
	})

	t.Run("Success Get", func(t *testing.T) {
		data.On("GetByID", 1).Return(counselingDuration, nil).Once()

		res, err := service.GetByID(1)

		assert.Nil(t, err)
		assert.NotNil(t, res)
		data.AssertExpectations(t)
	})
}
