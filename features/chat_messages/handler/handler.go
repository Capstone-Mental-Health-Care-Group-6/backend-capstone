package handler

import (
	root "FinalProject/features/chat_messages"
	"FinalProject/features/chat_messages/dto"
	"FinalProject/helper"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type MessageHandler struct {
	srv root.MessageServiceInterface
}

func New(srv root.MessageServiceInterface) root.MessageHandlerInterface {
	return &MessageHandler{
		srv: srv,
	}
}

func (h *MessageHandler) Index() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		chat, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			response := helper.ApiResponse[any]{
				Status:  http.StatusBadRequest,
				Message: "invalid chat id",
			}
			return ctx.JSON(http.StatusBadRequest, response)
		}
		response := helper.ApiResponse[any]{
			Status:  http.StatusOK,
			Message: "success",
			Data:    h.srv.GetMessages(ctx, chat),
		}
		return ctx.JSON(http.StatusOK, response)
	}
}

func (h *MessageHandler) Observe() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		chat, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			response := helper.ApiResponse[any]{
				Status:  http.StatusBadRequest,
				Message: "invalid chat id",
			}
			return ctx.JSON(http.StatusBadRequest, response)
		}
		message, err := strconv.Atoi(ctx.Param("message"))
		if err != nil {
			response := helper.ApiResponse[any]{
				Status:  http.StatusBadRequest,
				Message: "invalid message id",
			}
			return ctx.JSON(http.StatusBadRequest, response)
		}
		response := helper.ApiResponse[any]{
			Status:  http.StatusOK,
			Message: "success",
			Data:    h.srv.GetMessage(ctx, chat, message),
		}
		return ctx.JSON(http.StatusOK, response)
	}
}

func (h *MessageHandler) Store() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		chat, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			response := helper.ApiResponse[any]{
				Status:  http.StatusBadRequest,
				Message: "invalid chat id",
			}
			return ctx.JSON(http.StatusBadRequest, response)
		}
		request := &dto.Request{}
		if err := ctx.Bind(request); err != nil || strings.TrimSpace(request.Text) == "" {
			response := helper.ApiResponse[any]{
				Status:  http.StatusBadRequest,
				Message: "invalid message data payload",
			}
			return ctx.JSON(http.StatusBadRequest, response)
		}
		result := h.srv.CreateMessage(ctx, chat, request)
		if ctx.Get("jwt.token.error") != nil {
			response := helper.ApiResponse[any]{
				Status:  http.StatusBadRequest,
				Message: "jwt token invalid",
			}
			return ctx.JSON(http.StatusBadRequest, response)
		}
		if result != nil {
			response := helper.ApiResponse[any]{
				Status:  http.StatusCreated,
				Message: "success",
				Data:    result,
			}
			return ctx.JSON(http.StatusCreated, response)
		}
		return ctx.JSON(http.StatusInternalServerError, nil)
	}
}

func (h *MessageHandler) Edit() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		chat, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			response := helper.ApiResponse[any]{
				Status:  http.StatusBadRequest,
				Message: "invalid chat id",
			}
			return ctx.JSON(http.StatusBadRequest, response)
		}
		message, err := strconv.Atoi(ctx.Param("message"))
		if err != nil {
			response := helper.ApiResponse[any]{
				Status:  http.StatusBadRequest,
				Message: "invalid message id",
			}
			return ctx.JSON(http.StatusBadRequest, response)
		}
		request := &dto.Request{}
		if err := ctx.Bind(request); err != nil || strings.TrimSpace(request.Text) == "" {
			response := helper.ApiResponse[any]{
				Status:  http.StatusBadRequest,
				Message: "invalid message data payload",
			}
			return ctx.JSON(http.StatusBadRequest, response)
		}
		result := h.srv.UpdateMessage(ctx, chat, message, request)
		if ctx.Get("jwt.token.error") != nil {
			response := helper.ApiResponse[any]{
				Status:  http.StatusBadRequest,
				Message: "jwt token invalid",
			}
			return ctx.JSON(http.StatusBadRequest, response)
		}
		if result != nil {
			response := helper.ApiResponse[any]{
				Status:  http.StatusOK,
				Message: "success",
				Data:    result,
			}
			return ctx.JSON(http.StatusOK, response)
		}
		return ctx.JSON(http.StatusInternalServerError, nil)
	}
}

func (h *MessageHandler) Destroy() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		chat, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			response := helper.ApiResponse[any]{
				Status:  http.StatusBadRequest,
				Message: "invalid chat id",
			}
			return ctx.JSON(http.StatusBadRequest, response)
		}
		message, err := strconv.Atoi(ctx.Param("message"))
		if err != nil {
			response := helper.ApiResponse[any]{
				Status:  http.StatusBadRequest,
				Message: "invalid message id",
			}
			return ctx.JSON(http.StatusBadRequest, response)
		}
		if h.srv.DeleteMessage(ctx, chat, message) {
			return ctx.NoContent(http.StatusNoContent)
		}
		return ctx.JSON(http.StatusInternalServerError, nil)
	}
}
