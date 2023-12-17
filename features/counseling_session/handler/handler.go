package handler

import (
	counselingsession "FinalProject/features/counseling_session"
	"FinalProject/helper"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type CounselingSessionHandler struct {
	s   counselingsession.CounselingSessionServiceInterface
	jwt helper.JWTInterface
}

func New(s counselingsession.CounselingSessionServiceInterface, jwt helper.JWTInterface) counselingsession.CounselingSessionHandlerInterface {
	return &CounselingSessionHandler{
		s:   s,
		jwt: jwt,
	}
}

func (h *CounselingSessionHandler) GetAllCounseling() echo.HandlerFunc {
	return func(c echo.Context) error {
		// role := h.jwt.CheckRole(c)

		// if role != "Admin" {
		// 	return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Only admin can access this page", nil))
		// }

		result, err := h.s.GetAllCounseling()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse(err.Error(), nil))
		}
		if len(result) == 0 {
			return c.JSON(http.StatusOK, helper.FormatResponse("Success", "Data is Empty"))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success", result))
	}
}

func (h *CounselingSessionHandler) CreateCounseling() echo.HandlerFunc {
	return func(c echo.Context) error {
		// role := h.jwt.CheckRole(c)

		// if role != "Admin" {
		// 	return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Only admin can access this page", nil))
		// }

		var req InputRequest

		if err := c.Bind(&req); err != nil {
			c.Logger().Info("Handler : Bind Input Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid Format Request", nil))
		}

		//isValid, errors := helper.ValidateForm(req)
		//if !isValid {
		//	return c.JSON(http.StatusBadRequest, helper.FormatResponseValidation("Invalid Format Request", errors))
		//}

		var serviceInput = new(counselingsession.CounselingSession)
		serviceInput.TransactionID = req.TransactionID
		serviceInput.Date = req.Date
		serviceInput.Time = req.Time
		serviceInput.Duration = req.Duration
		serviceInput.Status = "process"
		serviceInput.Alasan = ""
		serviceInput.DetailAlasan = ""

		//CREATE COUNSELING HANDLER

		result, err := h.s.CreateCounseling(*serviceInput)
		if err != nil {
			c.Logger().Info("Handler : Input Process Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Failed to create bundle", nil))
		}

		var response = new(InputResponse)
		response.TransactionID = result.TransactionID
		response.Date = result.Date
		response.Time = result.Time
		response.Duration = result.Duration
		response.Status = result.Status
		serviceInput.Alasan = ""
		serviceInput.DetailAlasan = ""

		return c.JSON(http.StatusCreated, helper.FormatResponse("Success to create counseling session", response))
	}
}

func (h *CounselingSessionHandler) GetCounseling() echo.HandlerFunc {
	return func(c echo.Context) error {
		// role := h.jwt.CheckRole(c)

		// if role != "Admin" {
		// 	return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Only admin can access this page", nil))
		// }

		var paramID = c.Param("id")
		id, err := strconv.Atoi(paramID)

		if err != nil {
			c.Logger().Info("Handler : Param ID Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", "Invalid ID"))
		}

		result, err := h.s.GetCounseling(id)

		if err != nil {
			c.Logger().Info("Handler : Get By ID Process Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Fail to get counseling data", nil))
		}

		if result.TransactionID == 0 {
			return c.JSON(http.StatusNotFound, helper.FormatResponse("Success", "Data not found"))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success to get counseling data", result))
	}
}

func (h *CounselingSessionHandler) GetCounselingByUserID() echo.HandlerFunc {
	return func(c echo.Context) error {
		// role := h.jwt.CheckRole(c)

		// if role != "Admin" {
		// 	return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Only admin can access this page", nil))
		// }

		var paramID = c.Param("id")
		id, err := strconv.Atoi(paramID)

		if err != nil {
			c.Logger().Info("Handler : Param ID Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", "Invalid ID"))
		}

		result, err := h.s.GetAllCounselingByUserID(id)

		if err != nil {
			c.Logger().Info("Handler : Get By ID Process Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Fail to get counseling data", nil))
		}

		if len(result) == 0 {
			return c.JSON(http.StatusNotFound, helper.FormatResponse("Success", "Data not found"))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success to get counseling data", result))
	}
}

