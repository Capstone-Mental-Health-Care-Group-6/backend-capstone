package handler

import (
	"FinalProject/features/articles"
	"FinalProject/helper"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ArticleHandler struct {
	s   articles.ArticleServiceInterface
	jwt helper.JWTInterface
}

func NewHandler(service articles.ArticleServiceInterface, jwt helper.JWTInterface) articles.ArticleHandlerInterface {
	return &ArticleHandler{
		s:   service,
		jwt: jwt,
	}
}

func (ah *ArticleHandler) GetArticlesByDoctorID() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := ah.jwt.GetID(c)
		if err != nil {
			c.Logger().Error("Handler : Get ID With JWT Error : ", err)
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Cannot Get ID With JWT", nil))
		}

		result, err := ah.s.GetArticleByDoctor(int(id))

		if err != nil {
			c.Logger().Error("Handler : Get Article By Doctor ID Process Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Get Article By Doctor ID", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success Get Data", result))
	}
}

func (ah *ArticleHandler) GetArticles() echo.HandlerFunc {
	return func(c echo.Context) error {
		name := c.QueryParam("name")
		kategori := c.QueryParam("kategori")
		publication := c.QueryParam("publication")
		limit := c.QueryParam("limit")

		limitInt, _ := strconv.Atoi(limit)
		var timePublication int

		switch publication {
		case "7Days":
			timePublication = 1
			break
		case "30Days":
			timePublication = 2
			break
		default:
			timePublication = -1
		}

		result, err := ah.s.GetArticles(name, kategori, timePublication, limitInt)

		if err != nil {
			c.Logger().Error("Handler : Get All Process Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Get All Process Error", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success Get All Data", result))
	}
}

func (ah *ArticleHandler) GetArticle() echo.HandlerFunc {
	return func(c echo.Context) error {
		var paramID = c.Param("id")
		id, err := strconv.Atoi(paramID)
		if err != nil {
			c.Logger().Error("Handler : Param ID Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid User Input Param ID", nil))
		}

		result, err := ah.s.GetArticle(id)

		if err != nil {
			c.Logger().Error("Handler : Get By ID Process Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Get By ID Process Error", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success Get Data", result))
	}
}

func (ah *ArticleHandler) CreateArticle() echo.HandlerFunc {
	return func(c echo.Context) error {
		role := ah.jwt.CheckRole(c)
		id, _ := ah.jwt.GetID(c)

		if role != "Doctor" {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Only doctor can access this page", nil))
		}
		var input = new(InputRequest)
		if err := c.Bind(input); err != nil {
			c.Logger().Error("Handler : Bind Input Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid User Input", nil))
		}

		var serviceInput = new(articles.Article)
		serviceInput.UserID = id
		serviceInput.CategoryID = input.CategoryID
		serviceInput.Title = input.Title
		serviceInput.Content = input.Content
		serviceInput.Status = "Pending"

		formHeaderThumbnail, err := c.FormFile("thumbnail")
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Failed, Select a File for Thumbnail", nil))
		}

		formThumbnail, err := formHeaderThumbnail.Open()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Failed to get Thumbnail", nil))
		}

		// uploadUrlThumbnail, err := ah.s.ThumbnailUpload(articles.ThumbnailDataModel{ThumbnailPhoto: formThumbnail})

		serviceInput.ThumbnailFile = formThumbnail

		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse(err.Error(), nil))
		}

		result, err := ah.s.CreateArticle(*serviceInput)

		if err != nil {
			c.Logger().Error("Handler : Input Process Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Input Process Error", nil))
		}

		var response = new(InputResponse)
		response.UserID = result.UserID
		response.CategoryID = result.CategoryID
		response.Title = result.Title
		response.Content = result.Content
		response.Thumbnail = result.ThumbnailUrl
		response.Status = result.Status
		response.Slug = result.Slug

		return c.JSON(http.StatusCreated, helper.FormatResponse("Success Create Data", response))
	}
}

func (ah *ArticleHandler) UpdateArticle() echo.HandlerFunc {
	return func(c echo.Context) error {
		role := ah.jwt.CheckRole(c)

		if role != "Doctor" {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Only doctor can access this page", nil))
		}
		var paramID = c.Param("id")
		id, err := strconv.Atoi(paramID)
		if err != nil {
			c.Logger().Error("Handler : Param ID Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid User Input Param ID", nil))
		}

		var input = new(UpdateRequest)
		if err := c.Bind(input); err != nil {
			c.Logger().Error("Handler : Bind Input Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid User Input", nil))
		}

		var serviceUpdate = new(articles.UpdateArticle)
		serviceUpdate.Title = input.Title
		serviceUpdate.Content = input.Content
		formHeaderThumbnail, err := c.FormFile("thumbnail")
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Failed, Select a File for Thumbnail", nil))
		}

		formThumbnail, err := formHeaderThumbnail.Open()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Failed to get Thumbnail", nil))
		}

		// uploadUrlThumbnail, err := ah.s.ThumbnailUpload(articles.ThumbnailDataModel{ThumbnailPhoto: formThumbnail})

		serviceUpdate.ThumbnailFile = formThumbnail

		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse(err.Error(), nil))
		}

		result, err := ah.s.UpdateArticle(*serviceUpdate, id)

		if err != nil {
			c.Logger().Error("Handler : Update Process Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Update Process Error", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success Update Data", result))
	}
}

func (ah *ArticleHandler) RejectArticle() echo.HandlerFunc {
	return func(c echo.Context) error {
		role := ah.jwt.CheckRole(c)

		if role != "Admin" {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Only admin can access this page", nil))
		}
		var paramID = c.Param("id")
		id, err := strconv.Atoi(paramID)

		if err != nil {
			c.Logger().Error("Handler : Param ID Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid User Input Param ID", nil))
		}

		_, err = ah.s.RejectArticle(id)

		if err != nil {
			c.Logger().Error("Handler : Reject Article Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse(err.Error(), nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success Reject Article", nil))
	}
}

func (ah *ArticleHandler) PublishArticle() echo.HandlerFunc {
	return func(c echo.Context) error {
		role := ah.jwt.CheckRole(c)

		if role != "Admin" {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Only admin can access this page", nil))
		}
		var paramID = c.Param("id")
		id, err := strconv.Atoi(paramID)

		if err != nil {
			c.Logger().Error("Handler : Param ID Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid User Input Param ID", nil))
		}

		_, err = ah.s.PublishArticle(id)

		if err != nil {
			c.Logger().Error("Handler : Publish Article Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse(err.Error(), nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success Publish Article", nil))
	}
}

func (ah *ArticleHandler) ArticleDashboard() echo.HandlerFunc {
	return func(c echo.Context) error {
		role := ah.jwt.CheckRole(c)
		if role != "Admin" {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Unauthorized", nil))
		}

		res, err := ah.s.ArticleDashboard()

		if err != nil {
			c.Logger().Error("Handler: Article Dashboard Process Error: ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Article Dashboard Process Error", nil))
		}

		var response = new(DashboardResponse)
		response.TotalArticle = res.TotalArticle
		response.TotalArticleBaru = res.TotalArticleBaru
		response.TotalArticlePending = res.TotalArticlePending

		return c.JSON(http.StatusOK, helper.FormatResponse("Success Get Article Dashboard", response))
	}
}
