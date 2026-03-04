package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	OpenAIKey  string
	ServerPort string
}

func Load() Config {
	err := godotenv.Load("/workspaces/ai-dev-project/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	cfg := Config{
		OpenAIKey:  os.Getenv("OPENAI_API_KEY2"),
		ServerPort: os.Getenv("PORT"),
	}

	if cfg.ServerPort == "" {
		cfg.ServerPort = ":8080"
	}

	if cfg.OpenAIKey == "" {
		log.Fatal("OPENAI_API_KEY is not set")
	}

	return cfg
}
