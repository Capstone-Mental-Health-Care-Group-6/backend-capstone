package service

import (
	transaction "FinalProject/features/transaction"
	"errors"
	"fmt"
)

type TransactionService struct {
	d transaction.TransactionDataInterface
}

func New(data transaction.TransactionDataInterface) transaction.TransactionServiceInterface {
	return &TransactionService{
		d: data,
	}
}

func (as *TransactionService) GetTransactions() ([]transaction.TransactionInfo, error) {
	result, err := as.d.GetAll()
	if err != nil {
		return nil, errors.New("Get All Transactions Failed")
	}
	return result, nil
}

func (as *TransactionService) GetTransaction(id int) ([]transaction.TransactionInfo, error) {
	result, err := as.d.GetByID(id)
	if err != nil {
		return nil, errors.New("Get By ID Process Failed")
	}
	return result, nil
}

func (as *TransactionService) CreateTransaction(newData transaction.Transaction) (*transaction.Transaction, error) {
	result, err := as.d.Insert(newData)
	fmt.Println("Ini new data: ", newData)
	if err != nil {
		return nil, errors.New("Insert Process Failed")
	}
	return result, nil
}

func (as *TransactionService) UpdateTransaction(newData transaction.UpdateTransaction, id string) (bool, error) {
	result, err := as.d.GetAndUpdate(newData, id)

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
