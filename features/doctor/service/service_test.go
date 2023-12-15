package service

import (
	"FinalProject/features/doctor"
	"FinalProject/features/doctor/mocks"
	mockHelper "FinalProject/helper/mocks"
	mockUtil "FinalProject/utils/mocks"
	"errors"
	"mime/multipart"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetDoctors(t *testing.T) {
	data := mocks.NewDoctorDataInterface(t)
	cld := mockUtil.NewCloudinaryInterface(t)
	email := mockHelper.NewEmailInterface(t)
	meet := mockHelper.NewMeetInterface(t)
	service := NewDoctor(data, cld, email, meet)
	dataDoctor := []doctor.DoctorAll{}

	t.Run("Server Error", func(t *testing.T) {
		data.On("GetAll", "name").Return(nil, errors.New("Get All Process Failed")).Once()

		res, err := service.GetDoctors("name")

		assert.Error(t, err)
		assert.Nil(t, res)
		assert.EqualError(t, err, "Get All Process Failed")
	})

	t.Run("Success Get", func(t *testing.T) {
		data.On("GetAll", "name").Return(dataDoctor, nil).Once()

		res, err := service.GetDoctors("name")

		assert.Nil(t, err)
		assert.NotNil(t, res)
		data.AssertExpectations(t)
	})
}

func TestGetDoctor(t *testing.T) {
	data := mocks.NewDoctorDataInterface(t)
	cld := mockUtil.NewCloudinaryInterface(t)
	email := mockHelper.NewEmailInterface(t)
	meet := mockHelper.NewMeetInterface(t)
	service := NewDoctor(data, cld, email, meet)
	var dataDoctor doctor.DoctorAll

	t.Run("Server Error", func(t *testing.T) {
		data.On("GetByID", 1).Return(nil, errors.New("Get By ID Doctor Process Failed")).Once()

		res, err := service.GetDoctor(1)

		assert.Error(t, err)
		assert.Nil(t, res)
		assert.EqualError(t, err, "Get By ID Doctor Process Failed")
	})

	t.Run("Success Get", func(t *testing.T) {
		data.On("GetByID", 1).Return(&dataDoctor, nil).Once()

		res, err := service.GetDoctor(1)

		assert.Nil(t, err)
		assert.NotNil(t, res)
		data.AssertExpectations(t)
	})
}

func TestGetDoctorByUserId(t *testing.T) {
	data := mocks.NewDoctorDataInterface(t)
	cld := mockUtil.NewCloudinaryInterface(t)
	email := mockHelper.NewEmailInterface(t)
	meet := mockHelper.NewMeetInterface(t)
	service := NewDoctor(data, cld, email, meet)
	var dataDoctor doctor.DoctorAll

	t.Run("Server Error", func(t *testing.T) {
		data.On("GetDoctorByUserId", 1).Return(nil, errors.New("Get Doctor By User ID Process Failed")).Once()

		res, err := service.GetDoctorByUserId(1)

		assert.Error(t, err)
		assert.Nil(t, res)
		assert.EqualError(t, err, "Get Doctor By User ID Process Failed")
	})

	t.Run("Success Get", func(t *testing.T) {
		data.On("GetDoctorByUserId", 1).Return(&dataDoctor, nil).Once()

		res, err := service.GetDoctorByUserId(1)

		assert.Nil(t, err)
		assert.NotNil(t, res)
		data.AssertExpectations(t)
	})
}

func TestGetDoctorWorkadays(t *testing.T) {
	data := mocks.NewDoctorDataInterface(t)
	cld := mockUtil.NewCloudinaryInterface(t)
	email := mockHelper.NewEmailInterface(t)
	meet := mockHelper.NewMeetInterface(t)
	service := NewDoctor(data, cld, email, meet)
	dataDoctor := []doctor.DoctorWorkdays{}

	t.Run("Server Error", func(t *testing.T) {
		data.On("GetByIDWorkadays", 1).Return(nil, errors.New("Get By ID Workadays Process Failed")).Once()

		res, err := service.GetDoctorWorkadays(1)

		assert.Error(t, err)
		assert.Nil(t, res)
		assert.EqualError(t, err, "Get By ID Workadays Process Failed")
	})

	t.Run("Success Get", func(t *testing.T) {
		data.On("GetByIDWorkadays", 1).Return(dataDoctor, nil).Once()

		res, err := service.GetDoctorWorkadays(1)

		assert.Nil(t, err)
		assert.NotNil(t, res)
		data.AssertExpectations(t)
	})
}

func TestGetDoctorEducation(t *testing.T) {
	data := mocks.NewDoctorDataInterface(t)
	cld := mockUtil.NewCloudinaryInterface(t)
	email := mockHelper.NewEmailInterface(t)
	meet := mockHelper.NewMeetInterface(t)
	service := NewDoctor(data, cld, email, meet)
	dataDoctor := []doctor.DoctorEducation{}

	t.Run("Server Error", func(t *testing.T) {
		data.On("GetByIDEducation", 1).Return(nil, errors.New("Get By ID Education Process Failed")).Once()

		res, err := service.GetDoctorEducation(1)

		assert.Error(t, err)
		assert.Nil(t, res)
		assert.EqualError(t, err, "Get By ID Education Process Failed")
	})

	t.Run("Success Get", func(t *testing.T) {
		data.On("GetByIDEducation", 1).Return(dataDoctor, nil).Once()

		res, err := service.GetDoctorEducation(1)

		assert.Nil(t, err)
		assert.NotNil(t, res)
		data.AssertExpectations(t)
	})
}

func TestGetDoctorExperience(t *testing.T) {
	data := mocks.NewDoctorDataInterface(t)
	cld := mockUtil.NewCloudinaryInterface(t)
	email := mockHelper.NewEmailInterface(t)
	meet := mockHelper.NewMeetInterface(t)
	service := NewDoctor(data, cld, email, meet)
	dataDoctor := []doctor.DoctorExperience{}

	t.Run("Server Error", func(t *testing.T) {
		data.On("GetByIDExperience", 1).Return(nil, errors.New("Get By ID Experience Process Failed")).Once()

		res, err := service.GetDoctorExperience(1)

		assert.Error(t, err)
		assert.Nil(t, res)
		assert.EqualError(t, err, "Get By ID Experience Process Failed")
	})

	t.Run("Success Get", func(t *testing.T) {
		data.On("GetByIDExperience", 1).Return(dataDoctor, nil).Once()

		res, err := service.GetDoctorExperience(1)

		assert.Nil(t, err)
		assert.NotNil(t, res)
		data.AssertExpectations(t)
	})
}

func TestCreateDoctor(t *testing.T) {
	data := mocks.NewDoctorDataInterface(t)
	cld := mockUtil.NewCloudinaryInterface(t)
	email := mockHelper.NewEmailInterface(t)
	meet := mockHelper.NewMeetInterface(t)
	service := NewDoctor(data, cld, email, meet)
	dataDoctor := doctor.Doctor{
		UserID:            1,
		DoctorName:        "Hau",
		DoctorNIK:         "1234",
		DoctorDOB:         "25-09-1999",
		DoctorProvinsi:    "Jawa Wireng",
		DoctorKota:        "Kota Cinta",
		DoctorNumberPhone: "0898654861",
		DoctorGender:      "L",
		DoctorAvatar:      "https://",
		DoctorDescription: "description",
		DoctorMeetLink:    "meet link",
		DoctorSIPP:        "sipp",
		DoctorSIPPFile:    "file",
	}
	var mockEmail = "irvanhau@gmail.com"

	t.Run("Server Error", func(t *testing.T) {
		data.On("Insert", dataDoctor).Return(nil, errors.New("Insert Doctor Process Failed")).Once()

		res, err := service.CreateDoctor(dataDoctor)

		assert.Error(t, err)
		assert.Nil(t, res)
		assert.EqualError(t, err, "Insert Doctor Process Failed")
	})

	t.Run("Find Email Error", func(t *testing.T) {
		data.On("Insert", dataDoctor).Return(&dataDoctor, nil).Once()
		data.On("FindEmail", dataDoctor.UserID).Return(&mockEmail, errors.New("Find Email Process Failed")).Once()

		res, err := service.CreateDoctor(dataDoctor)

		assert.Error(t, err)
		assert.Nil(t, res)
		assert.EqualError(t, err, "Find Email Process Failed")
	})

	t.Run("Send Email Error", func(t *testing.T) {
		data.On("Insert", dataDoctor).Return(&dataDoctor, nil).Once()
		data.On("FindEmail", dataDoctor.UserID).Return(&mockEmail, nil).Once()
		email.On("HtmlBodyRegistDoctor", dataDoctor.DoctorName).Return("header", "htmlbody").Once()
		email.On("SendEmail", mockEmail, "header", "htmlbody").Return(errors.New("Send Email Error")).Once()

		res, err := service.CreateDoctor(dataDoctor)

		assert.Error(t, err)
		assert.Nil(t, res)
		assert.EqualError(t, err, "Send Email Error")
	})

	t.Run("Success Create", func(t *testing.T) {
		data.On("Insert", dataDoctor).Return(&dataDoctor, nil).Once()
		data.On("FindEmail", dataDoctor.UserID).Return(&mockEmail, nil).Once()
		email.On("HtmlBodyRegistDoctor", dataDoctor.DoctorName).Return("header", "htmlbody").Once()
		email.On("SendEmail", mockEmail, "header", "htmlbody").Return(nil).Once()

		res, err := service.CreateDoctor(dataDoctor)

		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, res.UserID, dataDoctor.UserID)
		assert.Equal(t, res.DoctorName, dataDoctor.DoctorName)
		assert.Equal(t, res.DoctorNIK, dataDoctor.DoctorNIK)
		assert.Equal(t, res.DoctorDOB, dataDoctor.DoctorDOB)
		assert.Equal(t, res.DoctorProvinsi, dataDoctor.DoctorProvinsi)
		assert.Equal(t, res.DoctorKota, dataDoctor.DoctorKota)
		assert.Equal(t, res.DoctorNumberPhone, dataDoctor.DoctorNumberPhone)
		assert.Equal(t, res.DoctorGender, dataDoctor.DoctorGender)
		assert.Equal(t, res.DoctorAvatar, dataDoctor.DoctorAvatar)
		assert.Equal(t, res.DoctorDescription, dataDoctor.DoctorDescription)
		assert.Equal(t, res.DoctorMeetLink, dataDoctor.DoctorMeetLink)
		assert.Equal(t, res.DoctorSIPP, dataDoctor.DoctorSIPP)
		assert.Equal(t, res.DoctorSIPPFile, dataDoctor.DoctorSIPPFile)
		data.AssertExpectations(t)
		email.AssertExpectations(t)
	})
}

