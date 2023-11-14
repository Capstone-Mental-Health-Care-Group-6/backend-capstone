package data

import (
	"FinalProject/features/transaction"
	"fmt"

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

func (ad *TransactionData) GetAndUpdate(newData transaction.Transaction, id string) (bool, error) {

	var transaction Transaction
	db := ad.db // Assuming ad.db is the GORM database instance
	db.Where("midtrans_id = ?", id).First(&transaction)
	// If you want to get the ID specifically, you can access it like this:
	fmt.Println("This is the id: ", transaction.ID)
	transactionID := transaction.ID

	fmt.Println("This is the new payment status: ", newData.PaymentStatus)

	// Now you can use transactionID in your update query
	qry := db.Table("transactions").Where("id = ?", transactionID).Updates(Transaction{
		// MidtransID: newData.MidtransID,
		// PriceResult: newData.PriceResult,
		PaymentStatus: newData.PaymentStatus,
	})

	// Check for errors in the update query
	if qry.Error != nil {
		// Handle the error
		return false, nil
	}
	// var qry = ad.db.Table("transaction").Where("id = ?", id).Updates(Transaction{nil})

	if err := qry.Error; err != nil {
		return false, err
	}

	// if dataCount := qry.RowsAffected; dataCount < 1 {
	// 	return false, nil
	// }

	return true, nil
}

func (ad *TransactionData) GetAll() ([]transaction.TransactionInfo, error) {

	return nil, nil
}

func (ad *TransactionData) GetByID(id int) ([]transaction.TransactionInfo, error) {
	return nil, nil
}

func (ad *TransactionData) Insert(newData transaction.Transaction) (*transaction.Transaction, error) {
	var dbData = new(Transaction)

	dbData.UserID = newData.UserID
	dbData.MidtransID = *newData.MidtransID

	dbData.PriceResult = *newData.PriceResult

	dbData.PaymentStatus = *&newData.PaymentStatus
	dbData.PaymentType = *&newData.PaymentType

	if err := ad.db.Create(dbData).Error; err != nil {
		return nil, err
	}

	return &newData, nil
}
