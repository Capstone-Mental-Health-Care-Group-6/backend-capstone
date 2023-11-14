package handler

import (
	"FinalProject/features/transaction"
	"FinalProject/helper"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/midtrans/midtrans-go/example"
)

var cApi coreapi.Client

type TransactionHandler struct {
	s transaction.TransactionServiceInterface
}

func NewTransactionHandler(service transaction.TransactionServiceInterface) transaction.TransactionHandlerInterface {
	return &TransactionHandler{
		s: service,
	}
}

func (th *TransactionHandler) NotifTransaction() echo.HandlerFunc {
	return func(c echo.Context) error {

		midtrans.ServerKey = "SB-Mid-server-VXw9IjVeH_fZSL4IZykw3LR4"
		midtrans.Environment = midtrans.Sandbox

		cApi.New(midtrans.ServerKey, midtrans.Sandbox)

		var notificationPayload map[string]interface{}

		err := json.NewDecoder(c.Request().Body).Decode(&notificationPayload)
		if err != nil {
			return err
		}

		if err != nil {
			c.Logger().Fatal("Handler : Input Process Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Fail", nil))
		}

		fmt.Println("Notification Payload:", notificationPayload)

		orderId, exists := notificationPayload["order_id"].(string)
		if !exists {
			return echo.NewHTTPError(http.StatusBadRequest, "order_id not found")
		}

		transactionStatusResp, e := cApi.CheckTransaction(orderId)
		if e != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, e.GetMessage())
		} else {
			if transactionStatusResp != nil {
				if transactionStatusResp.TransactionStatus == "capture" {

					if transactionStatusResp.FraudStatus == "challenge" {
						fmt.Println("Payment status challenged")
						var serviceUpdate = new(transaction.UpdateTransaction)
						serviceUpdate.PaymentStatus = 1 //CHALLENGE

						th.s.UpdateTransaction(*serviceUpdate, transactionStatusResp.OrderID)

					} else if transactionStatusResp.FraudStatus == "accept" {
						var serviceUpdate = new(transaction.UpdateTransaction)
						serviceUpdate.PaymentStatus = 2 //ACCEPT

						th.s.UpdateTransaction(*serviceUpdate, transactionStatusResp.OrderID)
						fmt.Println("Payment received")
						// TODO set transaction status on your database to 'success'
					}
				} else if transactionStatusResp.TransactionStatus == "settlement" {
					var serviceUpdate = new(transaction.UpdateTransaction)
					serviceUpdate.PaymentStatus = 2 //ACCEPT

					th.s.UpdateTransaction(*serviceUpdate, transactionStatusResp.OrderID)
					fmt.Println("Payment status settlement")

					// TODO set transaction status on your databaase to 'success'
				} else if transactionStatusResp.TransactionStatus == "deny" {
					fmt.Println("Payment status denied")

					var serviceUpdate = new(transaction.UpdateTransaction)
					serviceUpdate.PaymentStatus = 3 //DENIED

					th.s.UpdateTransaction(*serviceUpdate, transactionStatusResp.OrderID)

					// TODO you can ignore 'deny', because most of the time it allows payment retries
					// and later can become success
				} else if transactionStatusResp.TransactionStatus == "cancel" || transactionStatusResp.TransactionStatus == "expire" {
					fmt.Println("Payment status failure")

					var serviceUpdate = new(transaction.UpdateTransaction)
					serviceUpdate.PaymentStatus = 4 //FAILURE

					th.s.UpdateTransaction(*serviceUpdate, transactionStatusResp.OrderID)

					// TODO set transaction status on your databaase to 'failure'
				} else if transactionStatusResp.TransactionStatus == "pending" {
					fmt.Println("Payment status pending")
					var serviceUpdate = new(transaction.UpdateTransaction)
					serviceUpdate.PaymentStatus = 5 //WAITING

					th.s.UpdateTransaction(*serviceUpdate, transactionStatusResp.OrderID)
				}
			}
		}

		return c.String(http.StatusOK, "ok")
	}
}

