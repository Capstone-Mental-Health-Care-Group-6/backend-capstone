package handler

import (
	counselingtopics "FinalProject/features/counseling_topics"
	"FinalProject/helper"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CounselingTopicHandler struct {
	service counselingtopics.CounselingTopicServiceInterface
}

func NewHandler(service counselingtopics.CounselingTopicServiceInterface) counselingtopics.CounselingTopicHandlerInterface {
	return &CounselingTopicHandler{
		service: service,
	}
}

func (cmh *CounselingTopicHandler) GetCounselingTopics() echo.HandlerFunc {
	return func(c echo.Context) error {
		result, err := cmh.service.GetAll()

		if err != nil {
			c.Logger().Error("Handler : Get All Process Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Get All Process Error", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success Get All Data", result))
	}
}
func (cmh *CounselingTopicHandler) GetCounselingTopic() echo.HandlerFunc {
	return func(c echo.Context) error {
		paramId := c.Param("id")
		id, err := strconv.Atoi(paramId)

		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid User Input Param ID", nil))
		}

		result, err := cmh.service.GetByID(id)

		if err != nil {
			c.Logger().Error("Handler : Get By ID Process Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Get By ID Process Error", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success Get Data", result))
	}
}
