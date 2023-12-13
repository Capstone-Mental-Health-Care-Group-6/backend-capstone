package withdraw

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Withdraw struct {
	ID            uint      `json:"id"`
	DoctorID      uint      `json:"doctor_id"`
	ConfirmByID   uint      `json:"confirm_by_id"`
	BalanceBefore uint      `json:"balance_before"`
	BalanceAfter  uint      `json:"balance_after"`
	BalanceReq    uint      `json:"balance_req"`
	PaymentMethod string    `json:"payment_method"`
	PaymentNumber string    `json:"payment_number"`
	PaymentName   string    `json:"payment_name"`
	DateConfirmed time.Time `json:"date_confirmed"`
	Status        string    `json:"status"`
}

type WithdrawInfo struct {
	ID            uint      `json:"id"`
	DoctorName    string    `json:"doctor_name"`
	ConfirmName   string    `json:"confirm_name"`
	BalanceBefore uint      `json:"balance_before"`
	BalanceAfter  uint      `json:"balance_after"`
	BalanceReq    uint      `json:"balance_req"`
	PaymentMethod string    `json:"payment_method"`
	PaymentNumber string    `json:"payment_number"`
	PaymentName   string    `json:"payment_name"`
	DateConfirmed time.Time `json:"date_confirmed"`
	Status        string    `json:"status"`
}

type WithdrawHandlerInterface interface {
	GetAllWithdraw() echo.HandlerFunc
	GetAllWithdrawDokter() echo.HandlerFunc
	CreateWithdraw() echo.HandlerFunc
	GetWithdraw() echo.HandlerFunc
	UpdateStatus() echo.HandlerFunc
}

type WithdrawServiceInterface interface {
	GetAllWithdraw() ([]WithdrawInfo, error)
	GetAllWithdrawDokter(id uint) ([]WithdrawInfo, error)
	GetUserDoctor(id uint) (uint, error)
	CreateWithdraw(newData Withdraw) (*Withdraw, error)
	GetBalance(idDoctor uint) (uint, error)
	GetByID(id int) (*WithdrawInfo, error)
	UpdateStatus(id int, newData Withdraw) (bool, error)
}

type WithdrawDataInterface interface {
	GetAll() ([]WithdrawInfo, error)
	GetAllDoctor(id uint) ([]WithdrawInfo, error)
	GetUserDoctor(id uint) (uint, error)
	Insert(newData Withdraw) (*Withdraw, error)
	GetBalance(idDoctor uint) (uint, error)
	LessBalance(idDoctor uint, balance uint) (bool, error)
	GetByID(id int) (*WithdrawInfo, error)
	UpdateStatus(id int, newData Withdraw) (bool, error)
}
