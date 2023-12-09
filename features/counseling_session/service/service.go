package service

import (
	counselingsession "FinalProject/features/counseling_session"
	"FinalProject/utils/cloudinary"
)

type CounselingSessionService struct {
	d   counselingsession.CounselingSessionDataInterface
	cld cloudinary.CloudinaryInterface
}

func New(data counselingsession.CounselingSessionDataInterface, cld cloudinary.CloudinaryInterface) counselingsession.CounselingSessionServiceInterface {
	return &CounselingSessionService{
		d:   data,
		cld: cld,
	}
}

func (s *CounselingSessionService) GetAllCounseling() ([]counselingsession.CounselingSession, error) {
	result, err := s.d.GetAll()
	if err != nil {
		return nil, err
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
		return nil, err
	}

	return result, nil
}

func (s *CounselingSessionService) GetCounseling(id int) (*counselingsession.CounselingSession, error) {
	result, err := s.d.GetById(id)
	if err != nil {
		return nil, err
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
		return false, err
	}

	return result, nil
}

func (s *CounselingSessionService) DeleteCounseling(id int) (bool, error) {
	result, err := s.d.Delete(id)
	if err != nil {
		return false, err
	}

	return result, nil
}
