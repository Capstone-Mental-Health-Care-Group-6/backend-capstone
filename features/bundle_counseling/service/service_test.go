package service

import (
	bundlecounseling "FinalProject/features/bundle_counseling"
	"FinalProject/features/bundle_counseling/mocks"
	mockUtil "FinalProject/utils/mocks"
	"errors"
	"mime/multipart"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllBundle(t *testing.T) {
	data := mocks.NewBundleCounselingDataInterface(t)
	cld := mockUtil.NewCloudinaryInterface(t)
	service := New(data, cld)
	dataBundle := []bundlecounseling.BundleCounselingInfo{}

	t.Run("Server Error", func(t *testing.T) {
		data.On("GetAll").Return(nil, errors.New("Get All Process Failed")).Once()

		res, err := service.GetAllBundle()

		assert.Error(t, err)
		assert.EqualError(t, err, "Get All Process Failed")
		assert.Nil(t, res)
	})

	t.Run("Success Get", func(t *testing.T) {
		data.On("GetAll").Return(dataBundle, nil).Once()

		res, err := service.GetAllBundle()

		assert.Nil(t, err)
		assert.NotNil(t, res)
		data.AssertExpectations(t)
	})
}

func TestCreateBundle(t *testing.T) {
	data := mocks.NewBundleCounselingDataInterface(t)
	cld := mockUtil.NewCloudinaryInterface(t)
	service := New(data, cld)
	var file multipart.File
	dataBundle := bundlecounseling.BundleCounseling{
		Name:         "name",
		Sessions:     1,
		Type:         "type",
		Price:        120000,
		Description:  "description",
		ActivePriode: 1,
		Avatar:       "https://",
	}
	fileBundle := bundlecounseling.BundleCounselingFile{
		Avatar: file,
	}

	t.Run("Upload Failed", func(t *testing.T) {
		cld.On("UploadImageHelper", fileBundle.Avatar).Return("", errors.New("Upload Failed")).Once()

		res, err := service.CreateBundle(dataBundle, fileBundle)

		assert.Error(t, err)
		assert.Nil(t, res)
		assert.EqualError(t, err, "Upload Failed")
	})

	t.Run("Server Error", func(t *testing.T) {
		data.On("Create", dataBundle).Return(nil, errors.New("Create Process Failed")).Once()
		cld.On("UploadImageHelper", fileBundle.Avatar).Return(dataBundle.Avatar, nil).Once()

		res, err := service.CreateBundle(dataBundle, fileBundle)

		assert.Error(t, err)
		assert.Nil(t, res)
		assert.EqualError(t, err, "Create Process Failed")
	})

	t.Run("Success Create", func(t *testing.T) {
		data.On("Create", dataBundle).Return(&dataBundle, nil).Once()
		cld.On("UploadImageHelper", fileBundle.Avatar).Return(dataBundle.Avatar, nil).Once()

		res, err := service.CreateBundle(dataBundle, fileBundle)

		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, res.Name, dataBundle.Name)
		assert.Equal(t, res.Description, dataBundle.Description)
		assert.Equal(t, res.Sessions, dataBundle.Sessions)
		assert.Equal(t, res.Type, dataBundle.Type)
		assert.Equal(t, res.Price, dataBundle.Price)
		assert.Equal(t, res.ActivePriode, dataBundle.ActivePriode)
		assert.Equal(t, res.Avatar, dataBundle.Avatar)
		data.AssertExpectations(t)
		cld.AssertExpectations(t)
	})
}

func TestGetBundle(t *testing.T) {
	data := mocks.NewBundleCounselingDataInterface(t)
	cld := mockUtil.NewCloudinaryInterface(t)
	service := New(data, cld)
	var dataBundle bundlecounseling.BundleCounseling

	t.Run("Server Error", func(t *testing.T) {
		data.On("GetById", 1).Return(nil, errors.New("Get By ID Process Failed")).Once()

		res, err := service.GetBundle(1)

		assert.Error(t, err)
		assert.Nil(t, res)
		assert.EqualError(t, err, "Get By ID Process Failed")
	})

	t.Run("Success Get", func(t *testing.T) {
		data.On("GetById", 1).Return(&dataBundle, nil).Once()

		res, err := service.GetBundle(1)

		assert.Nil(t, err)
		assert.NotNil(t, res)
		data.AssertExpectations(t)
	})
}

