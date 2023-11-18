package service

import (
	"FinalProject/features/doctor"
	"FinalProject/helper"
	"FinalProject/utils/cloudinary"
	"errors"
)

type DoctorService struct {
	data doctor.DoctorDataInterface
	cld  cloudinary.CloudinaryInterface
	jwt  helper.JWTInterface
}

func NewDoctor(data doctor.DoctorDataInterface, cloudinary cloudinary.CloudinaryInterface) doctor.DoctorServiceInterface {
	return &DoctorService{
		data: data,
		cld:  cloudinary,
	}
}

func (psvc *DoctorService) GetDoctors() ([]doctor.Doctor, error) {
	result, err := psvc.data.GetAll()
	if err != nil {
		return nil, errors.New("get All Process Failed")
	}
	return result, nil
}

func (psvc *DoctorService) GetDoctor(id int) ([]doctor.Doctor, error) {
	result, err := psvc.data.GetByID(id)
	if err != nil {
		return nil, errors.New("get By ID Process Failed")
	}
	return result, nil
}

func (psvc *DoctorService) CreateDoctor(newData doctor.Doctor) (*doctor.Doctor, error) {
	result, err := psvc.data.Insert(newData)
	if err != nil {
		return nil, errors.New("insert Process Failed")
	}
	return result, nil
}

func (psvc *DoctorService) DoctorAvatarUpload(newData doctor.DoctorAvatarPhoto) (string, error) {
	uploadUrl, err := psvc.cld.UploadImageHelper(newData.DoctorAvatar)
	if err != nil {
		return "", errors.New("Upload Avatar Failed")
	}
	return uploadUrl, nil
}

func (psvc *DoctorService) DoctorSTRUpload(newData doctor.DoctorSTRFileDataModel) (string, error) {
	uploadUrl, err := psvc.cld.UploadImageHelper(newData.DoctorSTRFile)
	if err != nil {
		return "", errors.New("Upload Surat Tanda Registrasi Failed")
	}
	return uploadUrl, nil
}

func (psvc *DoctorService) DoctorSIPPUpload(newData doctor.DoctorSIPPFileDataModel) (string, error) {
	uploadUrl, err := psvc.cld.UploadImageHelper(newData.DoctorSIPPFile)
	if err != nil {
		return "", errors.New("Upload SIPP Failed")
	}
	return uploadUrl, nil
}

func (psvc *DoctorService) DoctorCVUpload(newData doctor.DoctorCVDataModel) (string, error) {
	uploadUrl, err := psvc.cld.UploadImageHelper(newData.DoctorCV)
	if err != nil {
		return "", errors.New("Upload CV Failed")
	}
	return uploadUrl, nil
}

func (psvc *DoctorService) DoctorIjazahUpload(newData doctor.DoctorIjazahDataModel) (string, error) {
	uploadUrl, err := psvc.cld.UploadImageHelper(newData.DoctorIjazah)
	if err != nil {
		return "", errors.New("Upload Ijazah Failed")
	}
	return uploadUrl, nil
}
