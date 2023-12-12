package data

import (
	counselingtopics "FinalProject/features/counseling_topics"

	"gorm.io/gorm"
)

type CounselingTopicData struct {
	db *gorm.DB
}

func New(db *gorm.DB) counselingtopics.CounselingTopicDataInterface {
	return &CounselingTopicData{
		db: db,
	}
}

func (cmd *CounselingTopicData) GetAll() ([]counselingtopics.CounselingTopicInfo, error) {
	var data = []counselingtopics.CounselingTopicInfo{}
	if err := cmd.db.Table("counseling_topics").Scan(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}
func (cmd *CounselingTopicData) GetByID(id int) ([]counselingtopics.CounselingTopicInfo, error) {
	var data = []counselingtopics.CounselingTopicInfo{}
	if err := cmd.db.Table("counseling_topics").Where("id = ?", id).Scan(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}
