package data

import (
	counselingsession "FinalProject/features/counseling_session"
	"errors"

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

	var qry = bc.db.Table("counseling_session").
		Select("counseling_session.*").
		Where("counseling_session.deleted_at is null").
		Scan(&listCounselingSession)

	if err := qry.Error; err != nil {
		return nil, err
	}

	return listCounselingSession, nil
}

func (bc *CounselingSessionData) Create(input counselingsession.CounselingSession) (*counselingsession.CounselingSession, error) {

	var newData = new(counselingsession.CounselingSession)
	newData.TransactionID = input.TransactionID
	newData.Date = input.Date
	newData.Time = input.Time
	newData.Duration = input.Duration
	newData.Status = input.Status
	newData.Alasan = input.Alasan
	newData.DetailAlasan = input.DetailAlasan
	// MASUKIN DATA

	if err := bc.db.Table("counseling_session").Create(newData).Error; err != nil {
		return nil, err
	}

	return &input, nil
}

func (bc *CounselingSessionData) GetById(id int) (*counselingsession.CounselingSession, error) {
	var result = new(counselingsession.CounselingSession)

	var qry = bc.db.Table("counseling_session").
		Select("counseling_session.*").
		Where("counseling_session.deleted_at is null").
		Where("counseling_session.id = ?", id).
		Scan(&result)

	if err := qry.Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (bc *CounselingSessionData) GetAllCounselingByUserID(userID int) ([]counselingsession.CounselingSession, error) {
	var listCounselingSession = []counselingsession.CounselingSession{}

	var qry = bc.db.Table("counseling_session").
		Select("counseling_session.*").
		Where("counseling_session.user_id = ?", userID).
		Where("counseling_session.deleted_at is null").
		Scan(&listCounselingSession)

	if err := qry.Error; err != nil {
		return nil, err
	}

	return listCounselingSession, nil
}

// func (bc *CounselingSessionData) GetByUserId(id int) ([]counselingsession.CounselingSessionInfo, error) {
// 	// var result = new(counselingsession.CounselingSession)
// 	var counselingInfo []counselingsession.CounselingSessionInfo

// 	// var qry = bc.db.Table("counseling_session").
// 	// 	Select("counseling_session.*").
// 	// 	Where("counseling_session.deleted_at is null").
// 	// 	Where("counseling_session.id = ?", id).
// 	// 	Scan(&result)

// 		qry := bc.db.Table("counseling_session").
// 		Select("counseling_session.*").
// 		// Joins("LEFT JOIN doctors_expertise_relation ON doctors.id = doctors_expertise_relation.doctor_id").
// 		Where("counseling_session.deleted_at IS NULL").
// 		Where("counseling_session.id = ?", id).
// 		Scan(&doctors)

// 	if err := qry.Error; err != nil {
// 		return nil, err
// 	}

// 	return result, nil
// }

func (bc *CounselingSessionData) Update(id int, input counselingsession.CounselingSession) (bool, error) {
	var newData = map[string]interface{}{
		"transaction_id": input.TransactionID,
		"date":           input.Date,
		"time":           input.Time,
		"duration":       input.Duration,
		"status":         input.Status,
	}

	//UPDATE DATA

	if err := bc.db.Table("counseling_session").Where("id = ?", id).Updates(newData).Error; err != nil {
		return false, err
	}

	return true, nil
}

func (bc *CounselingSessionData) Delete(id int) (bool, error) {
	var deleteData = new(counselingsession.CounselingSession)

	if err := bc.db.Table("counseling_session").Where("id = ?", id).Delete(&deleteData).Error; err != nil {
		return false, err
	}

	return true, nil
}

func (bc *CounselingSessionData) ApprovePatient(id int) (bool, error) {
	var qry = bc.db.Table("counseling_session").Where("id = ?", id).Updates(CounselingSession{
		Status: "not_finished",
	})

	if err := qry.Error; err != nil {
		return false, err
	}

	if dataCount := qry.RowsAffected; dataCount < 1 {
		return false, errors.New("Update Data Error, No Data Affected")
	}

	return true, nil
}

func (bc *CounselingSessionData) RejectPatient(id int, newData counselingsession.StatusUpdate) (bool, error) {
	var qry = bc.db.Table("counseling_session").Where("id = ?", id).Updates(CounselingSession{
		Status: "rejected",
		Alasan: newData.Alasan,
	})

	if err := qry.Error; err != nil {
		return false, err
	}

	if dataCount := qry.RowsAffected; dataCount < 1 {
		return false, errors.New("Update Data Error, No Data Affected")
	}

	return true, nil
}
