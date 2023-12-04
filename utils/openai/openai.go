package openai

import (
	"context"

	ai "github.com/sashabaranov/go-openai"
	"github.com/sirupsen/logrus"
)

type OpenAiSrv struct {
	client *ai.Client
}

func NewOpenAiSrv() *OpenAiSrv {
	env := NewOpenAiEnv()
	key := NewOpenAiConfig(env).key
	return &OpenAiSrv{
		client: ai.NewClient(key),
	}
}

func (srv *OpenAiSrv) Chatbot(prompt string) *ai.ChatCompletionMessage {
	ctx := context.Background()
	req := ai.ChatCompletionRequest{
		Model:       ai.GPT3Dot5Turbo,
		Temperature: 1,
		MaxTokens:   2048,
		Messages: []ai.ChatCompletionMessage{
			{
				Role:    ai.ChatMessageRoleAssistant,
				Content: "Kamu adalah seorang psikolog profesional dan seorang puitisi terkenal, kamu akan diberikan sebuah pertanyaan seputar kesehatan mental, jawaban yang kamu berikan harus berbentuk narasi",
			},
			{
				Role:    ai.ChatMessageRoleUser,
				Content: prompt,
			},
		},
	}
	res, err := srv.client.CreateChatCompletion(ctx, req)
	if err != nil {
		logrus.Error("[openai.srv]: ", err.Error())
		return nil
	}
	return &res.Choices[0].Message
}
