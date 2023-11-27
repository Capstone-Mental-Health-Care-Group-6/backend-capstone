package service

import (
	"FinalProject/features/patients"
	"FinalProject/helper"
	"FinalProject/utils/cloudinary"
	"errors"
	"strings"
)

type PatientService struct {
	data patients.PatientDataInterface
	cld  cloudinary.CloudinaryInterface
	jwt  helper.JWTInterface
}

func NewPatient(data patients.PatientDataInterface, cloudinary cloudinary.CloudinaryInterface, jwt helper.JWTInterface) patients.PatientServiceInterface {
	return &PatientService{
		data: data,
		cld:  cloudinary,
		jwt:  jwt,
	}
}

func (psvc *PatientService) GetPatients() ([]patients.Patientdetail, error) {
	result, err := psvc.data.GetAll()
	if err != nil {
		return nil, errors.New("get All Process Failed")
	}
	return result, nil
}

func (psvc *PatientService) GetPatient(id int) (patients.Patientdetail, error) {
	result, err := psvc.data.GetByID(id)
	if err != nil {
		return result, errors.New("get By ID Process Failed")
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

func (psvc *PatientService) UpdatePatient(id int, newData patients.UpdateProfile) (bool, error) {
	result, err := psvc.data.Update(id, newData)
	if err != nil {
		return false, errors.New("update Process Failed")
	}
	return result, nil
}

func (psvc *PatientService) LoginPatient(email, password string) (*patients.PatientCredential, error) {
	result, err := psvc.data.LoginPatient(email, password)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return nil, errors.New("data not found")
		}
		return nil, errors.New("login Process Failed")
	}

	tokenData := psvc.jwt.GenerateJWT(result.ID, result.Role, result.Status)

	if tokenData == nil {
		return nil, errors.New("token Process Failed")
	}

	response := new(patients.PatientCredential)
	response.Name = result.Name
	response.Email = result.Email
	response.Access = tokenData

	return response, nil
}

func (psvc *PatientService) UpdatePassword(id int, newData patients.UpdatePassword) (bool, error) {
	result, err := psvc.data.UpdatePassword(id, newData)
	if err != nil {
		return false, errors.New("update Process Failed")
	}
	return result, nil
}