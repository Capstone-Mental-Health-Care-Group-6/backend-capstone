package data

import (
	counselingsession "FinalProject/features/counseling_session"

	"gorm.io/gorm"
)

type CounselingSessionData struct {
	db *gorm.DB
}

func New(db *gorm.DB) counselingsession.CounselingSessionDataInterface {
	return &CounselingSessionData{
		db: db,
	}
}

func (bc *CounselingSessionData) GetAll() ([]counselingsession.CounselingSession, error) {
	var listCounselingSession = []counselingsession.CounselingSession{}

	var qry = bc.db.Table("bundle_counseling").
		Select("bundle_counseling.*").
		Where("bundle_counseling.deleted_at is null").
		Where("bundle_counseling.type = ?", "PREMIUM").
		Scan(&listCounselingSession)

	if err := qry.Error; err != nil {
		return nil, err
	}

	return listCounselingSession, nil
}

func (bc *CounselingSessionData) Create(input counselingsession.CounselingSession) (*counselingsession.CounselingSession, error) {
	// var newCounselingSession = &CounselingSession.CounselingSession{
	// 	Name:         input.Name,
	// 	Sessions:     input.Sessions,
	// 	Type:         input.Type,
	// 	Price:        input.Price,
	// 	Description:  input.Description,
	// 	ActivePriode: input.ActivePriode,
	// 	Avatar:       input.Avatar,
	// }

	var newData = new(counselingsession.CounselingSession)
	// MASUKIN DATA

	if err := bc.db.Table("bundle_counseling").Create(newData).Error; err != nil {
		return nil, err
	}

	return &input, nil
}

func (bc *CounselingSessionData) GetById(id int) (*counselingsession.CounselingSession, error) {
	var result = new(counselingsession.CounselingSession)

	var qry = bc.db.Table("bundle_counseling").
		Select("bundle_counseling.*").
		Where("bundle_counseling.deleted_at is null").
		Where("bundle_counseling.id = ?", id).
		Scan(&result)

	if err := qry.Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (bc *CounselingSessionData) Update(id int, input counselingsession.CounselingSession) (bool, error) {
	// var newData = map[string]interface{}{
	// 	"name":          input.Name,
	// 	"sessions":      input.Sessions,
	// 	"type":          input.Type,
	// 	"price":         input.Price,
	// 	"description":   input.Description,
	// 	"active_priode": input.ActivePriode,
	// }

	// if input.Avatar != "" {
	// 	newData["avatar"] = input.Avatar
	// }

	//UPDATE DATA

	// if err := bc.db.Table("counseling_session").Where("id = ?", id).Updates(newData).Error; err != nil {
	// 	return false, err
	// }

	return true, nil
}

func (bc *CounselingSessionData) Delete(id int) (bool, error) {
	var deleteData = new(counselingsession.CounselingSession)

	if err := bc.db.Table("counseling_session").Where("id = ?", id).Delete(&deleteData).Error; err != nil {
		return false, err
	}

	return true, nil
}