func TestCreateDoctorExpertise(t *testing.T) {
	data := mocks.NewDoctorDataInterface(t)
	cld := mockUtil.NewCloudinaryInterface(t)
	email := mockHelper.NewEmailInterface(t)
	meet := mockHelper.NewMeetInterface(t)
	service := NewDoctor(data, cld, email, meet)
	dataDoctor := doctor.DoctorExpertiseRelation{
		DoctorID:    1,
		ExpertiseID: 2,
	}

	t.Run("Server Error", func(t *testing.T) {
		data.On("InsertExpertise", dataDoctor).Return(nil, errors.New("Insert Doctor Expertise Process Failed")).Once()

		res, err := service.CreateDoctorExpertise(dataDoctor)

		assert.Error(t, err)
		assert.Nil(t, res)
		assert.EqualError(t, err, "Insert Doctor Expertise Process Failed")
	})

	t.Run("Success Insert", func(t *testing.T) {
		data.On("InsertExpertise", dataDoctor).Return(&dataDoctor, nil).Once()

		res, err := service.CreateDoctorExpertise(dataDoctor)

		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, res.DoctorID, dataDoctor.DoctorID)
		assert.Equal(t, res.ExpertiseID, dataDoctor.ExpertiseID)
		data.AssertExpectations(t)
	})
}

