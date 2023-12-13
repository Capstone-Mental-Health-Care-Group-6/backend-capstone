package service

import (
	"FinalProject/features/chatbot"
	"FinalProject/features/chatbot/mocks"
	mockUtil "FinalProject/utils/mocks"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetAllChatbot(t *testing.T) {
	data := mocks.NewChatbotDataInterface(t)
	openai := mockUtil.NewOpenAIInterface(t)
	service := New(data, openai)
	dataChat := []chatbot.Chatbot{}

	t.Run("Server Error", func(t *testing.T) {
		data.On("GetAllChatBot", 1).Return(nil, errors.New("Get All Process Failed")).Once()

		res, err := service.GetAllChatBot(1)

		assert.Error(t, err)
		assert.Nil(t, res)
		assert.EqualError(t, err, "Get All Process Failed")
	})

	t.Run("Success Get", func(t *testing.T) {
		data.On("GetAllChatBot", 1).Return(dataChat, nil).Once()

		res, err := service.GetAllChatBot(1)

		assert.Nil(t, err)
		assert.NotNil(t, res)
		data.AssertExpectations(t)
	})
}

func TestInsertChatBot(t *testing.T) {
	data := mocks.NewChatbotDataInterface(t)
	openai := mockUtil.NewOpenAIInterface(t)
	service := New(data, openai)
	dataChat := chatbot.Chatbot{
		UserID:       1,
		Prompt:       "prompt",
		ResultPrompt: "result",
		Date:         time.Now(),
	}

	t.Run("Generate Text Error", func(t *testing.T) {
		openai.On("GenerateText", dataChat.Prompt).Return("", errors.New("Generate Text Error")).Once()

		res, err := service.InsertChatBot(dataChat)

		assert.Error(t, err)
		assert.NotNil(t, res)
		assert.EqualError(t, err, "Generate Text Error")
	})

	t.Run("Server Error", func(t *testing.T) {
		openai.On("GenerateText", dataChat.Prompt).Return(dataChat.ResultPrompt, nil).Once()
		data.On("InsertChatBot", dataChat).Return(dataChat, errors.New("Insert Process Failed")).Once()

		res, err := service.InsertChatBot(dataChat)

		assert.Error(t, err)
		assert.NotNil(t, res)
		assert.EqualError(t, err, "Insert Process Failed")
	})

	t.Run("Success Insert", func(t *testing.T) {
		openai.On("GenerateText", dataChat.Prompt).Return(dataChat.ResultPrompt, nil).Once()
		data.On("InsertChatBot", dataChat).Return(dataChat, nil).Once()

		res, err := service.InsertChatBot(dataChat)

		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, res.UserID, dataChat.UserID)
		assert.Equal(t, res.Prompt, dataChat.Prompt)
		assert.Equal(t, res.ResultPrompt, dataChat.ResultPrompt)
		assert.Equal(t, res.Date, dataChat.Date)
		data.AssertExpectations(t)
		openai.AssertExpectations(t)
	})
}
