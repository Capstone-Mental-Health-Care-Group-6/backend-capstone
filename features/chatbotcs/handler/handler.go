package handler

import (
	"FinalProject/features/chatbotcs"
	"FinalProject/helper"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ChatbotCsHandler struct {
	srv chatbotcs.ChatbotCsServiceInterface
}

func New(srv chatbotcs.ChatbotCsServiceInterface) chatbotcs.ChatbotCsHandlerInterface {
	return &ChatbotCsHandler{
		srv: srv,
	}
}

func (ch *ChatbotCsHandler) ChatBotCs() echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Content-Type", "text/event-stream")
		c.Response().Header().Set("Cache-Control", "no-cache")
		c.Response().Header().Set("Connection", "keep-alive")

		messages := ch.srv.JoinGroup()
		defer ch.srv.LeaveGroup(messages)
		fmt.Fprintf(c.Response().Writer, "\n")
		c.Response().Flush()

		for {
			select {
			case message := <-messages:
				fmt.Fprintf(c.Response().Writer, "data: %s\n\n", message)
				c.Response().Flush()

			case <-c.Request().Context().Done():
				return nil
			}
		}
	}
}

func (ch *ChatbotCsHandler) CreateMessage() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req = new(InputRequest)
		if err := c.Bind(req); err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", err.Error()))
		}

		isValid, errors := helper.ValidateJSON(req)
		if !isValid {
			return c.JSON(http.StatusBadRequest, helper.FormatResponseValidation("Invalid Format Request", errors))
		}

		req.Type = "user"

		getResponBot := ch.srv.GetAnswer(req.Message)

		ch.srv.CreateMsg(req.Message)
		ch.srv.CreateMsg(getResponBot)

		return nil
	}
}
