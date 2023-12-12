package handler

import (
	counselingduration "FinalProject/features/counseling_durations"
	"FinalProject/helper"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CounselingDurationHandler struct {
	service counselingduration.CounselingDurationServiceInterface
}

func NewHandler(s counselingduration.CounselingDurationServiceInterface) counselingduration.CounselingDurationHandlerInterface {
	return &CounselingDurationHandler{
		service: s,
	}
}

func (cdh *CounselingDurationHandler) GetCounselingDurations() echo.HandlerFunc {
	return func(c echo.Context) error {
		result, err := cdh.service.GetAll()

		if err != nil {
			c.Logger().Error("Handler : Get All Process Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Get All Process Error", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success Get All Data", result))
	}
}
func (cdh *CounselingDurationHandler) GetCounselingDuration() echo.HandlerFunc {
	return func(c echo.Context) error {
		paramId := c.Param("id")
		id, err := strconv.Atoi(paramId)

		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid User Input Param ID", nil))
		}

		result, err := cdh.service.GetByID(id)

		if err != nil {
			c.Logger().Error("Handler : Get By ID Process Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Get By ID Process Error", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success Get Data", result))
	}
}
