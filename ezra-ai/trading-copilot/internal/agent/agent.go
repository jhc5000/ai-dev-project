package agent

import "ezra-ai/trading-copilot/internal/llm"

type TradingAgent struct {
	llm llm.Client
}

func NewTradingAgent(llm llm.Client) *TradingAgent {
	return &TradingAgent{llm: llm}
}

func (a *TradingAgent) Analyze(snapshot MarketSnapshot) (string, error) {
	prompt := BuildPrompt(snapshot)
	return a.llm.Generate(prompt)
}
