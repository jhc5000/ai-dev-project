package types

type Intent string

const (
	IntentTrade   Intent = "trade"
	IntentCoach   Intent = "coach"
	IntentJournal Intent = "journal"
	IntentFriend  Intent = "friend"
)
