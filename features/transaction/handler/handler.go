package handler

import (
	"FinalProject/features/transaction"
	"FinalProject/helper"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type TransactionHandler struct {
	s   transaction.TransactionServiceInterface
	jwt helper.JWTInterface
}

func NewTransactionHandler(service transaction.TransactionServiceInterface, jwt helper.JWTInterface) transaction.TransactionHandlerInterface {
	// mt.InitMidtrans(c)
	return &TransactionHandler{
		s:   service,
		jwt: jwt,
	}
}

func (th *TransactionHandler) NotifTransaction() echo.HandlerFunc {
	return func(c echo.Context) error {
		var notificationPayload map[string]interface{}

		err := json.NewDecoder(c.Request().Body).Decode(&notificationPayload)

		fmt.Println("Notification Payload:", notificationPayload)

		if err != nil {
			if strings.Contains(err.Error(), "Order ID Not Found") {
				return c.JSON(http.StatusBadRequest, helper.FormatResponse("Order ID Not Found", nil))
			}

			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Midtrans POST method error", nil))
		}

		var serviceUpdate = new(transaction.UpdateTransaction)

		res, err := th.s.UpdateTransaction(notificationPayload, *serviceUpdate)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Midtrans cannot update the database", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success Update", res))
	}
}

func (th *TransactionHandler) CreateTransaction() echo.HandlerFunc {
	return func(c echo.Context) error {
		getID, err := th.jwt.GetID(c)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Fail, cant get ID from JWT", nil))
		}

		var input = new(InputRequest)
		if err := c.Bind(input); err != nil {
			logrus.Info("Handler : Bind Input Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail to bind request", nil))
		}

		var serviceInput = new(transaction.Transaction)

		serviceInput.PriceMethod = input.PriceMethod
		serviceInput.PriceDuration = input.PriceDuration
		serviceInput.PriceCounseling = input.PriceCounseling

		serviceInput.PriceResult = input.PriceMethod + input.PriceDuration + input.PriceCounseling

		serviceInput.UserID = getID
		serviceInput.PaymentStatus = 5

		serviceInput.TopicID = input.TopicID
		serviceInput.PatientID = getID
		serviceInput.DoctorID = input.DoctorID
		serviceInput.MethodID = input.MethodID
		serviceInput.DurationID = input.DurationID

		serviceInput.CounselingID = input.CounselingID
		serviceInput.CounselingSession = input.CounselingSession
		serviceInput.CounselingType = input.CounselingType
		serviceInput.PaymentType = input.PaymentType

		if input.PaymentType == "manual" {

			formHeaderPaymentProof, err := c.FormFile("payment_proof")
			if err != nil {
				return c.JSON(http.StatusBadRequest, helper.FormatResponse("Failed, Select a File for Upload Payment Proof", nil))
			}

			formPaymentProof, err := formHeaderPaymentProof.Open()
			if err != nil {
				return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Failed to get payment proof", nil))
			}

			uploadUrlPaymentProof, err := th.s.PaymentProofUpload(transaction.PaymentProofDataModel{PaymentProofPhoto: formPaymentProof})
			if err != nil {
				return c.JSON(http.StatusBadRequest, helper.FormatResponse("Failed to Upload Payment Proof", nil))
			}

			serviceInput.PaymentProof = uploadUrlPaymentProof

			logrus.Info("Ini payment proof: ", uploadUrlPaymentProof)

			result, err := th.s.CreateManualTransaction(*serviceInput)

			if err != nil {
				return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Failed to create manual transaction", nil))
			}

			var response = new(ManualTransactionResponse)
			response.PriceResult = result.PriceResult
			response.UserID = result.UserID
			response.MidtransID = result.MidtransID
			response.PaymentStatus = result.PaymentStatus
			response.PaymentProof = result.PaymentProof
			response.PaymentType = result.PaymentType

			return c.JSON(http.StatusCreated, helper.FormatResponse("Success Create Manual Transaction", response))

		} else {

			serviceInput.PaymentProof = "midtrans_payment"
			_, response, err := th.s.CreateTransaction(*serviceInput)

			if err != nil {
				return c.JSON(http.StatusInternalServerError, helper.FormatResponse(err.Error(), nil))
			}

			return c.JSON(http.StatusCreated, helper.FormatResponse("Success Create Midtrans Transaction", response))

		}

	}
}

