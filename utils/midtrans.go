package utils

import (
	"FinalProject/configs"

	"github.com/midtrans/midtrans-go"
)

func InitMidtrans(c configs.ProgrammingConfig) {

	midtrans.ServerKey = c.MidtransServerKey
	midtrans.ClientKey = c.MidtransClientKey

	if c.MidtransEnvironment == "production" {
		midtrans.Environment = midtrans.Production
	} else {
		midtrans.Environment = midtrans.Sandbox

	}

}
