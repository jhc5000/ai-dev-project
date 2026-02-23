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
