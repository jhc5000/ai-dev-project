package types

type ChatRequest struct {
	Message string         `json:"message"`
	Market  MarketSnapshot `json:"market"`
}

type ChatResponse struct {
	Reply string `json:"reply"`
}
