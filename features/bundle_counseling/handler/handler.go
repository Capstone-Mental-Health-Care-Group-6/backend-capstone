package handler

import (
	bundlecounseling "FinalProject/features/bundle_counseling"
	"FinalProject/helper"
	"mime/multipart"
	"net/http"
	"strconv"

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
			c.Logger().Info("Handler : Bind Input Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid Format Request", nil))
		}

		isValid, errors := helper.ValidateForm(req)
		if !isValid {
			return c.JSON(http.StatusBadRequest, helper.FormatResponseValidation("Invalid Format Request", errors))
		}

		file, _ := c.FormFile("avatar")

		isValidFile, errorsFile := helper.ValidateFile(file, 5*1024*1024, "image/jpeg", "image/png")
		if !isValidFile {
			return c.JSON(http.StatusBadRequest, helper.FormatResponseValidation("Invalid Format Request", errorsFile))
		}

		openFile, err := file.Open()
		if err != nil {
			c.Logger().Info("Handler : Open File Error : ", err.Error())
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
			c.Logger().Info("Handler : Input Process Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Failed to create bundle", nil))
		}

		return c.JSON(http.StatusCreated, helper.FormatResponse("Success to create bundle", result))
	}
}

func (h *BundleCounselingHandler) GetBundle() echo.HandlerFunc {
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

		result, err := h.s.GetBundle(id)

		if err != nil {
			c.Logger().Info("Handler : Get By ID Process Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Fail to get bundle", nil))
		}

		if result.ID == 0 {
			return c.JSON(http.StatusNotFound, helper.FormatResponse("Success", "Data not found"))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success to get bundle", result))
	}
}

func (h *BundleCounselingHandler) UpdateBundle() echo.HandlerFunc {
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

		var req InputRequest

		if err := c.Bind(&req); err != nil {
			c.Logger().Info("Handler : Bind Input Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid Format Request", nil))
		}

		isValid, errors := helper.ValidateForm(req)
		if !isValid {
			return c.JSON(http.StatusBadRequest, helper.FormatResponseValidation("Invalid Format Request", errors))
		}

		var openFile multipart.File
		file, _ := c.FormFile("avatar")
		if file != nil {
			isValidFile, errorsFile := helper.ValidateFile(file, 5*1024*1024, "image/jpeg", "image/png")
			if !isValidFile {
				return c.JSON(http.StatusBadRequest, helper.FormatResponseValidation("Invalid Format Request", errorsFile))
			}

			openFileForm, err := file.Open()
			if err != nil {
				c.Logger().Info("Handler : Open File Error : ", err.Error())
				return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Failed to create bundle", nil))
			}

			openFile = openFileForm
		}

		var serviceInput = new(bundlecounseling.BundleCounseling)
		serviceInput.Name = req.Name
		serviceInput.Sessions = req.Sessions
		serviceInput.Price = req.Price
		serviceInput.Type = req.Type
		serviceInput.Description = req.Description
		serviceInput.ActivePriode = req.ActivePriode

		var bundleFile bundlecounseling.BundleCounselingFile
		if openFile != nil {
			bundleFile.Avatar = openFile
		}

		result, err := h.s.UpdateBundle(id, *serviceInput, bundleFile)

		if err != nil {
			c.Logger().Info("Handler : Input Process Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Failed to update bundle", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success to update bundle", result))
	}
}

func (h *BundleCounselingHandler) DeleteBundle() echo.HandlerFunc {
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

		result, err := h.s.DeleteBundle(id)

		if err != nil {
			c.Logger().Info("Handler : Delete Process Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Failed to delete bundle", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success to delete bundle", result))
	}
}

func (h *BundleCounselingHandler) GetAllBundleFilter() echo.HandlerFunc {
	return func(c echo.Context) error {

		jenisPaket := c.QueryParam("type")

		if jenisPaket == "" {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Type param is required", nil))
		}

		metodeID := c.QueryParam("metode")
		metodeIDInt, _ := strconv.Atoi(metodeID)

		if metodeID == "" {
			metodeIDInt = 1
		}

		durasiID := c.QueryParam("durasi")
		durasiIDInt, _ := strconv.Atoi(durasiID)

		if durasiID == "" {
			durasiIDInt = 1
		}

		result, err := h.s.GetAllBundleFilter(jenisPaket, metodeIDInt, durasiIDInt)

		if err != nil {
			c.Logger().Info("Handler : Get All Process Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Failed to get all bundle", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success to get all bundle", result))
	}
}
