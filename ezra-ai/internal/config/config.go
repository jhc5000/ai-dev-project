package config

import (
	"log"
	"os"
)

type Config struct {
	OpenAIKey  string
	ServerPort string
}

func Load() Config {
	cfg := Config{
		OpenAIKey:  os.Getenv("OPENAI_API_KEY2"),
		ServerPort: os.Getenv("SERVER_PORT"),
	}

	if cfg.ServerPort == "" {
		cfg.ServerPort = ":8080"
	}

	if cfg.OpenAIKey == "" {
		log.Fatal("OPENAI_API_KEY is not set")
	}

	return cfg
}
