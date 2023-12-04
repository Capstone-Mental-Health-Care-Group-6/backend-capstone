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

func (ah *ArticleHandler) GetArticles() echo.HandlerFunc {
	return func(c echo.Context) error {
		name := c.QueryParam("name")
		kategori := c.QueryParam("kategori")
		publication := c.QueryParam("publication")
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

		result, err := ah.s.GetArticles(name, kategori, timePublication)

		if err != nil {
			c.Logger().Fatal("Handler : Get All Process Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Fail", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success", result))
	}
}

func (ah *ArticleHandler) GetArticle() echo.HandlerFunc {
	return func(c echo.Context) error {
		var paramID = c.Param("id")
		id, err := strconv.Atoi(paramID)
		if err != nil {
			c.Logger().Fatal("Handler : Param ID Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", nil))
		}

		result, err := ah.s.GetArticle(id)

		if err != nil {
			c.Logger().Fatal("Handler : Get By ID Process Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Fail", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success", result))
	}
}

func (ah *ArticleHandler) CreateArticle() echo.HandlerFunc {
	return func(c echo.Context) error {
		role := ah.jwt.CheckRole(c)
		id, _ := ah.jwt.GetID(c)

		if role != "Doctor" {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Only admin can access this page", nil))
		}
		var input = new(InputRequest)
		if err := c.Bind(&input); err != nil {
			c.Logger().Fatal("Handler : Bind Input Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", nil))
		}

		var serviceInput = new(articles.Article)
		serviceInput.UserID = id
		serviceInput.CategoryID = input.CategoryID
		serviceInput.Title = input.Title
		serviceInput.Content = input.Content
		serviceInput.Thumbnail = input.Thumbnail
		serviceInput.Status = "Pending"
		serviceInput.Slug = input.Slug

		result, err := ah.s.CreateArticle(*serviceInput)

		if err != nil {
			c.Logger().Fatal("Handler : Input Process Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Fail", nil))
		}

		var response = new(InputResponse)
		response.UserID = result.UserID
		response.CategoryID = result.CategoryID
		response.Title = result.Title
		response.Content = result.Content
		response.Thumbnail = result.Thumbnail
		response.Status = result.Status
		response.Slug = result.Slug

		return c.JSON(http.StatusCreated, helper.FormatResponse("Success", response))
	}
}

func (ah *ArticleHandler) UpdateArticle() echo.HandlerFunc {
	return func(c echo.Context) error {
		role := ah.jwt.CheckRole(c)

		if role != "Doctor" {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Only admin can access this page", nil))
		}
		var paramID = c.Param("id")
		id, err := strconv.Atoi(paramID)
		if err != nil {
			c.Logger().Fatal("Handler : Param ID Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", nil))
		}

		var input = new(UpdateRequest)
		if err := c.Bind(&input); err != nil {
			c.Logger().Fatal("Handler : Bind Input Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", nil))
		}

		var serviceUpdate = new(articles.UpdateArticle)
		serviceUpdate.Title = input.Title
		serviceUpdate.Content = input.Content
		serviceUpdate.Thumbnail = input.Thumbnail
		serviceUpdate.Slug = input.Slug

		result, err := ah.s.UpdateArticle(*serviceUpdate, id)

		if err != nil {
			c.Logger().Fatal("Handler : Input Process Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Fail", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success", result))
	}
}

func (ah *ArticleHandler) DenyArticle() echo.HandlerFunc {
	return func(c echo.Context) error {
		role := ah.jwt.CheckRole(c)

		if role != "Admin" {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Only admin can access this page", nil))
		}
		var paramID = c.Param("id")
		id, err := strconv.Atoi(paramID)

		if err != nil {
			c.Logger().Fatal("Handler : Param ID Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", nil))
		}

		_, err = ah.s.DenyArticle(id)

		if err != nil {
			c.Logger().Fatal("Handler : Deny Article Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Fail", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success Denny Article", nil))
	}
}

func (ah *ArticleHandler) ApproveArticle() echo.HandlerFunc {
	return func(c echo.Context) error {
		role := ah.jwt.CheckRole(c)

		if role != "Admin" {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Only admin can access this page", nil))
		}
		var paramID = c.Param("id")
		id, err := strconv.Atoi(paramID)

		if err != nil {
			c.Logger().Fatal("Handler : Param ID Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", nil))
		}

		_, err = ah.s.ApproveArticle(id)

		if err != nil {
			c.Logger().Fatal("Handler : Approve Article Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Fail", nil))
		}

		return c.JSON(http.StatusNoContent, helper.FormatResponse("Success Approve Article", nil))
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
			c.Logger().Error("Handler: Callback process error: ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse(err.Error(), nil))
		}

		var response = new(DashboardResponse)
		response.TotalArticle = res.TotalArticle
		response.TotalArticleBaru = res.TotalArticleBaru
		response.TotalArticlePending = res.TotalArticlePending

		return c.JSON(http.StatusOK, helper.FormatResponse("Success get article dashboard", response))
	}
}
