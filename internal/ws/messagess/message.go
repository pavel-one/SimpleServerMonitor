package messagess

type Message struct {
	Event string `json:"event"`
}

func (m *Message) GetEvent() string {
	return m.Event
}