func TestCreateDoctorWorkadays(t *testing.T) {
	data := mocks.NewDoctorDataInterface(t)
	cld := mockUtil.NewCloudinaryInterface(t)
	email := mockHelper.NewEmailInterface(t)
	meet := mockHelper.NewMeetInterface(t)
	service := NewDoctor(data, cld, email, meet)
	dataDoctor := doctor.DoctorWorkdays{
		DoctorID:  1,
		WorkdayID: 1,
		StartTime: time.Now(),
		EndTime:   time.Now(),
	}

	t.Run("Server Error", func(t *testing.T) {
		data.On("InsertWorkadays", dataDoctor).Return(nil, errors.New("Insert Doctor Workadays Process Failed")).Once()

		res, err := service.CreateDoctorWorkadays(dataDoctor)

		assert.Error(t, err)
		assert.Nil(t, res)
		assert.EqualError(t, err, "Insert Doctor Workadays Process Failed")
	})

	t.Run("Success Insert", func(t *testing.T) {
		data.On("InsertWorkadays", dataDoctor).Return(&dataDoctor, nil).Once()

		res, err := service.CreateDoctorWorkadays(dataDoctor)

		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, res.DoctorID, dataDoctor.DoctorID)
		assert.Equal(t, res.WorkdayID, dataDoctor.WorkdayID)
		assert.Equal(t, res.StartTime, dataDoctor.StartTime)
		assert.Equal(t, res.EndTime, dataDoctor.EndTime)
		data.AssertExpectations(t)
	})
}

func TestCreateDoctorEducation(t *testing.T) {
	data := mocks.NewDoctorDataInterface(t)
	cld := mockUtil.NewCloudinaryInterface(t)
	email := mockHelper.NewEmailInterface(t)
	meet := mockHelper.NewMeetInterface(t)
	service := NewDoctor(data, cld, email, meet)
	dataDoctor := doctor.DoctorEducation{
		DoctorID:           1,
		DoctorUniversity:   "univ",
		DoctorStudyProgram: "study",
		DoctorEnrollYear:   time.Now(),
		DoctorGraduateYear: time.Now(),
	}

	t.Run("Server Error", func(t *testing.T) {
		data.On("InsertEducation", dataDoctor).Return(nil, errors.New("Insert Doctor Education Process Failed")).Once()

		res, err := service.CreateDoctorEducation(dataDoctor)

		assert.Error(t, err)
		assert.Nil(t, res)
		assert.EqualError(t, err, "Insert Doctor Education Process Failed")
	})

	t.Run("Success Insert", func(t *testing.T) {
		data.On("InsertEducation", dataDoctor).Return(&dataDoctor, nil).Once()

		res, err := service.CreateDoctorEducation(dataDoctor)

		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, res.DoctorID, dataDoctor.DoctorID)
		assert.Equal(t, res.DoctorUniversity, dataDoctor.DoctorUniversity)
		assert.Equal(t, res.DoctorStudyProgram, dataDoctor.DoctorStudyProgram)
		assert.Equal(t, res.DoctorEnrollYear, dataDoctor.DoctorEnrollYear)
		assert.Equal(t, res.DoctorGraduateYear, dataDoctor.DoctorGraduateYear)
		data.AssertExpectations(t)
	})
}

