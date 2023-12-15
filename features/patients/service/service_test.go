package service

import (
	"FinalProject/features/patients"
	"FinalProject/features/patients/mocks"
	mockHelper "FinalProject/helper/mocks"
	mockUtil "FinalProject/utils/mocks"
	"errors"
	"mime/multipart"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPatients(t *testing.T) {
	data := mocks.NewPatientDataInterface(t)
	cld := mockUtil.NewCloudinaryInterface(t)
	enkrip := mockHelper.NewHashInterface(t)
	jwt := mockHelper.NewJWTInterface(t)
	service := NewPatient(data, cld, jwt, enkrip)
	patient := []patients.Patientdetail{}

	t.Run("Server Error", func(t *testing.T) {
		data.On("GetAll", "status", "name").Return(nil, errors.New("Get All Process Failed")).Once()

		res, err := service.GetPatients("status", "name")

		assert.Error(t, err)
		assert.Nil(t, res)
		assert.EqualError(t, err, "Get All Process Failed")
	})

	t.Run("Success Get", func(t *testing.T) {
		data.On("GetAll", "status", "name").Return(patient, nil).Once()

		res, err := service.GetPatients("status", "name")

		assert.Nil(t, err)
		assert.NotNil(t, res)
		data.AssertExpectations(t)
	})
}

func TestGetPatient(t *testing.T) {
	data := mocks.NewPatientDataInterface(t)
	cld := mockUtil.NewCloudinaryInterface(t)
	enkrip := mockHelper.NewHashInterface(t)
	jwt := mockHelper.NewJWTInterface(t)
	service := NewPatient(data, cld, jwt, enkrip)
	var patient patients.Patientdetail

	t.Run("Server Error", func(t *testing.T) {
		data.On("GetByID", 1).Return(patient, errors.New("Get By ID Process Failed")).Once()

		res, err := service.GetPatient(1)

		assert.Error(t, err)
		assert.NotNil(t, res)
		assert.EqualError(t, err, "Get By ID Process Failed")
	})

	t.Run("Success Get", func(t *testing.T) {
		data.On("GetByID", 1).Return(patient, nil).Once()

		res, err := service.GetPatient(1)

		assert.Nil(t, err)
		assert.NotNil(t, res)
		data.AssertExpectations(t)
	})
}

func TestCreatePatient(t *testing.T) {
	data := mocks.NewPatientDataInterface(t)
	cld := mockUtil.NewCloudinaryInterface(t)
	enkrip := mockHelper.NewHashInterface(t)
	jwt := mockHelper.NewJWTInterface(t)
	service := NewPatient(data, cld, jwt, enkrip)
	patient := patients.Patiententity{
		Name:        "Kamal",
		Email:       "kamal@uwu.com",
		Password:    "password",
		DateOfBirth: "25-09-2009",
		Gender:      "L",
		Avatar:      "https://",
		Phone:       "081324687",
		Role:        "Patient",
		Status:      "Active",
	}

	t.Run("Success Create", func(t *testing.T) {
		enkrip.On("HashPassword", patient.Password).Return("password", nil).Once()
		data.On("Insert", patient).Return(&patient, nil).Once()

		res, err := service.CreatePatient(patient)

		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, res.Name, patient.Name)
		assert.Equal(t, res.Email, patient.Email)
		assert.Equal(t, res.Password, patient.Password)
		assert.Equal(t, res.DateOfBirth, patient.DateOfBirth)
		assert.Equal(t, res.Gender, patient.Gender)
		assert.Equal(t, res.Avatar, patient.Avatar)
		assert.Equal(t, res.Phone, patient.Phone)
		assert.Equal(t, res.Role, patient.Role)
		assert.Equal(t, res.Status, patient.Status)
		data.AssertExpectations(t)
		enkrip.AssertExpectations(t)
	})

	t.Run("Hash Password Error", func(t *testing.T) {
		enkrip.On("HashPassword", patient.Password).Return("", errors.New("Hash Password Error")).Once()

		res, err := service.CreatePatient(patient)

		assert.Error(t, err)
		assert.Nil(t, res)
		assert.EqualError(t, err, "Hash Password Error")
	})

	t.Run("Insert Process Failed", func(t *testing.T) {
		enkrip.On("HashPassword", patient.Password).Return("password", nil).Once()
		data.On("Insert", patient).Return(nil, errors.New("Insert Process Failed")).Once()

		res, err := service.CreatePatient(patient)

		assert.Error(t, err)
		assert.Nil(t, res)
		assert.EqualError(t, err, "Insert Process Failed")
	})
}

