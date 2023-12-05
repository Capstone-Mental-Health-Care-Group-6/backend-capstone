package openai

import (
	"FinalProject/configs"
	"context"

	"github.com/sashabaranov/go-openai"
)

type OpenAIInterface interface {
	GenerateText(prompt string) (string, error)
}

type OpenAI struct {
	cfg configs.ProgrammingConfig
}

func InitOpenAI(config configs.ProgrammingConfig) OpenAIInterface {
	return &OpenAI{
		cfg: config,
	}
}

func (o *OpenAI) GenerateText(prompt string) (string, error) {
	client := openai.NewClient(o.cfg.OpenAI)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleAssistant,
					Content: "Kamu adalah seorang psikolog profesional dan seorang puitisi terkenal, kamu akan diberikan sebuah pertanyaan seputar kesehatan mental, jawaban yang kamu berikan harus berbentuk narasi",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
		},
	)
	if err != nil {
		return "", err
	}
	return resp.Choices[0].Message.Content, nil
}
