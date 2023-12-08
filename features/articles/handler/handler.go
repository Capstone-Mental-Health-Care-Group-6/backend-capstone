package handler

import (
	"FinalProject/features/articles"
	"FinalProject/helper"
	"fmt"
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
		role := ah.jwt.CheckRole(c)

		if role != "Admin" {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Only admin can access this page", nil))
		}

		result, err := ah.s.GetArticles()

		if err != nil {
			c.Logger().Info("Handler : Get All Process Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Fail", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success", result))
	}
}

func (ah *ArticleHandler) GetArticle() echo.HandlerFunc {
	return func(c echo.Context) error {
		role := ah.jwt.CheckRole(c)

		if role != "Admin" {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Only admin can access this page", nil))
		}
		var paramID = c.Param("id")
		id, err := strconv.Atoi(paramID)
		if err != nil {
			c.Logger().Info("Handler : Param ID Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", nil))
		}

		result, err := ah.s.GetArticle(id)

		if err != nil {
			c.Logger().Info("Handler : Get By ID Process Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Fail", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success", result))
	}
}

func (ah *ArticleHandler) CreateArticle() echo.HandlerFunc {
	return func(c echo.Context) error {
		role := ah.jwt.CheckRole(c)

		if role != "Admin" {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Only admin can access this page", nil))
		}
		var input = new(InputRequest)
		if err := c.Bind(&input); err != nil {
			c.Logger().Info("Handler : Bind Input Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", nil))
		}

		var serviceInput = new(articles.Article)
		serviceInput.UserID = input.AdminID
		serviceInput.CategoryID = input.CategoryID
		serviceInput.Title = input.Title
		serviceInput.Content = input.Content
		serviceInput.Thumbnail = input.Thumbnail
		serviceInput.Status = "Active"
		serviceInput.Slug = input.Slug

		result, err := ah.s.CreateArticle(*serviceInput)

		if err != nil {
			c.Logger().Info("Handler : Input Process Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Fail", nil))
		}

		var response = new(InputResponse)
		response.AdminID = result.UserID
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

		if role != "Admin" {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Only admin can access this page", nil))
		}
		var paramID = c.Param("id")
		id, err := strconv.Atoi(paramID)
		if err != nil {
			c.Logger().Info("Handler : Param ID Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", nil))
		}

		var input = new(UpdateRequest)
		if err := c.Bind(&input); err != nil {
			c.Logger().Info("Handler : Bind Input Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", nil))
		}

		var serviceUpdate = new(articles.UpdateArticle)
		serviceUpdate.Title = input.Title
		serviceUpdate.Content = input.Content
		serviceUpdate.Thumbnail = input.Thumbnail
		serviceUpdate.Slug = input.Slug

		result, err := ah.s.UpdateArticle(*serviceUpdate, id)

		if err != nil {
			c.Logger().Info("Handler : Input Process Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Fail", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success", result))
	}
}

func (ah *ArticleHandler) DeleteArticle() echo.HandlerFunc {
	return func(c echo.Context) error {
		role := ah.jwt.CheckRole(c)

		if role != "Admin" {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Only admin can access this page", nil))
		}
		var paramID = c.Param("id")
		id, err := strconv.Atoi(paramID)

		if err != nil {
			c.Logger().Info("Handler : Param ID Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", nil))
		}

		result, err := ah.s.DeleteArticle(id)

		if err != nil {
			c.Logger().Info("Handler : Delete Process Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Fail", nil))
		}

		fmt.Println(result)

		return c.JSON(http.StatusNoContent, helper.FormatResponse("Success", result))
	}
}
