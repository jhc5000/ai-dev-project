package types

import "time"

type TradeJournalEntry struct {
	Timestamp time.Time `json:"timestamp"`
	Bias      string    `json:"bias"`
	Strategy  string    `json:"strategy"`
	Entry     float64   `json:"entry"`
	StopLoss float64   `json:"stop_loss"`
	Target   float64   `json:"target"`
	Notes    string    `json:"notes"`
}