func (h *CounselingSessionHandler) UpdateCounseling() echo.HandlerFunc {
	return func(c echo.Context) error {
		// role := h.jwt.CheckRole(c)

		// if role != "Admin" {
		// 	return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Only admin can access this page", nil))
		// }

		var paramID = c.Param("id")
		id, err := strconv.Atoi(paramID)

		if err != nil {
			c.Logger().Info("Handler : Param ID Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", "Invalid ID"))
		}

		var req InputRequestUpdate

		if err := c.Bind(&req); err != nil {
			c.Logger().Info("Handler : Bind Input Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid Format Request", nil))
		}

		isValid, errors := helper.ValidateForm(req)
		if !isValid {
			return c.JSON(http.StatusBadRequest, helper.FormatResponseValidation("Invalid Format Request", errors))
		}

		var serviceInput = new(counselingsession.CounselingSession)
		serviceInput.TransactionID = req.TransactionID
		serviceInput.Date = req.Date
		serviceInput.Time = req.Time
		serviceInput.Duration = req.Duration
		serviceInput.Status = req.Status

		result, err := h.s.UpdateCounseling(id, *serviceInput)

		if err != nil {
			c.Logger().Info("Handler : Input Process Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Failed to update bundle", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success to update bundle", result))
	}
}

func (h *CounselingSessionHandler) DeleteCounseling() echo.HandlerFunc {
	return func(c echo.Context) error {
		role := h.jwt.CheckRole(c)

		if role != "Admin" {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Only admin can access this page", nil))
		}

		var paramID = c.Param("id")
		id, err := strconv.Atoi(paramID)

		if err != nil {
			c.Logger().Info("Handler : Param ID Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", "Invalid ID"))
		}

		result, err := h.s.DeleteCounseling(id)

		if err != nil {
			c.Logger().Info("Handler : Delete Process Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Failed to delete counseling", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success to delete bundle", result))
	}
}

func (h *CounselingSessionHandler) ApprovePatient() echo.HandlerFunc {
	return func(c echo.Context) error {
		role := h.jwt.CheckRole(c)

		fmt.Println(role)
		if role != "Doctor" {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Unauthorized", nil))
		}

		doctor_id, _ := h.jwt.GetID(c)
		doctorID := int(doctor_id)

		var paramID = c.Param("id")
		id, err := strconv.Atoi(paramID)
		if err != nil {
			c.Logger().Error("Handler : Param ID Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid Param ID", nil))
		}

		result, err := h.s.ApprovePatient(id, doctorID)

		if err != nil {
			if strings.Contains(err.Error(), "No data counseling found") {
				return c.JSON(http.StatusInternalServerError, helper.FormatResponse("No data counseling found", nil))
			}

			if strings.Contains(err.Error(), "No transaction data found") {
				return c.JSON(http.StatusInternalServerError, helper.FormatResponse("No transaction data found", nil))
			}

			if strings.Contains(err.Error(), "Unauthorized permission for this doctor") {
				return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Unauthorized permission for this doctor", nil))
			}
			c.Logger().Info("Handler : Update Status Process Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Update Status Process Failed", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success Update Status", result))
	}
}

func (h *CounselingSessionHandler) RejectPatient() echo.HandlerFunc {
	return func(c echo.Context) error {
		role := h.jwt.CheckRole(c)

		fmt.Println(role)
		if role != "Doctor" {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Unauthorized", nil))
		}

		doctor_id, _ := h.jwt.GetID(c)
		doctorID := int(doctor_id)

		var paramID = c.Param("id")
		id, err := strconv.Atoi(paramID)
		if err != nil {
			c.Logger().Error("Handler : Param ID Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid Param ID", nil))
		}

		var input = new(RequestStatusUpdate)
		if err := c.Bind(input); err != nil {
			c.Logger().Error("Handler : Bind Input Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid User Input", nil))
		}

		var serviceUpdate = new(counselingsession.StatusUpdate)
		serviceUpdate.Alasan = input.Alasan

		result, err := h.s.RejectPatient(id, doctorID, *serviceUpdate)

		if err != nil {
			if strings.Contains(err.Error(), "No data counseling found") {
				return c.JSON(http.StatusInternalServerError, helper.FormatResponse("No data counseling found", nil))
			}

			if strings.Contains(err.Error(), "No transaction data found") {
				return c.JSON(http.StatusInternalServerError, helper.FormatResponse("No transaction data found", nil))
			}

			if strings.Contains(err.Error(), "Unauthorized permission for this doctor") {
				return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Unauthorized permission for this doctor", nil))
			}
			c.Logger().Info("Handler : Update Status Process Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Update Status Process Failed", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success Update Status", result))
	}
}