func (th *TransactionHandler) GetTransactions() echo.HandlerFunc {
	return func(c echo.Context) error {
		sortByPaymentType := c.QueryParam("payment_type")

		if sortByPaymentType != "" {

			result, err := th.s.GetTransactions(sortByPaymentType)

			if err != nil {
				logrus.Info("Handler : Get All Process Error : ", err.Error())
				return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Fail to retrieve transaction sorted data", nil))
			}

			if len(result) == 0 {
				return c.JSON(http.StatusOK, helper.FormatResponse("Success no data", nil))
			}

			return c.JSON(http.StatusOK, helper.FormatResponse("Success to retrieve data", result))

		}

		blank := ""
		result, err := th.s.GetTransactions(blank)

		if err != nil {
			logrus.Info("Handler : Get All Process Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Fail to retrieve transaction data", nil))
		}

		if len(result) == 0 {
			return c.JSON(http.StatusOK, helper.FormatResponse("Success no data", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success", result))
	}
}

func (th *TransactionHandler) GetTransaction() echo.HandlerFunc {
	return func(c echo.Context) error {
		sortByPaymentType := c.QueryParam("payment_type")
		var paramID = c.Param("id")
		id, err := strconv.Atoi(paramID)
		if err != nil {
			logrus.Info("Handler : Param User ID Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail to get user id", nil))
		}

		if sortByPaymentType != "" {

			//BASED ON USER ID

			result, err := th.s.GetTransaction(id, sortByPaymentType)

			if err != nil {
				logrus.Info("Handler : Get transactions by User ID Process Error : ", err.Error())
				return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Fail to retrieve transaction sorted data", nil))
			}

			if len(result) == 0 {
				return c.JSON(http.StatusOK, helper.FormatResponse("Success no data", nil))
			}

			return c.JSON(http.StatusOK, helper.FormatResponse("Success", result))

		}

		blank := ""

		result, err := th.s.GetTransaction(id, blank)

		if err != nil {
			logrus.Info("Handler : Get By User ID Process Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Fail to retrieve transactions data", nil))
		}

		if len(result) == 0 {
			return c.JSON(http.StatusOK, helper.FormatResponse("Success no data", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success", result))
	}
}

func (th *TransactionHandler) GetTransactionByPatientID() echo.HandlerFunc {
	return func(c echo.Context) error {
		sortByPaymentType := c.QueryParam("payment_type")
		var paramID = c.Param("id")
		id, err := strconv.Atoi(paramID)
		if err != nil {
			logrus.Info("Handler : Param ID Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail to get patientID", nil))
		}

		if sortByPaymentType != "" {

			result, err := th.s.GetTransactionByPatientID(id, sortByPaymentType)

			if err != nil {
				logrus.Info("Handler : Get All Process Error : ", err.Error())
				return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Fail to retrieve transaction sorted data", nil))
			}

			if len(result) == 0 {
				return c.JSON(http.StatusOK, helper.FormatResponse("Success no data", nil))
			}

			return c.JSON(http.StatusOK, helper.FormatResponse("Success", result))

		}

		blank := ""

		result, err := th.s.GetTransactionByPatientID(id, blank)

		if err != nil {
			logrus.Info("Handler : Get By ID Process Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Fail", nil))
		}

		if len(result) == 0 {
			return c.JSON(http.StatusOK, helper.FormatResponse("Success no data", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success", result))
	}
}

func (th *TransactionHandler) GetTransactionByMidtransID() echo.HandlerFunc {
	return func(c echo.Context) error {
		var paramID = c.Param("id")

		result, err := th.s.GetByIDMidtrans(paramID)

		if err != nil {
			logrus.Info("Handler : Get By ID Process Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Fail", nil))
		}

		if result == nil {
			return c.JSON(http.StatusOK, helper.FormatResponse("Success no data", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success", result))
	}
}

func (th *TransactionHandler) UpdateTransaction() echo.HandlerFunc {
	return func(c echo.Context) error {
		var paramID = c.Param("id")

		var input = new(UpdateRequest)
		if err := c.Bind(&input); err != nil {
			c.Logger().Info("Handler : Bind Input Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", nil))
		}

		var serviceUpdate = new(transaction.UpdateTransactionManual)
		serviceUpdate.UserID = input.UserID
		serviceUpdate.PriceMethod = input.PriceMethod
		serviceUpdate.PriceDuration = input.PriceDuration
		serviceUpdate.PriceCounseling = input.PriceCounseling
		serviceUpdate.PriceResult = input.PriceResult
		serviceUpdate.PaymentStatus = input.PaymentStatus

		result, err := th.s.UpdateTransactionManual(*serviceUpdate, paramID)

		if err != nil {
			// c.Logger().Info("Handler : Input Process Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse(err.Error(), nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success", result))
	}
}

func (th *TransactionHandler) DeleteTransaction() echo.HandlerFunc {
	return func(c echo.Context) error {
		var paramID = c.Param("id")
		id, err := strconv.Atoi(paramID)

		if err != nil {
			logrus.Info("Handler : Param ID Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", nil))
		}

		result, err := th.s.DeleteTransaction(id)

		if err != nil {
			logrus.Info("Handler : Delete Process Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Fail", nil))
		}

		fmt.Println(result)

		if result {
			return c.JSON(http.StatusOK, helper.FormatResponse("Transaction deleted success", nil))
		} else {
			return c.JSON(http.StatusNotFound, helper.FormatResponse("Transaction not found", result))
		}

	}
}
