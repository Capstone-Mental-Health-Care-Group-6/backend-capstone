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
		return nil, err
	}

	return result, nil
}

func (s *WithdrawService) GetBalance(idDoctor uint) (uint, error) {
	return s.wd.GetBalance(idDoctor)
}

func (s *WithdrawService) CreateWithdraw(newData withdraw.Withdraw) (*withdraw.Withdraw, error) {
	result, err := s.wd.Insert(newData)
	if err != nil {
		return nil, err
	}

	isTrue, err := s.wd.LessBalance(newData.DoctorID, newData.BalanceReq)
	if err != nil {
		return nil, err
	}

	if !isTrue {
		return nil, err
	}

	return result, nil
}

func (s *WithdrawService) GetByID(id int) (*withdraw.WithdrawInfo, error) {
	result, err := s.wd.GetByID(id)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *WithdrawService) UpdateStatus(id int, newData withdraw.Withdraw) (bool, error) {
	cekData, err := s.wd.GetByID(id)
	if err != nil {
		return false, err
	}

	if cekData.ID == 0 {
		return false, errors.New("data not found")
	}

	result, err := s.wd.UpdateStatus(id, newData)
	if err != nil {
		return false, err
	}

	return result, nil
}
