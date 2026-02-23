package agent

import "strings"

func DetectIntent(message string) string {
	msg := strings.ToLower(message)

	switch {
	case strings.Contains(msg, "journal"):
		return "journal"
	case strings.Contains(msg, "feel") ||
		strings.Contains(msg, "overtrading") ||
		strings.Contains(msg, "discipline"):
		return "coach"
	case strings.Contains(msg, "hey") ||
		strings.Contains(msg, "what's up") ||
		strings.Contains(msg, "lol") ||
		strings.Contains(msg, "just chatting"):
		return "friend"
	case strings.Contains(msg, "trade") ||
		strings.Contains(msg, "trading") ||
		strings.Contains(msg, "stock market") ||
		strings.Contains(msg, "spy symbol") ||
		strings.Contains(msg, "wall street") ||
		strings.Contains(msg, "position") ||
		strings.Contains(msg, "setup strategy") ||
		strings.Contains(msg, "options") ||
		strings.Contains(msg, "put option") ||
		strings.Contains(msg, "call option") ||
		strings.Contains(msg, "strike"):
		return "trade"
	default:
		return "friend"
	}
}
