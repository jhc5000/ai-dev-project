package main

import (
	"log"

	"ezra-ai/internal/api"
	"ezra-ai/internal/chat"
	"ezra-ai/internal/config"
	"ezra-ai/internal/llm"
)

func main() {
	// 1️⃣ Load config (env vars, keys, ports)
	cfg := config.Load()

	// 2️⃣ Initialize LLM client
	llmClient, err := llm.NewClient(cfg)
	if err != nil {
		log.Fatalf("failed to initialize LLM client: %v", err)
	}

	// 3️⃣ Initialize chat memory
	memory := chat.NewInMemoryMemory()

	// 4️⃣ Initialize chat service
	chatService := chat.NewService(llmClient, memory)

	// 5️⃣ Initialize HTTP handlers
	chatHandler := api.NewChatHandler(chatService)

	// 6️⃣ Initialize and start router
	router := api.NewRouter(chatHandler)

	log.Println("🚀 EzraAI is running...")
	router.Start(cfg.ServerPort)
}
