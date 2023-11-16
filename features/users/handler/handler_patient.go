package handler

import (
	"FinalProject/features/users"
	"FinalProject/helper"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type PatientHandler struct {
	svc users.PatientServiceInterface
}

func NewHandlerPatient(service users.PatientServiceInterface) users.PatientHandlerInterface {
	return &PatientHandler{
		svc: service,
	}
}

func (mdl *PatientHandler) GetPatients() echo.HandlerFunc {
	return func(c echo.Context) error {
		result, err := mdl.svc.GetPatients()

		if err != nil {
			c.Logger().Fatal("Handler : Get All Process Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Fail", nil))
		}

		if len(result) == 0 {
			return c.JSON(http.StatusOK, helper.FormatResponse("Success", "Data is Empty"))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success", result))
	}
}

func (mdl *PatientHandler) GetPatient() echo.HandlerFunc {
	return func(c echo.Context) error {
		var paramID = c.Param("id")
		id, err := strconv.Atoi(paramID)
		if err != nil {
			c.Logger().Info("Handler : Param ID Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", "Invalid ID"))
		}

		result, err := mdl.svc.GetPatient(id)

		if err != nil {
			c.Logger().Fatal("Handler : Get By ID Process Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Fail", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success", result))
	}
}

func (mdl *PatientHandler) CreatePatient() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = new(PatientRequest)

		if err := c.Bind(input); err != nil {
			c.Logger().Fatal("Handler : Bind Input Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", nil))
		}

		formHeaderPhoto, err := c.FormFile("avatar")
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Failed, Select a File for Upload", nil))
		}

		formPhoto, err := formHeaderPhoto.Open()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Failed", nil))
		}

		uploadUrlPhoto, err := mdl.svc.PhotoUpload(users.AvatarPhoto{Avatar: formPhoto})
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Failed", nil))
		}

		var serviceInput = new(users.Patiententity)
		serviceInput.Name = input.Name
		serviceInput.UserID = input.UserID
		serviceInput.DateOfBirth = input.DateOfBirth
		serviceInput.PlaceOfBirth = input.PlaceOfBirth
		serviceInput.Gender = input.Gender
		serviceInput.MarriageStatus = input.MarriageStatus
		serviceInput.Avatar = uploadUrlPhoto
		serviceInput.Address = input.Address

		result, err := mdl.svc.CreatePatient(*serviceInput)

		//handling error untuk duplicate user id dan user id tidak ada di table users belum fix
		if err != nil {
			c.Logger().Info("Handler : Input Process Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", nil))
		}

		var response = new(PatientResponse)
		response.Name = result.Name
		response.UserID = result.UserID
		response.DateOfBirth = result.DateOfBirth
		response.PlaceOfBirth = result.PlaceOfBirth
		response.Gender = result.Gender
		response.MarriageStatus = result.MarriageStatus
		response.Avatar = result.Avatar
		response.Address = result.Address

		return c.JSON(http.StatusOK, helper.FormatResponse("Success", result))
	}
}
