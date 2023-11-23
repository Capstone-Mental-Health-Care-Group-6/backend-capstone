package handler

import (
	"FinalProject/features/withdraw"
	"FinalProject/helper"
	"net/http"

	"github.com/labstack/echo/v4"
)

type WithdrawHandler struct {
	s   withdraw.WithdrawServiceInterface
	jwt helper.JWTInterface
}

func New(service withdraw.WithdrawServiceInterface, jwt helper.JWTInterface) withdraw.WithdrawHandlerInterface {
	return &WithdrawHandler{
		s:   service,
		jwt: jwt,
	}
}

func (wh *WithdrawHandler) GetAllWithdraw() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := wh.s.GetAllWithdraw()

		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse(err.Error(), nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success get all data", res))
	}
}

func (wh *WithdrawHandler) CreateWithdraw() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req = new(InputRequest)

		if err := c.Bind(req); err != nil {
			c.Logger().Fatal("Handler : Bind Input Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid Format Request", nil))
		}

		isValid, errors := helper.ValidateJSON(req)
		if !isValid {
			return c.JSON(http.StatusBadRequest, helper.FormatResponseValidation("Invalid Format Request", errors))
		}

		idJwt, err := wh.jwt.GetID(c)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Token is not valid", nil))
		}

		var serviceInput = new(withdraw.Withdraw)
		serviceInput.DoctorID = idJwt
		serviceInput.BalanceReq = req.BalanceReq
		serviceInput.PaymentMethod = req.PaymentMethod
		serviceInput.PaymentNumber = req.PaymentNumber
		serviceInput.PaymentName = req.PaymentName

		balance, err := wh.s.GetBalance(serviceInput.DoctorID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse(err.Error(), nil))
		}

		if balance < serviceInput.BalanceReq {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Balance not enough", nil))
		}

		serviceInput.BalanceBefore = balance
		serviceInput.BalanceAfter = balance - serviceInput.BalanceReq

		result, err := wh.s.CreateWithdraw(*serviceInput)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse(err.Error(), nil))
		}

		var ress = new(WithdrawResponse)
		ress.DoctorID = result.DoctorID
		ress.BalanceReq = result.BalanceReq
		ress.BalanceBefore = result.BalanceBefore
		ress.BalanceAfter = result.BalanceAfter
		ress.PaymentMethod = result.PaymentMethod
		ress.PaymentNumber = result.PaymentNumber
		ress.PaymentName = result.PaymentName

		return c.JSON(http.StatusCreated, helper.FormatResponse("Success create new withdraw", ress))
	}
}
