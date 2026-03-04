package llm

import (
	"context"
	"ezra-ai/internal/config"

	openai "github.com/sashabaranov/go-openai"
)

type Client interface {
	Chat(messages []openai.ChatCompletionMessage) (string, error)
}

type OpenAIClient struct {
	client *openai.Client
}

func NewOpenAIClient() *OpenAIClient {
	cfg := config.Load()
	// c, d := fmt.Println(cfg.OpenAIKey)
	// fmt.Println("checking GetEnv:", c, d)

	// a, b := os.LookupEnv(cfg.OpenAIKey)
	// fmt.Println("checking key:", cfg.OpenAIKey)
	return &OpenAIClient{
		client: openai.NewClient(cfg.OpenAIKey),
	}
}

func (o *OpenAIClient) Chat(messages []openai.ChatCompletionMessage) (string, error) {
	resp, err := o.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:       openai.GPT4oMini,
			Messages:    messages,
			Temperature: 0.4,
		},
	)
	if err != nil {
		return "", err
	}
	return resp.Choices[0].Message.Content, nil
}
