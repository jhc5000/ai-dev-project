package llm

import (
	"context"
	"errors"
	"time"

	openai "github.com/sashabaranov/go-openai"

	"ezra-ai/internal/config"
)

type OpenAIClient struct {
	client *openai.Client
	model  string
}

func NewClient(cfg config.Config) (*OpenAIClient, error) {
	if cfg.OpenAIKey == "" {
		return nil, errors.New("missing OPENAI_API_KEY")
	}

	client := openai.NewClient(cfg.OpenAIKey)

	return &OpenAIClient{
		client: client,
		model:  "gpt-4o-mini", // fast + cheap for personal use
	}, nil
}

func (c *OpenAIClient) Chat(messages []Message) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	req := openai.ChatCompletionRequest{
		Model:       c.model,
		Temperature: 0.3, // disciplined, not creative
		Messages:    toOpenAIMessages(messages),
	}

	resp, err := c.client.CreateChatCompletion(ctx, req)
	if err != nil {
		return "", err
	}

	if len(resp.Choices) == 0 {
		return "", errors.New("empty LLM response")
	}

	return resp.Choices[0].Message.Content, nil
}

func toOpenAIMessages(messages []Message) []openai.ChatCompletionMessage {
	out := make([]openai.ChatCompletionMessage, 0, len(messages))

	for _, m := range messages {
		out = append(out, openai.ChatCompletionMessage{
			Role:    m.Role,
			Content: m.Content,
		})
	}

	return out
}
