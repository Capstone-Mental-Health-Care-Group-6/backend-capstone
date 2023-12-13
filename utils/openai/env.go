package openai

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
	log "github.com/sirupsen/logrus"
)

type OpenAiEnv map[string]any

func NewOpenAiEnv() (env OpenAiEnv) {
	// create OpenAiEnv object
	env = make(OpenAiEnv)

	// check OPENAI_API_KEY environment variable
	if value, found := os.LookupEnv("OPENAI_API_KEY"); !found {
		log.Fatal("env: OPENAI_API_KEY not found")
	} else {
		env["OPENAI_API_KEY"] = value
	}

	return
}
