package data

import (
	CounselingSession "FinalProject/features/counseling_session"
	"FinalProject/features/doctor/data"
	"FinalProject/features/transaction"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type TransactionData struct {
	db *gorm.DB
}

func New(db *gorm.DB) transaction.TransactionDataInterface {
	return &TransactionData{
		db: db,
	}
}

func (ad *TransactionData) GetAndUpdate(newData transaction.UpdateTransaction, id string) (bool, error) {

	var transaction Transaction
	// db := ad.db

	ad.db.Where("midtrans_id = ?", id).First(&transaction)
	fmt.Println("This is the id: ", transaction.ID)
	transactionID := transaction.ID

	fmt.Println("This is the new payment status: ", newData.PaymentStatus)

	qry := ad.db.Table("transactions").Where("id = ?", transactionID).Updates(Transaction{
		PaymentStatus: newData.PaymentStatus,
	})

	if qry.Error != nil {
		return false, nil
	}

	if err := qry.Error; err != nil {
		return false, err
	}

	if newData.PaymentStatus == 2 {
		if transaction.DoctorID == 0 {
			return false, errors.New("Doctor ID not found in transaction")
		}

		fmt.Println("This is the existing DoctorID: ", transaction.DoctorID)

		existingDataDoctor := data.Doctor{}
		if err := ad.db.Table("doctors").Where("id = ?", transaction.DoctorID).First(&existingDataDoctor).Error; err != nil {
			fmt.Printf("Error fetching doctor data: %v\n", err)
			return false, err
		}

		existingDataDoctorRelation := data.DoctorExpertiseRelation{}
		if err := ad.db.Table("doctors_expertise_relation").Where("doctor_id = ?", transaction.DoctorID).First(&existingDataDoctorRelation).Error; err != nil {
			fmt.Printf("Error fetching doctor data: %v\n", err)
			return false, err
		}

		newDoctorBalance := existingDataDoctor.DoctorBalance + transaction.PriceResult
		fmt.Println("This is the new Update Balance: ", newDoctorBalance)

		qryToDoctor := ad.db.Table("doctors").Where("id = ?", transaction.DoctorID).Updates(data.Doctor{
			DoctorBalance: newDoctorBalance,
		})

		if err := qryToDoctor.Error; err != nil {
			fmt.Printf("Error updating doctor balance: %v\n", err)
			return false, err
		}

		if dataCount := qryToDoctor.RowsAffected; dataCount < 1 {
			return false, errors.New("Update Data Error, No Data Affected")
		}

		var newData = new(CounselingSession.CounselingSession)
		newData.TransactionID = transaction.ID
		newData.DoctorAvatar = existingDataDoctor.DoctorAvatar
		newData.DoctorExpertise = existingDataDoctorRelation.ExpertiseID
		newData.DoctorName = existingDataDoctor.DoctorName
		newData.UserID = transaction.UserID
		newData.Date = time.Now()
		newData.Time = time.Now()
		newData.Duration = transaction.DurationID
		newData.Status = "pending"
		// MASUKIN DATA

		if err := ad.db.Table("counseling_session").Create(newData).Error; err != nil {
			return false, err
		}
	}

	return true, nil
}

func (ad *TransactionData) GetByIDMidtrans(id string) ([]transaction.TransactionInfo, error) {
	var transactionInfos []transaction.TransactionInfo
	// var qry = ad.db.Table("transactions").Where("midtrans_id = ?", id).Select("user_id, midtrans_id, payment_status, payment_type, price_result").Find(&transactionInfos)
	var qry = ad.db.Table("transactions").Select(`
        transactions.*,
        counseling_topics.name as topic_name,
        patient_accounts.name as patient_name,
        patient_accounts.avatar as patient_avatar,
        doctors.doctor_name as doctor_name,
        counseling_methods.name as method_name,
        counseling_durations.name as duration_name,
        transactions.created_at,
        transactions.updated_at,
        doctor_ratings.id as doctor_rating_id,
        doctor_ratings.doctor_star_rating,
        doctor_ratings.doctor_review
    `).
		Joins("LEFT JOIN counseling_topics ON counseling_topics.id = transactions.topic_id").
		Joins("LEFT JOIN patient_accounts ON patient_accounts.id = transactions.patient_id").
		Joins("LEFT JOIN doctors ON doctors.id = transactions.doctor_id").
		Joins("LEFT JOIN counseling_methods ON counseling_methods.id = transactions.method_id").
		Joins("LEFT JOIN counseling_durations ON counseling_durations.id = transactions.duration_id").
		Joins("LEFT JOIN doctor_ratings ON doctor_ratings.doctor_id = transactions.doctor_id").
		Where("transactions.midtrans_id = ?", id).
		Where("transactions.deleted_at is null")

	if qry.Error != nil {
		return nil, qry.Error
	}

	fmt.Println("Json Response for query:", transactionInfos)

	return transactionInfos, nil
}

// func (ad *TransactionData) GetAll() ([]transaction.TransactionInfo, error) {
// 	var listTransactions []transaction.TransactionInfo            // Change to a slice to hold multiple transactions
// 	var qry = ad.db.Table("transactions").Find(&listTransactions) // Fetch all transactions data from the table

// 	if qry.Error != nil {
// 		return nil, qry.Error
// 	}

// 	return listTransactions, nil
// }

func (ad *TransactionData) GetAll(sort string) ([]transaction.TransactionInfo, error) {
	var listTransactions []transaction.TransactionInfo
	var qry = ad.db.Table("transactions")

	if sort != "" {
		qry = qry.Where("payment_type = ?", sort)
	}

	qry = qry.Find(&listTransactions)

	if qry.Error != nil {
		return nil, qry.Error
	}

	return listTransactions, nil
}

func (ad *TransactionData) GetByID(id int, sort string) ([]transaction.TransactionInfo, error) {
	var transactionInfo []transaction.TransactionInfo

	var qry = ad.db.Table("transactions").Select(`
        transactions.*,
        counseling_topics.name as topic_name,
        patient_accounts.name as patient_name,
        patient_accounts.avatar as patient_avatar,
        doctors.doctor_name as doctor_name,
        counseling_methods.name as method_name,
        counseling_durations.name as duration_name,
        transactions.created_at,
        transactions.updated_at,
        doctor_ratings.id as doctor_rating_id,
        doctor_ratings.doctor_star_rating,
        doctor_ratings.doctor_review
    `).
		Joins("LEFT JOIN counseling_topics ON counseling_topics.id = transactions.topic_id").
		Joins("LEFT JOIN patient_accounts ON patient_accounts.id = transactions.patient_id").
		Joins("LEFT JOIN doctors ON doctors.id = transactions.doctor_id").
		Joins("LEFT JOIN counseling_methods ON counseling_methods.id = transactions.method_id").
		Joins("LEFT JOIN counseling_durations ON counseling_durations.id = transactions.duration_id").
		Joins("LEFT JOIN doctor_ratings ON doctor_ratings.doctor_id = transactions.doctor_id").
		Where("transactions.user_id = ?", id).
		Where("transactions.deleted_at is null")

	if sort != "" {
		qry = qry.Where("transactions.payment_type = ?", sort)
	}

	qry = qry.Find(&transactionInfo)

	if qry.Error != nil {
		return nil, qry.Error
	}

	return transactionInfo, nil
}

func (ad *TransactionData) Insert(newData transaction.Transaction) (*transaction.Transaction, error) {
	if newData.DoctorID == 0 {
		return nil, errors.New("Doctor ID not found")
	}

	fmt.Println("This is the existing DoctorID: ", newData.DoctorID)

	existingDataDoctor := data.Doctor{}
	if err := ad.db.Table("doctors").Where("id = ?", newData.DoctorID).First(&existingDataDoctor).Error; err != nil {
		return nil, errors.New("Doctor ID not found")
	}

	var dbData = new(Transaction)

	dbData.TopicID = newData.TopicID
	dbData.PatientID = newData.PatientID
	dbData.DoctorID = newData.DoctorID
	dbData.MethodID = newData.MethodID
	dbData.DurationID = newData.DurationID
	dbData.CounselingID = newData.CounselingID
	dbData.UserID = newData.UserID
	dbData.MidtransID = newData.MidtransID

	dbData.CounselingSession = newData.CounselingSession
	dbData.CounselingType = newData.CounselingType

	dbData.PriceMethod = newData.PriceMethod
	dbData.PriceDuration = newData.PriceDuration
	dbData.PriceCounseling = newData.PriceCounseling
	dbData.PriceResult = newData.PriceResult
	dbData.PaymentProof = newData.PaymentProof

	dbData.PaymentStatus = newData.PaymentStatus
	dbData.PaymentType = newData.PaymentType

	fmt.Println("Ive succeed create payment status", newData.PaymentStatus)

	if err := ad.db.Create(dbData).Error; err != nil {
		return nil, err
	}

	var dbDataRating = new(data.DoctorRating)
	dbDataRating.DoctorID = newData.DoctorID
	dbDataRating.DoctorReview = ""
	dbDataRating.DoctorStarRating = 0
	dbDataRating.PatientID = newData.PatientID

	if err := ad.db.Create(dbDataRating).Error; err != nil {
		return nil, err
	}

	return &newData, nil
}

func (ad *TransactionData) Delete(id int) (bool, error) {
	var transactionInfo transaction.Transaction
	var qry = ad.db.Table("transactions").Where("id = ?", id).Delete(&transactionInfo)

	if qry.Error != nil {
		return false, qry.Error
	}

	return true, nil
}

func (ad *TransactionData) Update(newData transaction.UpdateTransactionManual, id int) (bool, error) {
	existingData := Transaction{}
	if err := ad.db.Table("transactions").Where("id = ?", id).First(&existingData).Error; err != nil {
		return false, err
	}

	// Check if the existing PaymentStatus is 4 or 2
	if existingData.PaymentStatus == 4 || existingData.PaymentStatus == 2 {
		fmt.Println("Error: You cannot update a transaction that's already finished.")
		return false, errors.New("Transaction already finished")
	}

	var qry = ad.db.Table("transactions").Where("id = ?", id).Updates(Transaction{
		UserID:          newData.UserID,
		PriceMethod:     newData.PriceMethod,
		PriceDuration:   newData.PriceDuration,
		PriceCounseling: newData.PriceCounseling,
		PriceResult:     newData.PriceResult,
		PaymentStatus:   newData.PaymentStatus,
		PaymentType:     newData.PaymentType})

	if err := qry.Error; err != nil {
		return false, err
	}

	if dataCount := qry.RowsAffected; dataCount < 1 {
		return false, errors.New("Update Data Error, No Data Affected")
	}

	if newData.PaymentStatus == 2 {
		if existingData.DoctorID == 0 {
			return false, errors.New("Doctor ID not found in transaction")
		}

		fmt.Println("This is the existing DoctorID: ", existingData.DoctorID)

		existingDataDoctor := data.Doctor{}
		if err := ad.db.Table("doctors").Where("id = ?", existingData.DoctorID).First(&existingDataDoctor).Error; err != nil {
			fmt.Printf("Error fetching doctor data: %v\n", err)
			return false, err
		}

		existingDataDoctorRelation := data.DoctorExpertiseRelation{}
		if err := ad.db.Table("doctors_expertise_relation").Where("doctor_id = ?", existingData.DoctorID).First(&existingDataDoctorRelation).Error; err != nil {
			fmt.Printf("Error fetching doctor data: %v\n", err)
			return false, err
		}

		newDoctorBalance := existingDataDoctor.DoctorBalance + existingData.PriceResult
		fmt.Println("This is the new Update Balance: ", newDoctorBalance)

		qryToDoctor := ad.db.Table("doctors").Where("id = ?", existingData.DoctorID).Updates(data.Doctor{
			DoctorBalance: newDoctorBalance,
		})

		if err := qryToDoctor.Error; err != nil {
			fmt.Printf("Error updating doctor balance: %v\n", err)
			return false, err
		}

		if dataCount := qryToDoctor.RowsAffected; dataCount < 1 {
			return false, errors.New("Update Data Error, No Data Affected")
		}

		var newData = new(CounselingSession.CounselingSession)
		newData.TransactionID = existingData.ID
		newData.DoctorAvatar = existingDataDoctor.DoctorAvatar
		newData.DoctorExpertise = existingDataDoctorRelation.ExpertiseID
		newData.DoctorName = existingDataDoctor.DoctorName
		newData.UserID = existingData.UserID
		newData.Date = time.Now()
		newData.Time = time.Now()
		newData.Duration = existingData.DurationID
		newData.Status = "pending"
		// MASUKIN DATA

		if err := ad.db.Table("counseling_session").Create(newData).Error; err != nil {
			return false, err
		}
	}

	return true, nil
}

func (ad *TransactionData) UpdateWithTrxID(newData transaction.UpdateTransactionManual, id string) (bool, error) {
	// Fetch the existing transaction data
	existingData := Transaction{}
	if err := ad.db.Table("transactions").Where("midtrans_id = ?", id).First(&existingData).Error; err != nil {
		return false, err
	}

	// Check if the existing PaymentStatus is 4 or 2
	if existingData.PaymentStatus == 4 || existingData.PaymentStatus == 2 {
		fmt.Println("Error: You cannot update a transaction that's already finished.")
		return false, errors.New("Transaction already finished")
	}

	// Perform the update
	qry := ad.db.Table("transactions").Where("midtrans_id = ?", id).Updates(Transaction{
		UserID:          newData.UserID,
		PriceMethod:     newData.PriceMethod,
		PriceDuration:   newData.PriceDuration,
		PriceCounseling: newData.PriceCounseling,
		PriceResult:     newData.PriceResult,
		PaymentStatus:   newData.PaymentStatus,
		PaymentType:     newData.PaymentType,
	})

	if err := qry.Error; err != nil {
		fmt.Printf("Error updating transaction: %v\n", err)
		return false, err
	}

	if dataCount := qry.RowsAffected; dataCount < 1 {
		return false, errors.New("Update Data Error, No Data Affected")
	}

	if newData.PaymentStatus == 2 {
		if existingData.DoctorID == 0 {
			return false, errors.New("Doctor ID not found in transaction")
		}

		fmt.Println("This is the existing DoctorID: ", existingData.DoctorID)

		existingDataDoctor := data.Doctor{}
		if err := ad.db.Table("doctors").Where("id = ?", existingData.DoctorID).First(&existingDataDoctor).Error; err != nil {
			fmt.Printf("Error fetching doctor data: %v\n", err)
			return false, err
		}

		newDoctorBalance := existingDataDoctor.DoctorBalance + existingData.PriceResult
		fmt.Println("This is the new Update Balance: ", newDoctorBalance)

		qryToDoctor := ad.db.Table("doctors").Where("id = ?", existingData.DoctorID).Updates(data.Doctor{
			DoctorBalance: newDoctorBalance,
		})

		if err := qryToDoctor.Error; err != nil {
			fmt.Printf("Error updating doctor balance: %v\n", err)
			return false, err
		}

		if dataCount := qryToDoctor.RowsAffected; dataCount < 1 {
			return false, errors.New("Update Data Error, No Data Affected")
		}
	}

	return true, nil
}
