package chat

import (
	"sync"
)

type Memory interface {
	Get(conversationID string) []Message
	Append(conversationID string, msg Message)
}

type InMemoryMemory struct {
	mu      sync.Mutex
	history map[string][]Message
}

func NewInMemoryMemory() *InMemoryMemory {
	return &InMemoryMemory{
		history: make(map[string][]Message),
	}
}

func (m *InMemoryMemory) Get(conversationID string) []Message {
	m.mu.Lock()
	defer m.mu.Unlock()

	return append([]Message{}, m.history[conversationID]...)
}

func (m *InMemoryMemory) Append(conversationID string, msg Message) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.history[conversationID] = append(m.history[conversationID], msg)
}
