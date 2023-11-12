package handler

import (
	"FinalProject/features/transaction"
	"FinalProject/helper"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/midtrans/midtrans-go/example"
)

type TransactionHandler struct {
	s transaction.TransactionServiceInterface
}

func NewTransactionHandler(service transaction.TransactionServiceInterface) transaction.TransactionHandlerInterface {
	return &TransactionHandler{
		s: service,
	}
}

func setupGlobalMidtransConfigApi() {
	midtrans.ServerKey = "SB-Mid-server-VXw9IjVeH_fZSL4IZykw3LR4"
	midtrans.ClientKey = "SB-Mid-client-hNK8kns-lS0o6nFn"
	// change value to `midtrans.Production`, if you want change the env to production
	midtrans.Environment = midtrans.Sandbox
}

func Insert() error {

	chargeReq := &coreapi.ChargeReq{
		PaymentType:  "bank_transfer",
		BankTransfer: &coreapi.BankTransferDetails{Bank: midtrans.BankBca},
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  example.Random(),
			GrossAmt: 100000,
		},
	}

	chargeResp, err := coreapi.ChargeTransaction(chargeReq)
	if err != nil {
		fmt.Println("Error: ", err)
		return err
	}

	var vaAccount string
	for _, va := range chargeResp.VaNumbers {
		if va.Bank == "bca" {
			vaAccount = va.VANumber
			break
		}
	}

	response := make(map[string]interface{})
	response["va_account"] = vaAccount

	fmt.Println("Success resp: ", response)

	return err
}

func (th *TransactionHandler) CreateTransaction() echo.HandlerFunc {
	return func(c echo.Context) error {

		setupGlobalMidtransConfigApi()

		chargeReq := &coreapi.ChargeReq{
			PaymentType:  "bank_transfer",
			BankTransfer: &coreapi.BankTransferDetails{Bank: midtrans.BankBca},
			TransactionDetails: midtrans.TransactionDetails{
				OrderID:  example.Random(),
				GrossAmt: 100000,
			},
		}

		chargeResp, err := coreapi.ChargeTransaction(chargeReq)
		if err != nil {
			fmt.Println("Error: ", err)
			return err
		}

		var vaAccount string
		for _, va := range chargeResp.VaNumbers {
			if va.Bank == "bca" {
				vaAccount = va.VANumber
				break
			}
		}

		var serviceInput = new(transaction.Transaction)
		serviceInput.TopicID = nil
		serviceInput.PatientID = nil
		serviceInput.DoctorID = nil
		serviceInput.MethodID = nil
		serviceInput.DurationID = nil
		serviceInput.CounselingID = nil
		serviceInput.UserID = 1
		serviceInput.MidtransID = &chargeResp.TransactionID
		// Assuming chargeResp.GrossAmount is *string
		grossAmount, _ := strconv.Atoi(chargeResp.GrossAmount)
		// Assuming serviceInput.PriceResult is *int
		result := grossAmount
		serviceInput.PriceResult = &result

		serviceInput.CounselingSession = nil
		serviceInput.CounselingType = nil

		serviceInput.PriceMethod = nil
		serviceInput.PriceDuration = nil
		serviceInput.PriceCounseling = nil

		serviceInput.PaymentStatus = 0
		serviceInput.PaymentType = "Bank BCA"

		th.s.CreateTransaction(*serviceInput)

		response := make(map[string]interface{})
		response["va_account"] = vaAccount
		response["data"] = serviceInput

		return c.JSON(http.StatusCreated, helper.FormatResponse("Success", response))
	}
}

func (th *TransactionHandler) GetTransactions() echo.HandlerFunc {
	return nil
}

func (th *TransactionHandler) GetTransaction() echo.HandlerFunc {

	return nil
}

// func (h *TransactionHandler) ChargeTransaction(c echo.Context) error {
// 	// Get the request body
// 	request := ChargeTransactionRequest{}
// 	if err := c.Bind(request); err != nil {
// 		return err
// 	}

// 	// Charge the transaction
// 	response, err := h.transactionService.ChargeTransaction(context.Background(), request)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, err)
// 	}

// 	// Return the response
// 	return c.JSON(http.StatusOK, response)
// }
