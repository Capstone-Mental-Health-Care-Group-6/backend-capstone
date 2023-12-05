package handler

import (
	root "FinalProject/features/chats"
	"FinalProject/features/chats/dto"
	"FinalProject/helper"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ChatHttpHandler struct {
	srv root.ChatServiceInterface
}

func NewChatHttpHandler(srv root.ChatServiceInterface) root.ChatHandlerInterface {
	return &ChatHttpHandler{
		srv: srv,
	}
}

func (h *ChatHttpHandler) Establish() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		user, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			response := helper.ApiResponse[any]{
				Status:  http.StatusBadRequest,
				Message: "invalid user id",
			}
			return ctx.JSON(http.StatusBadRequest, response)
		}
		if h.srv.SocketEstablish(ctx, user) != nil {
			fmt.Printf("INFO: user#%d connected\n", user)
			return nil
		}
		return ctx.JSON(http.StatusInternalServerError, nil)
	}
}

func (h *ChatHttpHandler) Index() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		user, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			response := helper.ApiResponse[any]{
				Status:  http.StatusBadRequest,
				Message: "invalid user id",
			}
			return ctx.JSON(http.StatusBadRequest, response)
		}
		result := h.srv.GetChats(ctx, user)
		if ctx.Get("websocket.connection.error") != nil {
			response := helper.ApiResponse[any]{
				Status:  http.StatusUpgradeRequired,
				Message: "websocket connection not yet established",
			}
			return ctx.JSON(http.StatusUpgradeRequired, response)
		}
		response := helper.ApiResponse[any]{
			Status:  http.StatusOK,
			Message: "success",
			Data:    result,
		}
		return ctx.JSON(http.StatusOK, response)
	}
}

func (h *ChatHttpHandler) Store() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		request := &dto.CreateChatRequest{}
		if err := ctx.Bind(request); err != nil || request.Patient == request.Doctor {
			response := helper.ApiResponse[any]{
				Status:  http.StatusBadRequest,
				Message: "invalid create chat data payload",
			}
			return ctx.JSON(http.StatusBadRequest, response)
		}
		result := h.srv.CreateChat(ctx, request)
		if ctx.Get("websocket.connection.error") != nil {
			response := helper.ApiResponse[any]{
				Status:  http.StatusUpgradeRequired,
				Message: "websocket connection not yet established",
			}
			return ctx.JSON(http.StatusUpgradeRequired, response)
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

func (h *ChatHttpHandler) Edit() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		chat, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			response := helper.ApiResponse[any]{
				Status:  http.StatusBadRequest,
				Message: "invalid chat id",
			}
			return ctx.JSON(http.StatusBadRequest, response)
		}
		request := &dto.UpdateChatRequest{}
		if err := ctx.Bind(request); err != nil {
			response := helper.ApiResponse[any]{
				Status:  http.StatusBadRequest,
				Message: "invalid update chat data payload",
			}
			return ctx.JSON(http.StatusBadRequest, response)
		}
		result := h.srv.UpdateChat(ctx, chat, request)
		if ctx.Get("websocket.connection.error") != nil {
			response := helper.ApiResponse[any]{
				Status:  http.StatusUpgradeRequired,
				Message: "websocket connection not yet established",
			}
			return ctx.JSON(http.StatusUpgradeRequired, response)
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

func (h *ChatHttpHandler) Destroy() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		chat, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			response := helper.ApiResponse[any]{
				Status:  http.StatusBadRequest,
				Message: "invalid chat id",
			}
			return ctx.JSON(http.StatusBadRequest, response)
		}
		result := h.srv.DeleteChat(ctx, chat)
		if ctx.Get("websocket.connection.error") != nil {
			response := helper.ApiResponse[any]{
				Status:  http.StatusUpgradeRequired,
				Message: "websocket connection not yet established",
			}
			return ctx.JSON(http.StatusUpgradeRequired, response)
		}
		if result {
			return ctx.NoContent(http.StatusNoContent)
		}
		return ctx.JSON(http.StatusInternalServerError, nil)
	}
}
