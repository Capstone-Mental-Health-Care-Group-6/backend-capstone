package openai

type OpenAiConfig struct {
	key string
}

func NewOpenAiConfig(env OpenAiEnv) *OpenAiConfig {
	return &OpenAiConfig{
		key: env["OPENAI_API_KEY"].(string),
	}
}
