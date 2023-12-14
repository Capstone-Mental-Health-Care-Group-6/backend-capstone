package handler

import (
	"FinalProject/features/withdraw"
	"FinalProject/helper"
	"net/http"
	"strconv"
	"time"

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
		role := wh.jwt.CheckRole(c)

		if role != "Admin" {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Only admin can access this page", nil))
		}

		res, err := wh.s.GetAllWithdraw()

		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse(err.Error(), nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success get all data", res))
	}
}

func (wh *WithdrawHandler) GetAllWithdrawDokter() echo.HandlerFunc {
	return func(c echo.Context) error {
		role := wh.jwt.CheckRole(c)

		if role != "Doctor" {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Only doctor can access this page", nil))
		}

		id, _ := wh.jwt.GetID(c)

		idDoctor, err := wh.s.GetUserDoctor(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse(err.Error(), nil))
		}
		res, err := wh.s.GetAllWithdrawDokter(idDoctor)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse(err.Error(), nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success get balance", res))
	}
}

func (wh *WithdrawHandler) CreateWithdraw() echo.HandlerFunc {
	return func(c echo.Context) error {
		role := wh.jwt.CheckRole(c)

		if role != "Doctor" {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Only doctor can access this page", nil))
		}

		var req = new(InputRequest)

		if err := c.Bind(req); err != nil {
			c.Logger().Info("Handler : Bind Input Error : ", err.Error())
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

		idDoctor, err := wh.s.GetUserDoctor(idJwt)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse(err.Error(), nil))
		}

		var serviceInput = new(withdraw.Withdraw)
		serviceInput.DoctorID = idDoctor
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

func (wh *WithdrawHandler) GetWithdraw() echo.HandlerFunc {
	return func(c echo.Context) error {
		var paramID = c.Param("id")
		id, err := strconv.Atoi(paramID)

		if err != nil {
			c.Logger().Info("Handler : Param ID Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Param ID Error", nil))
		}

		res, err := wh.s.GetByID(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse(err.Error(), nil))
		}

		if res.ID == 0 {
			return c.JSON(http.StatusNotFound, helper.FormatResponse("Data not found", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success get data", res))
	}
}

func (wh *WithdrawHandler) UpdateStatus() echo.HandlerFunc {
	return func(c echo.Context) error {
		role := wh.jwt.CheckRole(c)

		if role != "Admin" {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Only admin can access this page", nil))
		}

		var paramID = c.Param("id")
		id, err := strconv.Atoi(paramID)

		if err != nil {
			c.Logger().Info("Handler : Param ID Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Param ID Error", nil))
		}

		var req = new(UpdateStatusRequest)
		if err := c.Bind(req); err != nil {
			c.Logger().Info("Handler : Bind Input Error : ", err.Error())
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
		serviceInput.ConfirmByID = idJwt
		serviceInput.Status = req.Status
		serviceInput.DateConfirmed = time.Now()

		ress, err := wh.s.UpdateStatus(id, *serviceInput)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse(err.Error(), nil))
		}

		if ress != true {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Data not found", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success update status to "+req.Status, nil))
	}
}
