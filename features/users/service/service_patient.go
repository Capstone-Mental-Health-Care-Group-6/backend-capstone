package service

import (
	"FinalProject/features/users"
	"FinalProject/helper"
	"FinalProject/utils/cloudinary"
	"errors"
)

type PatientService struct {
	data users.PatientDataInterface
	cld  cloudinary.CloudinaryInterface
	jwt  helper.JWTInterface
}

func NewPatient(data users.PatientDataInterface, cloudinary cloudinary.CloudinaryInterface) users.PatientServiceInterface {
	return &PatientService{
		data: data,
		cld:  cloudinary,
	}
}

func (psvc *PatientService) GetPatients() ([]users.Patiententity, error) {
	result, err := psvc.data.GetAll()
	if err != nil {
		return nil, errors.New("get All Process Failed")
	}
	return result, nil
}

func (psvc *PatientService) GetPatient(id int) ([]users.Patiententity, error) {
	result, err := psvc.data.GetByID(id)
	if err != nil {
		return nil, errors.New("get By ID Process Failed")
	}
	return result, nil
}

func (psvc *PatientService) CreatePatient(newData users.Patiententity) (*users.Patiententity, error) {
	result, err := psvc.data.Insert(newData)
	if err != nil {
		return nil, errors.New("insert Process Failed")
	}
	return result, nil
}

func (psvc *PatientService) PhotoUpload(newData users.AvatarPhoto) (string, error) {
	uploadUrl, err := psvc.cld.UploadImageHelper(newData.Avatar)
	if err != nil {
		return "", errors.New("upload Avatar Failed")
	}
	return uploadUrl, nil
}
