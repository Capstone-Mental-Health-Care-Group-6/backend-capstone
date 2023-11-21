package withdraw

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Withdraw struct {
	ID            uint
	DoctorID      uint
	ConfirmByID   uint
	BalanceBefore uint
	BalanceAfter  uint
	BalanceReq    uint
	PaymentMethod string
	PaymentNumber string
	DateConfirmed time.Time
	Status        string
}

type WithdrawInfo struct {
	ID            uint
	DoctorName    string
	ConfirmName   string
	BalanceBefore uint
	BalanceAfter  uint
	BalanceReq    uint
	PaymentMethod string
	PaymentNumber string
	DateConfirmed time.Time
	Status        string
}

type WithdrawHandlerInterface interface {
	GetAllWithdraw() echo.HandlerFunc
}

type WithdrawServiceInterface interface {
	GetAllWithdraw() ([]WithdrawInfo, error)
}

type WithdrawDataInterface interface {
	GetAll() ([]WithdrawInfo, error)
}
