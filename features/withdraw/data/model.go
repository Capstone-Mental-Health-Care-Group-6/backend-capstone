package data

import (
	"time"

	"gorm.io/gorm"
)

type Withdraw struct {
	*gorm.Model
	DoctorID      uint      `gorm:"column:doctor_id"`
	ConfirmByID   uint      `gorm:"column:confirm_by_id;default:NULL"`
	BalanceBefore uint      `gorm:"column:balance_before;default:NULL"`
	BalanceAfter  uint      `gorm:"column:balance_after;default:NULL"`
	BalanceReq    uint      `gorm:"column:balance_req"`
	PaymentMethod string    `gorm:"column:payment_method"`
	PaymentNumber string    `gorm:"column:payment_number"`
	PaymentName   string    `gorm:"column:payment_name"`
	DateConfirmed time.Time `gorm:"column:date_confirmed;default:NULL"`
	Status        string    `gorm:"column:status;type:enum('PENDING','PROCESS','DONE');default:'PENDING'"`
}

func (Withdraw) TableName() string {
	return "withdraws"
}
