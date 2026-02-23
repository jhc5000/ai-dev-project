package llm

// Message mirrors OpenAI / Anthropic chat formats
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// ChatRequest is what we send to the LLM
type ChatRequest struct {
	Model       string    `json:"model"`
	Messages    []Message `json:"messages"`
	Temperature float32   `json:"temperature"`
}

// ChatResponse is the normalized output we care about
type ChatResponse struct {
	Content string
}
