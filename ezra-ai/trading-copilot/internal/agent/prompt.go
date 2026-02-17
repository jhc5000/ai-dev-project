package agent

import "fmt"

func BuildPrompt(snapshot MarketSnapshot) string {
	return fmt.Sprintf(`
You are a professional SPY options trading assistant.
Analyze the 15-minute chart and recommend ONE high-probability setup.

Rules:
- Trade only SPY
- Timeframe: 15-minute
- Favor trend continuation unless clearly ranging
- Risk must be clearly defined
- No overtrading

Market Data:
Symbol: %s
Price: %.2f
VWAP: %.2f
RSI: %.2f
EMA20: %.2f
EMA50: %.2f
EMA200: %.2f
Trend: %s
Session: %s

Return your response in JSON with:
bias, strategy, entry, stop_loss, target, confidence, rationale
`, snapshot.Symbol,
		snapshot.Price,
		snapshot.VWAP,
		snapshot.RSI,
		snapshot.EMA20,
		snapshot.EMA50,
		snapshot.EMA200,
		snapshot.Trend,
		snapshot.Session,
	)
}