func TestCreateDoctorExperience(t *testing.T) {
	data := mocks.NewDoctorDataInterface(t)
	cld := mockUtil.NewCloudinaryInterface(t)
	email := mockHelper.NewEmailInterface(t)
	meet := mockHelper.NewMeetInterface(t)
	service := NewDoctor(data, cld, email, meet)
	dataDoctor := doctor.DoctorExperience{
		DoctorID:             1,
		DoctorCompany:        "company",
		DoctorTitle:          "title",
		DoctorCompanyAddress: "address",
		DoctorStartDate:      time.Now(),
		DoctorEndDate:        time.Now(),
	}

	t.Run("Server Error", func(t *testing.T) {
		data.On("InsertExperience", dataDoctor).Return(nil, errors.New("Insert Doctor Experience Process Failed")).Once()

		res, err := service.CreateDoctorExperience(dataDoctor)

		assert.Error(t, err)
		assert.Nil(t, res)
		assert.EqualError(t, err, "Insert Doctor Experience Process Failed")
	})

	t.Run("Success Insert", func(t *testing.T) {
		data.On("InsertExperience", dataDoctor).Return(&dataDoctor, nil).Once()

		res, err := service.CreateDoctorExperience(dataDoctor)

		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, res.DoctorID, dataDoctor.DoctorID)
		assert.Equal(t, res.DoctorCompany, dataDoctor.DoctorCompany)
		assert.Equal(t, res.DoctorTitle, dataDoctor.DoctorTitle)
		assert.Equal(t, res.DoctorCompanyAddress, dataDoctor.DoctorCompanyAddress)
		assert.Equal(t, res.DoctorStartDate, dataDoctor.DoctorStartDate)
		assert.Equal(t, res.DoctorEndDate, dataDoctor.DoctorEndDate)
		data.AssertExpectations(t)
	})
}

func TestDoctorAvatarUpload(t *testing.T) {
	data := mocks.NewDoctorDataInterface(t)
	cld := mockUtil.NewCloudinaryInterface(t)
	email := mockHelper.NewEmailInterface(t)
	meet := mockHelper.NewMeetInterface(t)
	service := NewDoctor(data, cld, email, meet)
	var mockFile multipart.File
	dataDoctor := doctor.DoctorAvatarPhoto{
		DoctorAvatar: mockFile,
	}

	t.Run("Server Error", func(t *testing.T) {
		cld.On("UploadImageHelper", dataDoctor.DoctorAvatar).Return("", errors.New("Upload Avatar Failed")).Once()

		res, err := service.DoctorAvatarUpload(dataDoctor)

		assert.Error(t, err)
		assert.NotNil(t, res)
		assert.EqualError(t, err, "Upload Avatar Failed")
	})

	t.Run("Success Insert", func(t *testing.T) {
		cld.On("UploadImageHelper", dataDoctor.DoctorAvatar).Return("https://", nil).Once()

		res, err := service.DoctorAvatarUpload(dataDoctor)

		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, res, "https://")
		cld.AssertExpectations(t)
	})
}

func TestDoctorSTRUpload(t *testing.T) {
	data := mocks.NewDoctorDataInterface(t)
	cld := mockUtil.NewCloudinaryInterface(t)
	email := mockHelper.NewEmailInterface(t)
	meet := mockHelper.NewMeetInterface(t)
	service := NewDoctor(data, cld, email, meet)
	var mockFile multipart.File
	dataDoctor := doctor.DoctorSTRFileDataModel{
		DoctorSTRFile: mockFile,
	}

	t.Run("Server Error", func(t *testing.T) {
		cld.On("UploadImageHelper", dataDoctor.DoctorSTRFile).Return("", errors.New("Upload Surat Tanda Registrasi Failed")).Once()

		res, err := service.DoctorSTRUpload(dataDoctor)

		assert.Error(t, err)
		assert.NotNil(t, res)
		assert.EqualError(t, err, "Upload Surat Tanda Registrasi Failed")
	})

	t.Run("Success Insert", func(t *testing.T) {
		cld.On("UploadImageHelper", dataDoctor.DoctorSTRFile).Return("https://", nil).Once()

		res, err := service.DoctorSTRUpload(dataDoctor)

		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, res, "https://")
		cld.AssertExpectations(t)
	})
}

func TestDoctorSIPPUpload(t *testing.T) {
	data := mocks.NewDoctorDataInterface(t)
	cld := mockUtil.NewCloudinaryInterface(t)
	email := mockHelper.NewEmailInterface(t)
	meet := mockHelper.NewMeetInterface(t)
	service := NewDoctor(data, cld, email, meet)
	var mockFile multipart.File
	dataDoctor := doctor.DoctorSIPPFileDataModel{
		DoctorSIPPFile: mockFile,
	}

	t.Run("Server Error", func(t *testing.T) {
		cld.On("UploadImageHelper", dataDoctor.DoctorSIPPFile).Return("", errors.New("Upload SIPP Failed")).Once()

		res, err := service.DoctorSIPPUpload(dataDoctor)

		assert.Error(t, err)
		assert.NotNil(t, res)
		assert.EqualError(t, err, "Upload SIPP Failed")
	})

	t.Run("Success Insert", func(t *testing.T) {
		cld.On("UploadImageHelper", dataDoctor.DoctorSIPPFile).Return("https://", nil).Once()

		res, err := service.DoctorSIPPUpload(dataDoctor)

		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, res, "https://")
		cld.AssertExpectations(t)
	})
}

func TestDoctorCVUpload(t *testing.T) {
	data := mocks.NewDoctorDataInterface(t)
	cld := mockUtil.NewCloudinaryInterface(t)
	email := mockHelper.NewEmailInterface(t)
	meet := mockHelper.NewMeetInterface(t)
	service := NewDoctor(data, cld, email, meet)
	var mockFile multipart.File
	dataDoctor := doctor.DoctorCVDataModel{
		DoctorCV: mockFile,
	}

	t.Run("Server Error", func(t *testing.T) {
		cld.On("UploadImageHelper", dataDoctor.DoctorCV).Return("", errors.New("Upload CV Failed")).Once()

		res, err := service.DoctorCVUpload(dataDoctor)

		assert.Error(t, err)
		assert.NotNil(t, res)
		assert.EqualError(t, err, "Upload CV Failed")
	})

	t.Run("Success Insert", func(t *testing.T) {
		cld.On("UploadImageHelper", dataDoctor.DoctorCV).Return("https://", nil).Once()

		res, err := service.DoctorCVUpload(dataDoctor)

		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, res, "https://")
		cld.AssertExpectations(t)
	})
}

