package service

import (
	"FinalProject/features/withdraw"
	"errors"
)

type WithdrawService struct {
	wd withdraw.WithdrawDataInterface
}

func New(data withdraw.WithdrawDataInterface) withdraw.WithdrawServiceInterface {
	return &WithdrawService{
		wd: data,
	}
}

func (s *WithdrawService) GetAllWithdraw() ([]withdraw.WithdrawInfo, error) {
	result, err := s.wd.GetAll()
	if err != nil {
		return nil, errors.New("Get All Process Failed")
	}

	return result, nil
}

func (s *WithdrawService) GetAllWithdrawDokter(id uint) ([]withdraw.WithdrawInfo, error) {
	result, err := s.wd.GetAllDoctor(id)
	if err != nil {
		return nil, errors.New("Get All Withdraw Doctor Failed")
	}

	return result, nil
}

func (s *WithdrawService) GetBalance(idDoctor uint) (uint, error) {
	balance, err := s.wd.GetBalance(idDoctor)
	if err != nil {
		return 0, errors.New("Get Balance Error")
	}
	return balance, nil
}

func (s *WithdrawService) GetUserDoctor(id uint) (uint, error) {
	id, err := s.wd.GetUserDoctor(id)
	if err != nil {
		return 0, errors.New("Get User Doctor Error")
	}
	return id, nil
}

func (s *WithdrawService) CreateWithdraw(newData withdraw.Withdraw) (*withdraw.Withdraw, error) {
	result, err := s.wd.Insert(newData)
	if err != nil {
		return nil, errors.New("Insert Process Failed")
	}

	isTrue, err := s.wd.LessBalance(newData.DoctorID, newData.BalanceReq)
	if err != nil {
		return nil, errors.New("Less Balance Process Failed")
	}

	if !isTrue {
		return nil, errors.New("Less Balance")
	}

	return result, nil
}

func (s *WithdrawService) GetByID(id int) (*withdraw.WithdrawInfo, error) {
	result, err := s.wd.GetByID(id)
	if err != nil {
		return nil, errors.New("Get By ID Process Failed")
	}

	return result, nil
}

func (s *WithdrawService) UpdateStatus(id int, newData withdraw.Withdraw) (bool, error) {
	cekData, err := s.wd.GetByID(id)
	if err != nil {
		return false, errors.New("Get By ID Process Failed")
	}

	if cekData.ID == 0 {
		return false, errors.New("Data Not Found")
	}

	result, err := s.wd.UpdateStatus(id, newData)
	if err != nil {
		return false, errors.New("Update Status Process Failed")
	}

	return result, nil
}
