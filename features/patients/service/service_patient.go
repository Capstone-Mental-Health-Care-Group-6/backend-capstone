package service

import (
	"FinalProject/features/patients"
	"FinalProject/helper"
	"FinalProject/utils/cloudinary"
	"errors"
)

type PatientService struct {
	data patients.PatientDataInterface
	cld  cloudinary.CloudinaryInterface
	jwt  helper.JWTInterface
}

func New(data patients.PatientDataInterface, cloudinary cloudinary.CloudinaryInterface) patients.PatientServiceInterface {
	return &PatientService{
		data: data,
		cld:  cloudinary,
	}
}

func (psvc *PatientService) GetPatients() ([]patients.Patiententity, error) {
	result, err := psvc.data.GetAll()
	if err != nil {
		return nil, errors.New("get All Process Failed")
	}
	return result, nil
}

func (psvc *PatientService) GetPatient(id int) ([]patients.Patiententity, error) {
	result, err := psvc.data.GetByID(id)
	if err != nil {
		return nil, errors.New("get By ID Process Failed")
	}
	return result, nil
}

func (psvc *PatientService) CreatePatient(newData patients.Patiententity) (*patients.Patiententity, error) {
	result, err := psvc.data.Insert(newData)
	if err != nil {
		return nil, errors.New("insert Process Failed")
	}
	return result, nil
}

func (psvc *PatientService) PhotoUpload(newData patients.AvatarPhoto) (string, error) {
	uploadUrl, err := psvc.cld.UploadImageHelper(newData.Avatar)
	if err != nil {
		return "", errors.New("upload Avatar Failed")
	}
	return uploadUrl, nil
}