func TestDoctorIjazahUpload(t *testing.T) {
	data := mocks.NewDoctorDataInterface(t)
	cld := mockUtil.NewCloudinaryInterface(t)
	email := mockHelper.NewEmailInterface(t)
	meet := mockHelper.NewMeetInterface(t)
	service := NewDoctor(data, cld, email, meet)
	var mockFile multipart.File
	dataDoctor := doctor.DoctorIjazahDataModel{
		DoctorIjazah: mockFile,
	}

	t.Run("Server Error", func(t *testing.T) {
		cld.On("UploadImageHelper", dataDoctor.DoctorIjazah).Return("", errors.New("Upload Ijazah Failed")).Once()

		res, err := service.DoctorIjazahUpload(dataDoctor)

		assert.Error(t, err)
		assert.NotNil(t, res)
		assert.EqualError(t, err, "Upload Ijazah Failed")
	})

	t.Run("Success Insert", func(t *testing.T) {
		cld.On("UploadImageHelper", dataDoctor.DoctorIjazah).Return("https://", nil).Once()

		res, err := service.DoctorIjazahUpload(dataDoctor)

		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, res, "https://")
		cld.AssertExpectations(t)
	})
}

func TestUpdateDoctorDatapokok(t *testing.T) {
	data := mocks.NewDoctorDataInterface(t)
	cld := mockUtil.NewCloudinaryInterface(t)
	email := mockHelper.NewEmailInterface(t)
	meet := mockHelper.NewMeetInterface(t)
	service := NewDoctor(data, cld, email, meet)
	dataDoctor := doctor.DoctorDatapokokUpdate{
		UserID:            1,
		DoctorName:        "Hau",
		DoctorNIK:         "1234",
		DoctorDOB:         "25-09-1999",
		DoctorProvinsi:    "Jawa Wireng",
		DoctorKota:        "Kota Cinta",
		DoctorNumberPhone: "0898654861",
		DoctorGender:      "L",
		DoctorAvatar:      "https://",
		DoctorDescription: "description",
		DoctorMeetLink:    "meet link",
		DoctorSIPP:        "sipp",
		DoctorSIPPFile:    "file",
	}

	t.Run("Server Error", func(t *testing.T) {
		data.On("UpdateDoctorDatapokok", 1, dataDoctor).Return(false, errors.New("Update Datapokok Dokter Failed")).Once()

		res, err := service.UpdateDoctorDatapokok(1, dataDoctor)

		assert.Error(t, err)
		assert.Equal(t, false, res)
		assert.EqualError(t, err, "Update Datapokok Dokter Failed")
	})

	t.Run("Success Update", func(t *testing.T) {
		data.On("UpdateDoctorDatapokok", 1, dataDoctor).Return(true, nil).Once()

		res, err := service.UpdateDoctorDatapokok(1, dataDoctor)

		assert.Nil(t, err)
		assert.Equal(t, true, res)
		data.AssertExpectations(t)
	})
}

func TestUpdateDoctorEducation(t *testing.T) {
	data := mocks.NewDoctorDataInterface(t)
	cld := mockUtil.NewCloudinaryInterface(t)
	email := mockHelper.NewEmailInterface(t)
	meet := mockHelper.NewMeetInterface(t)
	service := NewDoctor(data, cld, email, meet)
	dataDoctor := doctor.DoctorEducation{
		DoctorID:           1,
		DoctorUniversity:   "univ",
		DoctorStudyProgram: "study",
		DoctorEnrollYear:   time.Now(),
		DoctorGraduateYear: time.Now(),
	}
	t.Run("Server Error", func(t *testing.T) {
		data.On("UpdateDoctorEducation", 1, 1, dataDoctor).Return(false, errors.New("Update Education Dokter Failed")).Once()

		res, err := service.UpdateDoctorEducation(1, 1, dataDoctor)

		assert.Error(t, err)
		assert.Equal(t, false, res)
		assert.EqualError(t, err, "Update Education Dokter Failed")
	})

	t.Run("Success Update", func(t *testing.T) {
		data.On("UpdateDoctorEducation", 1, 1, dataDoctor).Return(true, nil).Once()

		res, err := service.UpdateDoctorEducation(1, 1, dataDoctor)

		assert.Nil(t, err)
		assert.Equal(t, true, res)
		data.AssertExpectations(t)
	})
}

