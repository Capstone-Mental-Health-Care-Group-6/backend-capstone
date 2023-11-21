package service

import "FinalProject/features/withdraw"

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
