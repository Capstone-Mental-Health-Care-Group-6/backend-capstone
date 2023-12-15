package service

import (
	counselingsession "FinalProject/features/counseling_session"
	"errors"
)

type CounselingSessionService struct {
	d counselingsession.CounselingSessionDataInterface
}

func New(data counselingsession.CounselingSessionDataInterface) counselingsession.CounselingSessionServiceInterface {
	return &CounselingSessionService{
		d: data,
	}
}

func (s *CounselingSessionService) GetAllCounseling() ([]counselingsession.CounselingSession, error) {
	result, err := s.d.GetAll()
	if err != nil {
		return nil, errors.New("Get All Process Failed")
	}
	return result, nil
}

func (s *CounselingSessionService) CreateCounseling(input counselingsession.CounselingSession) (*counselingsession.CounselingSession, error) {

	newData := counselingsession.CounselingSession{
		TransactionID: input.TransactionID,
		Date:          input.Date,
		Time:          input.Time,
		Duration:      input.Duration,
		Status:        input.Status,
	}

	result, err := s.d.Create(newData)
	if err != nil {
		return nil, errors.New("Create Process Failed")
	}

	return result, nil
}

func (s *CounselingSessionService) GetCounseling(id int) (*counselingsession.CounselingSession, error) {
	result, err := s.d.GetById(id)
	if err != nil {
		return nil, errors.New("Get By ID Process Failed")
	}

	return result, nil
}

func (s *CounselingSessionService) GetAllCounselingByUserID(userID int) ([]counselingsession.CounselingSession, error) {
	result, err := s.d.GetAllCounselingByUserID(userID)
	if err != nil {
		return nil, errors.New("Get All By ID Process Failed")
	}

	return result, nil
}
func (s *CounselingSessionService) UpdateCounseling(id int, input counselingsession.CounselingSession) (bool, error) {

	newData := counselingsession.CounselingSession{
		TransactionID: input.TransactionID,
		Date:          input.Date,
		Time:          input.Time,
		Duration:      input.Duration,
		Status:        input.Status,
	}

	result, err := s.d.Update(id, newData)
	if err != nil {
		return false, errors.New("Update Process Failed")
	}

	return result, nil
}

func (s *CounselingSessionService) DeleteCounseling(id int) (bool, error) {
	result, err := s.d.Delete(id)
	if err != nil {
		return false, errors.New("Delete Process Failed")
	}

	return result, nil
}

func (s *CounselingSessionService) ApprovePatient(id, doctorID int) (bool, error) {
	err := s.d.CheckPatient(id, doctorID)

	if err != nil {
		return false, err
	}

	res, err := s.d.ApprovePatient(id)
	if err != nil {
		return false, errors.New("Approve Patient Process Failed")
	}
	return res, nil
}

func (s *CounselingSessionService) RejectPatient(id, doctorID int, newData counselingsession.StatusUpdate) (bool, error) {
	err := s.d.CheckPatient(id, doctorID)

	if err != nil {
		return false, err
	}

	res, err := s.d.RejectPatient(id, newData)
	if err != nil {
		return false, errors.New("Reject Patient Process Failed")
	}
	return res, nil
}
