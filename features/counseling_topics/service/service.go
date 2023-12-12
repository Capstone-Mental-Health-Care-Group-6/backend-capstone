package service

import (
	counselingtopics "FinalProject/features/counseling_topics"
	"errors"
)

type CounselingTopicService struct {
	data counselingtopics.CounselingTopicDataInterface
}

func New(data counselingtopics.CounselingTopicDataInterface) counselingtopics.CounselingTopicServiceInterface {
	return &CounselingTopicService{
		data: data,
	}
}

func (cms *CounselingTopicService) GetAll() ([]counselingtopics.CounselingTopicInfo, error) {
	res, err := cms.data.GetAll()

	if err != nil {
		return nil, errors.New("Get All Process Failed")
	}

	return res, nil
}

func (cms *CounselingTopicService) GetByID(id int) ([]counselingtopics.CounselingTopicInfo, error) {
	res, err := cms.data.GetByID(id)

	if err != nil {
		return nil, errors.New("Get By ID Process Failed")
	}

	return res, nil
}
