package events

type Interface interface {
	GetEvent() string
	GetData() []byte
}

type event struct {
	Event string `json:"event"`
	Data  []byte
}

func (e *event) GetEvent() string {
	return e.Event
}

func (e *event) GetData() []byte {
	return e.Data
}

type Chan chan Interface
