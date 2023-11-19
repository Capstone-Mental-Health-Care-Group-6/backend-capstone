package handler

import (
	"FinalProject/features/doctor"
	"FinalProject/helper"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type DoctorHandler struct {
	svc doctor.DoctorServiceInterface
}

func NewHandlerDoctor(service doctor.DoctorServiceInterface) doctor.DoctorHandlerInterface {
	return &DoctorHandler{
		svc: service,
	}
}

func (mdl *DoctorHandler) GetDoctors() echo.HandlerFunc {
	return func(c echo.Context) error {
		result, err := mdl.svc.GetDoctors()

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

func (mdl *DoctorHandler) GetDoctor() echo.HandlerFunc {
	return func(c echo.Context) error {
		var paramID = c.Param("id")
		id, err := strconv.Atoi(paramID)
		if err != nil {
			c.Logger().Info("Handler : Param ID Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", "Invalid ID"))
		}

		result, err := mdl.svc.GetDoctor(id)

		if err != nil {
			c.Logger().Fatal("Handler : Get By ID Process Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Fail", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success", result))
	}

}

func (mdl *DoctorHandler) CreateDoctor() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = new(DoctorRequest)

		if err := c.Bind(input); err != nil {
			c.Logger().Fatal("Handler : Bind Input Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", nil))
		}

		// authorizationHeader := c.Request().Header.Get("Authorization")

		// // Check if the authorization header is valid
		// if !strings.Contains(authorizationHeader, "Bearer") {
		// 	return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Invalid auth token", nil))

		// }

		// jwtMapClaims, err := mdl.svc.JwtExtractToken(authorizationHeader)
		// if err != nil {
		// 	return c.JSON(http.StatusBadRequest, helper.FormatResponse("Failed to extract", nil))
		// }

		// if jwtMapClaims.Role == 0 || jwtMapClaims.Role == 1 {
		// 	return c.JSON(http.StatusBadRequest, helper.FormatResponse("Not allowed to register as doctor", nil))
		// }

		formHeaderPhoto, err := c.FormFile("doctor_avatar")
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Failed, Select a File for Upload Avatar", nil))
		}

		formHeaderSIPP, err := c.FormFile("doctor_sipp_file")
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Failed, Select a File for Upload SIPP", nil))
		}

		formHeaderSTR, err := c.FormFile("doctor_str_file")
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Failed, Select a File for Upload STR", nil))
		}

		formHeaderCV, err := c.FormFile("doctor_cv")
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Failed, Select a File for Upload CV", nil))
		}

		formHeaderIjazah, err := c.FormFile("doctor_ijazah")
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Failed, Select a File for Upload Ijazah", nil))
		}

		formPhoto, err := formHeaderPhoto.Open()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Failed FormHeaderPhoto", nil))
		}

		formSIPP, err := formHeaderSIPP.Open()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Failed FormHeaderSIPP", nil))
		}

		formSTR, err := formHeaderSTR.Open()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Failed FormHeaderSTR", nil))
		}

		formCV, err := formHeaderCV.Open()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Failed FormHeaderCV", nil))
		}

		formIjazah, err := formHeaderIjazah.Open()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Failed FormHeaderIjazah", nil))
		}

		uploadUrlPhoto, err := mdl.svc.DoctorAvatarUpload(doctor.DoctorAvatarPhoto{DoctorAvatar: formPhoto})
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Failed Upload Avatar", nil))
		}

		uploadUrlSIPP, err := mdl.svc.DoctorSIPPUpload(doctor.DoctorSIPPFileDataModel{DoctorSIPPFile: formSIPP})
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Failed Upload SIPP", nil))
		}

		uploadUrlSTR, err := mdl.svc.DoctorSTRUpload(doctor.DoctorSTRFileDataModel{DoctorSTRFile: formSTR})
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Failed Upload STR", nil))
		}

		uploadUrlCV, err := mdl.svc.DoctorCVUpload(doctor.DoctorCVDataModel{DoctorCV: formCV})
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Failed Upload CV", nil))
		}

		uploadUrlIjazah, err := mdl.svc.DoctorIjazahUpload(doctor.DoctorIjazahDataModel{DoctorIjazah: formIjazah})
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Failed Upload Ijazah", nil))
		}

		var serviceInput = new(doctor.Doctor)

		serviceInput.UserID = 3
		serviceInput.DoctorName = input.DoctorName
		serviceInput.DoctorExperience = input.DoctorExperience
		serviceInput.DoctorDescription = input.DoctorDescription
		serviceInput.DoctorAvatar = uploadUrlPhoto

		serviceInput.DoctorOfficeName = input.DoctorOfficeName
		serviceInput.DoctorOfficeAddress = input.DoctorOfficeAddress
		serviceInput.DoctorOfficeCity = input.DoctorOfficeCity
		serviceInput.DoctorMeetLink = input.DoctorMeetLink

		serviceInput.DoctorSIPPFile = uploadUrlSIPP
		serviceInput.DoctorSTRFile = uploadUrlSTR
		serviceInput.DoctorCV = uploadUrlCV
		serviceInput.DoctorIjazah = uploadUrlIjazah

		serviceInput.DoctorBalance = 0
		serviceInput.DoctorStatus = "active"
		//INPUT REQUEST

		result, err := mdl.svc.CreateDoctor(*serviceInput)

		//handling error untuk duplicate user id dan user id tidak ada di table users belum fix
		if err != nil {
			c.Logger().Info("Handler : Input Process Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", nil))
		}

		var response = new(DoctorResponse)

		response.ID = result.ID
		response.UserID = result.UserID
		response.DoctorName = result.DoctorName
		response.DoctorExperience = result.DoctorExperience
		response.DoctorDescription = result.DoctorDescription
		response.DoctorAvatar = result.DoctorAvatar
		response.DoctorOfficeName = result.DoctorOfficeName
		response.DoctorOfficeAddress = result.DoctorOfficeAddress
		response.DoctorOfficeCity = result.DoctorOfficeCity
		response.DoctorMeetLink = result.DoctorMeetLink
		response.DoctorSIPPFile = result.DoctorSIPPFile
		response.DoctorSTRFile = result.DoctorSTRFile
		response.DoctorCV = result.DoctorCV
		response.DoctorIjazah = result.DoctorIjazah
		response.DoctorBalance = result.DoctorBalance
		response.DoctorStatus = result.DoctorStatus

		//INPUT RESPONSE

		return c.JSON(http.StatusOK, helper.FormatResponse("Success", response))
	}
}
