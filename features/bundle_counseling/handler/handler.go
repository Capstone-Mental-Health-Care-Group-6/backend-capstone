package handler

import (
	bundlecounseling "FinalProject/features/bundle_counseling"
	"FinalProject/helper"
	"net/http"

	"github.com/labstack/echo/v4"
)

type BundleCounselingHandler struct {
	s   bundlecounseling.BundleCounselingServiceInterface
	jwt helper.JWTInterface
}

func New(s bundlecounseling.BundleCounselingServiceInterface, jwt helper.JWTInterface) bundlecounseling.BundleCounselingHandlerInterface {
	return &BundleCounselingHandler{
		s:   s,
		jwt: jwt,
	}
}

func (h *BundleCounselingHandler) GetAllBundle() echo.HandlerFunc {
	return func(c echo.Context) error {
		role := h.jwt.CheckRole(c)

		if role != "Admin" {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Only admin can access this page", nil))
		}

		result, err := h.s.GetAllBundle()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse(err.Error(), nil))
		}
		if len(result) == 0 {
			return c.JSON(http.StatusOK, helper.FormatResponse("Success", "Data is Empty"))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success", result))
	}
}

func (h *BundleCounselingHandler) CreateBundle() echo.HandlerFunc {
	return func(c echo.Context) error {
		role := h.jwt.CheckRole(c)

		if role != "Admin" {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Only admin can access this page", nil))
		}

		var req InputRequest

		if err := c.Bind(&req); err != nil {
			c.Logger().Fatal("Handler : Bind Input Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid Format Request", nil))
		}

		isValid, errors := helper.ValidateForm(req)
		if !isValid {
			return c.JSON(http.StatusBadRequest, helper.FormatResponseValidation("Invalid Format Request", errors))
		}

		file, _ := c.FormFile("avatar")

		isValidFile, errorsFile := helper.ValidateFile(file, 1*1024*1024, "image/jpeg", "image/png")
		if !isValidFile {
			return c.JSON(http.StatusBadRequest, helper.FormatResponseValidation("Invalid Format Request", errorsFile))
		}

		openFile, err := file.Open()
		if err != nil {
			c.Logger().Fatal("Handler : Open File Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Failed to create bundle", nil))
		}

		var serviceInput = new(bundlecounseling.BundleCounseling)
		serviceInput.Name = req.Name
		serviceInput.Sessions = req.Sessions
		serviceInput.Price = req.Price
		serviceInput.Type = req.Type
		serviceInput.Description = req.Description
		serviceInput.ActivePriode = req.ActivePriode

		result, err := h.s.CreateBundle(*serviceInput, bundlecounseling.BundleCounselingFile{Avatar: openFile})
		if err != nil {
			c.Logger().Fatal("Handler : Input Process Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Failed to create bundle", nil))
		}

		return c.JSON(http.StatusCreated, helper.FormatResponse("Success to create bundle", result))
	}
}