func TestUpdateBundle(t *testing.T) {
	data := mocks.NewBundleCounselingDataInterface(t)
	cld := mockUtil.NewCloudinaryInterface(t)
	service := New(data, cld)
	var file multipart.File
	dataBundle := bundlecounseling.BundleCounseling{
		Name:         "name",
		Sessions:     1,
		Type:         "type",
		Price:        120000,
		Description:  "description",
		ActivePriode: 1,
	}
	fileBundle := bundlecounseling.BundleCounselingFile{
		Avatar: file,
	}

	t.Run("Server Error", func(t *testing.T) {
		data.On("Update", 1, dataBundle).Return(false, errors.New("Update Process Failed")).Once()

		res, err := service.UpdateBundle(1, dataBundle, fileBundle)

		assert.Error(t, err)
		assert.Equal(t, res, false)
		assert.EqualError(t, err, "Update Process Failed")
	})

	t.Run("Success Update", func(t *testing.T) {
		// cld.On("UploadImageHelper", fileBundle.Avatar).Return(dataBundle.Avatar, nil).Once()
		data.On("Update", 1, dataBundle).Return(true, nil).Once()

		res, err := service.UpdateBundle(1, dataBundle, fileBundle)

		assert.Nil(t, err)
		assert.Equal(t, res, true)
		data.AssertExpectations(t)
		cld.AssertExpectations(t)
	})
}

func TestDeleteBundle(t *testing.T) {
	data := mocks.NewBundleCounselingDataInterface(t)
	cld := mockUtil.NewCloudinaryInterface(t)
	service := New(data, cld)

	t.Run("Server Error", func(t *testing.T) {
		data.On("Delete", 1).Return(false, errors.New("Delete Process Failed")).Once()

		res, err := service.DeleteBundle(1)

		assert.Error(t, err)
		assert.Equal(t, res, false)
		assert.EqualError(t, err, "Delete Process Failed")
	})

	t.Run("Success Delete", func(t *testing.T) {
		data.On("Delete", 1).Return(true, nil).Once()

		res, err := service.DeleteBundle(1)

		assert.Nil(t, err)
		assert.Equal(t, res, true)
		data.AssertExpectations(t)
	})
}

func TestGetAllBundleFilter(t *testing.T) {
	data := mocks.NewBundleCounselingDataInterface(t)
	cld := mockUtil.NewCloudinaryInterface(t)
	service := New(data, cld)
	bundle := []bundlecounseling.BundleCounselingInfo{}

	t.Run("Error Harga Metode", func(t *testing.T) {
		data.On("HargaMetode", 1).Return(uint(1), errors.New("Get Harga Metode Failed")).Once()

		res, err := service.GetAllBundleFilter("INSTAN", 1, 1)

		assert.Error(t, err)
		assert.NotNil(t, res)
		assert.EqualError(t, err, "Get Harga Metode Failed")
	})

	t.Run("Error Harga Durasi", func(t *testing.T) {
		data.On("HargaMetode", 1).Return(uint(1), nil).Once()
		data.On("HargaDurasi", 1).Return(uint(1), errors.New("Get Harga Durasi Failed")).Once()

		res, err := service.GetAllBundleFilter("INSTAN", 1, 1)

		assert.Error(t, err)
		assert.NotNil(t, res)
		assert.EqualError(t, err, "Get Harga Durasi Failed")
	})

	t.Run("Error Get All", func(t *testing.T) {
		data.On("HargaMetode", 1).Return(uint(1), nil).Once()
		data.On("HargaDurasi", 1).Return(uint(1), nil).Once()
		data.On("GetAllFilter", "INSTAN").Return(nil, errors.New("Get All Filter Process Failed")).Once()

		res, err := service.GetAllBundleFilter("INSTAN", 1, 1)

		assert.Error(t, err)
		assert.NotNil(t, res)
		assert.EqualError(t, err, "Get All Filter Process Failed")
	})

	t.Run("Success Get", func(t *testing.T) {
		data.On("HargaMetode", 1).Return(uint(1), nil).Once()
		data.On("HargaDurasi", 1).Return(uint(1), nil).Once()
		data.On("GetAllFilter", "INSTAN").Return(bundle, nil).Once()

		res, err := service.GetAllBundleFilter("INSTAN", 1, 1)

		assert.Nil(t, err)
		assert.NotNil(t, res)
		data.AssertExpectations(t)
	})
}
