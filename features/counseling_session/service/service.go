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

	// uploadUrl, err := s.cld.UploadImageHelper(file.Avatar)
	// if err != nil {
	// 	return nil, errors.New("Upload Failed")
	// }

	// newData := counselingsession.counselingsession{
	// 	Name:         input.Name,
	// 	Sessions:     input.Sessions,
	// 	Type:         input.Type,
	// 	Price:        input.Price,
	// 	Description:  input.Description,
	// 	ActivePriode: input.ActivePriode,
	// 	Avatar:       uploadUrl,
	// }

	// result, err := s.d.Create(newData)
	// if err != nil {
	// 	return nil, err
	// }

	return nil, nil
}

func (s *CounselingSessionService) GetCounseling(id int) (*counselingsession.CounselingSession, error) {
	result, err := s.d.GetById(id)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *CounselingSessionService) UpdateCounseling(id int, input counselingsession.CounselingSession) (bool, error) {

	// newData := counselingsession.counselingsession{
	// 	Name:         input.Name,
	// 	Sessions:     input.Sessions,
	// 	Type:         input.Type,
	// 	Price:        input.Price,
	// 	Description:  input.Description,
	// 	ActivePriode: input.ActivePriode,
	// }

	// if file.Avatar != nil {
	// 	uploadUrl, err := s.cld.UploadImageHelper(file.Avatar)
	// 	if err != nil {
	// 		return false, errors.New("Upload Failed")
	// 	}

	// 	newData.Avatar = uploadUrl
	// }

	// result, err := s.d.Update(id, newData)
	// if err != nil {
	// 	return false, err
	// }

	return true, nil
}

func (s *CounselingSessionService) DeleteCounseling(id int) (bool, error) {
	result, err := s.d.Delete(id)
	if err != nil {
		return false, err
	}

	return result, nil
}