func TestPhotoUpload(t *testing.T) {
	data := mocks.NewPatientDataInterface(t)
	cld := mockUtil.NewCloudinaryInterface(t)
	enkrip := mockHelper.NewHashInterface(t)
	jwt := mockHelper.NewJWTInterface(t)
	service := NewPatient(data, cld, jwt, enkrip)
	var mockFile multipart.File
	patient := patients.AvatarPhoto{
		Avatar: mockFile,
	}

	t.Run("Upload Avatar Failed", func(t *testing.T) {
		cld.On("UploadImageHelper", patient.Avatar).Return("", errors.New("Upload Avatar Failed")).Once()

		res, err := service.PhotoUpload(patient)

		assert.Error(t, err)
		assert.NotNil(t, res)
		assert.EqualError(t, err, "Upload Avatar Failed")
	})

	t.Run("Success Upload", func(t *testing.T) {
		cld.On("UploadImageHelper", patient.Avatar).Return("sukses", nil).Once()

		res, err := service.PhotoUpload(patient)

		assert.Nil(t, err)
		assert.Equal(t, res, "sukses")
		cld.AssertExpectations(t)
	})
}

func TestUpdatePatient(t *testing.T) {
	data := mocks.NewPatientDataInterface(t)
	cld := mockUtil.NewCloudinaryInterface(t)
	enkrip := mockHelper.NewHashInterface(t)
	jwt := mockHelper.NewJWTInterface(t)
	service := NewPatient(data, cld, jwt, enkrip)
	patient := patients.UpdateProfile{
		Name:        "Irvan",
		Email:       "irvanhau@gmail.com",
		DateOfBirth: "25-09-2009",
		Gender:      "L",
		Avatar:      "https://",
		Phone:       "12318654",
	}

	t.Run("Server Error", func(t *testing.T) {
		data.On("Update", 1, patient).Return(false, errors.New("Update Process Failed")).Once()

		res, err := service.UpdatePatient(1, patient)

		assert.Error(t, err)
		assert.Equal(t, res, false)
		assert.EqualError(t, err, "Update Process Failed")
	})

	t.Run("Success Update", func(t *testing.T) {
		data.On("Update", 1, patient).Return(true, nil).Once()

		res, err := service.UpdatePatient(1, patient)

		assert.Nil(t, err)
		assert.Equal(t, res, true)
		data.AssertExpectations(t)
	})
}

func TestLoginPatient(t *testing.T) {
	data := mocks.NewPatientDataInterface(t)
	cld := mockUtil.NewCloudinaryInterface(t)
	enkrip := mockHelper.NewHashInterface(t)
	jwt := mockHelper.NewJWTInterface(t)
	service := NewPatient(data, cld, jwt, enkrip)
	patient := patients.Patiententity{
		Name:        "Kamal",
		Email:       "kamal@uwu.com",
		Password:    "password",
		DateOfBirth: "25-09-2009",
		Gender:      "L",
		Avatar:      "https://",
		Phone:       "081324687",
		Role:        "Patient",
		Status:      "Active",
	}

	t.Run("Server Error", func(t *testing.T) {
		data.On("LoginPatient", "irvanhau@gmail.com", "password").Return(nil, errors.New("Login Process Failed")).Once()

		result, err := service.LoginPatient("irvanhau@gmail.com", "password")

		assert.Error(t, err)
		assert.EqualError(t, err, "Login Process Failed")
		assert.Nil(t, result)
	})

	t.Run("Password Incorrect", func(t *testing.T) {
		userFail := patients.Patiententity{
			Email:    "irvanhau",
			Password: "vanhau123",
		}

		data.On("LoginPatient", userFail.Email, userFail.Password).Return(nil, errors.New("Incorrect Password")).Once()

		result, err := service.LoginPatient(userFail.Email, userFail.Password)

		assert.Error(t, err)
		assert.EqualError(t, err, "Incorrect Password")
		assert.Nil(t, result)
	})

	t.Run("Not Found", func(t *testing.T) {
		userFail := patients.Patiententity{
			Email:    "irvanhau",
			Password: "vanhau123",
		}

		data.On("LoginPatient", userFail.Email, userFail.Password).Return(nil, errors.New("User Not Found / User Inactive")).Once()

		result, err := service.LoginPatient(userFail.Email, userFail.Password)

		assert.Error(t, err)
		assert.EqualError(t, err, "User Not Found / User Inactive")
		assert.Nil(t, result)
	})

	t.Run("Generate JWT Error", func(t *testing.T) {
		data.On("LoginPatient", "irvanhau@gmail.com", "password").Return(&patient, nil).Once()
		jwt.On("GenerateJWT", uint(0), patient.Role, patient.Status).Return(nil).Once()

		result, err := service.LoginPatient("irvanhau@gmail.com", "password")

		assert.Error(t, err)
		assert.EqualError(t, err, "Token Process Failed")
		assert.Nil(t, result)
	})

	t.Run("Success Login", func(t *testing.T) {
		jwtResult := map[string]any{"access_token": uint(0), "role": "mockToken"}
		data.On("LoginPatient", "irvanhau@gmail.com", "password").Return(&patient, nil).Once()
		jwt.On("GenerateJWT", uint(0), patient.Role, patient.Status).Return(jwtResult).Once()

		result, err := service.LoginPatient("irvanhau@gmail.com", "password")

		data.AssertExpectations(t)
		jwt.AssertExpectations(t)

		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, patient.Email, result.Email)
		assert.Equal(t, patient.Name, result.Name)
		assert.Equal(t, jwtResult, result.Access)
	})
}