func TestUpdateDoctorExperience(t *testing.T) {
	data := mocks.NewDoctorDataInterface(t)
	cld := mockUtil.NewCloudinaryInterface(t)
	email := mockHelper.NewEmailInterface(t)
	meet := mockHelper.NewMeetInterface(t)
	service := NewDoctor(data, cld, email, meet)
	dataDoctor := doctor.DoctorExperience{
		DoctorID:             1,
		DoctorCompany:        "company",
		DoctorTitle:          "title",
		DoctorCompanyAddress: "address",
		DoctorStartDate:      time.Now(),
		DoctorEndDate:        time.Now(),
	}
	t.Run("Server Error", func(t *testing.T) {
		data.On("UpdateDoctorExperience", 1, 1, dataDoctor).Return(false, errors.New("Update Experience Dokter Failed")).Once()

		res, err := service.UpdateDoctorExperience(1, 1, dataDoctor)

		assert.Error(t, err)
		assert.Equal(t, false, res)
		assert.EqualError(t, err, "Update Experience Dokter Failed")
	})

	t.Run("Success Update", func(t *testing.T) {
		data.On("UpdateDoctorExperience", 1, 1, dataDoctor).Return(true, nil).Once()

		res, err := service.UpdateDoctorExperience(1, 1, dataDoctor)

		assert.Nil(t, err)
		assert.Equal(t, true, res)
		data.AssertExpectations(t)
	})
}

func TestUpdateDoctorWorkdays(t *testing.T) {
	data := mocks.NewDoctorDataInterface(t)
	cld := mockUtil.NewCloudinaryInterface(t)
	email := mockHelper.NewEmailInterface(t)
	meet := mockHelper.NewMeetInterface(t)
	service := NewDoctor(data, cld, email, meet)
	dataDoctor := doctor.DoctorWorkdays{
		DoctorID:  1,
		WorkdayID: 1,
		StartTime: time.Now(),
		EndTime:   time.Now(),
	}
	t.Run("Server Error", func(t *testing.T) {
		data.On("UpdateDoctorWorkdays", 1, 1, dataDoctor).Return(false, errors.New("Update Workdays Dokter Failed")).Once()

		res, err := service.UpdateDoctorWorkdays(1, 1, dataDoctor)

		assert.Error(t, err)
		assert.Equal(t, false, res)
		assert.EqualError(t, err, "Update Workdays Dokter Failed")
	})

	t.Run("Success Update", func(t *testing.T) {
		data.On("UpdateDoctorWorkdays", 1, 1, dataDoctor).Return(true, nil).Once()

		res, err := service.UpdateDoctorWorkdays(1, 1, dataDoctor)

		assert.Nil(t, err)
		assert.Equal(t, true, res)
		data.AssertExpectations(t)
	})
}

func TestUpdateDoctorRating(t *testing.T) {
	data := mocks.NewDoctorDataInterface(t)
	cld := mockUtil.NewCloudinaryInterface(t)
	email := mockHelper.NewEmailInterface(t)
	meet := mockHelper.NewMeetInterface(t)
	service := NewDoctor(data, cld, email, meet)
	dataDoctor := doctor.DoctorRating{
		DoctorID:         1,
		PatientID:        1,
		TransactionID:    "ID-001",
		DoctorStarRating: 4,
		DoctorReview:     "Mantap",
	}
	t.Run("Server Error", func(t *testing.T) {
		data.On("UpdateDoctorRating", 1, 1, dataDoctor).Return(false, errors.New("Update Rating Dokter Failed")).Once()

		res, err := service.UpdateDoctorRating(1, 1, dataDoctor)

		assert.Error(t, err)
		assert.Equal(t, false, res)
		assert.EqualError(t, err, "Update Rating Dokter Failed")
	})

	t.Run("Success Update", func(t *testing.T) {
		data.On("UpdateDoctorRating", 1, 1, dataDoctor).Return(true, nil).Once()

		res, err := service.UpdateDoctorRating(1, 1, dataDoctor)

		assert.Nil(t, err)
		assert.Equal(t, true, res)
		data.AssertExpectations(t)
	})
}

func TestDeleteDoctor(t *testing.T) {
	data := mocks.NewDoctorDataInterface(t)
	cld := mockUtil.NewCloudinaryInterface(t)
	email := mockHelper.NewEmailInterface(t)
	meet := mockHelper.NewMeetInterface(t)
	service := NewDoctor(data, cld, email, meet)

	t.Run("Server Error", func(t *testing.T) {
		data.On("DeleteDoctor", 1).Return(false, errors.New("Delete Dokter Failed")).Once()

		res, err := service.DeleteDoctor(1)

		assert.Error(t, err)
		assert.Equal(t, false, res)
		assert.EqualError(t, err, "Delete Dokter Failed")
	})

	t.Run("Success Delete", func(t *testing.T) {
		data.On("DeleteDoctor", 1).Return(true, nil).Once()

		res, err := service.DeleteDoctor(1)

		assert.Nil(t, err)
		assert.Equal(t, true, res)
		data.AssertExpectations(t)
	})
}

func TestDeleteDoctorEducation(t *testing.T) {
	data := mocks.NewDoctorDataInterface(t)
	cld := mockUtil.NewCloudinaryInterface(t)
	email := mockHelper.NewEmailInterface(t)
	meet := mockHelper.NewMeetInterface(t)
	service := NewDoctor(data, cld, email, meet)

	t.Run("Server Error", func(t *testing.T) {
		data.On("DeleteDoctorEducation", 1).Return(false, errors.New("Delete Education Dokter Failed")).Once()

		res, err := service.DeleteDoctorEducation(1)

		assert.Error(t, err)
		assert.Equal(t, false, res)
		assert.EqualError(t, err, "Delete Education Dokter Failed")
	})

	t.Run("Success Delete", func(t *testing.T) {
		data.On("DeleteDoctorEducation", 1).Return(true, nil).Once()

		res, err := service.DeleteDoctorEducation(1)

		assert.Nil(t, err)
		assert.Equal(t, true, res)
		data.AssertExpectations(t)
	})
}

