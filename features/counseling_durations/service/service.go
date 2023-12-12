package service

import (
	counselingduration "FinalProject/features/counseling_durations"
	"errors"
)

type CounselingDurationService struct {
	data counselingduration.CounselingDurationDataInterface
}

func New(data counselingduration.CounselingDurationDataInterface) counselingduration.CounselingDurationServiceInterface {
	return &CounselingDurationService{
		data: data,
	}
}

func (cds *CounselingDurationService) GetAll() ([]counselingduration.CounselingDurationInfo, error) {
	res, err := cds.data.GetAll()

	if err != nil {
		return nil, errors.New("Get All Process Failed")
	}

	return res, nil
}
func (cds *CounselingDurationService) GetByID(id int) ([]counselingduration.CounselingDurationInfo, error) {
	res, err := cds.data.GetByID(id)

	if err != nil {
		return nil, errors.New("Get By ID Process Failed")
	}

	return res, nil
}
