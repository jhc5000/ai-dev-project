package api

import (
	"encoding/json"
	"net/http"

	"ezra-ai/internal/chat"
)

type ChatHandler struct {
	chatService *chat.Service
}

func NewChatHandler(chatService *chat.Service) *ChatHandler {
	return &ChatHandler{
		chatService: chatService,
	}
}

type chatRequest struct {
	ConversationID string `json:"conversation_id"`
	Message        string `json:"message"`
}

type chatResponse struct {
	ConversationID string `json:"conversation_id"`
	Reply          string `json:"reply"`
}

func (h *ChatHandler) Handle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req chatRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	resp, err := h.chatService.Handle(chat.ChatRequest{
		ConversationID: req.ConversationID,
		Message:        req.Message,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(chatResponse{
		ConversationID: resp.ConversationID,
		Reply:          resp.Reply,
	})
}
