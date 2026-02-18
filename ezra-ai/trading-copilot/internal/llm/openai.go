package llm

import (
	"context"
	"os"

	openai "github.com/sashabaranov/go-openai"
)

type OpenAIClient struct {
	client *openai.Client
}

func NewOpenAIClient() *OpenAIClient {
	return &OpenAIClient{
		client: openai.NewClient(os.Getenv("OPENAI_API_KEY")),
	}
}

func (o *OpenAIClient) Generate(prompt string) (string, error) {
	resp, err := o.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT4oMini,
			Messages: []openai.ChatCompletionMessage{
				{Role: "system", Content: "You are a disciplined trading assistant."},
				{Role: "user", Content: prompt},
			},
			Temperature: 0.2,
		},
	)

	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}