func TestDeleteDoctorExperience(t *testing.T) {
	data := mocks.NewDoctorDataInterface(t)
	cld := mockUtil.NewCloudinaryInterface(t)
	email := mockHelper.NewEmailInterface(t)
	meet := mockHelper.NewMeetInterface(t)
	service := NewDoctor(data, cld, email, meet)

	t.Run("Server Error", func(t *testing.T) {
		data.On("DeleteDoctorExperience", 1).Return(false, errors.New("Delete Experience Dokter Failed")).Once()

		res, err := service.DeleteDoctorExperience(1)

		assert.Error(t, err)
		assert.Equal(t, false, res)
		assert.EqualError(t, err, "Delete Experience Dokter Failed")
	})

	t.Run("Success Delete", func(t *testing.T) {
		data.On("DeleteDoctorExperience", 1).Return(true, nil).Once()

		res, err := service.DeleteDoctorExperience(1)

		assert.Nil(t, err)
		assert.Equal(t, true, res)
		data.AssertExpectations(t)
	})
}

func TestDeleteDoctorWorkdays(t *testing.T) {
	data := mocks.NewDoctorDataInterface(t)
	cld := mockUtil.NewCloudinaryInterface(t)
	email := mockHelper.NewEmailInterface(t)
	meet := mockHelper.NewMeetInterface(t)
	service := NewDoctor(data, cld, email, meet)

	t.Run("Server Error", func(t *testing.T) {
		data.On("DeleteDoctorWorkdays", 1).Return(false, errors.New("Delete Workdays Dokter Failed")).Once()

		res, err := service.DeleteDoctorWorkdays(1)

		assert.Error(t, err)
		assert.Equal(t, false, res)
		assert.EqualError(t, err, "Delete Workdays Dokter Failed")
	})

	t.Run("Success Delete", func(t *testing.T) {
		data.On("DeleteDoctorWorkdays", 1).Return(true, nil).Once()

		res, err := service.DeleteDoctorWorkdays(1)

		assert.Nil(t, err)
		assert.Equal(t, true, res)
		data.AssertExpectations(t)
	})
}

func TestDeleteDoctorRating(t *testing.T) {
	data := mocks.NewDoctorDataInterface(t)
	cld := mockUtil.NewCloudinaryInterface(t)
	email := mockHelper.NewEmailInterface(t)
	meet := mockHelper.NewMeetInterface(t)
	service := NewDoctor(data, cld, email, meet)

	t.Run("Server Error", func(t *testing.T) {
		data.On("DeleteDoctorRating", 1).Return(false, errors.New("Delete Rating Dokter Failed")).Once()

		res, err := service.DeleteDoctorRating(1)

		assert.Error(t, err)
		assert.Equal(t, false, res)
		assert.EqualError(t, err, "Delete Rating Dokter Failed")
	})

	t.Run("Success Delete", func(t *testing.T) {
		data.On("DeleteDoctorRating", 1).Return(true, nil).Once()

		res, err := service.DeleteDoctorRating(1)

		assert.Nil(t, err)
		assert.Equal(t, true, res)
		data.AssertExpectations(t)
	})
}

func TestGetMeetLink(t *testing.T) {
	data := mocks.NewDoctorDataInterface(t)
	cld := mockUtil.NewCloudinaryInterface(t)
	email := mockHelper.NewEmailInterface(t)
	meet := mockHelper.NewMeetInterface(t)
	service := NewDoctor(data, cld, email, meet)
	mockMeet := []string{
		"meet",
	}

	t.Run("Link Sudah Digunakan", func(t *testing.T) {
		meet.On("GetMeetLink").Return(mockMeet).Once()
		data.On("IsLinkUsed", mockMeet[0]).Return(true).Once()

		res, err := service.GetMeetLink()

		assert.Error(t, err)
		assert.Equal(t, res, "")
		assert.EqualError(t, err, "Semua link sudah digunakan")
	})

	t.Run("Success Get", func(t *testing.T) {
		meet.On("GetMeetLink").Return(mockMeet).Once()
		data.On("IsLinkUsed", mockMeet[0]).Return(false).Once()

		res, err := service.GetMeetLink()

		assert.Nil(t, err)
		assert.Equal(t, res, mockMeet[0])
		data.AssertExpectations(t)
		meet.AssertExpectations(t)
	})
}

func TestDoctorDashboard(t *testing.T) {
	data := mocks.NewDoctorDataInterface(t)
	cld := mockUtil.NewCloudinaryInterface(t)
	email := mockHelper.NewEmailInterface(t)
	meet := mockHelper.NewMeetInterface(t)
	service := NewDoctor(data, cld, email, meet)
	var dashboard doctor.DoctorDashboard

	t.Run("Server Error", func(t *testing.T) {
		data.On("DoctorDashboard", 1).Return(dashboard, errors.New("Process Failed")).Once()

		res, err := service.DoctorDashboard(1)

		assert.Error(t, err)
		assert.NotNil(t, res)
		assert.EqualError(t, err, "Process Failed")
	})

	t.Run("Success Get", func(t *testing.T) {
		data.On("DoctorDashboard", 1).Return(dashboard, nil).Once()

		res, err := service.DoctorDashboard(1)

		assert.Nil(t, err)
		assert.NotNil(t, res)
		data.AssertExpectations(t)
	})
}

