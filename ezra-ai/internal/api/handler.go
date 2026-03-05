package api

import (
	"encoding/json"
	"net/http"

	"ezra-ai/internal/agent"
	"ezra-ai/internal/types"
)

type Handler struct {
	chat *agent.ChatAgent
}

func NewHandler(chat *agent.ChatAgent) *Handler {
	return &Handler{chat: chat}
}

func (h *Handler) Chat(w http.ResponseWriter, r *http.Request) {
	var req types.ChatRequest
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Allow common methods and headers
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	// Handle preflight (OPTIONS) requests
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	reply, err := h.chat.Chat(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(types.ChatResponse{Reply: reply})
}
