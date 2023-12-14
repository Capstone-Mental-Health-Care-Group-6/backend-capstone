package handler

import (
	"FinalProject/features/patients"
	"FinalProject/helper"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type PatientHandler struct {
	svc patients.PatientServiceInterface
	jwt helper.JWTInterface
}

func NewHandlerPatient(service patients.PatientServiceInterface, jwt helper.JWTInterface) patients.PatientHandlerInterface {
	return &PatientHandler{
		svc: service,
		jwt: jwt,
	}
}

func (mdl *PatientHandler) GetPatients() echo.HandlerFunc {
	return func(c echo.Context) error {
		status := c.QueryParam("status")
		name := c.QueryParam("name")

		role := mdl.jwt.CheckRole(c)
		fmt.Println(role)
		if role != "Admin" && role != "Patient" && role != "Doctor" {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Unauthorized", nil))
		}

		result, err := mdl.svc.GetPatients(status, name)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Cannot process data", nil))
		}

		if len(result) == 0 {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Success", "Data is Empty"))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success", result))
	}
}

func (mdl *PatientHandler) CreatePatient() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = new(PatientRequest)

		if err := c.Bind(input); err != nil {
			c.Logger().Error("Handler : Bind Input Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", nil))
		}

		//formHeaderPhoto, err := c.FormFile("avatar")
		//if err != nil {
		//	return c.JSON(http.StatusBadRequest, helper.FormatResponse("Failed, Select a File for Upload", nil))
		//}

		//isValidFile, errorsFile := helper.ValidateFile(formHeaderPhoto, 5*1024*1024, "image/jpeg", "image/png")
		//if !isValidFile {
		//	return c.JSON(http.StatusBadRequest, helper.FormatResponseValidation("Invalid Format Request", errorsFile))
		//}

		//formPhoto, err := formHeaderPhoto.Open()
		//if err != nil {
		//	return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Failed", nil))
		//}

		//uploadUrlPhoto, err := mdl.svc.PhotoUpload(patients.AvatarPhoto{Avatar: formPhoto})
		//if err != nil {
		//	return c.JSON(http.StatusBadRequest, helper.FormatResponse("Failed", nil))
		//}

		var serviceInput = new(patients.Patiententity)

		isValid, errors := helper.ValidateForm(serviceInput)
		if !isValid {
			return c.JSON(http.StatusBadRequest, helper.FormatResponseValidation("Invalid Format Request", errors))
		}
		serviceInput.Name = input.Name
		serviceInput.Email = input.Email
		serviceInput.Password = input.Password
		serviceInput.DateOfBirth = input.DateOfBirth
		serviceInput.Gender = input.Gender
		serviceInput.Phone = input.Phone
		serviceInput.Avatar = "https://res.cloudinary.com/du87kowmp/image/upload/v1702560428/Avatar/fc2dgtshfu2w9hlhgqed.png"
		serviceInput.Role = "Patient"
		serviceInput.Status = "Active"
		result, err := mdl.svc.CreatePatient(*serviceInput)

		if err != nil {
			c.Logger().Info("Handler : Input Process Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", nil))
		}

		var response = new(PatientResponse)
		response.Name = result.Name
		response.Email = result.Email
		response.DateOfBirth = result.DateOfBirth
		response.Gender = result.Gender
		response.Phone = result.Phone
		response.Avatar = result.Avatar
		response.Role = result.Role
		response.Status = result.Status

		return c.JSON(http.StatusCreated, helper.FormatResponse("Success", response))
	}
}

