package agent

type MarketSnapshot struct {
	Symbol    string  `json:"symbol"`
	Timeframe string  `json:"timeframe"` // "15m"
	Price     float64 `json:"price"`
	VWAP      float64 `json:"vwap"`
	RSI       float64 `json:"rsi"`
	EMA20     float64 `json:"ema_20"`
	EMA50     float64 `json:"ema_50"`
	EMA200    float64 `json:"ema_200"`
	Trend     string  `json:"trend"`   // bullish / bearish / range
	Session   string  `json:"session"` // premarket / RTH / power hour
}

type TradeSetup struct {
	Bias       string  `json:"bias"`     // bullish / bearish / neutral
	Strategy   string  `json:"strategy"` // e.g. "Pullback Call"
	Entry      float64 `json:"entry"`
	StopLoss   float64 `json:"stop_loss"`
	Target     float64 `json:"target"`
	Confidence int     `json:"confidence"` // 1–10
	Rationale  string  `json:"rationale"`
}
