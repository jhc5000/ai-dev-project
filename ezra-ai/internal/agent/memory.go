package agent

type Memory struct {
	history []string
	limit   int
}

func NewMemory(limit int) *Memory {
	return &Memory{limit: limit}
}

func (m *Memory) Add(msg string) {
	m.history = append(m.history, msg)
	if len(m.history) > m.limit {
		m.history = m.history[1:]
	}
}

func (m *Memory) Context() string {
	out := ""
	for _, h := range m.history {
		out += "- " + h + "\n"
	}
	return out
}
