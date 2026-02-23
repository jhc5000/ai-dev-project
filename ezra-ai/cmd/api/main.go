package main

import (
	"log"
	"net/http"

	"ezra-ai/internal/agent"
	"ezra-ai/internal/api"
	"ezra-ai/internal/journal"
	"ezra-ai/internal/llm"
)

func main() {
	llmClient := llm.NewOpenAIClient()
	journalStore := journal.NewStore()

	chatAgent := agent.NewChatAgent(llmClient, journalStore)
	handler := api.NewHandler(chatAgent)

	http.HandleFunc("/chat", handler.Chat)

	log.Println("EzraAI running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
