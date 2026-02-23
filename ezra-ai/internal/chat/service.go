package chat

import (
	"errors"

	"github.com/google/uuid"
)

type LLMClient interface {
	Chat(messages []Message) (string, error)
}

type Service struct {
	llm    LLMClient
	memory Memory
}

func NewService(llm LLMClient, memory Memory) *Service {
	return &Service{
		llm:    llm,
		memory: memory,
	}
}

type ChatRequest struct {
	ConversationID string
	Message        string
}

type ChatResponse struct {
	ConversationID string `json:"conversation_id"`
	Reply          string `json:"reply"`
}

func (s *Service) Handle(req ChatRequest) (*ChatResponse, error) {
	if req.Message == "" {
		return nil, errors.New("message cannot be empty")
	}

	conversationID := req.ConversationID
	if conversationID == "" {
		conversationID = uuid.NewString()
	}

	// 1️⃣ Store user message
	userMsg := Message{
		Role:    "user",
		Content: req.Message,
	}
	s.memory.Append(conversationID, userMsg)

	// 2️⃣ Build full message context
	messages := []Message{
		SystemPrompt(),
	}

	history := s.memory.Get(conversationID)
	messages = append(messages, history...)

	// 3️⃣ Call LLM
	reply, err := s.llm.Chat(messages)
	if err != nil {
		return nil, err
	}

	assistantMsg := Message{
		Role:    "assistant",
		Content: reply,
	}

	// 4️⃣ Store assistant response
	s.memory.Append(conversationID, assistantMsg)

	// 5️⃣ Return response
	return &ChatResponse{
		ConversationID: conversationID,
		Reply:          reply,
	}, nil
}
