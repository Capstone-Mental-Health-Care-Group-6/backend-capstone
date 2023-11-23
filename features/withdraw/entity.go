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
	PaymentName   string
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
	PaymentName   string
	DateConfirmed time.Time
	Status        string
}

type WithdrawHandlerInterface interface {
	GetAllWithdraw() echo.HandlerFunc
	CreateWithdraw() echo.HandlerFunc
	GetWithdraw() echo.HandlerFunc
}

type WithdrawServiceInterface interface {
	GetAllWithdraw() ([]WithdrawInfo, error)
	CreateWithdraw(newData Withdraw) (*Withdraw, error)
	GetBalance(idDoctor uint) (uint, error)
	GetByID(id int) (*WithdrawInfo, error)
}

type WithdrawDataInterface interface {
	GetAll() ([]WithdrawInfo, error)
	Insert(newData Withdraw) (*Withdraw, error)
	GetBalance(idDoctor uint) (uint, error)
	LessBalance(idDoctor uint, balance uint) (bool, error)
	GetByID(id int) (*WithdrawInfo, error)
}
