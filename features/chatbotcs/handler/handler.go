package handler

import (
	"FinalProject/features/chatbotcs"
	"FinalProject/helper"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ChatbotCsHandler struct {
	srv chatbotcs.ChatbotCsServiceInterface
	jwt helper.JWTInterface
}

func New(srv chatbotcs.ChatbotCsServiceInterface, jwt helper.JWTInterface) chatbotcs.ChatbotCsHandlerInterface {
	return &ChatbotCsHandler{
		srv: srv,
		jwt: jwt,
	}
}

func (ch *ChatbotCsHandler) ChatBotCs() echo.HandlerFunc {
	return func(c echo.Context) error {
		ip := c.RealIP() // Mengambil IP dari pengguna

		c.Response().Header().Set("Content-Type", "text/event-stream")
		c.Response().Header().Set("Cache-Control", "no-cache")
		c.Response().Header().Set("Connection", "keep-alive")

		messages := ch.srv.JoinGroup(ip)
		defer ch.srv.LeaveGroup(ip)

		welcome := chatbotcs.ChatbotCs{
			Message: "Selamat datang di Aplikasi Kesehatan Mental kami! Saya akan dengan senang hati membantu Anda memahami fitur-fitur yang tersedia. Berikut beberapa hal yang dapat Anda lakukan:",
			Type:    "bot",
		}

		jsonWel, err := json.Marshal(welcome)
		if err != nil {
			return err
		}

		c.Response().Writer.Write([]byte("data: "))
		c.Response().Writer.Write(jsonWel)
		c.Response().Writer.Write([]byte("\n\n"))
		c.Response().Flush()

		for {
			select {
			case message := <-messages:
				// fmt.Fprintf(c.Response().Writer, "data: %s\n\n", message)
				// if message.Type == "bot" {
				// 	fmt.Fprintf(c.Response().Writer, "data: %s\n\n", message.Message)
				// 	c.Response().Flush()
				// } else {
				// 	fmt.Fprintf(c.Response().Writer, "data: %s\n\n", message.Message)
				// }
				// if err := json.NewEncoder(c.Response().Writer).Encode(message); err != nil {
				// 	return err
				// }

				jsonMessage, err := json.Marshal(message)
				if err != nil {
					return err
				}
				c.Response().Writer.Write([]byte("data: "))
				c.Response().Writer.Write(jsonMessage)
				c.Response().Writer.Write([]byte("\n\n"))
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

		getResponBot := ch.srv.GetAnswer(req.Message)

		ip := c.RealIP() // Mengambil IP dari pengguna

		chatUser := chatbotcs.ChatbotCs{
			Message: req.Message,
			Type:    "user",
		}

		responBot := chatbotcs.ChatbotCs{
			Message: getResponBot,
			Type:    "bot",
		}
		ch.srv.CreateMsg(ip, chatUser)
		ch.srv.CreateMsg(ip, responBot)

		return c.JSON(http.StatusOK, helper.FormatResponse("Success send message", true))
	}
}
