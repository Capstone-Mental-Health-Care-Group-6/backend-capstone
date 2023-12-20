package data

import (
	counselingsession "FinalProject/features/counseling_session"
	"FinalProject/features/doctor"
	"errors"
	"fmt"

	transaction "FinalProject/features/transaction"

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

	if qry.Error != nil {
		return nil, qry.Error
	}

	for i, result := range listCounselingSession {
		existingData := transaction.TransactionInfo{}

		var qryTransaction = bc.db.Table("transactions").Select(`
	        transactions.*,
	        counseling_topics.name as topic_name,
	        patient_accounts.name as patient_name,
	        patient_accounts.avatar as patient_avatar,
	        doctors.doctor_name as doctor_name,
			doctors.doctor_avatar as doctor_avatar,
	        counseling_methods.name as method_name,
	        counseling_durations.name as duration_name,
	        transactions.created_at,
	        transactions.updated_at,
	        doctors_rating.id as doctor_rating_id,
	        doctors_rating.doctor_star_rating as doctor_star_rating,
	        doctors_rating.doctor_review as doctor_review
	    `).
			Joins("LEFT JOIN counseling_topics ON counseling_topics.id = transactions.topic_id").
			Joins("LEFT JOIN patient_accounts ON patient_accounts.id = transactions.patient_id").
			Joins("LEFT JOIN doctors ON doctors.id = transactions.doctor_id").
			Joins("LEFT JOIN counseling_methods ON counseling_methods.id = transactions.method_id").
			Joins("LEFT JOIN counseling_durations ON counseling_durations.id = transactions.duration_id").
			Joins("LEFT JOIN doctors_rating ON doctors_rating.transaction_id = transactions.midtrans_id").
			Where("transactions.id = ?", result.TransactionID).
			Where("transactions.deleted_at is null").
			Scan(&existingData)

		if qryTransaction.Error != nil {
			return nil, qryTransaction.Error
		}

		listCounselingSession[i].CounselingType = existingData.CounselingType
		listCounselingSession[i].CounselingMethod = existingData.MethodName
		listCounselingSession[i].CounselingTopic = existingData.TopicName
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

	existingData := transaction.TransactionInfo{}

	var qryTransaction = bc.db.Table("transactions").Select(`
        transactions.*,
        counseling_topics.name as topic_name,
        patient_accounts.name as patient_name,
        patient_accounts.avatar as patient_avatar,
        doctors.doctor_name as doctor_name,
        doctors.doctor_avatar as doctor_avatar,
        counseling_methods.name as method_name,
        counseling_durations.name as duration_name,
        transactions.created_at,
        transactions.updated_at,
        doctors_rating.id as doctor_rating_id,
        doctors_rating.doctor_star_rating as doctor_star_rating,
        doctors_rating.doctor_review as doctor_review
    `).
		Joins("LEFT JOIN counseling_topics ON counseling_topics.id = transactions.topic_id").
		Joins("LEFT JOIN patient_accounts ON patient_accounts.id = transactions.patient_id").
		Joins("LEFT JOIN doctors ON doctors.id = transactions.doctor_id").
		Joins("LEFT JOIN counseling_methods ON counseling_methods.id = transactions.method_id").
		Joins("LEFT JOIN counseling_durations ON counseling_durations.id = transactions.duration_id").
		Joins("LEFT JOIN doctors_rating ON doctors_rating.transaction_id = transactions.midtrans_id").
		Where("transactions.id = ?", result.TransactionID).
		Where("transactions.deleted_at is null").
		Scan(&existingData)

	if qryTransaction.Error != nil {
		return nil, qry.Error
	}

	result.CounselingType = existingData.CounselingType
	result.CounselingMethod = existingData.MethodName
	result.CounselingTopic = existingData.TopicName

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

	if qry.Error != nil {
		return nil, qry.Error
	}

	for i, result := range listCounselingSession {
		existingData := transaction.TransactionInfo{}

		var qryTransaction = bc.db.Table("transactions").Select(`
	        transactions.*,
	        counseling_topics.name as topic_name,
	        patient_accounts.name as patient_name,
	        patient_accounts.avatar as patient_avatar,
	        doctors.doctor_name as doctor_name,
			doctors.doctor_avatar as doctor_avatar,
	        counseling_methods.name as method_name,
	        counseling_durations.name as duration_name,
	        transactions.created_at,
	        transactions.updated_at,
	        doctors_rating.id as doctor_rating_id,
	        doctors_rating.doctor_star_rating as doctor_star_rating,
	        doctors_rating.doctor_review as doctor_review
	    `).
			Joins("LEFT JOIN counseling_topics ON counseling_topics.id = transactions.topic_id").
			Joins("LEFT JOIN patient_accounts ON patient_accounts.id = transactions.patient_id").
			Joins("LEFT JOIN doctors ON doctors.id = transactions.doctor_id").
			Joins("LEFT JOIN counseling_methods ON counseling_methods.id = transactions.method_id").
			Joins("LEFT JOIN counseling_durations ON counseling_durations.id = transactions.duration_id").
			Joins("LEFT JOIN doctors_rating ON doctors_rating.transaction_id = transactions.midtrans_id").
			Where("transactions.id = ?", result.TransactionID).
			Where("transactions.deleted_at is null").
			Scan(&existingData)

		if qryTransaction.Error != nil {
			return nil, qryTransaction.Error
		}

		listCounselingSession[i].CounselingType = existingData.CounselingType
		listCounselingSession[i].CounselingMethod = existingData.MethodName
		listCounselingSession[i].CounselingTopic = existingData.TopicName
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

func (bc *CounselingSessionData) CheckPatient(id, doctorID int) error {
	var listCounselingSession = counselingsession.CounselingSession{}

	if err := bc.db.Table("counseling_session").Where("id = ?", id).First(&listCounselingSession).Error; err != nil {
		return errors.New("No data counseling found")
	}

	existingDataPatient := transaction.Transaction{}
	if err := bc.db.Table("transactions").Where("id = ?", listCounselingSession.TransactionID).Scan(&existingDataPatient).Error; err != nil {
		return errors.New("No transaction data found")
	}

	dataDoctor := doctor.Doctor{}
	if err := bc.db.Table("doctors").Where("user_id = ?", doctorID).First(&dataDoctor).Error; err != nil {
		return errors.New("No data doctor found")
	}

	fmt.Println(existingDataPatient.DoctorID)
	fmt.Println(dataDoctor.ID)

	if existingDataPatient.DoctorID == dataDoctor.ID {
		return nil
	}

	return errors.New("Unauthorized permission for this doctor")

}
