package service

import (
	transaction "FinalProject/features/transaction"
	"FinalProject/utils/midtrans"
	"errors"
	"fmt"
)

type TransactionService struct {
	d  transaction.TransactionDataInterface
	mt midtrans.MidtransService
}

func New(data transaction.TransactionDataInterface, mid midtrans.MidtransService) transaction.TransactionServiceInterface {
	return &TransactionService{
		d:  data,
		mt: mid,
	}
}

func (as *TransactionService) GetTransactions() ([]transaction.TransactionInfo, error) {
	result, err := as.d.GetAll()
	if err != nil {
		return nil, errors.New("Get All Transactions Failed")
	}
	return result, nil
}

func (as *TransactionService) GetTransaction(id int) ([]transaction.Transaction, error) {
	result, err := as.d.GetByID(id)
	if err != nil {
		return nil, errors.New("Get By ID Process Failed")
	}
	return result, nil
}

func (as *TransactionService) GetByIDMidtrans(id string) ([]transaction.TransactionInfo, error) {
	result, err := as.d.GetByIDMidtrans(id)
	if err != nil {
		return nil, errors.New("Get By ID Midtrans Process Failed")
	}
	return result, nil
}

func (as *TransactionService) CreateTransaction(newData transaction.Transaction) (*transaction.Transaction, map[string]interface{}, error) {
	totalPrice := newData.PriceMethod + newData.PriceDuration + newData.PriceCounseling

	chargeResp, response, err := as.mt.GenerateTransaction(int(totalPrice), newData.PaymentType)

	if err != nil {
		fmt.Println("Error: ", err)
		return nil, nil, errors.New("Generate Transaction Failed")
	}

	newData.MidtransID = chargeResp.OrderID

	result, err := as.d.Insert(newData)

	fmt.Println("Ini new data: ", newData)
	if err != nil {
		return nil, nil, errors.New("Insert Process Failed")
	}
	return result, response, nil
}

func (as *TransactionService) UpdateTransaction(notificationPayload map[string]interface{}, newData transaction.UpdateTransaction) (bool, error) {
	paymentStatus, orderId, err := as.mt.TransactionStatus(notificationPayload)
	newData.PaymentStatus = uint(paymentStatus)
	result, err := as.d.GetAndUpdate(newData, orderId)

	if err != nil {
		return false, errors.New("Update Process Failed")
	}

	return result, nil
}

func (as *TransactionService) DeleteTransaction(id int) (bool, error) {
	result, err := as.d.Delete(id)

	if err != nil {
		return false, errors.New("Delete Process Failed")
	}

	return result, nil
}
