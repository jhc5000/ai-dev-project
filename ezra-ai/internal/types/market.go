package types

type MarketSnapshot struct {
	Symbol    string  `json:"symbol"`
	Timeframe string  `json:"timeframe"`
	Price     float64 `json:"price"`
	VWAP      float64 `json:"vwap"`
	RSI       float64 `json:"rsi"`
	EMA20     float64 `json:"ema_20"`
	EMA50     float64 `json:"ema_50"`
	EMA200    float64 `json:"ema_200"`
	Trend     string  `json:"trend"`
	Session   string  `json:"session"`
}
