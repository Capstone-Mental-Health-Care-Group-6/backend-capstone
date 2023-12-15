package service

import (
	"FinalProject/features/patients"
	"FinalProject/helper"
	"FinalProject/helper/enkrip"
	"FinalProject/utils/cloudinary"
	"errors"
	"strings"
)

type PatientService struct {
	data   patients.PatientDataInterface
	cld    cloudinary.CloudinaryInterface
	jwt    helper.JWTInterface
	enkrip enkrip.HashInterface
}

func NewPatient(data patients.PatientDataInterface, cloudinary cloudinary.CloudinaryInterface, jwt helper.JWTInterface, enkrip enkrip.HashInterface) patients.PatientServiceInterface {
	return &PatientService{
		data:   data,
		cld:    cloudinary,
		jwt:    jwt,
		enkrip: enkrip,
	}
}

func (psvc *PatientService) GetPatients(status, name string) ([]patients.Patientdetail, error) {
	result, err := psvc.data.GetAll(status, name)
	if err != nil {
		return nil, errors.New("Get All Process Failed")
	}
	return result, nil
}

func (psvc *PatientService) GetPatient(id int) (patients.Patientdetail, error) {
	result, err := psvc.data.GetByID(id)
	if err != nil {
		return result, errors.New("Get By ID Process Failed")
	}
	return result, nil
}

func (psvc *PatientService) CreatePatient(newData patients.Patiententity) (*patients.Patiententity, error) {
	hashPassword, err := psvc.enkrip.HashPassword(newData.Password)
	if err != nil {
		return nil, errors.New("Hash Password Error")
	}

	newData.Password = hashPassword
	result, err := psvc.data.Insert(newData)
	if err != nil {
		return nil, errors.New("Insert Process Failed")
	}
	return result, nil
}

func (psvc *PatientService) PhotoUpload(newData patients.AvatarPhoto) (string, error) {
	uploadUrl, err := psvc.cld.UploadImageHelper(newData.Avatar)
	if err != nil {
		return "", errors.New("Upload Avatar Failed")
	}
	return uploadUrl, nil
}

func (psvc *PatientService) UpdatePatient(id int, newData patients.UpdateProfile) (bool, error) {
	result, err := psvc.data.Update(id, newData)
	if err != nil {
		return false, errors.New("Update Process Failed")
	}
	return result, nil
}

func (psvc *PatientService) LoginPatient(email, password string) (*patients.PatientCredential, error) {
	result, err := psvc.data.LoginPatient(email, password)
	if err != nil {
		if strings.Contains(err.Error(), "Incorrect Password") {
			return nil, errors.New("Incorrect Password")
		}
		if strings.Contains(err.Error(), "Not Found") {
			return nil, errors.New("User Not Found / User Inactive")
		}
		return nil, errors.New("Login Process Failed")
	}

	tokenData := psvc.jwt.GenerateJWT(result.ID, result.Role, result.Status)

	if tokenData == nil {
		return nil, errors.New("Token Process Failed")
	}

	response := new(patients.PatientCredential)
	response.Name = result.Name
	response.Email = result.Email
	response.Access = tokenData

	return response, nil
}

func (psvc *PatientService) UpdatePassword(id int, newData patients.UpdatePassword) (bool, error) {
	hashPassword, err := psvc.enkrip.HashPassword(newData.Password)
	if err != nil {
		return false, errors.New("Hash Password Failed")
	}

	newData.Password = hashPassword
	result, err := psvc.data.UpdatePassword(id, newData)
	if err != nil {
		return false, errors.New("Update Password Failed")
	}
	return result, nil
}

func (psvc *PatientService) PatientDashboard() (patients.PatientDashboard, error) {
	res, err := psvc.data.PatientDashboard()

	if err != nil {
		return res, errors.New("Process Failed")
	}

	return res, nil
}

func (psvc *PatientService) UpdateStatus(id int, newData patients.UpdateStatus) (bool, error) {
	result, err := psvc.data.UpdateStatus(id, newData)
	if err != nil {
		return false, errors.New("Update Status Process Failed")
	}
	return result, nil
}

func (psvc *PatientService) InactivateAccount(id int) (bool, error) {
	result, err := psvc.data.InactivateAccount(id)

	if err != nil {
		return false, errors.New("Inactivate Account Process Failed")
	}

	return result, nil
}
