package service

import (
	counselingmethod "FinalProject/features/counseling_methods"
	"errors"
)

type CounselingMethodService struct {
	data counselingmethod.CounselingMethodDataInterface
}

func New(data counselingmethod.CounselingMethodDataInterface) counselingmethod.CounselingMethodServiceInterface {
	return &CounselingMethodService{
		data: data,
	}
}

func (cms *CounselingMethodService) GetAll() ([]counselingmethod.CounselingMethodInfo, error) {
	res, err := cms.data.GetAll()

	if err != nil {
		return nil, errors.New("Get All Process Failed")
	}

	return res, nil
}

func (cms *CounselingMethodService) GetByID(id int) ([]counselingmethod.CounselingMethodInfo, error) {
	res, err := cms.data.GetByID(id)

	if err != nil {
		return nil, errors.New("Get By ID Process Failed")
	}

	return res, nil
}
