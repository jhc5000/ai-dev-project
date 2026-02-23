package agent

import (
	"fmt"

	openai "github.com/sashabaranov/go-openai"

	"ezra-ai/internal/journal"
	"ezra-ai/internal/llm"
	"ezra-ai/internal/types"
)

type ChatAgent struct {
	llm     llm.Client
	memory  *Memory
	journal *journal.Store
}

func NewChatAgent(llm llm.Client, journal *journal.Store) *ChatAgent {
	return &ChatAgent{
		llm:     llm,
		memory:  NewMemory(8),
		journal: journal,
	}
}

func (a *ChatAgent) Chat(req types.ChatRequest) (string, error) {
	intent := DetectIntent(req.Message)

	a.memory.Add("User: " + req.Message)

	system := openai.ChatCompletionMessage{
		Role: "system",
		Content: fmt.Sprintf(
			"You are EzraAI.\nConversation context:\n%s",
			a.memory.Context(),
		),
	}

	var userPrompt string

	switch intent {
	case "friend":
		userPrompt = buildFriendPrompt(req.Message)
	case "coach":
		userPrompt = "User wants trading psychology help:\n" + req.Message
	case "journal":
		userPrompt = "User wants to journal or reflect:\n" + req.Message
	default:
		userPrompt = buildTradePrompt(req)
	}

	user := openai.ChatCompletionMessage{
		Role:    "user",
		Content: userPrompt,
	}

	reply, err := a.llm.Chat([]openai.ChatCompletionMessage{system, user})
	if err != nil {
		return "", err
	}

	a.memory.Add("EzraAI: " + reply)
	return reply, nil
}