func (th *TransactionHandler) CreateTransaction() echo.HandlerFunc {
	return func(c echo.Context) error {

		midtrans.ServerKey = "SB-Mid-server-VXw9IjVeH_fZSL4IZykw3LR4"
		midtrans.ClientKey = "SB-Mid-client-hNK8kns-lS0o6nFn"
		midtrans.Environment = midtrans.Sandbox

		var input = new(InputRequest)
		if err := c.Bind(&input); err != nil {
			c.Logger().Fatal("Handler : Bind Input Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", nil))
		}

		var serviceInput = new(transaction.Transaction)

		serviceInput.PriceMethod = input.PriceMethod
		serviceInput.PriceDuration = input.PriceDuration
		serviceInput.PriceCounseling = input.PriceCounseling
		serviceInput.PriceResult = input.PriceMethod + input.PriceDuration + input.PriceCounseling

		serviceInput.TopicID = input.TopicID
		serviceInput.PatientID = input.PatientID
		serviceInput.DoctorID = input.DoctorID
		serviceInput.MethodID = input.MethodID
		serviceInput.DurationID = input.DurationID

		serviceInput.CounselingID = input.CounselingID
		serviceInput.CounselingSession = input.CounselingSession
		serviceInput.CounselingType = input.CounselingType
		serviceInput.PaymentType = input.PaymentType

		result := input.PriceMethod + input.PriceDuration + input.PriceCounseling

		response := make(map[string]interface{})

		if input.PaymentType == "gopay" {
			chargeReq := &coreapi.ChargeReq{
				Gopay: &coreapi.GopayDetails{
					EnableCallback: true,
				},
				PaymentType: "gopay",
				TransactionDetails: midtrans.TransactionDetails{
					OrderID:  "G-" + example.Random(),
					GrossAmt: int64(result),
				},
			}

			chargeResp, err := coreapi.ChargeTransaction(chargeReq)
			if err != nil {
				fmt.Println("Error: ", err)
				return err
			}

			serviceInput.UserID = input.UserID
			serviceInput.MidtransID = chargeResp.OrderID
			serviceInput.PaymentStatus = 0

			if len(chargeResp.Actions) > 0 {
				for _, action := range chargeResp.Actions {
					if action.Name == "deeplink-redirect" {
						deepLinkURL := action.URL
						response["callback_url"] = deepLinkURL
						// c.DepositService.InsertPaymentToken(Deposit.ID, chargeResp.TransactionID, "-", deepLinkURL)
						break
					}
				}
			}
			response["payment_type"] = input.PaymentType

		} else if input.PaymentType == "bca" || input.PaymentType == "bni" || input.PaymentType == "bri" {

			var midtransBank midtrans.Bank
			switch input.PaymentType {
			case "bca":
				midtransBank = midtrans.BankBca
			case "bri":
				midtransBank = midtrans.BankBri
			case "bni":
				midtransBank = midtrans.BankBni
			default:
				midtransBank = midtrans.BankBca
			}

			chargeReq := &coreapi.ChargeReq{
				PaymentType:  "bank_transfer",
				BankTransfer: &coreapi.BankTransferDetails{Bank: midtransBank},
				TransactionDetails: midtrans.TransactionDetails{
					OrderID:  "B-" + example.Random(),
					GrossAmt: int64(result),
				},
			}

			chargeResp, err := coreapi.ChargeTransaction(chargeReq)
			if err != nil {
				fmt.Println("Error: ", err)
				return err
			}

			serviceInput.UserID = input.UserID
			serviceInput.MidtransID = chargeResp.OrderID
			serviceInput.PaymentStatus = 0

			fmt.Println("This is the data", chargeResp.OrderID, serviceInput.PaymentStatus, serviceInput.PaymentType)

			th.s.CreateTransaction(*serviceInput)

			var vaAccount string
			for _, va := range chargeResp.VaNumbers {
				if va.Bank == input.PaymentType {
					vaAccount = va.VANumber
					break
				}
			}
			response["payment_type"] = input.PaymentType
			response["va_account"] = vaAccount

		} else {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Unsupported payment type", nil))
		}

		return c.JSON(http.StatusCreated, helper.FormatResponse("Success", response))

	}
}

func (th *TransactionHandler) GetTransactions() echo.HandlerFunc {
	return nil
}

func (th *TransactionHandler) GetTransaction() echo.HandlerFunc {

	return nil
}