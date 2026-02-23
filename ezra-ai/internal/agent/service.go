package agent

import (
	"fmt"

	"ezra-ai/internal/types"
)

func buildTradePrompt(req types.ChatRequest) string {
	m := req.Market

	return fmt.Sprintf(`
Analyze SPY 15-minute chart and provide ONE setup.

Market:
Price: %.2f
VWAP: %.2f
RSI: %.2f
EMA20: %.2f
EMA50: %.2f
EMA200: %.2f
Trend: %s
Session: %s
`,
		m.Price,
		m.VWAP,
		m.RSI,
		m.EMA20,
		m.EMA50,
		m.EMA200,
		m.Trend,
		m.Session,
	)
}