func (mdl *PatientHandler) UpdatePatient() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := mdl.jwt.CheckID(c)
		userIdInt := int(id.(float64))
		var input = new(UpdateProfile)
		if err := c.Bind(input); err != nil {
			c.Logger().Info("Handler : Bind Input Error : ", err.Error())
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

		uploadUrlPhoto, err := mdl.svc.PhotoUpload(patients.AvatarPhoto{Avatar: formPhoto})
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Failed", nil))
		}

		var serviceUpdate = new(patients.UpdateProfile)
		serviceUpdate.Name = input.Name
		serviceUpdate.Email = input.Email
		serviceUpdate.DateOfBirth = input.DateOfBirth
		serviceUpdate.Gender = input.Gender
		serviceUpdate.Phone = input.Phone
		serviceUpdate.Avatar = uploadUrlPhoto

		result, err := mdl.svc.UpdatePatient(userIdInt, *serviceUpdate)

		if err != nil {
			c.Logger().Info("Handler : Input Process Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Fail", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success", result))
	}
}

func (mdl *PatientHandler) LoginPatient() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = new(LoginPatient)

		if err := c.Bind(input); err != nil {
			c.Logger().Error("Handler: Bind input error: ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", nil))
		}

		result, err := mdl.svc.LoginPatient(input.Email, input.Password)

		if err != nil {
			c.Logger().Error("Handler: Login process error: ", err.Error())
			if strings.Contains(err.Error(), "not found") {
				return c.JSON(http.StatusNotFound, helper.FormatResponse("Fail", nil))
			}
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", nil))
		}

		var response = new(PatientLoginResponse)
		response.Name = result.Name
		response.Email = result.Email
		response.Token = result.Access

		return c.JSON(http.StatusOK, helper.FormatResponse("Success", response))
	}
}

func (mdl *PatientHandler) UpdatePassword() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := mdl.jwt.CheckID(c)
		userIdInt := int(id.(float64))
		var input = new(UpdatePassword)
		if err := c.Bind(input); err != nil {
			c.Logger().Info("Handler : Bind Input Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", nil))
		}

		var serviceUpdate = new(patients.UpdatePassword)
		serviceUpdate.Password = input.Password

		result, err := mdl.svc.UpdatePassword(userIdInt, *serviceUpdate)

		if err != nil {
			c.Logger().Info("Handler : Input Process Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Fail", nil))
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
			c.Logger().Info("Handler : Get By ID Process Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Fail", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success", result))
	}
}

func (mdl *PatientHandler) PatientDashboard() echo.HandlerFunc {
	return func(c echo.Context) error {
		role := mdl.jwt.CheckRole(c)
		if role != "Admin" {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Unauthorized", nil))
		}

		res, err := mdl.svc.PatientDashboard()

		if err != nil {
			c.Logger().Error("Handler: Callback process error: ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse(err.Error(), nil))
		}

		var response = new(DashboardResponse)
		response.TotalUser = res.TotalUser
		response.TotalUserBaru = res.TotalUserBaru
		response.TotalUserActive = res.TotalUserActive
		response.TotalUserInactive = res.TotalUserInactive

		return c.JSON(http.StatusOK, helper.FormatResponse("Success get patient", response))
	}
}

func (mdl *PatientHandler) UpdateStatus() echo.HandlerFunc {
	return func(c echo.Context) error {
		role := mdl.jwt.CheckRole(c)
		fmt.Println(role)
		if role != "Admin" {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Unauthorized", nil))
		}
		var paramID = c.Param("id")
		id, err := strconv.Atoi(paramID)
		if err != nil {
			c.Logger().Info("Handler : Param ID Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", "Invalid ID"))
		}

		var req UpdateStatus

		if err := c.Bind(&req); err != nil {
			c.Logger().Info("Handler : Bind Input Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid Format Request", nil))
		}

		var serviceUpdate = new(patients.UpdateStatus)
		serviceUpdate.Status = req.Status

		result, err := mdl.svc.UpdateStatus(id, *serviceUpdate)

		if err != nil {
			c.Logger().Info("Handler : Input Process Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Fail", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success", result))
	}
}

func (mdl *PatientHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := mdl.jwt.CheckID(c)
		userIdInt := int(id.(float64))

		result, err := mdl.svc.Delete(userIdInt)

		if err != nil {
			c.Logger().Info("Handler : Delete Process Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Failed to delete bundle", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success to delete account", result))
	}
}
