package journal

import "ezra-ai/internal/types"

type Store struct {
	entries []types.TradeJournalEntry
}

func NewStore() *Store {
	return &Store{}
}

func (s *Store) Add(entry types.TradeJournalEntry) {
	s.entries = append(s.entries, entry)
}

func (s *Store) All() []types.TradeJournalEntry {
	return s.entries
}
