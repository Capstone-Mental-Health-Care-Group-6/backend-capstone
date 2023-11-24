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
	s transaction.TransactionServiceInterface
}

func NewTransactionHandler(service transaction.TransactionServiceInterface) transaction.TransactionHandlerInterface {
	// mt.InitMidtrans(c)
	return &TransactionHandler{
		s: service,
	}
}

func (th *TransactionHandler) NotifTransaction() echo.HandlerFunc {
	return func(c echo.Context) error {
		var notificationPayload map[string]interface{}

		err := json.NewDecoder(c.Request().Body).Decode(&notificationPayload)
		if err != nil {
			return err
		}

		if err != nil {
			logrus.Info("Handler : Input Process Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Fail", nil))
		}

		fmt.Println("Notification Payload:", notificationPayload)

		if err != nil {
			if strings.Contains(err.Error(), "Order ID Not Found") {
				return c.JSON(http.StatusBadRequest, helper.FormatResponse("Order ID Not Found", nil))
			}

			return c.JSON(http.StatusInternalServerError, helper.FormatResponse(err.Error(), nil))
		}

		var serviceUpdate = new(transaction.UpdateTransaction)

		res, err := th.s.UpdateTransaction(notificationPayload, *serviceUpdate)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse(err.Error(), nil))
		}
		return c.JSON(http.StatusOK, helper.FormatResponse("Success Update", res))
	}
}

func (th *TransactionHandler) CreateTransaction() echo.HandlerFunc {
	return func(c echo.Context) error {

		var input = new(InputRequest)
		if err := c.Bind(&input); err != nil {
			logrus.Info("Handler : Bind Input Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", nil))
		}

		var serviceInput = new(transaction.Transaction)

		serviceInput.PriceMethod = input.PriceMethod
		serviceInput.PriceDuration = input.PriceDuration
		serviceInput.PriceCounseling = input.PriceCounseling

		serviceInput.PriceResult = input.PriceMethod + input.PriceDuration + input.PriceCounseling

		serviceInput.UserID = input.UserID
		serviceInput.PaymentStatus = 0

		serviceInput.TopicID = input.TopicID
		serviceInput.PatientID = input.PatientID
		serviceInput.DoctorID = input.DoctorID
		serviceInput.MethodID = input.MethodID
		serviceInput.DurationID = input.DurationID

		serviceInput.CounselingID = input.CounselingID
		serviceInput.CounselingSession = input.CounselingSession
		serviceInput.CounselingType = input.CounselingType
		serviceInput.PaymentType = input.PaymentType

		_, response, err := th.s.CreateTransaction(*serviceInput)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse(err.Error(), nil))
		}

		return c.JSON(http.StatusCreated, helper.FormatResponse("Success Create Transaction", response))

	}
}

func (th *TransactionHandler) GetTransactions() echo.HandlerFunc {
	return func(c echo.Context) error {
		result, err := th.s.GetTransactions()

		if err != nil {
			logrus.Info("Handler : Get All Process Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Fail", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success", result))
	}
}

func (th *TransactionHandler) GetTransaction() echo.HandlerFunc {
	return func(c echo.Context) error {
		var paramID = c.Param("id")
		id, err := strconv.Atoi(paramID)
		if err != nil {
			logrus.Info("Handler : Param ID Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", nil))
		}

		result, err := th.s.GetTransaction(id)

		if err != nil {
			logrus.Info("Handler : Get By ID Process Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Fail", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success", result))
	}
}

func (th *TransactionHandler) GetTransactionByMidtransID() echo.HandlerFunc {
	return func(c echo.Context) error {
		var paramID = c.Param("id")
		// id, err := strconv.Atoi(paramID)
		// if err != nil {
		// 	logrus.Info("Handler : Param ID Error : ", err.Error())
		// 	return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", nil))
		// }

		result, err := th.s.GetByIDMidtrans(paramID)

		if err != nil {
			logrus.Info("Handler : Get By ID Process Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Fail", nil))
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
