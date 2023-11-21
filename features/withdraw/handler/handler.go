package handler

import (
	"FinalProject/features/withdraw"
	"FinalProject/helper"
	"net/http"

	"github.com/labstack/echo/v4"
)

type WithdrawHandler struct {
	s withdraw.WithdrawServiceInterface
}

func New(service withdraw.WithdrawServiceInterface) withdraw.WithdrawHandlerInterface {
	return &WithdrawHandler{
		s: service,
	}
}

func (wh *WithdrawHandler) GetAllWithdraw() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := wh.s.GetAllWithdraw()

		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse(err.Error(), nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success", res))
	}
}
