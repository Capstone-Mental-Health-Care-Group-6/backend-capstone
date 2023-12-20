package handler

import (
	root "FinalProject/features/chats"
	"FinalProject/features/chats/dto"
	"FinalProject/helper"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type ChatHandler struct {
	srv root.ChatServiceInterface
}

func New(srv root.ChatServiceInterface) root.ChatHandlerInterface {
	return &ChatHandler{
		srv: srv,
	}
}

func (h *ChatHandler) Establish() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		user, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			response := helper.ApiResponse[any]{
				Status:  http.StatusBadRequest,
				Message: "invalid user id",
			}
			return ctx.JSON(http.StatusBadRequest, response)
		}
		role := ctx.Param("role")
		if role != "doctor" && role != "patient" {
			response := helper.ApiResponse[any]{
				Status:  http.StatusBadRequest,
				Message: "invalid role user",
			}
			return ctx.JSON(http.StatusBadRequest, response)
		}
		h.srv.SocketEstablish(ctx, user, role)
		if msg := ctx.Get("ws.client.error"); msg != nil {
			logrus.Infof("[ws.establish]: %s@%d not found", role, user)
			response := helper.ApiResponse[any]{
				Status:  http.StatusNotFound,
				Message: msg.(string),
			}
			return ctx.JSON(http.StatusNotFound, response)
		}
		if client := ctx.Get("ws.connect"); client != nil {
			logrus.Infof("[ws.establish]: client@%s connected", client)
		}
		return nil
	}
}

func (h *ChatHandler) Index() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		result := h.srv.GetChats(ctx)
		if ctx.Get("jwt.token.error") != nil {
			response := helper.ApiResponse[any]{
				Status:  http.StatusBadRequest,
				Message: "jwt token invalid",
			}
			return ctx.JSON(http.StatusBadRequest, response)
		}
		if ctx.Get("ws.connect.error") != nil {
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

func (h *ChatHandler) Store() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		request := &dto.CreateChatRequest{}
		if err := ctx.Bind(request); err != nil {
			response := helper.ApiResponse[any]{
				Status:  http.StatusBadRequest,
				Message: "invalid create chat data payload",
			}
			return ctx.JSON(http.StatusBadRequest, response)
		}
		result := h.srv.CreateChat(ctx, request)
		if msg := ctx.Get("ws.client.error"); msg != nil {
			logrus.Infof("[ws.store]: %s", msg)
			response := helper.ApiResponse[any]{
				Status:  http.StatusNotFound,
				Message: msg.(string),
			}
			return ctx.JSON(http.StatusNotFound, response)
		}
		if msg := ctx.Get("ws.connect.error"); msg != nil {
			logrus.Infof("[ws.store]: %s", msg)
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

func (h *ChatHandler) Edit() echo.HandlerFunc {
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
		if ctx.Get("jwt.token.error") != nil {
			response := helper.ApiResponse[any]{
				Status:  http.StatusBadRequest,
				Message: "jwt token invalid",
			}
			return ctx.JSON(http.StatusBadRequest, response)
		}
		if msg := ctx.Get("ws.connect.error"); msg != nil {
			logrus.Infof("[ws.edit]: %s", msg)
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

func (h *ChatHandler) Destroy() echo.HandlerFunc {
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
		if ctx.Get("jwt.token.error") != nil {
			response := helper.ApiResponse[any]{
				Status:  http.StatusBadRequest,
				Message: "jwt token invalid",
			}
			return ctx.JSON(http.StatusBadRequest, response)
		}
		if msg := ctx.Get("ws.connect.error"); msg != nil {
			logrus.Infof("[ws.destroy]: %s", msg)
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