func TestDoctorDashboardPatient(t *testing.T) {
	data := mocks.NewDoctorDataInterface(t)
	cld := mockUtil.NewCloudinaryInterface(t)
	email := mockHelper.NewEmailInterface(t)
	meet := mockHelper.NewMeetInterface(t)
	service := NewDoctor(data, cld, email, meet)
	dashboard := []doctor.DoctorDashboardPatient{}

	t.Run("Server Error", func(t *testing.T) {
		data.On("DoctorDashboardPatient", 1).Return(dashboard, errors.New("Process Failed")).Once()

		res, err := service.DoctorDashboardPatient(1)

		assert.Error(t, err)
		assert.NotNil(t, res)
		assert.EqualError(t, err, "Process Failed")
	})

	t.Run("Success Get", func(t *testing.T) {
		data.On("DoctorDashboardPatient", 1).Return(dashboard, nil).Once()

		res, err := service.DoctorDashboardPatient(1)

		assert.Nil(t, err)
		assert.NotNil(t, res)
		data.AssertExpectations(t)
	})
}

func TestDashboardDoctorAdmin(t *testing.T) {
	data := mocks.NewDoctorDataInterface(t)
	cld := mockUtil.NewCloudinaryInterface(t)
	email := mockHelper.NewEmailInterface(t)
	meet := mockHelper.NewMeetInterface(t)
	service := NewDoctor(data, cld, email, meet)
	var dashboard doctor.DoctorDashboardAdmin

	t.Run("Server Error", func(t *testing.T) {
		data.On("DoctorDashboardAdmin").Return(dashboard, errors.New("Process Failed")).Once()

		res, err := service.DoctorDashboardAdmin()

		assert.Error(t, err)
		assert.NotNil(t, res)
		assert.EqualError(t, err, "Process Failed")
	})

	t.Run("Success Get", func(t *testing.T) {
		data.On("DoctorDashboardAdmin").Return(dashboard, nil).Once()

		res, err := service.DoctorDashboardAdmin()

		assert.Nil(t, err)
		assert.NotNil(t, res)
		data.AssertExpectations(t)
	})
}

func TestDenyDoctor(t *testing.T) {
	data := mocks.NewDoctorDataInterface(t)
	cld := mockUtil.NewCloudinaryInterface(t)
	email := mockHelper.NewEmailInterface(t)
	meet := mockHelper.NewMeetInterface(t)
	service := NewDoctor(data, cld, email, meet)
	var doctor doctor.DoctorAll

	t.Run("Get By ID Error", func(t *testing.T) {
		data.On("GetDoctorByUserId", 1).Return(nil, errors.New("Get Doctor By User ID Error")).Once()

		result, err := service.DenyDoctor(1)

		assert.Error(t, err)
		assert.EqualError(t, err, "Get Doctor By User ID Error")
		assert.Equal(t, false, result)
	})

	t.Run("Success Deny", func(t *testing.T) {
		data.On("GetDoctorByUserId", 1).Return(&doctor, nil).Once()
		data.On("DenyDoctor", int(doctor.ID)).Return(true, nil).Once()

		result, err := service.DenyDoctor(1)

		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, true, result)
		data.AssertExpectations(t)
		cld.AssertExpectations(t)
	})

	t.Run("Server Error", func(t *testing.T) {
		data.On("GetDoctorByUserId", 1).Return(&doctor, nil).Once()
		data.On("DenyDoctor", int(doctor.ID)).Return(false, errors.New("Deny Process Failed")).Once()

		result, err := service.DenyDoctor(1)

		assert.Error(t, err)
		assert.EqualError(t, err, "Deny Process Failed")
		assert.Equal(t, false, result)
	})
}

func TestApproveDoctor(t *testing.T) {
	data := mocks.NewDoctorDataInterface(t)
	cld := mockUtil.NewCloudinaryInterface(t)
	email := mockHelper.NewEmailInterface(t)
	meet := mockHelper.NewMeetInterface(t)
	service := NewDoctor(data, cld, email, meet)
	var doctor doctor.DoctorAll

	t.Run("Get By ID Error", func(t *testing.T) {
		data.On("GetDoctorByUserId", 1).Return(nil, errors.New("Get Doctor By User ID Error")).Once()

		result, err := service.ApproveDoctor(1)

		assert.Error(t, err)
		assert.EqualError(t, err, "Get Doctor By User ID Error")
		assert.Equal(t, false, result)
	})

	t.Run("Success Approve", func(t *testing.T) {
		data.On("GetDoctorByUserId", 1).Return(&doctor, nil).Once()
		data.On("ApproveDoctor", int(doctor.ID)).Return(true, nil).Once()

		result, err := service.ApproveDoctor(1)

		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, true, result)
		data.AssertExpectations(t)
		cld.AssertExpectations(t)
	})

	t.Run("Server Error", func(t *testing.T) {
		data.On("GetDoctorByUserId", 1).Return(&doctor, nil).Once()
		data.On("ApproveDoctor", int(doctor.ID)).Return(false, errors.New("Approve Process Failed")).Once()

		result, err := service.ApproveDoctor(1)

		assert.Error(t, err)
		assert.EqualError(t, err, "Approve Process Failed")
		assert.Equal(t, false, result)
	})
}
