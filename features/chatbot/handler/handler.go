package handler

import (
	"FinalProject/features/chatbot"
	"FinalProject/helper"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type ChatbotHandler struct {
	svc chatbot.ChatbotServiceInterface
	jwt helper.JWTInterface
}

func New(svc chatbot.ChatbotServiceInterface, jwt helper.JWTInterface) chatbot.ChatbotHandlerInterface {
	return &ChatbotHandler{
		svc: svc,
		jwt: jwt,
	}
}

func (h *ChatbotHandler) GetAllChatBot() echo.HandlerFunc {
	return func(c echo.Context) error {
		role := h.jwt.CheckRole(c)

		allowedRoles := []string{"Doctor", "Patient"}

		allowed := false
		for _, allowedRole := range allowedRoles {
			if role == allowedRole {
				allowed = true
				break
			}
		}

		if !allowed {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse("You don't have permission", nil))
		}

		idJwt, err := h.jwt.GetID(c)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Token is not valid", nil))
		}

		result, err := h.svc.GetAllChatBot(int(idJwt))

		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse(err.Error(), nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success to get chatbot", result))
	}
}

func (h *ChatbotHandler) CreateChatBot() echo.HandlerFunc {
	return func(c echo.Context) error {
		role := h.jwt.CheckRole(c)

		allowedRoles := []string{"Doctor", "Patient"}

		allowed := false
		for _, allowedRole := range allowedRoles {
			if role == allowedRole {
				allowed = true
				break
			}
		}

		if !allowed {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse("You don't have permission", nil))
		}

		idJwt, err := h.jwt.GetID(c)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Token is not valid", nil))
		}

		var req = new(InputRequest)
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse(err.Error(), nil))
		}

		isValid, errors := helper.ValidateJSON(req)
		if !isValid {
			return c.JSON(http.StatusBadRequest, helper.FormatResponseValidation("Invalid Format Request", errors))
		}

		var serviceInput = new(chatbot.Chatbot)
		serviceInput.UserID = idJwt
		serviceInput.Prompt = req.Prompt
		serviceInput.Date = time.Now()

		result, err := h.svc.InsertChatBot(*serviceInput)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse(err.Error(), nil))
		}

		return c.JSON(http.StatusCreated, helper.FormatResponse("Success to create chatbot", result))
	}
}
