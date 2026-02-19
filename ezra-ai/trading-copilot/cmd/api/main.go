package main

import (
	"ezra-ai/trading-copilot/internal/agent"
	"ezra-ai/trading-copilot/internal/llm"
	"fmt"
)

func main() {
	llmClient := llm.NewOpenAIClient()
	newAgent := agent.NewTradingAgent(llmClient)

	snapshot := agent.MarketSnapshot{
		Symbol:    "SPY",
		Timeframe: "15m",
		Price:     487.25,
		VWAP:      486.80,
		RSI:       58.4,
		EMA20:     487.10,
		EMA50:     485.90,
		EMA200:    480.30,
		Trend:     "bullish",
		Session:   "RTH",
	}

	result, err := newAgent.Analyze(snapshot)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}