func TestUpdatePassword(t *testing.T) {
	data := mocks.NewPatientDataInterface(t)
	cld := mockUtil.NewCloudinaryInterface(t)
	enkrip := mockHelper.NewHashInterface(t)
	jwt := mockHelper.NewJWTInterface(t)
	service := NewPatient(data, cld, jwt, enkrip)
	patient := patients.UpdatePassword{
		Password: "password",
	}

	t.Run("Hash Password Error", func(t *testing.T) {
		enkrip.On("HashPassword", patient.Password).Return("", errors.New("Hash Password Failed")).Once()

		res, err := service.UpdatePassword(1, patient)

		assert.Error(t, err)
		assert.Equal(t, res, false)
		assert.EqualError(t, err, "Hash Password Failed")
	})

	t.Run("Server Error", func(t *testing.T) {
		enkrip.On("HashPassword", patient.Password).Return("password", nil).Once()
		data.On("UpdatePassword", 1, patient).Return(false, errors.New("Update Password Failed")).Once()

		res, err := service.UpdatePassword(1, patient)

		assert.Error(t, err)
		assert.Equal(t, res, false)
		assert.EqualError(t, err, "Update Password Failed")
	})

	t.Run("Success Update", func(t *testing.T) {
		enkrip.On("HashPassword", patient.Password).Return("password", nil).Once()
		data.On("UpdatePassword", 1, patient).Return(true, nil).Once()

		res, err := service.UpdatePassword(1, patient)

		assert.Nil(t, err)
		assert.Equal(t, res, true)
		data.AssertExpectations(t)
		enkrip.AssertExpectations(t)
	})
}

func TestPatientDashboard(t *testing.T) {
	data := mocks.NewPatientDataInterface(t)
	cld := mockUtil.NewCloudinaryInterface(t)
	enkrip := mockHelper.NewHashInterface(t)
	jwt := mockHelper.NewJWTInterface(t)
	service := NewPatient(data, cld, jwt, enkrip)
	var dashboard patients.PatientDashboard

	t.Run("Server Error", func(t *testing.T) {
		data.On("PatientDashboard").Return(dashboard, errors.New("Process Failed")).Once()

		res, err := service.PatientDashboard()

		assert.Error(t, err)
		assert.NotNil(t, res)
		assert.EqualError(t, err, "Process Failed")
	})

	t.Run("Success Get", func(t *testing.T) {
		data.On("PatientDashboard").Return(dashboard, nil).Once()

		res, err := service.PatientDashboard()

		assert.Nil(t, err)
		assert.NotNil(t, res)
		data.AssertExpectations(t)
	})
}

func TestUpdateStatus(t *testing.T) {
	data := mocks.NewPatientDataInterface(t)
	cld := mockUtil.NewCloudinaryInterface(t)
	enkrip := mockHelper.NewHashInterface(t)
	jwt := mockHelper.NewJWTInterface(t)
	service := NewPatient(data, cld, jwt, enkrip)
	patient := patients.UpdateStatus{
		Status: "Active",
	}

	t.Run("Server Error", func(t *testing.T) {
		data.On("UpdateStatus", 1, patient).Return(false, errors.New("Update Status Process Failed")).Once()

		res, err := service.UpdateStatus(1, patient)

		assert.Error(t, err)
		assert.Equal(t, res, false)
		assert.EqualError(t, err, "Update Status Process Failed")
	})

	t.Run("Success Update", func(t *testing.T) {
		data.On("UpdateStatus", 1, patient).Return(true, nil).Once()

		res, err := service.UpdateStatus(1, patient)

		assert.Nil(t, err)
		assert.Equal(t, res, true)
		data.AssertExpectations(t)
	})
}

func TestInactivateAccount(t *testing.T) {
	data := mocks.NewPatientDataInterface(t)
	cld := mockUtil.NewCloudinaryInterface(t)
	enkrip := mockHelper.NewHashInterface(t)
	jwt := mockHelper.NewJWTInterface(t)
	service := NewPatient(data, cld, jwt, enkrip)

	t.Run("Server Error", func(t *testing.T) {
		data.On("InactivateAccount", 1).Return(false, errors.New("Inactivate Account Process Failed")).Once()

		res, err := service.InactivateAccount(1)

		assert.Error(t, err)
		assert.Equal(t, res, false)
		assert.EqualError(t, err, "Inactivate Account Process Failed")
	})

	t.Run("Success InactivateAccount", func(t *testing.T) {
		data.On("InactivateAccount", 1).Return(true, nil).Once()

		res, err := service.InactivateAccount(1)

		assert.Nil(t, err)
		assert.Equal(t, res, true)
		data.AssertExpectations(t)
	})
}
