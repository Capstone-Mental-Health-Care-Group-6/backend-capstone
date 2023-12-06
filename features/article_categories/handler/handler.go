package handler

import (
	articlecategories "FinalProject/features/article_categories"
	"FinalProject/helper"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ArticleCategoryHandler struct {
	s   articlecategories.ArticleCategoryServiceInterface
	jwt helper.JWTInterface
}

func NewHandler(service articlecategories.ArticleCategoryServiceInterface, j helper.JWTInterface) articlecategories.ArticleCategoryHandlerInterface {
	return &ArticleCategoryHandler{
		s:   service,
		jwt: j,
	}
}

func (ach *ArticleCategoryHandler) GetArticleCategories() echo.HandlerFunc {
	return func(c echo.Context) error {
		result, err := ach.s.GetArticleCategories()

		if err != nil {
			c.Logger().Fatal("Handler : Get All Process Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Fail", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success", result))
	}
}

func (ach *ArticleCategoryHandler) CreateArticleCategory() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = new(InputRequest)
		if err := c.Bind(&input); err != nil {
			c.Logger().Fatal("Handler : Bind Input Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", nil))
		}

		var serviceInput = new(articlecategories.ArticleCategory)
		serviceInput.Name = input.Name
		serviceInput.Slug = input.Slug

		result, err := ach.s.CreateArticleCategory(*serviceInput)
		if err != nil {
			c.Logger().Fatal("Handler : Input Process Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Fail", nil))
		}

		var response = new(InputResponse)
		response.Name = result.Name
		response.Slug = result.Slug

		return c.JSON(http.StatusCreated, helper.FormatResponse("Success", response))
	}
}

func (ach *ArticleCategoryHandler) GetArticleCategory() echo.HandlerFunc {
	return func(c echo.Context) error {
		var paramID = c.Param("id")
		id, err := strconv.Atoi(paramID)

		if err != nil {
			c.Logger().Fatal("Handler : Param ID Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", nil))
		}

		result, err := ach.s.GetArticleCategory(id)

		if err != nil {
			c.Logger().Fatal("Handler : Get Process Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Fail", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success", result))
	}
}

func (ach *ArticleCategoryHandler) UpdateArticleCategory() echo.HandlerFunc {
	return func(c echo.Context) error {
		var paramID = c.Param("id")
		id, err := strconv.Atoi(paramID)

		if err != nil {
			c.Logger().Fatal("Handler : Param ID Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", nil))
		}

		var input = new(InputRequest)
		if err := c.Bind(&input); err != nil {
			c.Logger().Fatal("Handler : Bind Input Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", nil))
		}

		var serviceInput = new(articlecategories.UpdateArticleCategory)
		serviceInput.Name = input.Name
		serviceInput.Slug = input.Slug

		result, err := ach.s.UpdateArticleCategory(*serviceInput, id)

		if err != nil {
			c.Logger().Fatal("Handler : Update Process Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Fail", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success", result))
	}
}

func (ach *ArticleCategoryHandler) DeleteArticleCategory() echo.HandlerFunc {
	return func(c echo.Context) error {
		var paramID = c.Param("id")
		id, err := strconv.Atoi(paramID)

		if err != nil {
			c.Logger().Fatal("Handler : Param ID Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", nil))
		}

		result, err := ach.s.DeleteArticleCategory(id)

		if err != nil {
			c.Logger().Fatal("Handler : Delete Process Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Fail", nil))
		}
		return c.JSON(http.StatusNoContent, helper.FormatResponse("Success", result))
	}
}
