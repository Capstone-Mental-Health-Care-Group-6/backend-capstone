package handler

import (
	"FinalProject/features/users"
	"FinalProject/helper"
	"FinalProject/utils/oauth"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	s     users.UserServiceInterface
	oauth oauth.OauthGoogleInterface
}

func NewHandler(service users.UserServiceInterface, oauth oauth.OauthGoogleInterface) users.UserHandlerInterface {
	return &UserHandler{
		s:     service,
		oauth: oauth,
	}
}

func (uh *UserHandler) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = new(RegisterInput)
		if err := c.Bind(&input); err != nil {
			c.Logger().Info("Handler: Bind input error: ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", nil))
		}

		isValid, errors := helper.ValidateForm(input)
		if !isValid {
			return c.JSON(http.StatusBadRequest, helper.FormatResponseValidation("Invalid Format Request", errors))
		}

		var serviceInput = new(users.User)
		serviceInput.Name = input.Name
		serviceInput.Email = input.Email
		serviceInput.Password = input.Password
		serviceInput.Role = input.Role

		result, err := uh.s.Register(*serviceInput)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse(err.Error(), nil))
		}

		var response = new(RegisterResponse)
		response.Email = result.Email
		response.Name = result.Name
		response.Role = result.Role
		return c.JSON(http.StatusCreated, helper.FormatResponse("Success", response))
	}
}
func (uh *UserHandler) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = new(LoginInput)

		if err := c.Bind(input); err != nil {
			c.Logger().Error("Handler: Bind input error: ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", nil))
		}

		result, err := uh.s.Login(input.Email, input.Password)

		if err != nil {
			c.Logger().Error("Handler: Login process error: ", err.Error())
			if strings.Contains(err.Error(), "not found") {
				return c.JSON(http.StatusNotFound, helper.FormatResponse("Fail", nil))
			}
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Fail", nil))
		}

		var response = new(LoginResponse)
		response.Name = result.Name
		response.Email = result.Email
		response.Token = result.Access

		return c.JSON(http.StatusOK, helper.FormatResponse("Success", response))
	}
}

func (uh *UserHandler) LoginGoogle() echo.HandlerFunc {
	return func(c echo.Context) error {
		state, _ := uh.oauth.GenerateState()
		url, _ := uh.oauth.AuthCodeURL(state)
		return c.Redirect(http.StatusTemporaryRedirect, url)
	}
}

func (uh *UserHandler) CallbackGoogle() echo.HandlerFunc {
	return func(c echo.Context) error {
		code := c.QueryParam("code")
		token, err := uh.oauth.Exchange(code)

		if err != nil {
			c.Logger().Error("Handler: Callback process error: ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Unable to exchange code for token", nil))
		}

		email, err := uh.oauth.GetEmail(token)

		if err != nil {
			c.Logger().Error("Handler: Callback process error: ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Unable to get user email", nil))
		}

		result, err := uh.s.GenerateJwt(email)
		if err != nil {
			c.Logger().Error("Handler: Callback process error: ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse(err.Error(), nil))
		}

		var response = new(LoginResponse)
		response.Name = result.Name
		response.Email = result.Email
		response.Token = result.Access

		return c.JSON(http.StatusOK, helper.FormatResponse("Success", response))
	}
}
